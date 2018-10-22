package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
	"time"
)

type Action interface {
	Apply(currentState *pogo.GameState) (*pogo.GameState, error)
}

var DefaultTurnActions = []Action{
	&IncrementTurnAction{},
	&SyncLastUpdatedAction{},
}

// IncrementTurnAction increases the Turn by 1
type IncrementTurnAction struct{}

func (a *IncrementTurnAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	nextState.Turn = nextState.Turn + 1
	return nextState, nil
}

// SyncLastUpdatedAction sets the Updated to the current epoch time
type SyncLastUpdatedAction struct{}

func (a *SyncLastUpdatedAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	nextState.Updated = time.Now().Unix() // seconds since epoch
	return nextState, nil
}

// TurnClockwise90Action rotate the Owner of this action by 90 degrees clockwise
type TurnClockwise90Action struct {
	Owner int
}

func (a *TurnClockwise90Action) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	player := nextState.GetPlayer(a.Owner)
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
	} else {
		player.Facing = player.Facing + 1
		if player.Facing > 3 {
			player.Facing = 0
		}
	}
	return nextState, nil
}

// TurnCounterClockwise90Action rotate the Owner of this action by 90 degrees counter-clockwise
type TurnCounterClockwise90Action struct {
	Owner int
}

func (a *TurnCounterClockwise90Action) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	player := nextState.GetPlayer(a.Owner)
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
	} else {
		player.Facing = player.Facing - 1
		if player.Facing < 0 {
			player.Facing = 3
		}
	}
	return nextState, nil
}
