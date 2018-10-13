package routes

import "github.com/gorilla/mux"

const (
	apiv1 = "/api/v1"
)

// RegisterAll will register all needed routes on the given router
func RegisterAll(r *mux.Router) {
	lobby := &LobbyRoutes{}
	lobby.AddRoutes(r)

	game := &GameRoutes{}
	game.AddRoutes(r)
}
