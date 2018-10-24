package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/gorilla/mux"
)

// LobbyRoute is the api route for all lobby interactions
const LobbyRoute = APIv1 + "/lobby"

// LobbyRoutes contains information about routes specific to lobby interaction
type LobbyRoutes struct {
	Services *services.MasterList
}

// AddRoutes will add all lobby routes to the given router
func (lr *LobbyRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(LobbyRoute, lr.lobbyCreateHandler).Methods("POST")
	r.HandleFunc(LobbyRoute+"/{lobbyID}/join", lr.lobbyJoinHandler).Methods("PUT")
	r.HandleFunc(LobbyRoute+"/{lobbyID}/start", lr.lobbyStartHandler).Methods("PUT")
}

func (lr *LobbyRoutes) lobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby := lr.Services.Lobby.CreateLobby()
	w.Write([]byte(lobby.ID.String()))
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

	msg := pogo.LobbyMsg{
		ID:      lobby.ID.String(),
		Players: lobby.Players,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "failed to build lobby message", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(bytes))

	// tell all pubsubbers
	for id, conn := range lobby.ActiveConnections {
		err = conn.WriteJSON(msg)
		if err != nil {
			fmt.Printf("Failed to tell player %v about lobby update: %v\n", id, err)
		}
	}

	return
}

func (lr *LobbyRoutes) lobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	lobby, found := lr.Services.Lobby.GetLobby(lobbyID)
	if !found {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
	}

	lr.Services.Game.CreateGame(lobby)

	lr.Services.Lobby.Close(lobbyID)
	lr.Services.Game.CreateGame(lobby)
}
