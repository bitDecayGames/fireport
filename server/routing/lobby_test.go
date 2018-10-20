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

	lobbies := svcs.Lobby.GetLobbies()
	assert.Len(t, lobbies, 0)

	// Create our lobby
	r, err := http.Post(fmt.Sprintf("http://127.0.0.1:%v%v", port, lobbyRoute), "application/json", bytes.NewBuffer([]byte("{}")))
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)

	body, err := ioutil.ReadAll(r.Body)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	lobbies = svcs.Lobby.GetLobbies()
	assert.Len(t, lobbies, 1)

	var lobbyID string
	var lobby *services.Lobby
	for id, l := range lobbies {
		lobbyID = id
		lobby = l
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
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/join", port, lobbyRoute, lobbyID),
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

	lobbies = svcs.Lobby.GetLobbies()
	if !assert.Len(t, lobbies[lobbyID].Players, 1) {
		t.Fatal("no lobbies")
	}
	assert.Equal(t, lobbies[lobbyID].Players[0], "TestPlayer1")

	// Create game from our lobby
	req, err = http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/start", port, lobbyRoute, lobbyID),
		bytes.NewBuffer([]byte("{}")),
	)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	r, err = http.DefaultClient.Do(req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)
	lobbies = svcs.Lobby.GetLobbies()
	if !assert.Len(t, lobbies, 0) {
		t.Fatal("expected lobby to be closed after game starts")
	}
}

func TestBadLobbyRequest(t *testing.T) {
	port, _ := startTestServer()

	// Join our lobby
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/join", port, lobbyRoute, "no-such-lobby"),
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

	assert.Contains(t, string(body), "no lobby found with Id 'no-such-lobby'")
}
