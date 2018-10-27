package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// MoveForwardAction move this player forward one space
type MoveForwardAction struct {
	Owner int
}

// Apply apply this action
func (a *MoveForwardAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	player := nextState.GetPlayer(a.Owner)
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
	}
	switch player.Facing {
	case 0: // going north
		player.Location = player.Location - nextState.BoardWidth
		break
	case 1: // going east
		player.Location = player.Location + 1
		break
	case 2: // going south
		player.Location = player.Location + nextState.BoardWidth
		break
	case 3: // going west
		player.Location = player.Location - 1
		break
	default:
		return nextState, fmt.Errorf("player %v with unknown facing %v", player.ID, player.Facing)
	}
	return nextState, nil
}

// GetOwner get the owner of this action
func (a *MoveForwardAction) GetOwner() int {
	return a.Owner
}

// TurnClockwise90Action rotate the Owner of this action by 90 degrees clockwise
type TurnClockwise90Action struct {
	Owner int
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

// GetOwner get the owner of this action
func (a *TurnClockwise90Action) GetOwner() int {
	return a.Owner
}

// TurnCounterClockwise90Action rotate the Owner of this action by 90 degrees counter-clockwise
type TurnCounterClockwise90Action struct {
	Owner int
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

// GetOwner get the owner of this action
func (a *TurnCounterClockwise90Action) GetOwner() int {
	return a.Owner
}
