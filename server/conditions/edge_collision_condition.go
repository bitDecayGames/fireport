package conditions

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// EdgeCollisionCondition address the case of two players moving into adjacent spaces through the same edge (passing through each other)
type EdgeCollisionCondition struct{}

// Apply applies the condition to the game state
func (c *EdgeCollisionCondition) Apply(gameState *pogo.GameState, actionGroup []actions.Action) error {
	return nil
}
