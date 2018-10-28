package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

const (
	gameRoute = APIv1 + "/game"
)

// GameRoutes contains information about routes specific to game interaction
type GameRoutes struct {
	Services *services.MasterList
}

// AddRoutes will add all game routes to the given router
func (gr *GameRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(gameRoute+"/{gameID}/turn/{turn}/player/{playerName}", gr.submitCardsHandler).Methods("PUT")
	r.HandleFunc(gameRoute+"/{gameID}/turn", gr.getCurrentTurnHandler).Methods("GET")
	r.HandleFunc(gameRoute+"/{gameID}/turn/{turn}/player/{playerName}", gr.getGameStateHandler).Methods("GET")
}

// submitCardsHandler handles a player request to submit cards for the given turn
func (gr *GameRoutes) submitCardsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success, though nothing has been recorded"))
}

// getCurrentTurnHandler returns the current turn for the requested game
func (gr *GameRoutes) getCurrentTurnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	//submissionTurn := vars["turn"]

	gameUUID, err := uuid.FromString(gameID)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v is not a valid game ID", gameID), http.StatusBadRequest)
		return
	}

	game, err := gr.Services.Game.GetActiveGame(gameUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to find game with id '%v'", gameID), http.StatusNotFound)
		return
	}

	msg := pogo.CurrentTurnMsg{
		GameID:      gameUUID.String(),
		CurrentTurn: game.CurrentTurn,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusNotFound)
		return
	}

	w.Write(bytes)
}

// getGameStateHandler returns the current state of the game for the requested turn
func (gr *GameRoutes) getGameStateHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}
