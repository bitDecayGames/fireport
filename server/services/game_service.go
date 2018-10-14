package services

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// GameService is responsible for managing our lobby list
type GameService interface {
	CreateGame(lobby *Lobby) *Game
	GetActiveGame(gameId uuid.UUID) (*Game, error)
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

// CreateGame creates a new Game from the lobby information and returns it
func (g *GameServiceImpl) CreateGame(lobby *Lobby) *Game {
	newGame := &Game{
		Name:    lobby.Name,
		ID:      lobby.ID,
		Players: lobby.Players,
	}
	g.activeGames = append(g.activeGames, newGame)
	return newGame
}

// GetGame returns a the active Game for an ID
func (g *GameServiceImpl) GetActiveGame(gameId uuid.UUID) (*Game, error) {
	for _, game := range g.activeGames {
		if game.ID == gameId {
			return game, nil
		}
	}
	return nil, fmt.Errorf("Please reconnect, game was dropped or ended.")
}
