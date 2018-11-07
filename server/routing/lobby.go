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
	r.HandleFunc(LobbyRoute+"/join", lr.lobbyJoinHandler).Methods("PUT")
	r.HandleFunc(LobbyRoute+"/{lobbyID}/ready", lr.lobbyReadyHandler).Methods("PUT")
	r.HandleFunc(LobbyRoute+"/{lobbyID}/start", lr.lobbyStartGameHandler).Methods("PUT")
}

func (lr *LobbyRoutes) lobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby := lr.Services.Lobby.CreateLobby()
	w.Write([]byte(lobby.ID))
}

func (lr *LobbyRoutes) lobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	joinMsg := pogo.LobbyJoinMsg{}
	err = json.Unmarshal(body, &joinMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lobby, err := lr.Services.Lobby.JoinLobby(joinMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	msg := pogo.LobbyMsg{
		ID:          lobby.ID,
		Players:     lobby.Players,
		ReadyStatus: lobby.PlayerReady,
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

func (lr *LobbyRoutes) lobbyReadyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	requestBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	readyMsg := pogo.PlayerReadyMsg{}
	err = json.Unmarshal(requestBytes, &readyMsg)

	lobby, err := lr.Services.Lobby.ReadyPlayer(lobbyID, readyMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	msg := pogo.LobbyMsg{
		ID:          lobby.ID,
		Players:     lobby.Players,
		ReadyStatus: lobby.PlayerReady,
	}

	w.Write([]byte("ready status updated"))

	// tell all pubsubbers
	for id, conn := range lobby.ActiveConnections {
		err = conn.WriteJSON(msg)
		if err != nil {
			fmt.Printf("Failed to tell player %v about lobby update: %v\n", id, err)
		}
	}
	return
}

func (lr *LobbyRoutes) lobbyStartGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	lobby, found := lr.Services.Lobby.Close(lobbyID)
	if !found {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
	}

	gameInstance := lr.Services.Game.CreateGame(lobby)

	msg := pogo.GameStartMsg{
		GameID:  gameInstance.ID,
		Players: gameInstance.Players,
		GameState: gameInstance.State,
	}

	// tell all pubsubbers
	for id, conn := range lobby.ActiveConnections {
		err := conn.WriteJSON(msg)
		if err != nil {
			fmt.Printf("Failed to tell player %v about the started game: %v\n", id, err)
		}
	}

	w.Write([]byte("The game has started."))

	return
}
