package main

import (
	"net/http"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/bitdecaygames/fireport/server/routing"
)

const (
	// Networking
	port = 8080
)

func main() {
	svcs := &services.MasterList{}
	routing.ServeGame(port, svcs)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
