package tests

import (
	"github.com/bitdecaygames/fireport/server/conditions"
	"testing"

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
