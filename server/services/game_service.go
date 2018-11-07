package services

import (
	"fmt"
	"sync"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/bitdecaygames/fireport/server/rules"
)

// GameService is responsible for managing our active games
type GameService interface {
	CreateGame(lobby Lobby) *GameInstance
	SubmitTurn(submit pogo.TurnSubmissionMsg) error
	GetCurrentTurn(gameID string) (int, error)
}

// GameInstance is a logical game instance which holds a game
// state as well as the information needed to communicate changes
// to the connected clients
type GameInstance struct {
	Lock              *sync.Mutex
	Name              string
	State			  pogo.GameState
	ID                string
	CurrentTurn       int
	Players           []string
	Rules             []rules.GameRule
	ActiveConnections map[string]PlayerConnection
	PlayerSubmissions map[string]pogo.TurnSubmissionMsg
	InputRules        []rules.InputRule
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
	newInstance := &GameInstance{
		Lock:              &sync.Mutex{},
		Name:              lobby.Name,
		ID:                lobby.ID,
		State:			   createInitialGameState(lobby),
		Players:           lobby.Players,
		Rules:             rules.DefaultGameRules,
		ActiveConnections: lobby.ActiveConnections,
		InputRules:        rules.DefaultInputRules,

		PlayerSubmissions: make(map[string]pogo.TurnSubmissionMsg),
	}
	g.activeGames[newInstance.ID] = newInstance
	return newInstance
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

// SubmitTurn will accept client input and step the game once all players have a submission
func (g *GameServiceImpl) SubmitTurn(submit pogo.TurnSubmissionMsg) error {
	game, err := g.lockActiveGame(submit.GameID)
	if err != nil {
		return err
	}
	defer game.Lock.Unlock()

	_, alreadySubmitted := game.PlayerSubmissions[submit.PlayerID]
	if alreadySubmitted {
		return fmt.Errorf("Player %v already has a pending turn submission", submit.PlayerID)
	}

	game.PlayerSubmissions[submit.PlayerID] = submit

	// TODO: If all player turns are submitted, step the game
	return nil
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
func createInitialGameState(lobby Lobby) pogo.GameState {
	var playerStates []pogo.PlayerState
	
	for i, player := range lobby.Players{
		playerStates = append(playerStates , createInitialPlayerStates(player, i ))
	}

	return pogo.GameState{
		Turn: 			0,
		Created: 		0,
		Updated: 		0,
		IDCounter:		0,
		BoardWidth:		3,
		BoardHeight: 	3,
		BoardSpaces: 	[]pogo.BoardSpace{
			{ID: 0, SpaceType: 0, State: 0},
			{ID: 1, SpaceType: 0, State: 0},
			{ID: 2, SpaceType: 0, State: 0},
			{ID: 3, SpaceType: 0, State: 0},
			{ID: 4, SpaceType: 0, State: 0},
			{ID: 5, SpaceType: 0, State: 0},
			{ID: 6, SpaceType: 0, State: 0},
			{ID: 7, SpaceType: 0, State: 0},
			{ID: 8, SpaceType: 0, State: 0},
		},
		Players: 		playerStates,
	}
}
// createInitialCards returns a slice of CardStates for the initial discard pile, can probably be refactored to pull a list of playable/implimented cards
func createInitialCards() []pogo.CardState {
	return []pogo.CardState{
		{ID: 101, CardType: pogo.MoveForwardOne},
		{ID: 102, CardType: pogo.MoveForwardOne},
		{ID: 103, CardType: pogo.MoveForwardTwo},
		{ID: 104, CardType: pogo.MoveForwardTwo},
		{ID: 105, CardType: pogo.MoveForwardThree},
		{ID: 106, CardType: pogo.MoveForwardThree},
		{ID: 107, CardType: pogo.MoveBackwardOne},
		{ID: 108, CardType: pogo.MoveBackwardOne},
		{ID: 109, CardType: pogo.MoveBackwardTwo},
		{ID: 110, CardType: pogo.MoveBackwardTwo},
		{ID: 111, CardType: pogo.MoveBackwardThree},
		{ID: 112, CardType: pogo.MoveBackwardThree},
		{ID: 113, CardType: pogo.TurnRight},
		{ID: 114, CardType: pogo.TurnRight},
	}
}
//createInitialPlayerStates creates the inital state for a player, probably needs a list of available starting locations
func createInitialPlayerStates(playerName string, playerId int) pogo.PlayerState {	
	return pogo.PlayerState{
		ID:			playerId,
		Name:		playerName,
		Location:	playerId,
		Hand:		[]pogo.CardState{},
		Deck:		[]pogo.CardState{},
		Discard:	createInitialCards(),
	}
}