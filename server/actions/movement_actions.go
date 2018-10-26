package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// TurnClockwise90Action rotate the Owner of this action by 90 degrees clockwise
type TurnClockwise90Action struct {
	Owner int
	*ActionTracker
}

// Apply apply this action
func (a *TurnClockwise90Action) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	player := nextState.GetPlayer(a.Owner)
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
	}
	player.Facing = player.Facing + 1
	if player.Facing > 3 {
		player.Facing = 0
	}
	return nextState, nil
}

// TurnCounterClockwise90Action rotate the Owner of this action by 90 degrees counter-clockwise
type TurnCounterClockwise90Action struct {
	Owner int
	*ActionTracker
}

// Apply apply this action
func (a *TurnCounterClockwise90Action) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	player := nextState.GetPlayer(a.Owner)
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
	}
	player.Facing = player.Facing - 1
	if player.Facing < 0 {
		player.Facing = 3
	}

	return nextState, nil
}
