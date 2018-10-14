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

func TestGameAPI(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	svcs := &services.MasterList{
		Game: &services.GameServiceImpl{},
	}

	go serveInternal(listener, svcs)

	games := svcs.Game.GetActiveGame()
	assert.Len(t, games, 0)

	r, err := http.Post(fmt.Sprintf("http://127.0.0.1:%v%v", port, gameRoute), "application/json", bytes.NewBuffer([]byte("{}")))
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "200 OK", r.Status)

	body, err := ioutil.ReadAll(r.Body)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	games = svcs.Game.GetActiveGame()
	assert.Len(t, games, 1)
	assert.Equal(t, games[0].ID.String(), string(body))
}
