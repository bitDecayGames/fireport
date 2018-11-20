package logic

import (
	"math/rand"
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
)

func TestIncrementTurn(t *testing.T) {
	curState := &pogo.GameState{
		RNG: rand.New(rand.NewSource(0)),
		Players: []pogo.PlayerState{
			{
				ID: 0,
				Hand: []pogo.CardState{
					{
						ID:       1,
						CardType: pogo.SkipTurn,
					},
				},
			},
		},
	}
	// this will actually increment the turn by 2 because of the DefaultTurnActions list being applied at the end
	nextState, err := StepGame(curState, []pogo.GameInputMsg{{CardID: 1, Owner: 0}})

	assert.Nil(t, err)
	assert.Equal(t, curState.Turn+2, nextState.Turn)
}
