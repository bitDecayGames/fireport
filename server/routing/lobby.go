package routing

import (
	"fmt"
	"io/ioutil"
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
	r.HandleFunc(lobbyRoute+"/{lobbyID}/join", lr.lobbyJoinHandler).Methods("PUT")
	r.HandleFunc(lobbyRoute+"/{lobbyID}/start", lr.lobbyStartHandler).Methods("PUT")
}

func (lr *LobbyRoutes) lobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby := lr.Service.CreateLobby()
	w.Write([]byte(lobby.ID.String()))
}

func (lr *LobbyRoutes) lobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	playerName, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	for _, lobby := range lr.Service.GetLobbies() {
		if lobby.ID.String() == lobbyID {
			lobby.Players = append(lobby.Players, string(playerName))
			return
		}
	}

	http.Error(w, fmt.Sprintf("No lobby found with ID '%v'", lobbyID), http.StatusNotFound)
}

func (lr *LobbyRoutes) lobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}
