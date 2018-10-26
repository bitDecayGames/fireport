package conditions

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/cards"
	"github.com/bitdecaygames/fireport/server/pogo"
	"sort"
)

// Condition checks each ActionGroup for a specific condition and modifies that ActionGroup if necessary
type Condition interface {
	Apply(gameState *pogo.GameState, actionGroup []actions.Action) error
}

// ProcessConditions with a GameState, Inputs, and Conditions, generate the necessary and valid list of actions to get to the next state
func ProcessConditions(currentState *pogo.GameState, inputs []pogo.GameInputMsg, conditions []Condition) (*pogo.GameState, error) {
	var nextState = currentState
	// group all of the inputs based on the order value they have
	var inputGroupsMap = make(map[int][]pogo.GameInputMsg)
	for _, input := range inputs {
		inputGroupsMap[input.Order] = append(inputGroupsMap[input.Order], input)
	}
	var order []int
	for key := range inputGroupsMap {
		order = append(order, key)
	}
	sort.Ints(order)
	for ord := range order {
		// these are all the inputs with order N
		var inputGroup = inputGroupsMap[ord]
		var cardGroup []cards.Card
		for _, input := range inputGroup {
			card, err := cards.GameInputToCard(input.CardID, input.Owner, nextState.GetCardType(input.CardID))
			if err != nil {
				return nextState, err
			}
			cardGroup = append(cardGroup, *card)
		}

		// try to group cards by type so that movement cards don't happen in parallel with attack cards
		var cardTypeGroupsMap = make(map[int][]cards.Card)
		for _, card := range cardGroup {
			var cardTypeGroupKey = int(card.CardType / 100)
			cardTypeGroupsMap[cardTypeGroupKey] = append(cardTypeGroupsMap[cardTypeGroupKey], card)
		}
		var cardTypeGroupKeysOrder []int
		for cardTypeGroupKey := range cardTypeGroupsMap {
			cardTypeGroupKeysOrder = append(cardTypeGroupKeysOrder, cardTypeGroupKey)
		}
		for _, cardTypeGroupKeyOrder := range cardTypeGroupKeysOrder {
			// these cards are now grouped by type
			var cardTypeGroup = cardTypeGroupsMap[cardTypeGroupKeyOrder]
			// now break up each cardTypeGroup into its own list of actions
			var actionLists [][]actions.Action
			var longest = 0
			for _, ctGroup := range cardTypeGroup {
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
			// loop through each action group and check it against every condition
			for _, actionGroup := range actionGroups {
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
