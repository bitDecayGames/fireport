package conditions

import (
	"fmt"
	"sort"

	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/cards"
	"github.com/bitdecaygames/fireport/server/pogo"
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
		var cardGroup []cards.Card
		for _, input := range inputGroup {
			card, err := cards.GameInputToCard(input.CardID, input.Owner, nextState.GetCardType(input.CardID))
			if err != nil {
				return nextState, err
			}
			cardGroup = append(cardGroup, *card)
		}
		//fmt.Printf("found %v cards in card group %v\n", len(cardGroup), cardGroup)
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
		sort.Ints(cardTypeGroupKeysOrder)
		//fmt.Printf("found %v types of cards in card type group %v\n", len(cardTypeGroupKeysOrder), cardTypeGroupsMap)
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

type playerTracker struct {
	PlayerA     *pogo.PlayerState
	PlayerB     *pogo.PlayerState
	ActionIndex int
	Moved       bool
	Collided    bool
}
