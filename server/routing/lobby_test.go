package routing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/stretchr/testify/assert"
)

func TestLobbyAPI(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	svcs := &services.MasterList{
		Lobby: &services.LobbyServiceImpl{},
	}

	go serveInternal(listener, svcs)

	lobbies := svcs.Lobby.GetLobbies()
	assert.Len(t, lobbies, 0)

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
	assert.Equal(t, lobbies[0].ID.String(), string(body))
}
