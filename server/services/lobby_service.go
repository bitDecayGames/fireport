package services

import (
	uuid "github.com/satori/go.uuid"
)

// LobbyService is responsible for managing our lobby list
type LobbyService interface {
	CreateLobby() *Lobby
	GetLobby(string) (*Lobby, bool)
	Close(string)
	GetLobbies() map[string]*Lobby
}

// PlayerConnection is a general connection that allows messages to be sent
type PlayerConnection interface {
	WriteJSON(interface{}) error
	Close() error
}

// Lobby is a group of players waiting to start a game
type Lobby struct {
	Name              string
	Id                uuid.UUID
	Players           []string
	ActiveConnections map[string]PlayerConnection
}

// LobbyServiceImpl is a concrete service
type LobbyServiceImpl struct {
	activeLobbies map[string]*Lobby
}

// NewLobbyService returns a new instance of the lobby service
func NewLobbyService() *LobbyServiceImpl {
	return &LobbyServiceImpl{
		activeLobbies: make(map[string]*Lobby),
	}
}

// CreateLobby creates a new lobby and returns it
func (l *LobbyServiceImpl) CreateLobby() *Lobby {
	newLobby := &Lobby{
		Id:                uuid.NewV4(),
		ActiveConnections: make(map[string]PlayerConnection),
	}

	l.activeLobbies[newLobby.Id.String()] = newLobby
	return newLobby
}

// GetLobby returns a map of lobbies currently active
func (l *LobbyServiceImpl) GetLobby(lobbyId string) (*Lobby, bool) {
	lobby, ok := l.activeLobbies[lobbyId]
	return lobby, ok
}

// GetLobbies returns a map of lobbies currently active
func (l *LobbyServiceImpl) GetLobbies() map[string]*Lobby {
	return l.activeLobbies
}

// Close closes the lobby with the provided Id. If no such lobby
// exists, this function does nothing
func (l *LobbyServiceImpl) Close(lobbyID string) {
	delete(l.activeLobbies, lobbyID)
}
