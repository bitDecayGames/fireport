package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

type Card struct {
	Id      int
	Owner   int
	Actions []Action
}

func (c *Card) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	// TODO: MW this method should validate that the card being played is in the owners hand
	var nextState *pogo.GameState = currentState
	for _, action := range c.Actions {
		nxt, err := action.Apply(nextState)
		if err != nil {
			return nextState, err
		}
		nextState = nxt
	}
	return nextState, nil
}

// Build a Card object with a GameInput object
func GameInputToCard(input *pogo.GameInput) (*Card, error) {
	switch input.CardId {
	case 0:
		return skipTurnCard(input.Owner), nil
	default:
		return nil, fmt.Errorf("no card found with the id '%v'", string(input.CardId))
	}
}

// Builds the skip turn card, just a test card to figure out how these all work together
func skipTurnCard(owner int) *Card {
	return &Card{Id: 0, Owner: owner, Actions: []Action{&IncrementTurnAction{Owner: owner}}}
}
