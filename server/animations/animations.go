package animations

import (
	"github.com/bitdecaygames/fireport/server/pogo"
)

// UIDs associated to animations
const (
	MoveForward            = 100
	MoveBackward           = 101
	TurnClockwise90        = 102
	TurnCounterClockwise90 = 103
	BumpedInto             = 104
	BumpInto               = 105
	FireCanon              = 301
	HitByCanon             = 302
)

//GetMoveForward animation
func GetMoveForward(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(MoveForward), Name: "MoveForward", Owner: owner}
}

//GetMoveBackward animation
func GetMoveBackward(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(MoveBackward), Name: "MoveBackward", Owner: owner}
}

//GetTurnClockwise90 animation
func GetTurnClockwise90(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(TurnClockwise90), Name: "TurnClockwise90", Owner: owner}
}

//GetTurnCounterClockwise90 animation
func GetTurnCounterClockwise90(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(TurnCounterClockwise90), Name: "TurnCounterClockwise90", Owner: owner}
}

//GetBumpedInto animation
func GetBumpedInto(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(BumpedInto), Name: "BumpedInto", Owner: owner}
}

//GetBumpInto animation
func GetBumpInto(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(BumpInto), Name: "BumpInto", Owner: owner}
}

//GetFireCanon animation
func GetFireCanon(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(FireCanon), Name: "FireCanon", Owner: owner}
}

//GetHitByCanon animation
func GetHitByCanon(owner int) pogo.AnimationAction {
	return pogo.AnimationAction{ID: int(HitByCanon), Name: "HitByCanon", Owner: owner}
}
