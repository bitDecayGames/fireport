package services

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/conditions"
	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/bitdecaygames/fireport/server/triggers"
	"sync"

	"github.com/bitdecaygames/fireport/server/rules"
	"github.com/satori/go.uuid"
)

// GameService is responsible for managing our active games
type GameService interface {
	CreateGame(lobby Lobby) *Game
	GetActiveGame(gameID uuid.UUID) (Game, error)
}

// Game is a group of players in a game
type Game struct {
	Name              string
	ID                uuid.UUID
	State			  pogo.GameState
	Players           []string
	Rules             []rules.GameRule
	ActiveConnections map[string]PlayerConnection
	InputRules        []rules.InputRule
}

var gameMutex = &sync.Mutex{}

// GameServiceImpl is a concrete service
type GameServiceImpl struct {
	activeGames []*Game
}

// CreateGame creates a new Game from the lobby information and returns it
func (g *GameServiceImpl) CreateGame(lobby Lobby) *Game {
	gameMutex.Lock()
	defer gameMutex.Unlock()

	newGame := &Game{
		Name:              lobby.Name,
		ID:                lobby.ID,
		State:			   createInitialGameState(lobby),
		Players:           lobby.Players,
		Rules:             rules.DefaultGameRules,
		ActiveConnections: lobby.ActiveConnections,
		InputRules:        rules.DefaultInputRules,
	}
	g.activeGames = append(g.activeGames, newGame)
	return newGame
}

// GetActiveGame returns a copy of the active Game for an ID or an error if the game is not found
func (g *GameServiceImpl) GetActiveGame(gameID uuid.UUID) (Game, error) {
	gameMutex.Lock()
	defer gameMutex.Unlock()

	for _, game := range g.activeGames {
		if game.ID == gameID {
			return *game, nil
		}
	}
	return Game{}, fmt.Errorf("no game found with uuid '%v'", gameID.String())
}

// StepGame moves the game state forward using a list of inputs
func (g *GameServiceImpl) StepGame(currentState *pogo.GameState, inputs []pogo.GameInputMsg) (*pogo.GameState, error) {
	// process conditions and apply all action groups
	var nextState, err = conditions.ProcessConditions(currentState, inputs, []conditions.Condition{ // TODO: MW Conditions should probably be a part of the Game struct
		&conditions.SpaceCollisionCondition{},
		&conditions.EdgeCollisionCondition{},
	})
	if err != nil {
		return nextState, err
	}

	// each step of the game should Apply the DefaultTurnActions list
	nextState, err = actions.ApplyActions(nextState, actions.DefaultTurnActions)
	if err != nil {
		return nextState, err
	}

	var cardIDs []int
	for i := range inputs {
		cardIDs = append(cardIDs, inputs[i].CardID)
	}

	// apply our post-step triggers
	nextState, err = triggers.ApplyTriggers(nextState, triggers.DefaultPostStepTriggers(5, cardIDs)) // TODO: MW magic number alert
	if err != nil {
		return nextState, err
	}

	// TODO: MW there needs to be a way to track what actions have been successfully applied each step
	return nextState, nil
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