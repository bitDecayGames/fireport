package actions

import (
	"fmt"
	"testing"

	"github.com/bitdecaygames/fireport/server/animations"
	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/stretchr/testify/assert"
)

func TestIncrementTurnAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &IncrementTurnAction{}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, a.Turn+1, b.Turn)
}

func TestSyncLastUpdatedAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &SyncLastUpdatedAction{}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.True(t, a.Updated < b.Updated)
}

func TestDrawCardAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &DrawCardAction{}

	var b, err = action.Apply(a)
	assert.Error(t, err)

	action.Owner = 100

	b, err = action.Apply(a)
	assert.NoError(t, err)

	assert.True(t, len(b.Players[0].Hand) > len(a.Players[0].Hand))
	assert.True(t, len(b.Players[0].Deck) < len(a.Players[0].Deck))

	b, err = action.Apply(b)
	assert.NoError(t, err)
	b, err = action.Apply(b)
	assert.NoError(t, err)
	b, err = action.Apply(b)
	assert.NoError(t, err)
	b, err = action.Apply(b)
	assert.NoError(t, err)
	b, err = action.Apply(b)
	assert.Error(t, err)
}

func TestDiscardCardAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &DiscardCardAction{}

	var b, err = action.Apply(a)
	assert.Error(t, err)

	action.Owner = 100

	b, err = action.Apply(a)
	assert.Error(t, err)

	action.CardID = 101

	b, err = action.Apply(a)
	assert.NoError(t, err)

	assert.True(t, len(b.Players[0].Discard) > len(a.Players[0].Discard))
	assert.True(t, len(b.Players[0].Hand) < len(a.Players[0].Hand))
}

func TestResetDiscardPileAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &ResetDiscardPileAction{}

	var b, err = action.Apply(a)
	assert.Error(t, err)

	action.Owner = 100

	b, err = action.Apply(a)
	assert.NoError(t, err)

	assert.True(t, len(b.Players[0].Deck) > len(a.Players[0].Deck))
	assert.True(t, len(b.Players[0].Discard) < len(a.Players[0].Discard))
	assert.Equal(t, 10, len(b.Players[0].Deck))
	assert.Equal(t, 0, len(b.Players[0].Discard))
	assert.Equal(t, a.Players[0].Deck[len(a.Players[0].Deck)-1].ID, b.Players[0].Deck[len(b.Players[0].Deck)-1].ID)
	assert.NotEqual(t, a.Players[0].Deck[0].ID, b.Players[0].Deck[0].ID)
}

func TestShuffleDeckAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &ShuffleDeckAction{}

	var b, err = action.Apply(a)
	assert.Error(t, err)

	action.Owner = 100

	b, err = action.Apply(a)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))
	for _, cardA := range a.Players[0].Deck {
		var found = false
		for _, cardB := range b.Players[0].Deck {
			if cardB.ID == cardA.ID {
				found = true
				break
			}
		}
		assert.True(t, found, fmt.Sprintf("could not find card %v in next deck state", cardA.ID))
	}
}

func TestShuffleDiscardAction(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &ShuffleDiscardAction{}

	var b, err = action.Apply(a)
	assert.Error(t, err)

	action.Owner = 100

	b, err = action.Apply(a)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Discard), len(b.Players[0].Discard))
	for _, cardA := range a.Players[0].Discard {
		var found = false
		for _, cardB := range b.Players[0].Discard {
			if cardB.ID == cardA.ID {
				found = true
				break
			}
		}
		assert.True(t, found, fmt.Sprintf("could not find card %v in next discard state", cardA.ID))
	}
}

func TestTurnClockwise90Action(t *testing.T) {
	var a = pogo.GetTestState()
	var action = &TurnClockwise90Action{Owner: a.Players[0].ID}
	a.AddEmptyAnimationSlice()

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, a.Players[0].Facing+1, b.Players[0].Facing)

	// Animation check
	for _, animation := range b.Animations[len(b.Animations)-1] {
		assert.Equal(t, animation.Name, animations.GetTurnClockwise90(b.Players[0].ID).Name)
		assert.Equal(t, animation.Owner, animations.GetTurnClockwise90(b.Players[0].ID).Owner)
	}
}

func TestTurnCounterClockwise90Action(t *testing.T) {
	var a = pogo.GetTestState()
	// set player facing to zero so we can test overflow handling
	a.Players[0].Facing = 0
	var action = &TurnCounterClockwise90Action{Owner: a.Players[0].ID}
	a.AddEmptyAnimationSlice()

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, 3, b.Players[0].Facing)

	// Animation check
	for _, animation := range b.Animations[len(b.Animations)-1] {
		assert.Equal(t, animation.Name, animations.GetTurnCounterClockwise90(b.Players[0].ID).Name)
		assert.Equal(t, animation.Owner, animations.GetTurnCounterClockwise90(b.Players[0].ID).Owner)
	}
}
