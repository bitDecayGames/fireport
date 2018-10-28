package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestLobbyPubSubHappyPath(t *testing.T) {
	port, _ := startTestServer()

	// Create our lobby
	resp, err := post(port, LobbyRoute, "")
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	lobbyID := string(resp)

	player1Name := "TestPlayer1"
	player2Name := "Player2"

	msg := pogo.LobbyJoinMsg{
		LobbyID:  lobbyID,
		PlayerID: player1Name,
	}

	resp, err = put(port, LobbyRoute+"/join", msg)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	initialLobby := pogo.LobbyMsg{}
	err = json.Unmarshal(resp, &initialLobby)

	assert.Equal(t, lobbyID, initialLobby.ID)
	assert.Len(t, initialLobby.Players, 1)
	assert.Equal(t, initialLobby.Players[0], player1Name)

	path := fmt.Sprintf("%v/%v/%v", pubsubRoute, lobbyID, player1Name)
	t.Logf("Path: %v", path)

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf(`localhost:%v`, port), Path: path}

	t.Logf("Connecting to: %v", u)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	go func() {
		msg := pogo.LobbyJoinMsg{
			LobbyID:  lobbyID,
			PlayerID: player2Name,
		}

		_, err := put(port, LobbyRoute+"/join", msg)
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}
	}()

	_, message, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	updatedLobby := pogo.LobbyMsg{}
	err = json.Unmarshal(message, &updatedLobby)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, lobbyID, updatedLobby.ID)
	assert.Len(t, updatedLobby.Players, 2)
	assert.Equal(t, updatedLobby.Players[0], player1Name)
	assert.Equal(t, updatedLobby.Players[1], player2Name)
}
