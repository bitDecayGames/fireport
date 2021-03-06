package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/bitdecaygames/fireport/server/files"

	"github.com/bitdecaygames/fireport/server/logic"
	"github.com/sirupsen/logrus"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/bitdecaygames/fireport/server/rules"
)

// GameService is responsible for managing our active games
type GameService interface {
	CreateGame(lobby Lobby) *GameInstance
	SubmitTurn(submit pogo.TurnSubmissionMsg) error
	GetCurrentTurn(gameID string) (int, error)
	GetCurrentState(gameID string) (*pogo.GameState, error)
	SubmitSimpleTestTurn(gameID string, playerName string, playerID int, cards []int) error
}

// GameInstance is a logical game instance which holds a game
// state as well as the information needed to communicate changes
// to the connected clients
type GameInstance struct {
	Lock              *sync.Mutex
	Name              string
	State             pogo.GameState
	ID                string
	CurrentTurn       int
	Players           []string
	Rules             []rules.GameRule
	InputRules        []rules.InputRule
	ActiveConnections map[string]PlayerConnection
	PlayerSubmissions map[string]pogo.TurnSubmissionMsg
	Log               *logrus.Logger
}

var gameMutex = &sync.Mutex{}

// GameServiceImpl is a concrete service
type GameServiceImpl struct {
	activeGames map[string]*GameInstance
}

// NewGameService returns a new instance of a GameService
func NewGameService() GameService {
	return &GameServiceImpl{
		activeGames: make(map[string]*GameInstance),
	}
}

// CreateGame creates a new Game from the lobby information and returns it
func (g *GameServiceImpl) CreateGame(lobby Lobby) *GameInstance {
	seed := time.Now().UnixNano()
	newInstance := &GameInstance{
		Lock:              &sync.Mutex{},
		Name:              lobby.Name,
		ID:                lobby.ID,
		State:             createInitialGameState(lobby, seed),
		Players:           lobby.Players,
		Rules:             rules.DefaultGameRules,
		InputRules:        rules.DefaultInputRules,
		ActiveConnections: lobby.ActiveConnections,
		PlayerSubmissions: make(map[string]pogo.TurnSubmissionMsg),
		Log:               getGameLogger(lobby.ID),
	}
	g.activeGames[newInstance.ID] = newInstance
	fmt.Println("Game created: ", newInstance.ID)
	newInstance.Log.Info("Game created with seed: ", seed)
	return newInstance
}

func getGameLogger(gameID string) *logrus.Logger {
	f, err := files.GetLogFile(gameID)
	if err != nil {
		fmt.Println("failed to make file logger for game ", gameID, ": ", err)
	}

	logger := logrus.New()
	logger.SetOutput(f)

	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "15:04:05"
	formatter.FullTimestamp = true
	logger.SetFormatter(formatter)

	return logger
}

// GetCurrentTurn returns the current turn of an active game, or an error
// if not game is found with the given ID
func (g *GameServiceImpl) GetCurrentTurn(gameID string) (int, error) {
	game, err := g.lockActiveGame(gameID)
	if err != nil {
		return -1, err
	}
	defer game.Lock.Unlock()

	return game.CurrentTurn, nil
}

// GetCurrentState returns the current state of an active game, or an error
// if not game is found with the given ID
func (g *GameServiceImpl) GetCurrentState(gameID string) (*pogo.GameState, error) {
	game, err := g.lockActiveGame(gameID)
	if err != nil {
		return nil, err
	}
	defer game.Lock.Unlock()

	return &game.State, nil
}

// SubmitTurn will accept client input and step the game once all players have a submission
func (g *GameServiceImpl) SubmitTurn(submit pogo.TurnSubmissionMsg) error {
	game, err := g.lockActiveGame(submit.GameID)
	if err != nil {
		game.sendMessageOverWebSocketConnections(&pogo.GameErrorMsg{Error: err.Error()})
		return err
	}
	defer game.Lock.Unlock()

	_, alreadySubmitted := game.PlayerSubmissions[submit.PlayerID]
	if alreadySubmitted {
		return fmt.Errorf("Player %v already has a pending turn submission", submit.PlayerID)
	}

	err = rules.ApplyInputRules(&game.State, submit.Inputs, game.InputRules)
	if err != nil {
		game.sendMessageOverWebSocketConnections(&pogo.GameErrorMsg{Error: err.Error()})
		return err
	}

	game.PlayerSubmissions[submit.PlayerID] = submit

	allTurnsSubmitted := true
	for _, pid := range game.Players {
		if _, found := game.PlayerSubmissions[pid]; !found {
			// still waiting on player turns to come in
			allTurnsSubmitted = false
			break
		}
	}

	if allTurnsSubmitted {
		game.Log.Infof("Stepping game for turn %v", game.CurrentTurn)
		allInputs := make([]pogo.GameInputMsg, 0)
		for _, msg := range game.PlayerSubmissions {
			allInputs = append(allInputs, msg.Inputs...)
		}
		logGameTurn(game, allInputs)
		game.PlayerSubmissions = map[string]pogo.TurnSubmissionMsg{}
		// TODO: Does it make sense to pass pointers through all the logic, or just structs?
		oldState := game.State
		newState, err := logic.StepGame(&game.State, allInputs)
		if err != nil {
			game.sendMessageOverWebSocketConnections(&pogo.GameErrorMsg{Error: err.Error()})
			return err
		}

		err = rules.ApplyGameRules(&oldState, newState, game.Rules)
		if err != nil {
			game.sendMessageOverWebSocketConnections(&pogo.GameErrorMsg{Error: err.Error()})
			return err
		}

		msg := &pogo.TurnResultMsg{
			GameID:        game.ID,
			PreviousState: oldState,
			CurrentState:  *newState,
		}

		game.State = *newState

		for pid, conn := range game.ActiveConnections {
			err = conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("failed to send state to player ", pid)
			}
		}
	}

	return nil
}

func (g *GameInstance) sendMessageOverWebSocketConnections(msg pogo.Typer) {
	for pid, conn := range g.ActiveConnections {
		err := conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("failed to send message to player ", pid)
		}
	}
}

// lockActiveGame locks the game with the given gameID and returns it, or returns an error
// if no game found with the given ID. BE SURE TO UNLOCK THE GAME WHEN FINISHED
func (g *GameServiceImpl) lockActiveGame(gameID string) (*GameInstance, error) {
	game, ok := g.activeGames[gameID]
	if !ok {
		return nil, fmt.Errorf("no game found with uuid '%v'", gameID)
	}

	// TODO: may want to have some sort of timeout here
	fmt.Println("locking game " + gameID)
	game.Lock.Lock()
	return game, nil
}

//createInitialGameState creates the initial state for the lobby, probably should call some board creation method to ensure width, height and tile types are set accordingly
func createInitialGameState(lobby Lobby, seedValue int64) pogo.GameState {
	var playerStates []pogo.PlayerState
	gameState := pogo.GameState{
		Turn:        0,
		RNG:         rand.New(rand.NewSource(seedValue)),
		Created:     time.Now().Unix(),
		Updated:     time.Now().Unix(),
		IDCounter:   0,
		BoardWidth:  6, // TODO: MW magic number alert
		BoardHeight: 6, // TODO: MW magic number alert
	}

	gameState.BoardSpaces = createBoard(&gameState)
	nums := getRandomPlayerSpaces(gameState.RNG, len(lobby.Players), len(gameState.BoardSpaces))

	for i, player := range lobby.Players {
		playerStates = append(playerStates, createInitialPlayerState(player, nums[i], &gameState))
	}

	gameState.Players = playerStates

	// TODO: MW maybe this shouldn't go here? Like maybe it should go one up from this method
	finalState, err := logic.StepGame(&gameState, nil)
	if err != nil {
		panic(err)
	}

	return *finalState
}

func getRandomPlayerSpaces(rand *rand.Rand, num, numRange int) []int {
	if numRange < num {
		panic("cannot have more players than tile spaces")
	}
	nums := make([]int, 0)
	valid := true
	for len(nums) < num {
		pick := rand.Intn(numRange)

		for _, i := range nums {
			if i == pick {
				valid = false
			}
		}
		
		if valid {
			nums = append(nums, pick)
		}
	}
	return nums
}

// createInitialCards returns a slice of CardStates for the initial discard pile, can probably be refactored to pull a list of playable/implimented cards
func createInitialCards(gameState *pogo.GameState) []pogo.CardState {
	return []pogo.CardState{
		{ID: gameState.GetNewID(), CardType: pogo.MoveForwardOne},
		{ID: gameState.GetNewID(), CardType: pogo.MoveForwardTwo},
		{ID: gameState.GetNewID(), CardType: pogo.MoveForwardThree},
		{ID: gameState.GetNewID(), CardType: pogo.MoveBackwardOne},
		{ID: gameState.GetNewID(), CardType: pogo.MoveBackwardTwo},
		{ID: gameState.GetNewID(), CardType: pogo.MoveBackwardThree},
		{ID: gameState.GetNewID(), CardType: pogo.TurnRight},
		{ID: gameState.GetNewID(), CardType: pogo.TurnRight},
		{ID: gameState.GetNewID(), CardType: pogo.TurnRight},
		{ID: gameState.GetNewID(), CardType: pogo.TurnLeft},
		{ID: gameState.GetNewID(), CardType: pogo.TurnLeft},
		{ID: gameState.GetNewID(), CardType: pogo.TurnLeft},
		{ID: gameState.GetNewID(), CardType: pogo.FireBasic},
		{ID: gameState.GetNewID(), CardType: pogo.FireBasic},
		{ID: gameState.GetNewID(), CardType: pogo.FireBasic},
		{ID: gameState.GetNewID(), CardType: pogo.FireBasic},
		{ID: gameState.GetNewID(), CardType: pogo.FireBasic},
	}
}

//createInitialPlayerStates creates the inital state for a player, probably needs a list of available starting locations
func createInitialPlayerState(playerName string, playerLocation int, gameState *pogo.GameState) pogo.PlayerState {
	return pogo.PlayerState{
		ID:       gameState.GetNewID(),
		Name:     playerName,
		Location: playerLocation,
		Health:   10, // TODO: MW magic number alert
		Hand:     []pogo.CardState{},
		Deck:     []pogo.CardState{},
		Discard:  createInitialCards(gameState),
	}
}

//createBoard creates a board, will need to accept some type of identifier down the line if we want multiple maps
func createBoard(gameState *pogo.GameState) []pogo.BoardSpace {
	var boardSpaces []pogo.BoardSpace
	for i := 0; i < gameState.BoardWidth*gameState.BoardHeight; i++ {
		boardSpaces = append(boardSpaces, pogo.BoardSpace{ID: gameState.GetNewID(), SpaceType: 0, State: 0})
	}
	return boardSpaces
}

func logGameTurn(game *GameInstance, inputs []pogo.GameInputMsg) {
	data, err := json.Marshal(game.State)
	if err != nil {
		game.Log.Error("Error marshalling game state for logging: ", err, game.State)
	} else {
		state := string(data)
		game.Log.Info("Game Turn: ")
		game.Log.Info(state)
		for _, input := range inputs {
			msgData, err := json.Marshal(input)
			if err != nil {
				game.Log.Error("Error marshalling input message for logging: ", err, input)
			} else {
				inputStr := string(msgData)
				game.Log.Info(inputStr)
			}
		}

	}
}
