package actions

import (
	"github.com/bitdecaygames/fireport/server/pogo"
)

// Action the smallest unit of modification to be made to a game state object
type Action interface {
	Apply(currentState *pogo.GameState) (*pogo.GameState, error)
	GetGroup() int
	SetGroup(group int)
	GetPlayed() bool
	SetPlayed(played bool)
}


// ActionTracker adds some helper methods and values to every action
type ActionTracker struct {
	Group  int
	Played bool
}

// GetGroup returns the group order that this action is a part of
func (a *ActionTracker) GetGroup() int {
	return a.Group
}

// SetGroup sets the group order that this action is a part of
func (a *ActionTracker) SetGroup(group int) {
	a.Group = group
}

// GetPlayed true if this card has been played
func (a *ActionTracker) GetPlayed() bool {
	return a.Played
}

// SetPlayed set to true if this card has been played
func (a *ActionTracker) SetPlayed(played bool) {
	a.Played = played
}




