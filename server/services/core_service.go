package services

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/conditions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// CoreService is a stateless service that generates new game states given a set of inputs
type CoreService interface {
	StepGame(currentState *pogo.GameState, inputs []pogo.GameInputMsg) (*pogo.GameState, error)
}

// CoreServiceImpl is a concrete service
type CoreServiceImpl struct {
}

// StepGame moves the game state forward using a list of inputs
func (g *CoreServiceImpl) StepGame(currentState *pogo.GameState, inputs []pogo.GameInputMsg) (*pogo.GameState, error) {
	var nextState, err = conditions.ProcessConditions(currentState, inputs, []conditions.Condition{
		&conditions.SpaceCollisionCondition{},
	})
	if err != nil {
		return nextState, err
	}
	// each step of the game should Apply the DefaultTurnActions list
	for _, defaultTurnAction := range actions.DefaultTurnActions {
		nxt, err := defaultTurnAction.Apply(nextState)
		if err != nil {
			return nextState, err
		}
		nextState = nxt
	}
	// TODO: MW there needs to be a way to track what actions have been successfully applied each step
	return nextState, nil
}
