package conditions

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/bitdecaygames/fireport/server/services"
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

/*
+++
+A+
+B+
+++
*/
func TestTwoPlayersMoveButDoNotCollide(t *testing.T) {
	var gameState = getTestState(3, 4, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 4, Facing: 2},
		{ID: 200, Name: "B", Location: 7, Facing: 1},
	})
	var inputs = []pogo.GameInputMsg{
		// A Move Forward
		// B Move Forward
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
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
++++
+AB+
*/
func TestTwoPlayersTurnIntoEachOther(t *testing.T) {
	var gameState = getTestState(4, 2, []pogo.PlayerState{
		{ID: 100, Name: "A", Location: 5, Facing: 0},
		{ID: 200, Name: "B", Location: 6, Facing: 0},
	})
	var inputs = []pogo.GameInputMsg{
		// A Turn Right
		// B Turn Left
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
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
		{ID: 100, Name: "A", Location: 9, Facing: 0},
		{ID: 200, Name: "B", Location: 10, Facing: 0},
		{ID: 300, Name: "C", Location: 1, Facing: 2},
	})
	var inputs = []pogo.GameInputMsg{
		// A Turn Left
		// B Turn Left
		// C Move Forward
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
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
		{ID: 100, Name: "A", Location: 9, Facing: 0},
		{ID: 200, Name: "B", Location: 10, Facing: 0},
		{ID: 300, Name: "C", Location: 0, Facing: 2},
	})
	var inputs = []pogo.GameInputMsg{
		// A Turn Left
		// B Turn Left
		// C Turn Left
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
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
		{ID: 100, Name: "A", Location: 12, Facing: 0},
		{ID: 200, Name: "B", Location: 8, Facing: 3},
		{ID: 300, Name: "C", Location: 2, Facing: 2},
		{ID: 400, Name: "D", Location: 6, Facing: 1},
	})
	var inputs = []pogo.GameInputMsg{
		// A Turn Right
		// B Turn Right
		// C Turn Right
		// D Turn Right
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
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
		{ID: 100, Name: "A", Location: 14, Facing: 0},
		{ID: 200, Name: "B", Location: 1, Facing: 2},
	})
	var inputs = []pogo.GameInputMsg{
		// A Turn Left
		// B Move Forward x3
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++++
		+B++
		++A+
		++++
	*/
	assert.Equal(t, 10, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 5, nextState.Players[1].Location)
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
		{ID: 100, Name: "A", Location: 14, Facing: 0},
		{ID: 200, Name: "B", Location: 2, Facing: 2},
	})
	var inputs = []pogo.GameInputMsg{
		// A Turn Left
		// B Move Forward x3
	} // TODO: fill out these inputs
	var core = &services.CoreServiceImpl{}

	var nextState, err = core.StepGame(gameState, inputs)
	assert.NoError(t, err)

	/*
		++++
		++++
		+A++
		++B+
	*/
	assert.Equal(t, 9, nextState.Players[0].Location)
	assert.Equal(t, 3, nextState.Players[0].Facing)
	assert.Equal(t, 14, nextState.Players[1].Location)
	assert.Equal(t, 2, nextState.Players[1].Facing)
}
