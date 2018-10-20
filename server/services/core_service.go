package services

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// CoreService is a stateless service that generates new game states given a set of inputs
type CoreService interface {
	StepGame(currentState *pogo.GameState, inputs []pogo.GameInputMsg) (*pogo.GameState, error)
}

// CoreServiceImpl is a concrete service
type CoreServiceImpl struct {
}

// CreateGame creates a new Game from the lobby information and returns it
func (g *CoreServiceImpl) StepGame(currentState *pogo.GameState, inputs []pogo.GameInputMsg) (*pogo.GameState, error) {
	var nextState = currentState
	for _, input := range inputs {
		card, err := actions.GameInputToCard(&input)
		if err != nil {
			return nextState, err
		}
		nxt, err := card.Apply(nextState)
		if err != nil {
			return nextState, err
		}
		nextState = nxt
	}
	// TODO: MW each step of the game should Apply the IncrementTurnAction
	// TODO: MW there needs to be a way to track what actions have been successfully applied each step
	return nextState, nil
}
