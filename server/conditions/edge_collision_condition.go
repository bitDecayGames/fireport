package conditions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// EdgeCollisionCondition address the case of two players moving into adjacent spaces through the same edge (passing through each other)
type EdgeCollisionCondition struct{}

// Apply applies the condition to the game state
func (c *EdgeCollisionCondition) Apply(gameState *pogo.GameState, actionGroup []actions.Action) error {
	// TODO: clean up all of the printf statements
	fmt.Printf("applying the edge collision condition to an action group of %v\n", len(actionGroup))
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
		for playerBIndex := range futureState.Players {
			if trackers[i].PlayerA.ID == futureState.Players[playerBIndex].ID {
				trackers[i].PlayerB = &futureState.Players[playerBIndex]
				break
			}
		}
	}
	var dirty = false
	for i := range trackers {
		var me = &trackers[i]
		for k := range trackers {
			var other = &trackers[k]
			if me.PlayerB.ID != other.PlayerB.ID && me.PlayerB.Location == other.PlayerA.Location && me.PlayerA.Location == other.PlayerB.Location {
				me.Collided = true
				other.Collided = true
				break
			}
		}

		fmt.Printf("trying to calculate tracker %v\n", trackers[i])
		if trackers[i].Collided && trackers[i].ActionIndex >= 0 {
			fmt.Printf("player %v %v was seen as in an edge collision due to action %v\n", i, trackers[i].PlayerB.ID, trackers[i].ActionIndex)
			actionGroup[trackers[i].ActionIndex] = &actions.BumpDamageSelfAction{Owner: trackers[i].PlayerB.ID}
			dirty = true
		}
	}
	fmt.Println("")
	if dirty {
		return c.Apply(gameState, actionGroup)
	}

	return nil
}
