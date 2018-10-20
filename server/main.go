package main

import (
	"github.com/bitdecaygames/fireport/server/services"

	"github.com/bitdecaygames/fireport/server/routing"
)

const (
	// Networking
	port = 8080
)

func main() {
	svcs := services.NewMasterList()
	routing.ServeGame(port, svcs)
}
