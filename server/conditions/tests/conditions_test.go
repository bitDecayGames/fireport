package tests

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/animations"

	"github.com/bitdecaygames/fireport/server/conditions"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
)

func getTestState(width int, height int, players []pogo.PlayerState) *pogo.GameState {
	var boardSpaces []pogo.BoardSpace
	for i := 0; i < width*height; i++ {
		boardSpaces = append(boardSpaces, pogo.BoardSpace{ID: 0, SpaceType: 0, State: 0})
	}
	return &pogo.GameState{
		Turn:        0,
		Created:     1000,
		Updated:     2000,
		IDCounter:   1000,
		BoardWidth:  width,
		BoardHeight: height,
		BoardSpaces: boardSpaces,
		Players:     players,
	}
}

func stepGame(currentState *pogo.GameState, inputs []pogo.GameInputMsg) (*pogo.GameState, error) {
	return conditions.ProcessConditions(currentState, inputs, []conditions.Condition{
		&conditions.SpaceCollisionCondition{},
		&conditions.EdgeCollisionCondition{},
		&conditions.BoundaryCollisionCondition{},
	})
}

/*
+++
+A+
+B+
+++
*/
func TestTwoPlayersAreMoving(t *testing.T) {
	var gameState = getTestState(3, 4, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 1, Facing: 2, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardOne}}},
		{ID: 200, Name: "B", Location: 7, Facing: 1, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.MoveForwardOne}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
		{CardID: 1001, Owner: 200, Order: 1}, // B Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		+A+
		++B
		+++
	*/
	assert.Equal(t, 4, nextState.Players[0].Location)
	assert.Equal(t, 2, nextState.Players[0].Facing)
	assert.Equal(t, 8, nextState.Players[1].Location)
	assert.Equal(t, 1, nextState.Players[1].Facing)

	//animation tests
	assert.Equal(t, len(nextState.Animations), 1)
	for _, animation := range nextState.Animations[0] {
		assert.Equal(t, animation.Name, animations.GetMoveForward(1).Name)
		assert.Equal(t, animation.ID, animations.GetMoveForward(1).ID)
		assert.Subset(t, []int{100, 200}, []int{animation.Owner})
	}
}

/*
+++
+A+
+B+
+++
*/
func TestTwoPlayersMoveButDoNotCollide(t *testing.T) {
	var gameState = getTestState(3, 4, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 2, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardOne}}},
		{ID: 200, Name: "B", Location: 7, Facing: 1, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.MoveForwardOne}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
		{CardID: 1001, Owner: 200, Order: 1}, // B Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		+++
		+AB
		+++
	*/
	assert.Equal(t, 7, nextState.Players[0].Location)
	assert.Equal(t, 2, nextState.Players[0].Facing)
	assert.Equal(t, 8, nextState.Players[1].Location)
	assert.Equal(t, 1, nextState.Players[1].Facing)

	//animation tests
	assert.Equal(t, len(nextState.Animations), 1)
	for _, animation := range nextState.Animations[0] {
		assert.Equal(t, animation.Name, animations.GetMoveForward(1).Name)
		assert.Equal(t, animation.ID, animations.GetMoveForward(1).ID)
		assert.Subset(t, []int{100, 200}, []int{animation.Owner})
	}
}

/*
+++
+A+
B++
+++
*/
func TestTwoPlayersSimpleSpaceCollide(t *testing.T) {
	var gameState = getTestState(3, 4, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 2, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardOne}}},
		{ID: 200, Name: "B", Location: 6, Facing: 1, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.MoveForwardOne}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
		{CardID: 1001, Owner: 200, Order: 1}, // B Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		+A+
		B++
		+++
	*/
	assert.Equal(t, 4, nextState.Players[0].Location)
	assert.Equal(t, 2, nextState.Players[0].Facing)
	assert.Equal(t, 6, nextState.Players[1].Location)
	assert.Equal(t, 1, nextState.Players[1].Facing)

	//animation tests
	assert.Equal(t, len(nextState.Animations), 1)
	for _, animation := range nextState.Animations[0] {
		assert.Equal(t, animation.Name, animations.GetBumpedInto(1).Name)
		assert.Equal(t, animation.ID, animations.GetBumpedInto(1).ID)
		assert.Subset(t, []int{100, 200}, []int{animation.Owner})
	}
}

/*
++++
+BA+
++++
*/
func TestTwoPlayersSimpleEdgeCollide(t *testing.T) {
	var gameState = getTestState(4, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 6, Facing: 3, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardOne}}},
		{ID: 200, Name: "B", Location: 5, Facing: 1, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.MoveForwardOne}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
		{CardID: 1001, Owner: 200, Order: 1}, // B Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++++
		+BA+
		++++
	*/
	assert.Equal(t, 6, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 5, nextState.Players[1].Location)
	assert.Equal(t, 1, nextState.Players[1].Facing)

	//animation tests
	assert.Equal(t, len(nextState.Animations), 1)
	for _, animation := range nextState.Animations[0] {
		assert.Equal(t, animation.Name, animations.GetBumpedInto(1).Name)
		assert.Equal(t, animation.ID, animations.GetBumpedInto(1).ID)
		assert.Subset(t, []int{100, 200}, []int{animation.Owner})
	}
}

/*
++++
+AB+
*/
func TestTwoPlayersTurnIntoEachOther(t *testing.T) {
	var gameState = getTestState(4, 2, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 5, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnRight}}},
		{ID: 200, Name: "B", Location: 6, Facing: 0, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.TurnLeft}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Right
		{CardID: 1001, Owner: 200, Order: 1}, // B Turn Left
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+AB+
		++++
	*/
	assert.Equal(t, 1, nextState.Players[0].Location)
	assert.Equal(t, 1, nextState.Players[0].Facing)
	assert.Equal(t, 2, nextState.Players[1].Location)
	assert.Equal(t, 3, nextState.Players[1].Facing)

	//animation tests
	assert.Equal(t, len(nextState.Animations), 3)
	for _, animation := range nextState.Animations[0] {
		assert.Equal(t, animation.Name, animations.GetMoveForward(1).Name)
		assert.Equal(t, animation.ID, animations.GetMoveForward(1).ID)
		assert.Subset(t, []int{100, 200}, []int{animation.Owner})
	}
	for _, animation := range nextState.Animations[1] {
		if animation.Owner == 100 {
			assert.Equal(t, animation.Name, animations.GetTurnClockwise90(1).Name)
			assert.Equal(t, animation.ID, animations.GetTurnClockwise90(1).ID)
		} else if animation.Owner == 200 {
			assert.Equal(t, animation.Name, animations.GetTurnCounterClockwise90(1).Name)
			assert.Equal(t, animation.ID, animations.GetTurnCounterClockwise90(1).ID)
		} else {
			assert.Equal(t, 1, 2, "Animations don't have correct owner")
		}
	}
	for _, animation := range nextState.Animations[2] {
		assert.Equal(t, animation.Name, animations.GetBumpedInto(1).Name)
		assert.Equal(t, animation.ID, animations.GetBumpedInto(1).ID)
		assert.Subset(t, []int{100, 200}, []int{animation.Owner})
	}
}

/*
+C++
++++
+AB+
*/
func TestTwoOutOfThreePlayersCollide(t *testing.T) {
	var gameState = getTestState(4, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 9, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnLeft}}},
		{ID: 200, Name: "B", Location: 10, Facing: 0, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.TurnLeft}}},
		{ID: 300, Name: "C", Location: 1, Facing: 2, Hand: []pogo.CardState{{ID: 1002, CardType: pogo.MoveForwardOne}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Left
		{CardID: 1001, Owner: 200, Order: 1}, // B Turn Left
		{CardID: 1002, Owner: 300, Order: 1}, // C Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+C++
		+B++
		A+++
	*/
	assert.Equal(t, 8, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 5, nextState.Players[1].Location)
	assert.Equal(t, 3, nextState.Players[1].Facing)
	assert.Equal(t, 1, nextState.Players[2].Location)
	assert.Equal(t, 2, nextState.Players[2].Facing)
}

/*
C+++
++++
+AB+
*/
func TestOneCollisionCausesAnother(t *testing.T) {
	var gameState = getTestState(4, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 9, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnLeft}}},
		{ID: 200, Name: "B", Location: 10, Facing: 0, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.TurnLeft}}},
		{ID: 300, Name: "C", Location: 0, Facing: 2, Hand: []pogo.CardState{{ID: 1002, CardType: pogo.TurnLeft}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Left
		{CardID: 1001, Owner: 200, Order: 1}, // B Turn Left
		{CardID: 1002, Owner: 300, Order: 1}, // C Turn Left
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++++
		CAB+
		++++
	*/
	assert.Equal(t, 5, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 6, nextState.Players[1].Location)
	assert.Equal(t, 3, nextState.Players[1].Facing)
	assert.Equal(t, 4, nextState.Players[2].Location)
	assert.Equal(t, 1, nextState.Players[2].Facing)
}

/*
++C++
+D+B+
++A++
*/
func TestTurningTanksCollide(t *testing.T) {
	var gameState = getTestState(5, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 12, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnRight}}},
		{ID: 200, Name: "B", Location: 8, Facing: 3, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.TurnRight}}},
		{ID: 300, Name: "C", Location: 2, Facing: 2, Hand: []pogo.CardState{{ID: 1002, CardType: pogo.TurnRight}}},
		{ID: 400, Name: "D", Location: 6, Facing: 1, Hand: []pogo.CardState{{ID: 1003, CardType: pogo.TurnRight}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Right
		{CardID: 1001, Owner: 200, Order: 1}, // B Turn Right
		{CardID: 1002, Owner: 300, Order: 1}, // C Turn Right
		{CardID: 1003, Owner: 400, Order: 1}, // D Turn Right
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+C+B+
		+++++
		+D+A+
	*/
	assert.Equal(t, 13, nextState.Players[0].Location)
	assert.Equal(t, 1, nextState.Players[0].Facing)
	assert.Equal(t, 3, nextState.Players[1].Location)
	assert.Equal(t, 0, nextState.Players[1].Facing)
	assert.Equal(t, 1, nextState.Players[2].Location)
	assert.Equal(t, 3, nextState.Players[2].Facing)
	assert.Equal(t, 11, nextState.Players[3].Location)
	assert.Equal(t, 2, nextState.Players[3].Facing)
}

/*
+B++
++++
++++
++A+
*/
func TestRotatingAndMovementPartI(t *testing.T) {
	var gameState = getTestState(4, 4, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 14, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnLeft}}},
		{ID: 200, Name: "B", Location: 1, Facing: 2, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.MoveForwardThree}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Left
		{CardID: 1001, Owner: 200, Order: 1}, // B Move Forward x3
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++++
		++++
		+A++
		+B++
	*/
	assert.Equal(t, 9, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 13, nextState.Players[1].Location)
	assert.Equal(t, 2, nextState.Players[1].Facing)
}

/*
++B+
++++
++++
++A+
*/
func TestRotatingAndMovementPartII(t *testing.T) {
	var gameState = getTestState(4, 4, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 14, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnLeft}}},
		{ID: 200, Name: "B", Location: 2, Facing: 2, Hand: []pogo.CardState{{ID: 1001, CardType: pogo.MoveForwardThree}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Left
		{CardID: 1001, Owner: 200, Order: 1}, // B Move Forward x3
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++++
		++++
		+AB+
		++++
	*/
	assert.Equal(t, 9, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 10, nextState.Players[1].Location)
	assert.Equal(t, 2, nextState.Players[1].Facing)
}

/*
+++
+A+
+++
*/
func TestCollisionWithTopSide(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardTwo}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+A+
		+++
		+++
	*/
	assert.Equal(t, 1, nextState.Players[0].Location)
	assert.Equal(t, gameState.Players[0].Health-1, nextState.Players[0].Health)
}

/*
+++
+A+
+++
*/
func TestCollisionWithBottomSide(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 2, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardTwo}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		+++
		+A+
	*/
	assert.Equal(t, 7, nextState.Players[0].Location)
	assert.Equal(t, gameState.Players[0].Health-1, nextState.Players[0].Health)
}

/*
+++
+A+
+++
*/
func TestCollisionWithLeftSide(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 3, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardTwo}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		A++
		+++
	*/
	assert.Equal(t, 3, nextState.Players[0].Location)
	assert.Equal(t, gameState.Players[0].Health-1, nextState.Players[0].Health)
}

/*
+++
+A+
+++
*/
func TestCollisionWithRightSide(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 1, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.MoveForwardTwo}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Move Forward
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		++A
		+++
	*/
	assert.Equal(t, 5, nextState.Players[0].Location)
	assert.Equal(t, gameState.Players[0].Health-1, nextState.Players[0].Health)
}

/*
A++
+++
+++
*/
func TestCollisionWithTopLeftCorner(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 0, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnLeft}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Left
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		A++
		+++
		+++
	*/
	assert.Equal(t, 0, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, gameState.Players[0].Health-2, nextState.Players[0].Health)
}

/*
++A
+++
+++
*/
func TestCollisionWithTopRightCorner(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 2, Facing: 0, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnRight}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Right
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++A
		+++
		+++
	*/
	assert.Equal(t, 2, nextState.Players[0].Location)
	assert.Equal(t, 1, nextState.Players[0].Facing)
	assert.Equal(t, gameState.Players[0].Health-2, nextState.Players[0].Health)
}

/*
+++
+++
++A
*/
func TestCollisionWithBottomRightCorner(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 8, Facing: 2, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnLeft}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Left
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		+++
		++A
	*/
	assert.Equal(t, 8, nextState.Players[0].Location)
	assert.Equal(t, 1, nextState.Players[0].Facing)
	assert.Equal(t, gameState.Players[0].Health-2, nextState.Players[0].Health)
}

/*
+++
+++
A++
*/
func TestCollisionWithBottomLeftCorner(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 6, Facing: 2, Hand: []pogo.CardState{{ID: 1000, CardType: pogo.TurnRight}}},
	})
	var inputs = []pogo.GameInputMsg{
		{CardID: 1000, Owner: 100, Order: 1}, // A Turn Right
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		+++
		A++
	*/
	assert.Equal(t, 6, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, gameState.Players[0].Health-2, nextState.Players[0].Health)
}

/*
+++
++B
++A
*/
func TestEdgeCollisionWithPlayersGoingSameDirectionPartI(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 8, Facing: 2, Hand: []pogo.CardState{
			{ID: 101, CardType: pogo.MoveForwardTwo},
		}},
		{ID: 200, Name: "B", Location: 5, Facing: 2, Hand: []pogo.CardState{
			{ID: 201, CardType: pogo.MoveForwardTwo},
		}},
	})
	var inputs []pogo.GameInputMsg

	for _, player := range gameState.Players {
		for order, card := range player.Hand {
			inputs = append(inputs, pogo.GameInputMsg{CardID: card.ID, Owner: player.ID, Order: order})
		}
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		++B
		++A
	*/
	assert.Equal(t, 8, nextState.Players[0].Location, "A is in the wrong location")
	assert.Equal(t, 2, nextState.Players[0].Facing, "A is facing the wrong way")
	assert.Equal(t, gameState.Players[0].Health-2, nextState.Players[0].Health, "A has the wrong health")
	assert.Equal(t, 5, nextState.Players[1].Location, "B is in the wrong location")
	assert.Equal(t, 2, nextState.Players[1].Facing, "B is facing the wrong way")
	assert.Equal(t, gameState.Players[1].Health-2, nextState.Players[1].Health, "B has the wrong health")
	assert.NotEqual(t, nextState.Players[0].Location, nextState.Players[1].Location, "A and B are in the same location")
}

/*
+++
++B
++A
*/
func TestEdgeCollisionWithPlayersGoingSameDirectionPartII(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 8, Facing: 2, Hand: []pogo.CardState{
			{ID: 101, CardType: pogo.MoveForwardTwo},
		}},
		{ID: 200, Name: "B", Location: 5, Facing: 2, Hand: []pogo.CardState{
			{ID: 201, CardType: pogo.MoveForwardOne},
		}},
	})
	var inputs []pogo.GameInputMsg

	for _, player := range gameState.Players {
		for order, card := range player.Hand {
			inputs = append(inputs, pogo.GameInputMsg{CardID: card.ID, Owner: player.ID, Order: order})
		}
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		++B
		++A
	*/
	assert.Equal(t, 8, nextState.Players[0].Location, "A is in the wrong location")
	assert.Equal(t, 2, nextState.Players[0].Facing, "A is facing the wrong way")
	assert.Equal(t, gameState.Players[0].Health-2, nextState.Players[0].Health, "A has the wrong health")
	assert.Equal(t, 5, nextState.Players[1].Location, "B is in the wrong location")
	assert.Equal(t, 2, nextState.Players[1].Facing, "B is facing the wrong way")
	assert.Equal(t, gameState.Players[1].Health-1, nextState.Players[1].Health, "B has the wrong health")
	assert.NotEqual(t, nextState.Players[0].Location, nextState.Players[1].Location, "A and B are in the same location")
}

/*
+++
++B
++A
*/
func TestEdgeCollisionWithPlayersGoingSameDirectionPartIII(t *testing.T) {
	var gameState = getTestState(3, 3, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 8, Facing: 2, Hand: []pogo.CardState{
			{ID: 101, CardType: pogo.MoveForwardOne},
		}},
		{ID: 200, Name: "B", Location: 5, Facing: 2, Hand: []pogo.CardState{
			{ID: 201, CardType: pogo.MoveForwardOne},
		}},
	})
	var inputs []pogo.GameInputMsg

	for _, player := range gameState.Players {
		for order, card := range player.Hand {
			inputs = append(inputs, pogo.GameInputMsg{CardID: card.ID, Owner: player.ID, Order: order})
		}
	}

	var nextState, err = stepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		+++
		++B
		++A
	*/
	assert.Equal(t, 8, nextState.Players[0].Location, "A is in the wrong location")
	assert.Equal(t, 2, nextState.Players[0].Facing, "A is facing the wrong way")
	assert.Equal(t, gameState.Players[0].Health-1, nextState.Players[0].Health, "A has the wrong health")
	assert.Equal(t, 5, nextState.Players[1].Location, "B is in the wrong location")
	assert.Equal(t, 2, nextState.Players[1].Facing, "B is facing the wrong way")
	assert.Equal(t, gameState.Players[1].Health-1, nextState.Players[1].Health, "B has the wrong health")
	assert.NotEqual(t, nextState.Players[0].Location, nextState.Players[1].Location, "A and B are in the same location")
}
