package pogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameState_DeepCopy(t *testing.T) {
	// var a = GetTestState()
	var a = GetTestStateAnimationActions()
	var b = a.DeepCopy()

	assert.Equal(t, a.Turn, b.Turn)
	b.Turn++
	assert.Equal(t, a.Turn+1, b.Turn)

	assert.Equal(t, len(a.Players), len(b.Players))

	var pA = a.Players[0]
	var pB = b.Players[0]

	assert.Equal(t, pA.ID, pB.ID)
	assert.Equal(t, pA.Name, pB.Name)
	assert.Equal(t, pA.Location, pB.Location)
	assert.Equal(t, pA.Facing, pB.Facing)
	assert.Equal(t, pA.Health, pB.Health)
	assert.Equal(t, pA.Health, pB.Health)
	assert.Equal(t, len(pA.Hand), len(pB.Hand))
	assert.Equal(t, len(pA.Deck), len(pB.Deck))
	assert.Equal(t, len(pA.Discard), len(pB.Discard))

	b.Players[0].ID++
	b.Players[0].Hand[0].ID++
	b.Players = append(b.Players, PlayerState{})

	assert.Equal(t, a.Players[0].ID+1, b.Players[0].ID)
	assert.Equal(t, a.Players[0].Hand[0].ID+1, b.Players[0].Hand[0].ID)
	assert.Equal(t, len(a.Players)+1, len(b.Players))

	assert.Equal(t, a.Animations, b.Animations)
}

func TestGameState_GetPlayer(t *testing.T) {
	var state = GetTestState()

	var a = state.GetPlayer(state.Players[0].ID)
	var b = state.GetPlayer(state.Players[0].ID)

	a.Facing = a.Facing + 1

	assert.Equal(t, a.Facing, b.Facing)
}
