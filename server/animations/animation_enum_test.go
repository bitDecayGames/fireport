package animations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getAllAnimationsAsMap() map[string]Animation {
	animations := make(map[string]Animation)
	animations["MoveForward"] = MoveForward
	animations["MoveBackward"] = MoveBackward
	animations["TurnClockwise90"] = TurnClockwise90
	animations["TurnCounterClockwise90"] = TurnCounterClockwise90
	animations["BumpInto"] = BumpInto
	animations["TakeDamage"] = TakeDamage
	animations["FireCanon"] = FireCanon
	return animations
}

func getAllAnimationsAsSlice() []Animation {
	animations := []Animation{
		MoveForward,
		MoveBackward,
		TurnClockwise90,
		TurnCounterClockwise90,
		BumpInto,
		TakeDamage,
		FireCanon}
	return animations
}

func TestStringFunction(t *testing.T) {
	animations := getAllAnimationsAsMap()
	for k, v := range animations {
		assert.Equal(t, k, v.String())
	}
}

func TestAnimationIDs(t *testing.T) {
	animations := getAllAnimationsAsSlice()
	for id, anim := range animations {
		fmt.Printf("The animation: %v", anim.UID())
		assert.Equal(t, id, anim.UID())
	}
}
