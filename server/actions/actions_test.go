package actions

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
		IdCounter:   300,
		BoardWidth:  2,
		BoardHeight: 2,
		BoardSpaces: []pogo.BoardSpace{
			{Id: 0, SpaceType: 0, State: 0},
			{Id: 1, SpaceType: 0, State: 0},
			{Id: 2, SpaceType: 0, State: 0},
			{Id: 3, SpaceType: 0, State: 0},
		},
		Players: []pogo.PlayerState{
			{
				Id:       100,
				Name:     "PlayerOne",
				Location: 0,
				Hand: []pogo.CardState{
					{Id: 101, CardType: 0},
					{Id: 102, CardType: 0},
					{Id: 103, CardType: 0},
					{Id: 104, CardType: 0},
					{Id: 105, CardType: 0},
				},
				Deck: []pogo.CardState{
					{Id: 106, CardType: 0},
					{Id: 107, CardType: 0},
					{Id: 108, CardType: 0},
					{Id: 109, CardType: 0},
					{Id: 110, CardType: 0},
				},
				Discard: []pogo.CardState{
					{Id: 111, CardType: 0},
					{Id: 112, CardType: 0},
					{Id: 113, CardType: 0},
					{Id: 114, CardType: 0},
					{Id: 115, CardType: 0},
				},
			},
			{
				Id:       200,
				Name:     "PlayerTwo",
				Location: 3,
				Hand: []pogo.CardState{
					{Id: 201, CardType: 0},
					{Id: 202, CardType: 0},
					{Id: 203, CardType: 0},
					{Id: 204, CardType: 0},
					{Id: 205, CardType: 0},
				},
				Deck: []pogo.CardState{
					{Id: 206, CardType: 0},
					{Id: 207, CardType: 0},
					{Id: 208, CardType: 0},
					{Id: 209, CardType: 0},
					{Id: 210, CardType: 0},
				},
				Discard: []pogo.CardState{
					{Id: 211, CardType: 0},
					{Id: 212, CardType: 0},
					{Id: 213, CardType: 0},
					{Id: 214, CardType: 0},
					{Id: 215, CardType: 0},
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

func TestTurnClockwise90Action(t *testing.T) {
	var a = getTestState()
	var action = &TurnClockwise90Action{Owner: a.Players[0].Id}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, a.Players[0].Facing+1, b.Players[0].Facing)
}

func TestTurnCounterClockwise90Action(t *testing.T) {
	var a = getTestState()
	var action = &TurnCounterClockwise90Action{Owner: a.Players[0].Id}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, 3, b.Players[0].Facing)
}
