package routing

import (
	"encoding/json"
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
)

func TestGameInteraction(t *testing.T) {
	port, svcs := startTestServer()

	lobby := svcs.Lobby.CreateLobby()
	lobby.Name = "Testing"
	lobby.Players = []string{
		"Player1",
		"Player2",
		"Player3",
	}

	game := svcs.Game.CreateGame(*lobby)
	ID := game.ID

	resp, err := get(port, gameRoute+"/"+ID+"/turn", []byte{})
	if err != nil {
		t.Fatal(err)
	}

	currentTurn := pogo.CurrentTurnMsg{}
	err = json.Unmarshal(resp, &currentTurn)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	player1Turn := pogo.TurnSubmissionMsg{
		GameID:   ID,
		PlayerID: "Player1",
	}

	// TODO: Actually play legal cards and test accordingly

	t.Logf("Sending player turn: \n%+v", player1Turn)
	resp, err = put(port,
		gameRoute+"/"+ID+"/turn/Player1",
		player1Turn,
	)
	if !assert.Nil(t, err) {
		t.Log(string(resp))
		t.Fatal(err)
	}

	resp, err = put(port,
		gameRoute+"/"+ID+"/turn/Player1",
		player1Turn,
	)
	if !assert.NotNil(t, err) {
		t.Fatal("A second submission was accepted")
	}
}
