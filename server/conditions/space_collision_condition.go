package conditions

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// SpaceCollisionCondition address the case of two players moving into the same space
type SpaceCollisionCondition struct{}

// Apply applies the condition to the game state
func (c *SpaceCollisionCondition) Apply(gameState *pogo.GameState, actionGroup []actions.Action) error {
	var futureState = gameState
	var trackers []playerTracker
	for playerAIndex := range gameState.Players {
		trackers = append(trackers, playerTracker{PlayerA: &gameState.Players[playerAIndex], ActionIndex: -1})
	}
	for actionIndex := range actionGroup {
		var actOwner = actionGroup[actionIndex].GetOwner()
		nxt, err := actionGroup[actionIndex].Apply(futureState)
		if err != nil {
			return err
		}
		futureState = nxt
		for i := range trackers {
			if trackers[i].PlayerA.ID == actOwner {
				trackers[i].ActionIndex = actionIndex
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
		for playerBIndex := range futureState.Players {
			if trackers[i].PlayerB.ID != futureState.Players[playerBIndex].ID && trackers[i].PlayerB.Location == futureState.Players[playerBIndex].Location {
				trackers[i].Collided = true
			}
		}
		trackers[i].Moved = trackers[i].PlayerA.Location != trackers[i].PlayerB.Location

		if trackers[i].Moved && trackers[i].Collided && trackers[i].ActionIndex >= 0 {
			actionGroup[trackers[i].ActionIndex] = &actions.BumpDamageSelfAction{Owner: trackers[i].PlayerB.ID}
			dirty = true
		}
	}
	if dirty {
		return c.Apply(gameState, actionGroup)
	}

	return nil
}
