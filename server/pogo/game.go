package pogo

import (
	"github.com/satori/go.uuid"
)

// GameStartMsg contains all information a client needs to start
// playing in a game
type GameStartMsg struct {
	GameID uuid.UUID
}

// TurnTimeoutMsg contains information for when a player's turn has timed out
type TurnTimeoutMsg struct {
	// Currently no needed information
}

// GameTimeoutMsg contains information for when a game has had no activity
// and will be ended by the server
type GameTimeoutMsg struct {
	// Currently no needed information
}

// TurnResultMsg contains information on game state changes that occurred
// in the latest turn
type TurnResultMsg struct {
	GameID        uuid.UUID
	PreviousState GameState
	CurrentState  GameState
	Actions       []PlayerAction
	PlayerOptions []PlayerAction
}

// GameState is an all-containing set of information for the state of a
// game at a given point in time
type GameState struct {
	TurnsTaken int
	// TODO: Fill out the rest of the state
}

type GameInput struct {
	CardId int
	Owner  int
}

type GameResultSummaryMsg struct {
	Winner string
	// TODO: attach stats and other summary information
}
