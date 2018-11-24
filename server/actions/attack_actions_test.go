package actions

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/stretchr/testify/assert"
)

func TestFireAction(t *testing.T) {
	var a = pogo.GetTestState()
	a.Players[0].Location = 0
	a.Players[0].Facing = SOUTH
	a.Players[0].Health = 1

	a.Players[1].Location = 2
	a.Players[1].Facing = NORTH
	a.Players[1].Health = 1

	var action = &FireBasicAction{Owner: a.Players[0].ID}

	var b, err = action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, 1, b.Players[0].Health)
	assert.Equal(t, 0, b.Players[1].Health)
}
