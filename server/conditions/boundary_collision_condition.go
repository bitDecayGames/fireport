package conditions

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// BoundaryCollisionCondition address the case of a player that runs into the outer boundary of the play space
type BoundaryCollisionCondition struct{}

// Apply applies the condition to the game state
func (c *BoundaryCollisionCondition) Apply(gameState *pogo.GameState, actionGroup []actions.Action) error {
	if gameState.BoardWidth == 0 {
		return nil
	}
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
		var me = &trackers[i]
		// check for collisions on the top side
		if me.PlayerB.Location < 0 {
			me.Collided = true
		}
		// check for collisions on the bottom side
		if me.PlayerB.Location >= len(gameState.BoardSpaces) {
			me.Collided = true
		}
		// check for collisions on the left side
		if me.PlayerA.Location%gameState.BoardWidth == 0 && me.PlayerB.Location%gameState.BoardWidth == gameState.BoardWidth-1 {
			me.Collided = true
		}
		// check for collisions on the right side
		if me.PlayerA.Location%gameState.BoardWidth == gameState.BoardWidth-1 && me.PlayerB.Location%gameState.BoardWidth == 0 {
			me.Collided = true
		}

		if trackers[i].Collided && trackers[i].ActionIndex >= 0 {
			actionGroup[trackers[i].ActionIndex] = &actions.BumpDamageSelfAction{Owner: trackers[i].PlayerB.ID}
			dirty = true
		}
	}
	if dirty {
		return c.Apply(gameState, actionGroup)
	}

	return nil
}
