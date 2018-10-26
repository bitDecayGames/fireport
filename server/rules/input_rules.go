package rules

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// InputRule is meant to validate the inputs for each player against the current state
type InputRule interface {
	Apply(gameState *pogo.GameState, player *pogo.PlayerState, inputs []pogo.GameInputMsg) error
}

// DefaultInputRules this is this list of default input rules for a game
var DefaultInputRules = []InputRule{
	&MinNumberOfInputsRule{NumberOfInputs: 0},
	&MaxNumberOfInputsRule{NumberOfInputs: 3},
	&MustHaveCardInHandToPlayRule{},
	&MaxNumberOfSwapsRule{NumberOfSwaps: 1},
}

// ApplyInputRules checks that the input for a game are valid
func ApplyInputRules(gameState *pogo.GameState, inputs []pogo.GameInputMsg, rules []InputRule) []error {
	var errors []error
	for _, player := range gameState.Players {
		var playerInputs []pogo.GameInputMsg
		for _, input := range inputs {
			if input.Owner == player.ID {
				playerInputs = append(playerInputs, input)
			}
		}
		for _, rule := range rules {
			err := rule.Apply(gameState, &player, playerInputs)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

// MinNumberOfInputsRule forces a player to play at least n number of cards
type MinNumberOfInputsRule struct {
	NumberOfInputs int
}

// Apply apply this rule
func (r *MinNumberOfInputsRule) Apply(gameState *pogo.GameState, player *pogo.PlayerState, inputs []pogo.GameInputMsg) error {
	if len(inputs) < r.NumberOfInputs {
		return fmt.Errorf("expected the player %v to have at least %v inputs instead found %v", player.ID, r.NumberOfInputs, len(inputs))
	}
	return nil
}

// MaxNumberOfInputsRule forces a player to play less than or equal to n number of cards
type MaxNumberOfInputsRule struct {
	NumberOfInputs int
}

// Apply apply this rule
func (r *MaxNumberOfInputsRule) Apply(gameState *pogo.GameState, player *pogo.PlayerState, inputs []pogo.GameInputMsg) error {
	if len(inputs) > r.NumberOfInputs {
		return fmt.Errorf("expected the player %v to have at most %v inputs instead found %v", player.ID, r.NumberOfInputs, len(inputs))
	}
	return nil
}

// MustHaveCardInHandToPlayRule a player is only allowed to play a card if it is in their hand
type MustHaveCardInHandToPlayRule struct{}

// Apply apply this rule
func (r *MustHaveCardInHandToPlayRule) Apply(gameState *pogo.GameState, player *pogo.PlayerState, inputs []pogo.GameInputMsg) error {
	for _, input := range inputs {
		var found = false
		for _, cardInHand := range player.Hand {
			if input.CardID == cardInHand.ID {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("player %v tried to play card %v that was not in their hand", player.ID, input.CardID)
		}
	}
	return nil
}

// MaxNumberOfSwapsRule limits the number of swaps a player is allowed to play per turn
type MaxNumberOfSwapsRule struct {
	NumberOfSwaps int
}

// Apply apply this rule
func (r *MaxNumberOfSwapsRule) Apply(gameState *pogo.GameState, player *pogo.PlayerState, inputs []pogo.GameInputMsg) error {
	var numOfSwaps = 0
	for _, input := range inputs {
		if input.Swap > 0 {
			numOfSwaps++
		}
	}
	if numOfSwaps > r.NumberOfSwaps {
		return fmt.Errorf("expected the player %v to have at most %v swaps instead found %v", player.ID, r.NumberOfSwaps, numOfSwaps)
	}
	return nil
}
