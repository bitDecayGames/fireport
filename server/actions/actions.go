package actions

import "github.com/bitdecaygames/fireport/server/pogo"

type Action interface {
	Apply(currentState *pogo.GameState) (*pogo.GameState, error)
}

// Increment the TurnsTaken by 1
type IncrementTurnAction struct {
	Owner int
}

func (a *IncrementTurnAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := pogo.GameState{TurnsTaken: currentState.TurnsTaken + 1} // TODO: MW need an easy way to copy... I think there are libraries for this
	return &nextState, nil
}
