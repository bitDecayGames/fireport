package actions

import "github.com/bitdecaygames/fireport/server/pogo"

type Action interface {
	Apply(currentState *pogo.GameState) (*pogo.GameState, error)
}

// Increment the Turn by 1
type IncrementTurnAction struct {
	Owner int
}
func (a *IncrementTurnAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := *currentState // copy the contents of the game state
	nextState.Turn = nextState.Turn + 1
	return &nextState, nil
}

