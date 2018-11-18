package actions

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
)

func getTestState() *pogo.GameState {
	return &pogo.GameState{
		Turn:        0,
		RNG:         rand.New(rand.NewSource(0)), // this may be useful as our test state will always generate the same sequence of things
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

func TestIncrementTurnAction(t *testing.T) {
	var a = getTestState()
	var action = &IncrementTurnAction{}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, a.Turn+1, b.Turn)
}

func TestSyncLastUpdatedAction(t *testing.T) {
	var a = getTestState()
	var action = &SyncLastUpdatedAction{}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.True(t, a.Updated < b.Updated)
}

func TestDrawCardAction(t *testing.T) {
	var a = getTestState()
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
	var a = getTestState()
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
	var a = getTestState()
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
	var a = getTestState()
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
	var a = getTestState()
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
	var a = getTestState()
	var action = &TurnClockwise90Action{Owner: a.Players[0].ID}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, a.Players[0].Facing+1, b.Players[0].Facing)
}

func TestTurnCounterClockwise90Action(t *testing.T) {
	var a = getTestState()
	var action = &TurnCounterClockwise90Action{Owner: a.Players[0].ID}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, 3, b.Players[0].Facing)
}
