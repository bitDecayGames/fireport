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

	_, found := s.Services.Lobby.IfLobbyExists(lobbyID, func(l *services.Lobby) {
		playerFound := false
		for _, id := range l.Players {
			if id == playerID {
				playerFound = true
				break
			}
		}
		if !playerFound {
			http.Error(w, fmt.Sprintf("player '%v' not found in lobby", playerID), http.StatusNotFound)
			return
		}

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to upgrade connection to websocket: '%v'", err), http.StatusInternalServerError)
			return
		}

		l.ActiveConnections[playerID] = c
	})

	if !found {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
		return
	}
}
