package animations

// UIDs associated to animations
const (
	MoveForward            = 100
	MoveBackward           = 101
	TurnClockwise90        = 102
	TurnCounterClockwise90 = 103
	BeBumpedInto           = 104
	DoBumpInto             = 105
	FireCanon              = 301
	HitByCanon             = 302
)

// AnimationAction tracks the specific animations required by the client to move from state A to state B
type AnimationAction struct {
	ID    int    // id for this specific action
	Name  string // name key for the type of animation
	Owner int    // UID of the player the animation is associated with
}

//GetMoveForward animation
func GetMoveForward(owner int) AnimationAction {
	return AnimationAction{ID: int(MoveForward), Name: "MoveForward", Owner: owner}
}

//GetMoveBackward animation
func GetMoveBackward(owner int) AnimationAction {
	return AnimationAction{ID: int(MoveBackward), Name: "MoveBackward", Owner: owner}
}

//GetTurnClockwise90 animation
func GetTurnClockwise90(owner int) AnimationAction {
	return AnimationAction{ID: int(TurnClockwise90), Name: "TurnClockwise90", Owner: owner}
}

//GetTurnCounterClockwise90 animation
func GetTurnCounterClockwise90(owner int) AnimationAction {
	return AnimationAction{ID: int(TurnCounterClockwise90), Name: "TurnCounterClockwise90", Owner: owner}
}

//GetBeBumpedInto animation
func GetBeBumpedInto(owner int) AnimationAction {
	return AnimationAction{ID: int(BeBumpedInto), Name: "BeBumpedInto", Owner: owner}
}

//GetDoBumpInto animation
func GetDoBumpInto(owner int) AnimationAction {
	return AnimationAction{ID: int(DoBumpInto), Name: "DoBumpInto", Owner: owner}
}

//GetFireCanon animation
func GetFireCanon(owner int) AnimationAction {
	return AnimationAction{ID: int(FireCanon), Name: "FireCanon", Owner: owner}
}

//GetHitByCanon animation
func GetHitByCanon(owner int) AnimationAction {
	return AnimationAction{ID: int(HitByCanon), Name: "HitByCanon", Owner: owner}
}
