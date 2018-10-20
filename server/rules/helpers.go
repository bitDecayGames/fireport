package rules

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
)

func doesPlayerHaveCardId(p *pogo.PlayerState, id int) bool {
	var allPlayerCards = collectAllPlayersCards(p)
	for _, card := range allPlayerCards {
		if card.Id == id {
			return true
		}
	}
	return false
}

func collectAllPlayersCards(p *pogo.PlayerState) []pogo.CardState {
	return append(append(p.Hand, p.Deck...), p.Discard...)
}

func collectAllGameStateIds(g *pogo.GameState) []idTracker {
	var ids []idTracker = nil

	for _, player := range g.Players {
		ids = append(ids, idTracker{Id:player.Id, Name:"player"})

		for _, card := range player.Hand {
			ids = append(ids, idTracker{Id:card.Id, Name:fmt.Sprintf("player %v hand", player.Id)})
		}
	}

	for _, space := range g.BoardSpaces {
		ids = append(ids, idTracker{Id:space.Id, Name:"board space"})
	}

	return ids
}

type idTracker struct {
	Id int
	Name string
}