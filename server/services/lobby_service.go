package services

import uuid "github.com/satori/go.uuid"

// LobbyService is responsible for managing our lobby list
type LobbyService interface {
	CreateLobby() *Lobby
	GetLobbies() []*Lobby
}

// Lobby is a group of players waiting to start a game
type Lobby struct {
	Name    string
	ID      uuid.UUID
	Players []string
}

// LobbyServiceImpl is a concrete service
type LobbyServiceImpl struct {
	activeLobbies []*Lobby
}

// CreateLobby creates a new lobby and returns it
func (l *LobbyServiceImpl) CreateLobby() *Lobby {
	newLobby := &Lobby{
		ID: uuid.NewV4(),
	}
	l.activeLobbies = append(l.activeLobbies, newLobby)
	return newLobby
}

// GetLobbies returns a list of lobbies currently active
func (l *LobbyServiceImpl) GetLobbies() []*Lobby {
	return l.activeLobbies
}
