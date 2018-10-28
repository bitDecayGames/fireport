package routing

import (
	"encoding/json"
	"strconv"
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
	ID := game.ID.String()

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
		GameID: ID,
	}

	// TODO: Actually play legal cards and test accordingly

	resp, err = put(port,
		gameRoute+"/"+ID+"/turn/"+strconv.Itoa(currentTurn.CurrentTurn)+"/player/Player1",
		player1Turn,
	)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
}
