package actions

import (
	"github.com/bitdecaygames/fireport/server/pogo"
)

// Action the smallest unit of modification to be made to a game state object
type Action interface {
	Apply(currentState *pogo.GameState) (*pogo.GameState, error)
	GetOwner() int
}

// ApplyActions applies a list of actions to a given state
func ApplyActions(currentState *pogo.GameState, actions []Action) (*pogo.GameState, error) {
	var nextState = currentState
	for _, action := range actions {
		nxt, err := action.Apply(nextState)
		if err != nil {
			return nextState, err
		}
		nextState = nxt
	}
	return nextState, nil
}