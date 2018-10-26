package conditions

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// SpaceCollisionCondition address the case of two players moving into the same space
type SpaceCollisionCondition struct{}

// Apply applies the condition to the game state
func (c *SpaceCollisionCondition) Apply(gameState *pogo.GameState, actionGroup []actions.Action) error {
	return nil
}
