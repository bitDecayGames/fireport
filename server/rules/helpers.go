package rules

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

func doesPlayerHaveCardID(p *pogo.PlayerState, id int) bool {
	var allPlayerCards = collectAllPlayersCards(p)
	fmt.Printf("all cards: %v\n", allPlayerCards)
	for _, card := range allPlayerCards {
		if card.ID == id {
			return true
		}
	}
	return false
}

func collectAllPlayersCards(p *pogo.PlayerState) []pogo.CardState {
	return append(append(p.Hand, p.Deck...), p.Discard...)
}

func collectAllGameStateIds(g *pogo.GameState) []idTracker {
	var ids []idTracker

	for _, player := range g.Players {
		ids = append(ids, idTracker{ID: player.ID, Name: "player"})

		for _, card := range player.Hand {
			ids = append(ids, idTracker{ID: card.ID, Name: fmt.Sprintf("player %v hand", player.ID)})
		}
	}

	for _, space := range g.BoardSpaces {
		ids = append(ids, idTracker{ID: space.ID, Name: "board space"})
	}

	return ids
}

type idTracker struct {
	ID   int
	Name string
}
