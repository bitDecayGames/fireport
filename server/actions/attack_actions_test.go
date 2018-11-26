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

	action := &FireBasicAction{Owner: a.Players[0].ID}

	b, err := action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, 1, b.Players[0].Health)
	assert.Equal(t, 0, b.Players[1].Health)

	b.Players[0].Location = 3
	b.Players[0].Facing = WEST

	b.Players[1].Location = 2
	b.Players[1].Health = 1

	c, err := action.Apply(b)
	assert.NoError(t, err)
	assert.Equal(t, 1, c.Players[0].Health)
	assert.Equal(t, 0, c.Players[1].Health)

	c.Players[0].Location = 3
	c.Players[0].Facing = NORTH

	c.Players[1].Location = 1
	c.Players[1].Health = 1

	d, err := action.Apply(c)
	assert.NoError(t, err)
	assert.Equal(t, 1, d.Players[0].Health)
	assert.Equal(t, 0, d.Players[1].Health)

	d.Players[0].Location = 0
	d.Players[0].Facing = EAST

	d.Players[1].Health = 1

	e, err := action.Apply(d)
	assert.NoError(t, err)
	assert.Equal(t, 1, e.Players[0].Health)
	assert.Equal(t, 0, e.Players[1].Health)
}
