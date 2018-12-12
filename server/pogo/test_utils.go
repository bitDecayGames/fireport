package pogo

import (
	"math/rand"

	"github.com/bitdecaygames/fireport/server/animations"
)

// Constants used in test states
const (
	PlayerOneID   = 100
	PlayerOneName = "PlayerOne"
	PlayerTwoID   = 200
	PlayerTwoName = "PlayerTwo"
)

// GetTestState is a base state for tests to run with
func GetTestState() *GameState {
	return &GameState{
		Turn:        0,
		RNG:         rand.New(rand.NewSource(0)), // this may be useful as our test state will always generate the same sequence of things
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
				ID:       PlayerOneID,
				Name:     PlayerOneName,
				Location: 0,
				Facing:   1,
				Health:   1,
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
				ID:       PlayerTwoID,
				Name:     PlayerTwoName,
				Location: 3,
				Facing:   2,
				Health:   2,
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

// GetTestStateAnimationActions is a base state for tests to run with including animation actions
func GetTestStateAnimationActions() *GameState {

	return &GameState{
		Turn:        1,
		RNG:         rand.New(rand.NewSource(0)), // this may be useful as our test state will always generate the same sequence of things
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
				ID:       PlayerOneID,
				Name:     PlayerOneName,
				Location: 0,
				Facing:   1,
				Health:   1,
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
				ID:       PlayerTwoID,
				Name:     PlayerTwoName,
				Location: 3,
				Facing:   2,
				Health:   2,
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
		Animations: [][]animations.AnimationAction{
			[]animations.AnimationAction{
				animations.GetMoveForward(PlayerOneID),
				animations.GetMoveBackward(PlayerTwoID),
			},
			[]animations.AnimationAction{
				animations.GetMoveForward(PlayerTwoID),
				animations.GetTurnClockwise90(PlayerOneID),
			},
			[]animations.AnimationAction{
				animations.GetFireCanon(PlayerOneID),
				animations.GetHitByCanon(PlayerTwoID),
			},
		},
	}
}
