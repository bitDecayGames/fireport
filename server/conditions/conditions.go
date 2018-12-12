package conditions

import (
	"fmt"

	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/animations"
	"github.com/bitdecaygames/fireport/server/cards"
	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/pkg/errors"
)

const (
	// MaxConditionSteps the maximum number of steps a condition will make before throwing an error
	MaxConditionSteps = 10
)

// Condition checks each ActionGroup for a specific condition and modifies that ActionGroup if necessary
type Condition interface {
	Apply(gameState *pogo.GameState, actionGroup []actions.Action, step int) (int, error)
}

// ProcessConditions with a GameState, Inputs, and Conditions, generate the necessary and valid list of actions to get to the next state
func ProcessConditions(currentState *pogo.GameState, inputs []pogo.GameInputMsg, conditions []Condition) (*pogo.GameState, error) {
	var nextState = currentState
	// [turnOrder][GameInputMsg's]
	// NOTE: Will it be possible for a player to submit turn 1 and 3, but not turn 2? May lead to cool strats
	turnGrouped := make([][]pogo.GameInputMsg, 0)
	for _, input := range inputs {
		for len(turnGrouped) < input.Order+1 {
			turnGrouped = append(turnGrouped, make([]pogo.GameInputMsg, 0))
		}
		turnGrouped[input.Order] = append(turnGrouped[input.Order], input)
	}

	for _, inputGroup := range turnGrouped {
		cardGroup, err := getCardsFromInputs(inputGroup, nextState)
		if err != nil {
			return nextState, errors.Wrap(err, "failed to parse cards from inputs")
		}

		// [cardPriority][list of cards with that priority]
		// NOTE: Some cardPriority lists may be empty (ex: there may be movements, and shots, but no utility)
		cardPriorities := make([][]cards.Card, 0)
		for _, card := range cardGroup {
			for len(cardPriorities) < card.CardType.Priority()+1 {
				cardPriorities = append(cardPriorities, make([]cards.Card, 0))
			}
			cardPriorities[card.CardType.Priority()] = append(cardPriorities[card.CardType.Priority()], card)
		}

		for _, cardPriorityGroup := range cardPriorities {
			nextState, err = applyCardsToState(cardPriorityGroup, nextState, conditions)
			if err != nil {
				return nextState, err
			}
		}
	}
	return nextState, nil
}

// getCardsFromInputs returns a slice of cards based on the given input list
func getCardsFromInputs(inputs []pogo.GameInputMsg, state *pogo.GameState) ([]cards.Card, error) {
	var cardGroup []cards.Card
	for _, input := range inputs {
		card, err := cards.GameInputToCard(input.CardID, input.Owner, state.GetCardType(input.CardID))
		if err != nil {
			return cardGroup, err
		}
		cardGroup = append(cardGroup, *card)
	}
	return cardGroup, nil
}

// applyCardsToState will apply all the cards to the game state with the given conditions, or an error otherwise
func applyCardsToState(cards []cards.Card, state *pogo.GameState, conditions []Condition) (*pogo.GameState, error) {
	// the group of action lists now has to be horizontally partitioned into action groups
	// In otherwords, we are grouping all of the cards' 1st actions together, 2nd actions together, etc
	var actionGroups [][]actions.Action

	for _, card := range cards {
		for i, action := range card.Actions {
			for len(actionGroups) < i+1 {
				actionGroups = append(actionGroups, make([]actions.Action, 0))
			}
			actionGroups[i] = append(actionGroups[i], action)
		}
	}

	//fmt.Printf("processing %v action groups\n", len(actionGroups))
	// loop through each action group and check it against every condition
	for _, actionGroup := range actionGroups {
		//fmt.Printf("processing action group with %v actions\n", len(actionGroup))
		dirty := true
		var loops = 0
		for dirty {
			dirty = false
			for _, cond := range conditions {
				// this is the step that actually checks each condition
				var steps, condErr = cond.Apply(state, actionGroup, 0)
				if condErr != nil {
					return state, condErr
				}
				if steps != 0 {
					dirty = true
					loops++
					break
				}
			}
			if loops > MaxConditionSteps {
				return state, fmt.Errorf("failed to apply conditions, reached maximum loop counter of %v", MaxConditionSteps)
			}
		}
		// here is where the actions are applied to the state to generate each next state
		animationGroup := []animations.AnimationAction{}
		for _, act := range actionGroup {
			var nxt, actErr = act.Apply(state)
			if actErr != nil {
				return nxt, actErr
			}
			state = nxt
			animationGroup = append(animationGroup, act.GetAnimation())
		}
		state.Animations = append(state.Animations, animationGroup)
	}

	return state, nil
}

type playerTracker struct {
	PlayerA     *pogo.PlayerState
	PlayerB     *pogo.PlayerState
	ActionIndex int
	Moved       bool
	Collided    bool
}
