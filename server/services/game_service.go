package services

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// GameService is responsible for managing our lobby list
type GameService interface {
	CreateGame() *Game
	GetActiveGame() []*Game
}

// Game is a group of players in a game
type Game struct {
	Name    string
	ID      uuid.UUID
	Players []string
}

// GameServiceImpl is a concrete service
type GameServiceImpl struct {
	activeGames []*Game
}

// CreateGame creates a new Game and returns it
func (g *GameServiceImpl) CreateGame() *Game {
	newGame := &Game{
		ID: uuid.NewV4(),
	}
	g.activeGames = append(g.activeGames, newGame)
	return newGame
}

// GetGame returns a the active Game for a player
func (g *GameServiceImpl) GetActiveGame(playerName string) (*Game, error) {
	for _, game := range g.activeGames {
		for _, player := range game.Players {
			if player == playerName {
				return game, nil
			}
		}
	}
	return nil, fmt.Errorf("WE FUCKED, please reconnect.")
}
