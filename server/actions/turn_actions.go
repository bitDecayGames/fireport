package actions

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
	"math/rand"
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

// DiscardCardAction move a card from the player's hand onto their discard
type DiscardCardAction struct {
	Owner  int
	CardID int
}

// Apply apply this action
func (a *DiscardCardAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	for i := range nextState.Players {
		if nextState.Players[i].ID == a.Owner {
			var discarded = false
			for k, card := range nextState.Players[i].Hand {
				if card.ID == a.CardID {
					nextState.Players[i].Discard = append(nextState.Players[i].Discard, nextState.Players[i].Hand[i])
					nextState.Players[i].Hand = append(nextState.Players[i].Hand[:k], nextState.Players[i].Hand[k+1:]...)
					discarded = true
				}
			}
			if discarded {
				return nextState, nil
			}
			return nextState, fmt.Errorf("player %v tried to discard a card %v that was not in their hand", a.Owner, a.CardID)
		}
	}
	return nextState, fmt.Errorf("failed to find player %v", a.Owner)
}

// GetOwner get the owner of this action
func (a *DiscardCardAction) GetOwner() int {
	return a.Owner
}

// ResetDiscardPileAction put all of the cards from a player's discard onto the bottom of their deck
type ResetDiscardPileAction struct {
	Owner int
}

// Apply apply this action
func (a *ResetDiscardPileAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	for i := range nextState.Players {
		if nextState.Players[i].ID == a.Owner {
			nextState.Players[i].Deck = append(nextState.Players[i].Discard, nextState.Players[i].Deck...)
			nextState.Players[i].Discard = nextState.Players[i].Discard[:0]
			return nextState, nil
		}
	}
	return nextState, fmt.Errorf("failed to find player %v", a.Owner)
}

// GetOwner get the owner of this action
func (a *ResetDiscardPileAction) GetOwner() int {
	return a.Owner
}

func shuffle(cards []pogo.CardState) []pogo.CardState {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]pogo.CardState, len(cards))
	perm := r.Perm(len(cards))
	for i, randIndex := range perm {
		ret[i] = cards[randIndex]
	}
	return ret
}

// ShuffleDeckAction randomly shuffle a player's deck
type ShuffleDeckAction struct {
	Owner int
}

// Apply apply this action
func (a *ShuffleDeckAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	for i := range nextState.Players {
		if nextState.Players[i].ID == a.Owner {
			nextState.Players[i].Deck = shuffle(nextState.Players[i].Deck)
			return nextState, nil
		}
	}
	return nextState, fmt.Errorf("failed to find player %v", a.Owner)
}

// GetOwner get the owner of this action
func (a *ShuffleDeckAction) GetOwner() int {
	return a.Owner
}

// ShuffleDiscardAction randomly shuffle a player's discard
type ShuffleDiscardAction struct {
	Owner int
}

// Apply apply this action
func (a *ShuffleDiscardAction) Apply(currentState *pogo.GameState) (*pogo.GameState, error) {
	nextState := currentState.DeepCopy()
	for i := range nextState.Players {
		if nextState.Players[i].ID == a.Owner {
			nextState.Players[i].Discard = shuffle(nextState.Players[i].Discard)
			return nextState, nil
		}
	}
	return nextState, fmt.Errorf("failed to find player %v", a.Owner)
}

// GetOwner get the owner of this action
func (a *ShuffleDiscardAction) GetOwner() int {
	return a.Owner
}
