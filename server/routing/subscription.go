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

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to upgrade connection to websocket: '%v'", err), http.StatusInternalServerError)
		return
	}

	err = s.Services.Lobby.RegisterConnection(lobbyID, playerID, &PlayerConnWrapper{con: c})

	if err != nil {
		c.Close()
		http.Error(w, fmt.Sprintf("unable to register connection '%v'", err), http.StatusNotFound)
		return
	}
}
