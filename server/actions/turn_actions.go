package actions

import (
	"github.com/bitdecaygames/fireport/server/pogo"
	"time"
)

// DefaultTurnActions the list of default actions that will be applied at the end of every turn
var DefaultTurnActions = []Action{
	&IncrementTurnAction{},
	&SyncLastUpdatedAction{},
}

// IncrementTurnAction increases the Turn by 1
type IncrementTurnAction struct{}

// Apply apply this action
func (a *IncrementTurnAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	nextState.Turn = nextState.Turn + 1
	return nextState, nil
}

// GetOwner get the owner of this action
func (a *IncrementTurnAction) GetOwner() int {
	return -1
}

// SyncLastUpdatedAction sets the Updated to the current epoch time
type SyncLastUpdatedAction struct{}

// Apply apply this action
func (a *SyncLastUpdatedAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	nextState.Updated = time.Now().Unix() // seconds since epoch
	return nextState, nil
}

// GetOwner get the owner of this action
func (a *SyncLastUpdatedAction) GetOwner() int {
	return -1
}
