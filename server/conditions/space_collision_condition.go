package conditions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// SpaceCollisionCondition address the case of two players moving into the same space
type SpaceCollisionCondition struct{}

// Apply applies the condition to the game state
func (c *SpaceCollisionCondition) Apply(gameState *pogo.GameState, actionGroup []actions.Action) error {
	// TODO: clean up all of the printf statements
	fmt.Printf("applying the space collision condition to an action group of %v\n", len(actionGroup))
	var futureState = gameState
	var trackers []playerTracker
	for playerAIndex := range gameState.Players {
		trackers = append(trackers, playerTracker{PlayerA: &gameState.Players[playerAIndex], ActionIndex: -1})
	}
	fmt.Printf("found trackers %v\n", trackers)
	for actionIndex := range actionGroup {
		var actOwner = actionGroup[actionIndex].GetOwner()
		fmt.Printf("applying action with owner %v at index %v\n", actOwner, actionIndex)
		nxt, err := actionGroup[actionIndex].Apply(futureState)
		if err != nil {
			return err
		}
		futureState = nxt
		for i := range trackers {
			if trackers[i].PlayerA.ID == actOwner {
				trackers[i].ActionIndex = actionIndex
				fmt.Printf("found tracker with playerA.ID %v %v\n", actOwner, trackers[i])
				break
			}
		}
	}
	for i := range trackers {
		for _, playerB := range futureState.Players {
			if trackers[i].PlayerA.ID == playerB.ID {
				trackers[i].PlayerB = &playerB
				break
			}
		}
	}
	for i := range trackers {
		for playerBIndex := range futureState.Players {
			if trackers[i].PlayerB.ID != futureState.Players[playerBIndex].ID && trackers[i].PlayerB.Location == futureState.Players[playerBIndex].Location {
				trackers[i].Collided = true
			}
		}
		trackers[i].Moved = trackers[i].PlayerA.Location != trackers[i].PlayerB.Location

		fmt.Printf("trying to calculate tracker %v\n", trackers[i])
		if trackers[i].Moved && trackers[i].Collided && trackers[i].ActionIndex >= 0 {
			fmt.Printf("player %v %v was seen as in a space collision due to action %v\n", i, trackers[i].PlayerB.ID, trackers[i].ActionIndex)
			actionGroup[trackers[i].ActionIndex] = &actions.BumpDamageSelfAction{Owner: trackers[i].PlayerB.ID}
		}
	}

	return nil
}
