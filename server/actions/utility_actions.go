package actions

import (
	"fmt"
	"log"

	"github.com/bitdecaygames/fireport/server/animations"
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
	// TODO: MW probably need a direction component on this
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

// GetAnimation get the animation of this action
func (a *BumpDamageSelfAction) GetAnimation() *pogo.AnimationAction {
	return animations.GetBumpedInto(a.GetOwner())
}

// WinGameAction marks the game as won by the owner oF this action
type WinGameAction struct {
	Owner int
}

// Apply apply this action
func (a *WinGameAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	if a.Owner >= 0 { // -1 means that it is a tie game
		player := nextState.GetPlayer(a.Owner)
		if player == nil {
			return nextState, fmt.Errorf("there is no player with id %v", a.Owner)
		}
	}
	nextState.Winner = a.Owner
	nextState.IsGameFinished = true
	return nextState, nil
}

// GetOwner get the owner of this action
func (a *WinGameAction) GetOwner() int {
	return a.Owner
}

// GetAnimation get the animation of this action
func (a *WinGameAction) GetAnimation() *pogo.AnimationAction {
	log.Fatal("Get Animation not yet implemented for WinGameAction")
	return nil
}
