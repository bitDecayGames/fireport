package services

import (
	"github.com/bitdecaygames/fireport/server/pogo"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Step a game one turn with one game input that points to the IncrementTurnAction
func TestIncrementTurn(t *testing.T) {
	coreSvc := &CoreServiceImpl{}

	curState := &pogo.GameState{Players: []pogo.PlayerState{{ID: 0, Hand: []pogo.CardState{{ID: 1, CardType: pogo.SkipTurn}}}}}
	// this will actually increment the turn by 2 because of the DefaultTurnActions list being applied at the end
	nextState, err := coreSvc.StepGame(curState, []pogo.GameInputMsg{{CardID: 1, Owner: 0}})

	assert.Nil(t, err)
	assert.Equal(t, curState.Turn+2, nextState.Turn)
}
