package routing

import (
	"net"

	"github.com/bitdecaygames/fireport/server/services"
)

func startTestServer() (int, *services.MasterList) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	svcs := &services.MasterList{
		Lobby: &services.LobbyServiceImpl{},
	}

	go serveInternal(listener, svcs)
	return port, svcs
}
