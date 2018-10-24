package routing

import (
	"fmt"
	"net/http"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	pubsubRoute = APIv1 + "/pubsub"
)

var upgrader = websocket.Upgrader{}

// Subscriber handles all event publishing
type Subscriber struct {
	Services *services.MasterList
}

// AddRoutes will add all game routes to the given router
func (s *Subscriber) AddRoutes(r *mux.Router) {
	r.HandleFunc(pubsubRoute+"/{lobbyID}/{playerID}", s.subscribeHandler)
}

func (s *Subscriber) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]
	playerID := vars["playerID"]

	lobby, ok := s.Services.Lobby.GetLobby(lobbyID)
	if !ok {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
		return
	}

	ok = false
	for _, id := range lobby.Players {
		if id == playerID {
			ok = true
			break
		}
	}
	if !ok {
		http.Error(w, fmt.Sprintf("player '%v' not found in lobby", playerID), http.StatusNotFound)
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to upgrade connection to websocket: '%v'", err), http.StatusInternalServerError)
		return
	}

	lobby.ActiveConnections[playerID] = c
}
