package services

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/conditions"
	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/bitdecaygames/fireport/server/triggers"
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
	nextState, err = triggers.ApplyTriggers(nextState, triggers.DefaultPostStepTriggers(5, cardIDs))
	if err != nil {
		return nextState, err
	}

	// TODO: MW there needs to be a way to track what actions have been successfully applied each step
	return nextState, nil
}
