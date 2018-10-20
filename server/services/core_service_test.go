package services

import (
	"github.com/bitdecaygames/fireport/server/pogo"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Step a game one turn with one game input that points to the IncrementTurnAction
func TestIncrementTurn(t *testing.T) {
	coreSvc := &CoreServiceImpl{}

	curState := &pogo.GameState{}
	nextState, err := coreSvc.StepGame(curState, []pogo.GameInput{pogo.GameInput{CardId: 0, Owner: 0}})

	assert.Nil(t, err)
	assert.Equal(t, curState.TurnsTaken+1, nextState.TurnsTaken)
}
