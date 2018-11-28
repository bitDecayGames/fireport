package services

import (
	"encoding/json"
	"github.com/bitdecaygames/fireport/server/pogo"
)

// SubmitSimpleTestTurn an easy way to test a specific turn submission for errors
func (g *GameServiceImpl) SubmitSimpleTestTurn(gameID string, playerName string, playerID int, cards []int) error {
	turnSubmission := pogo.TurnSubmissionMsg{
		GameID:   gameID,
		PlayerID: playerName,
		Inputs:   []pogo.GameInputMsg{},
	}

	for order, cardID := range cards {
		turnSubmission.Inputs = append(turnSubmission.Inputs, pogo.GameInputMsg{
			Owner:  playerID,
			CardID: cardID,
			Order:  order,
			Swap:   0,
		})
	}

	return g.SubmitTurn(turnSubmission)
}

// SetTestGameState unmarshals the string into a pogo.GameState and then sets that game state to this instance's .State
func (g *GameInstance) SetTestGameState(state string) error {
	obj := &pogo.GameState{}
	data := []byte(state)
	// wish this didn't have to be here... but circular dependencies...
	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	g.State = *obj
	g.Players = []string{}
	for _, player := range g.State.Players {
		g.Players = append(g.Players, player.Name)
	}
	return nil
}
