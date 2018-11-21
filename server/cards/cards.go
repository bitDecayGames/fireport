package cards

import (
	"fmt"

	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// Card defines the list of actions for a given card
type Card struct {
	ID       int
	CardType pogo.CardType
	Owner    int
	Actions  []actions.Action
}

// Apply apply the list of actions from this card to the game state
func (c *Card) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	var nextState = currentState
	for _, action := range c.Actions {
		nxt, err := action.Apply(nextState)
		if err != nil {
			return nextState, err
		}
		nextState = nxt
	}
	return nextState, nil
}

// GameInputToCard Build a Card object with a card id, player id, and card type
func GameInputToCard(cardID int, playerID int, cardType pogo.CardType) (*Card, error) {
	switch cardType {
	case pogo.SkipTurn:
		return skipTurnCard(cardID, playerID), nil
	case pogo.MoveForwardOne:
		return moveForwardOneCard(cardID, playerID), nil
	case pogo.MoveForwardTwo:
		return moveForwardTwoCard(cardID, playerID), nil
	case pogo.MoveForwardThree:
		return moveForwardThreeCard(cardID, playerID), nil
	case pogo.MoveBackwardOne:
		return moveBackwardOneCard(cardID, playerID), nil
	case pogo.MoveBackwardTwo:
		return moveBackwardTwoCard(cardID, playerID), nil
	case pogo.MoveBackwardThree:
		return moveBackwardThreeCard(cardID, playerID), nil
	case pogo.TurnRight:
		return turnRightCard(cardID, playerID), nil
	case pogo.TurnLeft:
		return turnLeftCard(cardID, playerID), nil
	case pogo.FireBasic:
		return fireBasicCard(cardID, playerID), nil
	default:
		return nil, fmt.Errorf("no card found with the type %v", cardType)
	}
}

// Builds the skip turn card, just a test card to figure out how these all work together
func skipTurnCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.SkipTurn, Actions: []actions.Action{&actions.IncrementTurnAction{}}}
}

func moveForwardOneCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.MoveForwardOne, Actions: []actions.Action{&actions.MoveForwardAction{Owner: owner}}}
}

func moveForwardTwoCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.MoveForwardTwo, Actions: []actions.Action{&actions.MoveForwardAction{Owner: owner}, &actions.MoveForwardAction{Owner: owner}}}
}

func moveForwardThreeCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.MoveForwardThree, Actions: []actions.Action{&actions.MoveForwardAction{Owner: owner}, &actions.MoveForwardAction{Owner: owner}, &actions.MoveForwardAction{Owner: owner}}}
}

func moveBackwardOneCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.MoveBackwardOne, Actions: []actions.Action{&actions.MoveBackwardAction{Owner: owner}}}
}

func moveBackwardTwoCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.MoveBackwardTwo, Actions: []actions.Action{&actions.MoveBackwardAction{Owner: owner}, &actions.MoveBackwardAction{Owner: owner}}}
}

func moveBackwardThreeCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.MoveBackwardThree, Actions: []actions.Action{&actions.MoveBackwardAction{Owner: owner}, &actions.MoveBackwardAction{Owner: owner}, &actions.MoveBackwardAction{Owner: owner}}}
}

func turnRightCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.TurnRight, Actions: []actions.Action{&actions.MoveForwardAction{Owner: owner}, &actions.TurnClockwise90Action{Owner: owner}, &actions.MoveForwardAction{Owner: owner}}}
}

func turnLeftCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.TurnLeft, Actions: []actions.Action{&actions.MoveForwardAction{Owner: owner}, &actions.TurnCounterClockwise90Action{Owner: owner}, &actions.MoveForwardAction{Owner: owner}}}
}

func fireBasicCard(id int, owner int) *Card {
	return &Card{ID: id, Owner: owner, CardType: pogo.FireBasic, Actions: []actions.Action{&actions.FireBasicAction{Owner: owner}}}
}
