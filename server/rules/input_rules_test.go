package rules

import (
	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestState() *pogo.GameState {
	return &pogo.GameState{
		Turn:        0,
		Created:     1000,
		Updated:     2000,
		IDCounter:   300,
		BoardWidth:  2,
		BoardHeight: 2,
		BoardSpaces: []pogo.BoardSpace{
			{ID: 0, SpaceType: 0, State: 0},
			{ID: 1, SpaceType: 0, State: 0},
			{ID: 2, SpaceType: 0, State: 0},
			{ID: 3, SpaceType: 0, State: 0},
		},
		Players: []pogo.PlayerState{
			{
				ID:       100,
				Name:     "PlayerOne",
				Location: 0,
				Hand: []pogo.CardState{
					{ID: 101, CardType: 0},
					{ID: 102, CardType: 0},
					{ID: 103, CardType: 0},
					{ID: 104, CardType: 0},
					{ID: 105, CardType: 0},
				},
				Deck: []pogo.CardState{
					{ID: 106, CardType: 0},
					{ID: 107, CardType: 0},
					{ID: 108, CardType: 0},
					{ID: 109, CardType: 0},
					{ID: 110, CardType: 0},
				},
				Discard: []pogo.CardState{
					{ID: 111, CardType: 0},
					{ID: 112, CardType: 0},
					{ID: 113, CardType: 0},
					{ID: 114, CardType: 0},
					{ID: 115, CardType: 0},
				},
			},
			{
				ID:       200,
				Name:     "PlayerTwo",
				Location: 3,
				Hand: []pogo.CardState{
					{ID: 201, CardType: 0},
					{ID: 202, CardType: 0},
					{ID: 203, CardType: 0},
					{ID: 204, CardType: 0},
					{ID: 205, CardType: 0},
				},
				Deck: []pogo.CardState{
					{ID: 206, CardType: 0},
					{ID: 207, CardType: 0},
					{ID: 208, CardType: 0},
					{ID: 209, CardType: 0},
					{ID: 210, CardType: 0},
				},
				Discard: []pogo.CardState{
					{ID: 211, CardType: 0},
					{ID: 212, CardType: 0},
					{ID: 213, CardType: 0},
					{ID: 214, CardType: 0},
					{ID: 215, CardType: 0},
				},
			},
		},
	}
}

func TestMinNumberOfInputsRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &MinNumberOfInputsRule{NumberOfInputs: 1}

	var err = rule.Apply(gameState, &gameState.Players[0], nil)
	assert.Error(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{}})
	assert.NoError(t, err)
}

func TestMaxNumberOfInputsRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &MaxNumberOfInputsRule{NumberOfInputs: 1}

	var err = rule.Apply(gameState, &gameState.Players[0], nil)
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{}, {}})
	assert.Error(t, err)
}

func TestMustHaveCardInHandToPlayRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &MustHaveCardInHandToPlayRule{}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{CardID: 7}})
	assert.Error(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{CardID: 101}})
	assert.NoError(t, err)
}

func TestMaxNumberOfSwapsRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &MaxNumberOfSwapsRule{NumberOfSwaps: 1}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Swap: 1}, {Swap: 1}})
	assert.Error(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Swap: 1}, {CardID: 101}})
	assert.NoError(t, err)
}

func TestOneCardPerOrderRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &OneCardPerOrderRule{}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 3}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 3}, {Order: 2}, {Order: 3}})
	assert.Error(t, err)
}

func TestMaxAllowedOrderRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &MaxAllowedOrderRule{OrderLimit: 10}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 3}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 30}})
	assert.Error(t, err)
}

func TestMinAllowedOrderRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &MinAllowedOrderRule{OrderLimit: 2}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 2}, {Order: 3}, {Order: 4}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 3}})
	assert.Error(t, err)
}

func TestCannotSkipOrderRule(t *testing.T) {
	var gameState = getTestState()
	var rule = &CannotSkipOrderRule{}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 2}, {Order: 3}, {Order: 4}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 3}, {Order: 4}})
	assert.Error(t, err)
}
