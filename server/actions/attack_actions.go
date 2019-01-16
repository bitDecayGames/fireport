package actions

import (
	"fmt"

	"github.com/bitdecaygames/fireport/server/animations"

	"github.com/bitdecaygames/fireport/server/pogo"
)

// FireBasicAction action for the owner to shoot
type FireBasicAction struct {
	Owner int
}

const (

	// NORTH facing
	NORTH = iota
	// EAST facing
	EAST
	// SOUTH facing
	SOUTH
	// WEST facing
	WEST
)

// Apply apply this action
func (a *FireBasicAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	nextState.AppendAnimation(animations.GetFireCanon(a.GetOwner()))
	player := nextState.GetPlayer(a.GetOwner())
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.GetOwner())
	}

	return shoot(nextState, player)
}

// GetOwner get the owner of this action
func (a *FireBasicAction) GetOwner() int {
	return a.Owner
}

func shoot(state *pogo.GameState, shooter *pogo.PlayerState) (*pogo.GameState, error) {
	// For reference:
	// x := shooter.Location % state.BoardWidth
	// y := shooter.Location / state.BoardWidth

	switch shooter.Facing {
	case NORTH:
		pos := shooter.Location
		for pos >= 0 {
			pos = pos - state.BoardWidth
			for k := range state.Players {
				if state.Players[k].Location == pos {
					// YOU SUNK MY BATTLESHIP
					state.Players[k].Health--
					state.AppendAnimation(animations.GetHitByCanon(state.Players[k].ID))
					return state, nil
				}
			}
		}
	case EAST:
		x := shooter.Location % state.BoardWidth
		checkSpaces := state.BoardWidth - x

		pos := shooter.Location
		for i := 1; i <= checkSpaces; i++ {
			for k := range state.Players {
				if state.Players[k].Location == pos+i {
					// YOU SUNK MY BATTLESHIP
					state.Players[k].Health--
					state.AppendAnimation(animations.GetHitByCanon(state.Players[k].ID))
					return state, nil
				}
			}
		}
	case SOUTH:
		pos := shooter.Location
		for pos < len(state.BoardSpaces) {
			pos = pos + state.BoardWidth
			for k := range state.Players {
				if state.Players[k].Location == pos {
					// YOU SUNK MY BATTLESHIP
					state.Players[k].Health--
					state.AppendAnimation(animations.GetHitByCanon(state.Players[k].ID))
					return state, nil
				}
			}
		}
	case WEST:
		x := shooter.Location % state.BoardWidth
		checkSpaces := state.BoardWidth - x
		pos := shooter.Location
		for i := 1; i <= checkSpaces; i++ {
			for k := range state.Players {
				if state.Players[k].Location == pos-i {
					// YOU SUNK MY BATTLESHIP
					state.Players[k].Health--
					state.AppendAnimation(animations.GetHitByCanon(state.Players[k].ID))
					return state, nil
				}
			}
		}
	default:
		return state, fmt.Errorf("player %v with unknown facing %v", shooter.ID, shooter.Facing)
	}

	// nothing was hit
	return state, nil
}
