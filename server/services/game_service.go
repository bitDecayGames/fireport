package services

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// GameService is responsible for managing our lobby list
type GameService interface {
	CreateGame(lobby *Lobby) *Game
	GetActiveGame(gameID uuid.UUID) (*Game, error)
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

// GetActiveGame returns a the active Game for an ID or an error if the game is not found
func (g *GameServiceImpl) GetActiveGame(gameID uuid.UUID) (*Game, error) {
	for _, game := range g.activeGames {
		if game.ID == gameID {
			return game, nil
		}
	}
	return nil, fmt.Errorf("no game found with uuid '%v'", gameID.String())
}
