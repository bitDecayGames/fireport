package triggers

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// DefaultPostStepTriggers the default post-step triggers
func DefaultPostStepTriggers(numberOfCardsInHand int, cardsToDiscard []int) []Trigger {
	return []Trigger{
		&WinTrigger{},
		&NotEnoughCardsInHandTrigger{RequiredCardsInHand: numberOfCardsInHand},
		&RefreshDeckTrigger{},
		&DiscardUsedCardsTrigger{Cards: cardsToDiscard},
	}
}

// WinTrigger the trigger for marking the game as finished
type WinTrigger struct {
	winner int
}

// Check check the current state to fire the trigger
func (t *WinTrigger) Check(currentState *pogo.GameState) bool {
	t.winner = -1
	if currentState.IsGameFinished {
		return false
	}
	var playersWithHealth []pogo.PlayerState
	for _, player := range currentState.Players {
		if player.Health > 0 {
			playersWithHealth = append(playersWithHealth, player)
		}
	}
	if len(playersWithHealth) == 1 {
		t.winner = playersWithHealth[0].ID
		return true
	} else if len(playersWithHealth) == 0 {
		return true
	}
	return false
}

// GetActions get the actions for this trigger
func (t *WinTrigger) GetActions() []actions.Action {
	return []actions.Action{&actions.WinGameAction{Owner: t.winner}}
}

// NotEnoughCardsInHandTrigger triggers when a player does not have enough cards in their hand
type NotEnoughCardsInHandTrigger struct {
	RequiredCardsInHand int
	playerID            int
}

// Check check the current state to fire the trigger
func (t *NotEnoughCardsInHandTrigger) Check(currentState *pogo.GameState) bool {
	t.playerID = -1
	for _, player := range currentState.Players {
		if len(player.Hand) < t.RequiredCardsInHand && len(player.Deck) > 0 {
			t.playerID = player.ID
			return true
		}
	}
	return false
}

// GetActions get the actions for this trigger
func (t *NotEnoughCardsInHandTrigger) GetActions() []actions.Action {
	return []actions.Action{&actions.DrawCardAction{Owner: t.playerID}}
}

// RefreshDeckTrigger triggers when a player has no cards in their deck
type RefreshDeckTrigger struct {
	playerID int
}

// Check check the current state to fire the trigger
func (t *RefreshDeckTrigger) Check(currentState *pogo.GameState) bool {
	t.playerID = -1
	for _, player := range currentState.Players {
		if len(player.Deck) == 0 && len(player.Discard) > 0 {
			t.playerID = player.ID
			return true
		}
	}
	return false
}

// GetActions get the actions for this trigger
func (t *RefreshDeckTrigger) GetActions() []actions.Action {
	return []actions.Action{&actions.ResetDiscardPileAction{Owner: t.playerID}, &actions.ShuffleDeckAction{Owner: t.playerID}}
}

// DiscardUsedCardsTrigger triggers when a player has used a card
type DiscardUsedCardsTrigger struct {
	Cards    []int
	playerID int
	cardID   int
}

// Check check the current state to fire the trigger
func (t *DiscardUsedCardsTrigger) Check(currentState *pogo.GameState) bool {
	t.playerID = -1
	t.cardID = -1
	for _, player := range currentState.Players {
		for _, card := range player.Hand {
			for i, cardID := range t.Cards {
				if card.ID == cardID {
					t.cardID = cardID
					t.Cards = append(t.Cards[:i], t.Cards[i+1:]...)
					t.playerID = player.ID
					return true
				}
			}
		}
	}
	return false
}

// GetActions get the actions for this trigger
func (t *DiscardUsedCardsTrigger) GetActions() []actions.Action {
	return []actions.Action{&actions.DiscardCardAction{Owner: t.playerID, CardID: t.cardID}}
}
