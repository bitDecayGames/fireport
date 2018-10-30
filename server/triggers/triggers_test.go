package triggers

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
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
				Health:   10,
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
				Health:   10,
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

func TestWinTrigger(t *testing.T) {
	var trigger = &WinTrigger{}
	var a = getTestState()

	assert.Equal(t, 0, a.Winner)
	assert.Equal(t, false, a.IsGameFinished)

	var b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, 0, b.Winner)
	assert.Equal(t, false, b.IsGameFinished)

	a.Players[1].Health = 0

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, 100, b.Winner)
	assert.Equal(t, true, b.IsGameFinished)

	a.Players[0].Health = 0

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, -1, b.Winner)
	assert.Equal(t, true, b.IsGameFinished)
}

func TestNotEnoughCardsInHandTrigger(t *testing.T) {
	var trigger = &NotEnoughCardsInHandTrigger{RequiredCardsInHand: 5}
	var a = getTestState()

	var b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand))
	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))

	var action = &actions.DiscardCardAction{Owner: 100, CardID: 101}
	a, err = action.Apply(a)
	assert.NoError(t, err)

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand)-1)
	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck)+1)
}

func TestRefreshDeckTrigger(t *testing.T) {
	var trigger = &RefreshDeckTrigger{}
	var a = getTestState()

	var b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))

	a, err = actions.ApplyActions(a, []actions.Action{
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
	})
	assert.NoError(t, err)

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck)-5)

	a, err = actions.ApplyActions(b, []actions.Action{
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
	})
	assert.NoError(t, err)

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))
}

func TestDrawAndDeckRefreshTriggersTogether(t *testing.T) {
	var triggers = []Trigger{&NotEnoughCardsInHandTrigger{RequiredCardsInHand: 5}, &RefreshDeckTrigger{}}
	var a = getTestState()

	var b, err = ApplyTriggers(a, triggers)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand))
	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))
	assert.Equal(t, len(a.Players[0].Discard), len(b.Players[0].Discard))

	a, err = actions.ApplyActions(a, []actions.Action{
		&actions.DiscardCardAction{Owner: 100, CardID: 101},
		&actions.DiscardCardAction{Owner: 100, CardID: 102},
		&actions.DiscardCardAction{Owner: 100, CardID: 103},
		&actions.DiscardCardAction{Owner: 100, CardID: 104},
		&actions.DiscardCardAction{Owner: 100, CardID: 105},
		&actions.DrawCardAction{Owner: 100},
		&actions.DiscardCardAction{Owner: 100, CardID: 110},
		&actions.DrawCardAction{Owner: 100},
	})
	assert.NoError(t, err)

	assert.Equal(t, 1, len(a.Players[0].Hand))
	assert.Equal(t, 3, len(a.Players[0].Deck))
	assert.Equal(t, 11, len(a.Players[0].Discard))

	b, err = ApplyTriggers(a, triggers)
	assert.NoError(t, err)

	assert.Equal(t, 5, len(b.Players[0].Hand))
	assert.Equal(t, 10, len(b.Players[0].Deck))
	assert.Equal(t, 0, len(b.Players[0].Discard))
}
