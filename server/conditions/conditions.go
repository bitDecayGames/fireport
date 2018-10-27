package conditions

import (
	"fmt"

	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/cards"
	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/pkg/errors"
)

// Condition checks each ActionGroup for a specific condition and modifies that ActionGroup if necessary
type Condition interface {
	Apply(gameState *pogo.GameState, actionGroup []actions.Action) error
}

// ProcessConditions with a GameState, Inputs, and Conditions, generate the necessary and valid list of actions to get to the next state
func ProcessConditions(currentState *pogo.GameState, inputs []pogo.GameInputMsg, conditions []Condition) (*pogo.GameState, error) {
	var nextState = currentState
	// [turnOrder][GameInputMsg's]
	turnGrouped := make([][]pogo.GameInputMsg, 0)
	for _, input := range inputs {
		for len(turnGrouped) < input.Order+1 {
			turnGrouped = append(turnGrouped, make([]pogo.GameInputMsg, 0))
		}
		turnGrouped[input.Order] = append(turnGrouped[input.Order], input)
	}

	for ord, inputGroup := range turnGrouped {
		fmt.Printf("Handling all inputs with order %v\n", ord)

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
			// now break up each cardPriorityGroup into its own list of actions
			var actionLists [][]actions.Action
			var longest = 0
			for _, ctGroup := range cardPriorityGroup {
				actionLists = append(actionLists, ctGroup.Actions)
				if len(ctGroup.Actions) > longest {
					longest = len(ctGroup.Actions)
				}
			}
			// the group of action lists now has to be horizontally partitioned into action groups
			var actionGroups [][]actions.Action
			for i := 0; i < longest; i++ {
				actionGroups = append(actionGroups, nil)
				for _, actionList := range actionLists {
					if i < len(actionList) {
						actionGroups[i] = append(actionGroups[i], actionList[i])
					}
				}
			}
			//fmt.Printf("processing %v action groups\n", len(actionGroups))
			// loop through each action group and check it against every condition
			for _, actionGroup := range actionGroups {
				//fmt.Printf("processing action group with %v actions\n", len(actionGroup))
				for _, cond := range conditions {
					// this is the step that actually checks each condition
					var condErr = cond.Apply(nextState, actionGroup)
					if condErr != nil {
						return nextState, condErr
					}
				}
				// here is where the actions are applied to the state to generate each next state
				for _, act := range actionGroup {
					var nxt, actErr = act.Apply(nextState)
					if actErr != nil {
						return nxt, actErr
					}
					nextState = nxt
				}
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

type playerTracker struct {
	PlayerA     *pogo.PlayerState
	PlayerB     *pogo.PlayerState
	ActionIndex int
	Moved       bool
	Collided    bool
}
