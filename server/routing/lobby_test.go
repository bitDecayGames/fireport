package routing

import (
	"fmt"
	"net/url"
	"testing"
	"encoding/json"
	"github.com/gorilla/websocket"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/bitdecaygames/fireport/server/services"
	"github.com/stretchr/testify/assert"
)

func TestLobbyAPI(t *testing.T) {
	TestPlayer1 := "TestPlayer1"
	TestPlayer2 := "TestPlayer2"

	port, svcs := startTestServer()

	lobbies := svcs.Lobby.GetLobbiesSnapshot()
	assert.Len(t, lobbies, 0)

	// Create our lobby
	body, err := post(port, LobbyRoute, []byte{})
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	lobbies = svcs.Lobby.GetLobbiesSnapshot()
	assert.Len(t, lobbies, 1)

	var lobbyID string
	var lobby *services.Lobby
	for id, l := range lobbies {
		lobbyID = id
		lobby = &l
		break
	}
	if lobby == nil {
		t.Fatal("no lobby found")
	}

	assert.Equal(t, lobbyID, string(body))
	assert.Len(t, lobby.Players, 0)

	// Join our lobby
	msg := pogo.LobbyJoinMsg{
		LobbyID:  lobbyID,
		PlayerID: TestPlayer1,
	}

	_, err = put(port, LobbyRoute+"/join", msg)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	msg = pogo.LobbyJoinMsg{
		LobbyID:  lobbyID,
		PlayerID: TestPlayer2,
	}

	_, err = put(port, LobbyRoute+"/join", msg)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	lobbies = svcs.Lobby.GetLobbiesSnapshot()
	if !assert.Len(t, lobbies[lobbyID].Players, 2) {
		t.Fatal("expected 2 players in game lobby")
	}
	assert.Equal(t, lobbies[lobbyID].Players[0], `TestPlayer1`)
	assert.Equal(t, lobbies[lobbyID].Players[1], `TestPlayer2`)

	// Ready player 1 in  our lobby
	readyMsg := pogo.PlayerReadyMsg{
		PlayerName: TestPlayer1,
		Ready:      true,
	}

	_, err = put(port, LobbyRoute+"/"+lobbyID+"/ready", readyMsg)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	// NotReady player 2 in  our lobby
	readyMsg = pogo.PlayerReadyMsg{
		PlayerName: TestPlayer2,
		Ready:      false,
	}

	_, err = put(port, LobbyRoute+"/"+lobbyID+"/ready", readyMsg)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	lobbies = svcs.Lobby.GetLobbiesSnapshot()
	if !assert.Len(t, lobbies[lobbyID].PlayerReady, 2) {
		t.Fatal("expected 2 players in with a ready status in game lobby")
	}
	assert.Equal(t, lobbies[lobbyID].PlayerReady[TestPlayer1], true)
	assert.Equal(t, lobbies[lobbyID].PlayerReady[TestPlayer2], false)

	path := fmt.Sprintf("%v/%v/%v", pubsubRoute, lobbyID, TestPlayer1)
	t.Logf("Path: %v", path)

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf(`localhost:%v`, port), Path: path}

	t.Logf("Connecting to: %v", u)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	go func() {
		//Starts Game
		_, err = put(port, LobbyRoute+"/"+lobbyID+"/start", nil)
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}
	}()

	_, message, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	newGameMsg := pogo.GameStartMsg{}
	err = json.Unmarshal(message, &newGameMsg)
	if err != nil {
		t.Fatal(err)
	}	
	assert.Equal(t, lobbyID, newGameMsg.GameID)
	assert.Len(t, newGameMsg.Players, 2)
	assert.Equal(t, newGameMsg.Players[0], TestPlayer1)
	assert.Equal(t, newGameMsg.Players[1], TestPlayer2)

	if !assert.NotNil(t,newGameMsg.GameState) {
		t.Fatal("expected initial game state to have something in it.")
	}

	p1Deck := newGameMsg.GameState.Players[0].Deck
	p2Deck := newGameMsg.GameState.Players[1].Deck
	decksEqual := true
	if len(p1Deck) != len(p2Deck) {
		decksEqual = false
	} else {
		for i, cardState := range p1Deck {
			if cardState.CardType != p2Deck[i].CardType {
				decksEqual = false
				break
			}
		}
	}
	
	if decksEqual {
		t.Fatal("players decks should not be the same")
	}

	lobbies = svcs.Lobby.GetLobbiesSnapshot()
	if !assert.Len(t, lobbies, 0) {
		t.Fatal("expected lobby to be closed after game starts")
	}
}

func TestBadLobbyRequest(t *testing.T) {
	port, _ := startTestServer()

	msg := pogo.LobbyJoinMsg{
		LobbyID:  "no-such-lobby",
		PlayerID: "TestPlayer1",
	}

	resp, err := put(port, LobbyRoute+"/join", msg)
	if !assert.Contains(t, err.Error(), "404 Not Found") {
		t.Fatal("Expected bad lobby join to fail")
	}
	assert.Contains(t, string(resp), "no lobby found with ID 'no-such-lobby'")
}
