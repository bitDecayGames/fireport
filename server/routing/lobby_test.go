package routing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/stretchr/testify/assert"
)

func TestLobbyAPI(t *testing.T) {
	port, svcs := startTestServer()

	lobbies := svcs.Lobby.GetLobbiesSnapshot()
	assert.Len(t, lobbies, 0)

	// Create our lobby
	r, err := http.Post(fmt.Sprintf("http://127.0.0.1:%v%v", port, LobbyRoute), "application/json", bytes.NewBuffer([]byte("{}")))
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)

	body, err := ioutil.ReadAll(r.Body)
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
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/join", port, LobbyRoute, lobbyID),
		bytes.NewBuffer([]byte("TestPlayer1")),
	)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	r, err = http.DefaultClient.Do(req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)

	lobbies = svcs.Lobby.GetLobbiesSnapshot()
	if !assert.Len(t, lobbies[lobbyID].Players, 1) {
		t.Fatal("expected 1 player in game lobby")
	}
	assert.Equal(t, lobbies[lobbyID].Players[0], "TestPlayer1")

	// Create game from our lobby
	req, err = http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/start", port, LobbyRoute, lobbyID),
		bytes.NewBuffer([]byte("{}")),
	)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	r, err = http.DefaultClient.Do(req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	if !assert.Equal(t, "200 OK", r.Status) {
		body, _ := ioutil.ReadAll(r.Body)
		t.Logf("Expected lobby: %v", lobbyID)
		t.Fatalf("Body from error message: %v", string(body))
	}

	lobbies = svcs.Lobby.GetLobbiesSnapshot()
	if !assert.Len(t, lobbies, 0) {
		t.Fatal("expected lobby to be closed after game starts")
	}
}

func TestBadLobbyRequest(t *testing.T) {
	port, _ := startTestServer()

	// Join our lobby
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/join", port, LobbyRoute, "no-such-lobby"),
		bytes.NewBuffer([]byte("TestPlayer1")),
	)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "404 Not Found", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	assert.Contains(t, string(body), "no lobby found with ID 'no-such-lobby'")
}
