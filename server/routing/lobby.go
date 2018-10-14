package routing

import (
	"net/http"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/gorilla/mux"
)

const lobbyRoute = apiv1 + "/lobby"

// LobbyRoutes contains information about routes specific to lobby interaction
type LobbyRoutes struct {
	Service services.LobbyService
}

// AddRoutes will add all lobby routes to the given router
func (lr *LobbyRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(lobbyRoute, lr.lobbyCreateHandler).Methods("POST")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/join", lr.lobbyJoinHandler).Methods("PUT")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/start", lr.lobbyStartHandler).Methods("PUT")
}

func (lr *LobbyRoutes) lobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby := lr.Service.CreateLobby()
	w.Write([]byte(lobby.ID.String()))
}

func (lr *LobbyRoutes) lobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}

func (lr *LobbyRoutes) lobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}
