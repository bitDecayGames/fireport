package pogo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestState() *GameState {
	return &GameState{
		Turn:        0,
		Created:     1000,
		Updated:     2000,
		IdCounter:   300,
		BoardWidth:  2,
		BoardHeight: 2,
		BoardSpaces: []BoardSpace{
			{Id: 0, SpaceType: 0, State: 0},
			{Id: 1, SpaceType: 0, State: 0},
			{Id: 2, SpaceType: 0, State: 0},
			{Id: 3, SpaceType: 0, State: 0},
		},
		Players: []PlayerState{
			{
				Id:       100,
				Name:     "PlayerOne",
				Location: 0,
				Hand: []CardState{
					{Id: 101, CardType: 0},
					{Id: 102, CardType: 0},
					{Id: 103, CardType: 0},
					{Id: 104, CardType: 0},
					{Id: 105, CardType: 0},
				},
				Deck: []CardState{
					{Id: 106, CardType: 0},
					{Id: 107, CardType: 0},
					{Id: 108, CardType: 0},
					{Id: 109, CardType: 0},
					{Id: 110, CardType: 0},
				},
				Discard: []CardState{
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
				Hand: []CardState{
					{Id: 201, CardType: 0},
					{Id: 202, CardType: 0},
					{Id: 203, CardType: 0},
					{Id: 204, CardType: 0},
					{Id: 205, CardType: 0},
				},
				Deck: []CardState{
					{Id: 206, CardType: 0},
					{Id: 207, CardType: 0},
					{Id: 208, CardType: 0},
					{Id: 209, CardType: 0},
					{Id: 210, CardType: 0},
				},
				Discard: []CardState{
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

func TestGameState_DeepCopy(t *testing.T) {
	var a = getTestState()
	var b = a.DeepCopy()

	assert.Equal(t, a.Turn, b.Turn)
	b.Turn += 1
	assert.Equal(t, a.Turn+1, b.Turn)

	assert.Equal(t, a.Players[0].Id, b.Players[0].Id)
	b.Players[0].Id += 1
	assert.Equal(t, a.Players[0].Id+1, b.Players[0].Id)

	assert.Equal(t, a.Players[0].Hand[0].Id, b.Players[0].Hand[0].Id)
	b.Players[0].Hand[0].Id += 1
	assert.Equal(t, a.Players[0].Hand[0].Id+1, b.Players[0].Hand[0].Id)

	assert.Equal(t, len(a.Players), len(b.Players))
	b.Players = append(b.Players, PlayerState{})
	assert.Equal(t, len(a.Players)+1, len(b.Players))
}

func TestGameState_GetPlayer(t *testing.T) {
	var state = getTestState()

	var a = state.GetPlayer(state.Players[0].Id)
	var b = state.GetPlayer(state.Players[0].Id)

	a.Facing = a.Facing + 1

	assert.Equal(t, a.Facing, b.Facing)
}