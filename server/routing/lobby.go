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
	r.HandleFunc(LobbyRoute+"/{lobbyID}/leave", lr.lobbyLeaveHandler).Methods("PUT")
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

	msg := &pogo.LobbyMsg{
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

	PublishMessage(msg, lobby)
	return
}

func (lr *LobbyRoutes) lobbyLeaveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	leaveMsg := pogo.LobbyLeaveMsg{}
	err = json.Unmarshal(body, &leaveMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lobby, err := lr.Services.Lobby.LeaveLobby(lobbyID, leaveMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	msg := &pogo.LobbyMsg{
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

	if len(lobby.Players) == 0 {
		// close empty lobbies
		lr.Services.Lobby.Close(lobbyID)
		return
	}

	PublishMessage(msg, lobby)
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
	msg := &pogo.LobbyMsg{
		ID:          lobby.ID,
		Players:     lobby.Players,
		ReadyStatus: lobby.PlayerReady,
	}

	w.Write([]byte("ready status updated"))

	PublishMessage(msg, lobby)
	return
}

func (lr *LobbyRoutes) lobbyStartGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyID := vars["lobbyID"]

	isReady, found := lr.Services.Lobby.IsReady(lobbyID)
	if !found {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
	}
	if !isReady {
		http.Error(w, fmt.Sprint("all players must be ready before you may start the game"), http.StatusExpectationFailed)
	}

	lobby, found := lr.Services.Lobby.Close(lobbyID)
	if !found {
		http.Error(w, fmt.Sprintf("no lobby found with ID '%v'", lobbyID), http.StatusNotFound)
	}

	gameInstance := lr.Services.Game.CreateGame(lobby)

	msg := &pogo.GameStartMsg{
		GameID:    gameInstance.ID,
		Players:   gameInstance.Players,
		GameState: gameInstance.State,
	}

	w.Write([]byte("The game has started."))

	PublishMessage(msg, lobby)
	return
}

func PublishMessage(msg pogo.Typer, lobby services.Lobby) {
	var err error
	for id, conn := range lobby.ActiveConnections {
		err = conn.WriteJSON(msg)
		if err != nil {
			fmt.Printf("Failed to tell player %v about lobby update: %v\n", id, err)
		}
	}
}
