package services

import (
	"fmt"
	"sync"

	"github.com/bitdecaygames/fireport/server/rules"
	"github.com/satori/go.uuid"
)

// GameService is responsible for managing our active games
type GameService interface {
	CreateGame(lobby Lobby) *Game
	GetActiveGame(gameID uuid.UUID) (Game, error)
}

// Game is a group of players in a game
type Game struct {
	Name              string
	ID                uuid.UUID
	Players           []string
	Rules             []rules.GameRule
	ActiveConnections map[string]PlayerConnection
	InputRules        []rules.InputRule
}

var gameMutex = &sync.Mutex{}

// GameServiceImpl is a concrete service
type GameServiceImpl struct {
	activeGames []*Game
}

// CreateGame creates a new Game from the lobby information and returns it
func (g *GameServiceImpl) CreateGame(lobby Lobby) *Game {
	gameMutex.Lock()
	defer gameMutex.Unlock()

	newGame := &Game{
		Name:              lobby.Name,
		ID:                lobby.ID,
		Players:           lobby.Players,
		Rules:             rules.DefaultGameRules,
		ActiveConnections: lobby.ActiveConnections,
		InputRules:        rules.DefaultInputRules,
	}
	g.activeGames = append(g.activeGames, newGame)
	return newGame
}

// GetActiveGame returns a copy of the active Game for an ID or an error if the game is not found
func (g *GameServiceImpl) GetActiveGame(gameID uuid.UUID) (Game, error) {
	gameMutex.Lock()
	defer gameMutex.Unlock()

	for _, game := range g.activeGames {
		if game.ID == gameID {
			return *game, nil
		}
	}
	return Game{}, fmt.Errorf("no game found with uuid '%v'", gameID.String())
}
