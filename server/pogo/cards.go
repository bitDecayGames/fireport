package pogo

// CardType an enum for card types
type CardType int

// this is the enum for CardTypes
const (
	Unknown CardType = -1

	SkipTurn CardType = 0

	MoveForwardOne   CardType = 100
	MoveForwardTwo   CardType = 101
	MoveForwardThree CardType = 102

	MoveBackwardOne   CardType = 103
	MoveBackwardTwo   CardType = 104
	MoveBackwardThree CardType = 105

	RotateRight CardType = 110
	RotateLeft  CardType = 111
	Rotate180   CardType = 112

	TurnRight CardType = 120
	TurnLeft  CardType = 121
)

// Priority is a helper func to return a numerical value for the priority of the card.
// Lower values indicate higher priority
func (c CardType) Priority() int {
	return int(c / 100)
}
