package rules

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
)

func TestMinNumberOfInputsRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &MinNumberOfInputsRule{NumberOfInputs: 1}

	var err = rule.Apply(gameState, &gameState.Players[0], nil)
	assert.Error(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{}})
	assert.NoError(t, err)
}

func TestMaxNumberOfInputsRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &MaxNumberOfInputsRule{NumberOfInputs: 1}

	var err = rule.Apply(gameState, &gameState.Players[0], nil)
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{}, {}})
	assert.Error(t, err)
}

func TestMustHaveCardInHandToPlayRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &MustHaveCardInHandToPlayRule{}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{CardID: 7}})
	assert.Error(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{CardID: 101}})
	assert.NoError(t, err)
}

func TestMaxNumberOfSwapsRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &MaxNumberOfSwapsRule{NumberOfSwaps: 1}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Swap: 1}, {Swap: 1}})
	assert.Error(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Swap: 1}, {CardID: 101}})
	assert.NoError(t, err)
}

func TestOneCardPerOrderRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &OneCardPerOrderRule{}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 3}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 3}, {Order: 2}, {Order: 3}})
	assert.Error(t, err)
}

func TestMaxAllowedOrderRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &MaxAllowedOrderRule{OrderLimit: 10}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 3}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 30}})
	assert.Error(t, err)
}

func TestMinAllowedOrderRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &MinAllowedOrderRule{OrderLimit: 2}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 2}, {Order: 3}, {Order: 4}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 2}, {Order: 3}})
	assert.Error(t, err)
}

func TestCannotSkipOrderRule(t *testing.T) {
	var gameState = pogo.GetTestState()
	var rule = &CannotSkipOrderRule{}

	var err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 2}, {Order: 3}, {Order: 4}})
	assert.NoError(t, err)

	err = rule.Apply(gameState, &gameState.Players[0], []pogo.GameInputMsg{{Order: 1}, {Order: 3}, {Order: 4}})
	assert.Error(t, err)
}
