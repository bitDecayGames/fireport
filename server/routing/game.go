package routing

import (
	"encoding/json"
	"net/http"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/gorilla/mux"
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
	r.HandleFunc(gameRoute+"/{gameID}/turn", gr.getCurrentTurnHandler).Methods("GET")
	r.HandleFunc(gameRoute+"/{gameID}/turn/{playerName}", gr.submitCardsHandler).Methods("PUT")
	r.HandleFunc(gameRoute+"/{gameID}/turn/{playerName}", gr.getGameStateHandler).Methods("GET")
}

// submitCardsHandler handles a player request to submit cards for the given turn
func (gr *GameRoutes) submitCardsHandler(w http.ResponseWriter, r *http.Request) {
	turnSubMsg := &pogo.TurnSubmissionMsg{}
	err := Unmarshal(r.Body, turnSubMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = gr.Services.Game.SubmitTurn(*turnSubMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getCurrentTurnHandler returns the current turn for the requested game
func (gr *GameRoutes) getCurrentTurnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	//submissionTurn := vars["turn"]

	turn, err := gr.Services.Game.GetCurrentTurn(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	msg := pogo.CurrentTurnMsg{
		GameID:      gameID,
		CurrentTurn: turn,
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
