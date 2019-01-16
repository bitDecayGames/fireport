package actions

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/animations"

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
	a.AddEmptyAnimationSlice()

	// Animation check for a state
	// GetTestState contains 3 animations by default.
	assert.Equal(t, len(a.Animations[len(a.Animations)-1]), 0)
	assert.Equal(t, len(a.Animations), 4)

	b, err := action.Apply(a)
	assert.NoError(t, err)
	assert.Equal(t, 1, b.Players[0].Health)
	assert.Equal(t, 0, b.Players[1].Health)

	// Animation check for b state
	for _, animation := range b.Animations[len(b.Animations)-1] {
		if animation.Owner == b.Players[0].ID {
			assert.Equal(t, animation.Name, animations.GetFireCanon(b.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetFireCanon(b.Players[0].ID).ID)

		} else if animation.Owner == b.Players[1].ID {
			assert.Equal(t, animation.Name, animations.GetHitByCanon(b.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetHitByCanon(b.Players[0].ID).ID)

		} else {
			assert.Fail(t, "Incorrect animation found within state.")
		}
	}
	b.AddEmptyAnimationSlice()
	assert.Equal(t, len(b.Animations), 5)

	b.Players[0].Location = 3
	b.Players[0].Facing = WEST

	b.Players[1].Location = 2
	b.Players[1].Health = 1

	c, err := action.Apply(b)
	assert.NoError(t, err)
	assert.Equal(t, 1, c.Players[0].Health)
	assert.Equal(t, 0, c.Players[1].Health)

	// Animation check for c state
	for _, animation := range c.Animations[len(c.Animations)-1] {
		if animation.Owner == c.Players[0].ID {
			assert.Equal(t, animation.Name, animations.GetFireCanon(c.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetFireCanon(c.Players[0].ID).ID)

		} else if animation.Owner == c.Players[1].ID {
			assert.Equal(t, animation.Name, animations.GetHitByCanon(c.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetHitByCanon(c.Players[0].ID).ID)

		} else {
			assert.Fail(t, "Incorrect animation found within state.")
		}
	}
	c.AddEmptyAnimationSlice()
	assert.Equal(t, len(c.Animations), 6)

	c.Players[0].Location = 3
	c.Players[0].Facing = NORTH

	c.Players[1].Location = 1
	c.Players[1].Health = 1

	d, err := action.Apply(c)
	assert.NoError(t, err)
	assert.Equal(t, 1, d.Players[0].Health)
	assert.Equal(t, 0, d.Players[1].Health)

	// Animation check for d state
	for _, animation := range d.Animations[len(d.Animations)-1] {
		if animation.Owner == d.Players[0].ID {
			assert.Equal(t, animation.Name, animations.GetFireCanon(d.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetFireCanon(d.Players[0].ID).ID)

		} else if animation.Owner == d.Players[1].ID {
			assert.Equal(t, animation.Name, animations.GetHitByCanon(d.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetHitByCanon(d.Players[0].ID).ID)

		} else {
			assert.Fail(t, "Incorrect animation found within state.")
		}
	}
	d.AddEmptyAnimationSlice()
	assert.Equal(t, len(d.Animations), 7)

	d.Players[0].Location = 0
	d.Players[0].Facing = EAST

	d.Players[1].Health = 1

	e, err := action.Apply(d)
	assert.NoError(t, err)
	assert.Equal(t, 1, e.Players[0].Health)
	assert.Equal(t, 0, e.Players[1].Health)

	// Animation check for e state
	for _, animation := range e.Animations[len(e.Animations)-1] {
		if animation.Owner == e.Players[0].ID {
			assert.Equal(t, animation.Name, animations.GetFireCanon(e.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetFireCanon(e.Players[0].ID).ID)

		} else if animation.Owner == e.Players[1].ID {
			assert.Equal(t, animation.Name, animations.GetHitByCanon(e.Players[0].ID).Name)
			assert.Equal(t, animation.ID, animations.GetHitByCanon(e.Players[0].ID).ID)

		} else {
			assert.Fail(t, "Incorrect animation found within state.")
		}
	}

	// Animation check for total size of animation slice
	assert.Equal(t, len(e.Animations), 7)
}
