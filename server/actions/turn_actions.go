package actions

import (
	"fmt"
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

// DrawCardAction draw a card from a player's deck and put it in their hand
type DrawCardAction struct {
	Owner int
}

// Apply apply this action
func (a *DrawCardAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	for i := range nextState.Players {
		if nextState.Players[i].ID == a.Owner {
			if len(nextState.Players[i].Deck) > 0 {
				nextState.Players[i].Hand = append(nextState.Players[i].Hand, nextState.Players[i].Deck[len(nextState.Players[i].Deck)-1])
				nextState.Players[i].Deck = nextState.Players[i].Deck[:len(nextState.Players[i].Deck)-1]
				return nextState, nil
			}
			return nextState, fmt.Errorf("player %v tried to draw a card from an empty deck", a.Owner)
		}
	}
	return nextState, fmt.Errorf("failed to find player %v", a.Owner)
}

// GetOwner get the owner of this action
func (a *DrawCardAction) GetOwner() int {
	return a.Owner
}
