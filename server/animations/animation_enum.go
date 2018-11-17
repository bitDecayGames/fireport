package animations

type Animation int // UIDs associated to animations

// UIDs associated with each animaion
const (
	MoveForward            Animation = 0
	MoveBackward           Animation = 1
	TurnClockwise90        Animation = 2
	TurnCounterClockwise90 Animation = 3
	BumpInto               Animation = 4
	TakeDamage             Animation = 5
	FireCanon              Animation = 6
)

// String returns the string name for an animation enum
func (anim Animation) String() string {
	animationNames := [...]string{
		"MoveForward",
		"MoveBackward",
		"TurnClockwise90",
		"TurnCounterClockwise90",
		"BumpInto",
		"TakeDamage",
		"FireCanon"}

	// prevent panicking in case anim is out of range
	if anim < MoveForward || anim > FireCanon {
		return "UnknownAnimation"
	}

	return animationNames[anim]
}

// UID returns the UID for the animation
func (anim Animation) UID() int {
	return int(anim)
}
