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
	Services *services.MasterList
}

// AddRoutes will add all lobby routes to the given router
func (lr *LobbyRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(lobbyRoute, lr.lobbyCreateHandler).Methods("POST")
	r.HandleFunc(lobbyRoute+"/{lobbyID}/join", lr.lobbyJoinHandler).Methods("PUT")
	r.HandleFunc(lobbyRoute+"/{lobbyID}/start", lr.lobbyStartHandler).Methods("PUT")
}

func (lr *LobbyRoutes) lobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby := lr.Services.Lobby.CreateLobby()
	w.Write([]byte(lobby.Id.String()))
}

func (lr *LobbyRoutes) lobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	playerName, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	lobby, ok := lr.Services.Lobby.GetLobbies()[lobbyID]
	if !ok {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
		return
	}

	lobby.Players = append(lobby.Players, string(playerName))
	return
}

func (lr *LobbyRoutes) lobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	lobby, found := lr.Services.Lobby.GetLobby(lobbyID)
	if !found {
		http.Error(w, fmt.Sprintf("no lobby found with Id '%v'", lobbyID), http.StatusNotFound)
	}

	lr.Services.Game.CreateGame(lobby)

	lr.Services.Lobby.Close(lobbyID)
	lr.Services.Game.CreateGame(lobby)
}
