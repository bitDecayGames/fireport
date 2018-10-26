package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// Card defines the list of actions for a given card
type Card struct {
	ID      int
	Owner   int
	Actions []Action
}

// Apply apply the list of actions from this card to the game state
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

// GameInputToCard Build a Card object with a GameInputMsg object
func GameInputToCard(input *pogo.GameInputMsg) (*Card, error) {
	switch input.CardID {
	case 0:
		return skipTurnCard(input.Owner), nil
	default:
		return nil, fmt.Errorf("no card found with the id '%v'", string(input.CardID))
	}
}

// Builds the skip turn card, just a test card to figure out how these all work together
func skipTurnCard(owner int) *Card {
	return &Card{ID: 0, Owner: owner, Actions: []Action{&IncrementTurnAction{}}}
}
