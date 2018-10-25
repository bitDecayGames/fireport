package services

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

// LobbyModFunc takes a lobby as an argument and is a thread-safe way of modifying a lobby
type LobbyModFunc = func(*Lobby)

// LobbyService is responsible for managing our lobby list
type LobbyService interface {
	CreateLobby() *Lobby
	IfLobbyExists(string, LobbyModFunc) (Lobby, bool)
	Close(string) (Lobby, bool)
	GetLobbiesSnapshot() map[string]Lobby
}

// PlayerConnection is a general connection that allows messages to be sent
type PlayerConnection interface {
	WriteJSON(interface{}) error
	Close() error
}

// Lobby is a group of players waiting to start a game
type Lobby struct {
	Name              string
	ID                uuid.UUID
	Players           []string
	ActiveConnections map[string]PlayerConnection
}

var lobbyMutex = &sync.Mutex{}

// LobbyServiceImpl is a concrete service
type LobbyServiceImpl struct {
	activeLobbies map[string]*Lobby
	mutex         *sync.Mutex
}

// NewLobbyService returns a new instance of the lobby service
func NewLobbyService() *LobbyServiceImpl {
	return &LobbyServiceImpl{
		activeLobbies: make(map[string]*Lobby),
		mutex:         &sync.Mutex{},
	}
}

// CreateLobby creates a new lobby and returns it
func (l *LobbyServiceImpl) CreateLobby() *Lobby {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newLobby := &Lobby{
		ID:                uuid.NewV4(),
		ActiveConnections: make(map[string]PlayerConnection),
	}

	l.activeLobbies[newLobby.ID.String()] = newLobby
	return newLobby
}

// IfLobbyExists calls the given function if the lobby is found, returning a copy of the
// modified lobby if it was found. A bool will be returned indicating if the lobby was found
func (l *LobbyServiceImpl) IfLobbyExists(lobbyID string, f LobbyModFunc) (Lobby, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobby, ok := l.activeLobbies[lobbyID]
	if ok {
		f(lobby)
		lobbyCopy := *lobby
		return lobbyCopy, true
	}

	return Lobby{}, false
}

// GetLobbiesSnapshot returns a copy of the map of currently active lobbies
func (l *LobbyServiceImpl) GetLobbiesSnapshot() map[string]Lobby {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobbiesCopy := make(map[string]Lobby)

	for id, l := range l.activeLobbies {
		lobbiesCopy[id] = *l
	}

	return lobbiesCopy
}

// Close closes the lobby with the provided ID and returns it.
// If no lobby exists, this function does nothing and returns false
func (l *LobbyServiceImpl) Close(lobbyID string) (Lobby, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobby, found := l.activeLobbies[lobbyID]
	if !found {
		return Lobby{}, false
	}

	delete(l.activeLobbies, lobbyID)
	return *lobby, true
}
