package pogo

// CardType an enum for card types
type CardType int

// this is the enum for CardTypes
// Cards in the 1xx are movement cards
// Cards in the 2xx are utility cards
// Cards in the 3xx are attack cards
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

	FireBasic CardType = 300
)

// Priority is a helper func to return a numerical value for the priority of the card.
// Lower values indicate higher priority
func (c CardType) Priority() int {
	return int(c / 100)
}
