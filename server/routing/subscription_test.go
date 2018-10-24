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
	r, err := http.Post(fmt.Sprintf("http://127.0.0.1:%v%v", port, LobbyRoute), "application/json", bytes.NewBuffer([]byte("{}")))
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)

	createLobbyBody, err := ioutil.ReadAll(r.Body)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	lobbyID := string(createLobbyBody)

	player1Name := "TestPlayer1"
	player2Name := "Player2"

	bodyBytes, err := joinLobby(port, lobbyID, player1Name)
	initialLobby, err := lobbyFromBytes(bodyBytes)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

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
		joinLobby(port, lobbyID, player2Name)
	}()

	_, message, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	updatedLobby, err := lobbyFromBytes(message)
	assert.Equal(t, lobbyID, updatedLobby.ID)
	assert.Len(t, updatedLobby.Players, 2)
	assert.Equal(t, updatedLobby.Players[0], player1Name)
	assert.Equal(t, updatedLobby.Players[1], player2Name)
}

func joinLobby(port int, lobbyID, playerName string) ([]byte, error) {
	// Join our lobby
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://127.0.0.1:%v%v/%v/join", port, LobbyRoute, lobbyID),
		bytes.NewBuffer([]byte(playerName)),
	)
	if err != nil {
		return nil, err
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if "200 OK" != r.Status {
		return nil, fmt.Errorf("got return code %v", r.Status)
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func lobbyFromBytes(bytes []byte) (*pogo.LobbyMsg, error) {
	initialLobby := &pogo.LobbyMsg{}
	err := json.Unmarshal(bytes, initialLobby)
	return initialLobby, err
}
