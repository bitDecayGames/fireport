package routing

import (
	"encoding/json"
	"testing"
	"time"

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
	println("") // TODO: MW if I comment out this line, the test fails to run... wtf?
}

func TestGameAnimations(t *testing.T) {
	port, svcs := startTestServer()

	lobby := svcs.Lobby.CreateLobby()
	lobby.Name = "Testing"
	lobby.Players = []string{
		"Player1",
		"Player2",
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
	t.Log(currentTurn)

	for i, p := range game.State.Players {

		var turnSubmission = pogo.TurnSubmissionMsg{
			GameID:   ID,
			PlayerID: p.Name,
			Inputs: []pogo.GameInputMsg{
				{
					CardID: p.Hand[0].ID,
					Owner:  p.ID,
					Order:  0,
				},
				{
					CardID: p.Hand[1].ID,
					Owner:  p.ID,
					Order:  1,
				},
			},
		}

		t.Logf("Sending player turn: \n%+v", turnSubmission)
		resp, err = put(port,
			gameRoute+"/"+ID+"/turn/"+p.Name,
			turnSubmission,
		)
		t.Log(string(resp))
		if !assert.Nil(t, err) {
			t.Log(string(resp))
			t.Fatal(err)
		}

		if i+1 < len(game.State.Players) {
			time.Sleep(2 * time.Second)
		} else {
			t.Log(string(resp))
		}
	}

	resp, err = get(port, gameRoute+"/"+ID+"/turn/state", []byte{})
	if err != nil {
		t.Fatal(err)
	}

	currentState := pogo.CurrentStateMsg{}
	err = json.Unmarshal(resp, &currentState)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	t.Log(currentState)

	assert.NotNil(t, currentState.CurrentState.Animations)
	assert.NotNil(t, currentState.CurrentState.Animations[0])

	println("") // TODO: MW if I comment out this line, the test fails to run... wtf?
}
