package routing

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/stretchr/testify/assert"
)

type LobbyTestSvc struct {
	createCalls int
}

func (l *LobbyTestSvc) CreateLobby() services.Lobby {
	l.createCalls++
	return services.Lobby{}
}

func (l *LobbyTestSvc) GetLobbies() []services.Lobby {
	return nil
}

func TestLobbyAPI(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	testSvc := &LobbyTestSvc{}
	svcs := &services.MasterList{
		Lobby: testSvc,
	}

	go serveInternal(listener, svcs)

	r, err := http.Post(fmt.Sprintf("http://127.0.0.1:%v%v", port, lobbyRoute), "application/json", bytes.NewBuffer([]byte("{}")))
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)
	assert.Equal(t, 1, testSvc.createCalls)
}
