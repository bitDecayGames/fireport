package actions

import (
	"github.com/bitdecaygames/fireport/server/pogo"
)

// Action the smallest unit of modification to be made to a game state object
type Action interface {
	Apply(currentState *pogo.GameState) (*pogo.GameState, error)
	GetOwner() int
}
