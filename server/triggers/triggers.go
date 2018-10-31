package triggers

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// Trigger encompasses actions that are only applied to the state if a condition is true
type Trigger interface {
	Check(currentState *pogo.GameState) bool
	GetActions() []actions.Action
}

// ApplyTriggers apply a list of triggers to a given game state
func ApplyTriggers(currentState *pogo.GameState, triggers []Trigger) (*pogo.GameState, error) {
	var nextState = currentState
	var dirty = false
	for _, trigger := range triggers {
		if trigger.Check(nextState) {
			dirty = true
			nxt, err := actions.ApplyActions(nextState, trigger.GetActions())
			if err != nil {
				return nextState, err
			}
			nextState = nxt
		}
	}
	if !dirty {
		return nextState, nil
	}
	return ApplyTriggers(nextState, triggers)
}
