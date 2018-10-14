package routing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

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

	lobbyID := lobbies[0].ID.String()
	assert.Equal(t, lobbyID, string(body))
	assert.Len(t, lobbies[0].Players, 0)

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
	if !assert.Len(t, lobbies[0].Players, 1) {
		t.Fatal("no lobbies")
	}
	assert.Equal(t, lobbies[0].Players[0], "TestPlayer1")
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

	assert.Contains(t, string(body), "No lobby found with ID 'no-such-lobby'")
}
