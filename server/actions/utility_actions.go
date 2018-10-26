package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// EmptyAction useful when you are modifying a list of actions but you need to maintain the indexes of the other actions
type EmptyAction struct {
	Owner int
}

// Apply apply this action
func (a *EmptyAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	return nextState, nil
}

// GetOwner get the owner of this action
func (a *EmptyAction) GetOwner() int {
	return a.Owner
}

// BumpDamageSelfAction damage myself because I bumped something solid
type BumpDamageSelfAction struct {
	Owner int
}

// Apply apply this action
func (a *BumpDamageSelfAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	player := nextState.GetPlayer(a.Owner)
	if player == nil {
		return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
	}
	player.Health--
	return nextState, nil
}

// GetOwner get the owner of this action
func (a *BumpDamageSelfAction) GetOwner() int {
	return a.Owner
}
