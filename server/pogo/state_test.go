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
		IDCounter:   300,
		BoardWidth:  2,
		BoardHeight: 2,
		BoardSpaces: []BoardSpace{
			{ID: 0, SpaceType: 0, State: 0},
			{ID: 1, SpaceType: 0, State: 0},
			{ID: 2, SpaceType: 0, State: 0},
			{ID: 3, SpaceType: 0, State: 0},
		},
		Players: []PlayerState{
			{
				ID:       100,
				Name:     "PlayerOne",
				Location: 0,
				Hand: []CardState{
					{ID: 101, CardType: 0},
					{ID: 102, CardType: 0},
					{ID: 103, CardType: 0},
					{ID: 104, CardType: 0},
					{ID: 105, CardType: 0},
				},
				Deck: []CardState{
					{ID: 106, CardType: 0},
					{ID: 107, CardType: 0},
					{ID: 108, CardType: 0},
					{ID: 109, CardType: 0},
					{ID: 110, CardType: 0},
				},
				Discard: []CardState{
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
				Hand: []CardState{
					{ID: 201, CardType: 0},
					{ID: 202, CardType: 0},
					{ID: 203, CardType: 0},
					{ID: 204, CardType: 0},
					{ID: 205, CardType: 0},
				},
				Deck: []CardState{
					{ID: 206, CardType: 0},
					{ID: 207, CardType: 0},
					{ID: 208, CardType: 0},
					{ID: 209, CardType: 0},
					{ID: 210, CardType: 0},
				},
				Discard: []CardState{
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

func TestGameState_DeepCopy(t *testing.T) {
	var a = getTestState()
	var b = a.DeepCopy()

	assert.Equal(t, a.Turn, b.Turn)
	b.Turn++
	assert.Equal(t, a.Turn+1, b.Turn)

	assert.Equal(t, a.Players[0].ID, b.Players[0].ID)
	b.Players[0].ID++
	assert.Equal(t, a.Players[0].ID+1, b.Players[0].ID)

	assert.Equal(t, a.Players[0].Hand[0].ID, b.Players[0].Hand[0].ID)
	b.Players[0].Hand[0].ID++
	assert.Equal(t, a.Players[0].Hand[0].ID+1, b.Players[0].Hand[0].ID)

	assert.Equal(t, len(a.Players), len(b.Players))
	b.Players = append(b.Players, PlayerState{})
	assert.Equal(t, len(a.Players)+1, len(b.Players))
}

func TestGameState_GetPlayer(t *testing.T) {
	var state = getTestState()

	var a = state.GetPlayer(state.Players[0].ID)
	var b = state.GetPlayer(state.Players[0].ID)

	a.Facing = a.Facing + 1

	assert.Equal(t, a.Facing, b.Facing)
}