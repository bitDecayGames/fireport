package routing

import (
	"net/http"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/gorilla/mux"
)

const (
	gameRoute = APIv1 + "/game"
)

// GameRoutes contains information about routes specific to game interaction
type GameRoutes struct {
	Service services.GameService
}

// AddRoutes will add all game routes to the given router
func (gr *GameRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(gameRoute+"/{gameName}/tick/{tick}/player/{playerName}/cards", gr.submitCardsHandler).Methods("PUT")
	r.HandleFunc(gameRoute+"/{gameName}/tick", gr.getCurrentTickHandler).Methods("GET")
	r.HandleFunc(gameRoute+"/{gameName}/tick/{tick}/player/{playerName}", gr.getGameStateHandler).Methods("GET")
}

// submitCardsHandler handles a player request to submit cards for the given tick
func (gr *GameRoutes) submitCardsHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}

// getCurrentTickHandler returns the current tick for the requested game
func (gr *GameRoutes) getCurrentTickHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}

// getGameStateHandler returns the current state of the game for the requested tick
func (gr *GameRoutes) getGameStateHandler(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}
