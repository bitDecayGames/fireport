package main

import (
	"fmt"
	"net/http"

	"github.com/bitdecaygames/fireport/server/routing"
)

const (
	// Networking
	port = 8080
)

func main() {
	bind := fmt.Sprintf(":%v", port)

	routing.ServeGame(bind)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
