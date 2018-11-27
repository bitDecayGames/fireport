package services

import (
	"fmt"
	"sync"

	"github.com/bitdecaygames/fireport/server/pogo"
)

// LobbyModFunc takes a lobby as an argument and is a thread-safe way of modifying a lobby
type LobbyModFunc = func(*Lobby)

// LobbyService is responsible for managing our lobby list
type LobbyService interface {
	CreateLobby() *Lobby
	JoinLobby(pogo.LobbyJoinMsg) (Lobby, error)
	ReadyPlayer(string, pogo.PlayerReadyMsg) (Lobby, error)
	RegisterConnection(string, string, PlayerConnection) error
	IsReady(string) (bool, bool)
	Close(string) (Lobby, bool)
	GetLobbiesSnapshot() map[string]Lobby
}

// PlayerConnection is a general connection that allows messages to be sent
type PlayerConnection interface {
	WriteJSON(pogo.Typer) error
	Close() error
}

// Lobby is a group of players waiting to start a game
type Lobby struct {
	Name              string
	ID                string
	Players           []string
	PlayerReady       map[string]bool
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
		//ID:                uuid.NewV4().String(), // TODO: MW uncomment this
		ID:                "GAME", // TODO: MW for dev play testing only
		PlayerReady:       make(map[string]bool),
		ActiveConnections: make(map[string]PlayerConnection),
	}

	l.activeLobbies[newLobby.ID] = newLobby
	fmt.Println("Lobby created: ", newLobby.ID)
	return newLobby
}

// JoinLobby will add the player to the lobby, if it exists, or an error
func (l *LobbyServiceImpl) JoinLobby(msg pogo.LobbyJoinMsg) (Lobby, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobby, ok := l.activeLobbies[msg.LobbyID]
	if !ok {
		return Lobby{}, fmt.Errorf("no lobby found with ID '%v'", msg.LobbyID)
	}

	lobby.Players = append(lobby.Players, msg.PlayerID)
	lobby.PlayerReady[msg.PlayerID] = false
	return *lobby, nil
}

// ReadyPlayer will set player's ready status to what ever they passed us(true or false), if they exist, or an error
func (l *LobbyServiceImpl) ReadyPlayer(lobbyID string, readyMsg pogo.PlayerReadyMsg) (Lobby, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobby, ok := l.activeLobbies[lobbyID]
	if !ok {
		return Lobby{}, fmt.Errorf("no lobby found with ID '%v'", lobbyID)
	}
	lobby.PlayerReady[readyMsg.PlayerName] = readyMsg.Ready

	return *lobby, nil
}

// IsReady checks the ready status of all players, if any are false, returns false
func (l *LobbyServiceImpl) IsReady(lobbyID string) (bool, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobby, ok := l.activeLobbies[lobbyID]
	if !ok {
		return false, false
	}
	if len(lobby.Players) < 2 {
		return false, true // games cannot be readied with less than 2 players
	}
	for _, playerName := range lobby.Players {
		if val, ok := lobby.PlayerReady[playerName]; ok {
			if !val {
				return false, true
			}
		}
	}
	return true, true
}

// RegisterConnection adds a connection the given lobby and returns true if it exists, or an error
// otherwise
func (l *LobbyServiceImpl) RegisterConnection(lobbyID string, playerID string, c PlayerConnection) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lobby, ok := l.activeLobbies[lobbyID]
	if !ok {
		return fmt.Errorf("no lobby found with ID '%v'", lobbyID)
	}

	for _, player := range lobby.Players {
		if playerID == player {
			lobby.ActiveConnections[playerID] = c
			return nil
		}
	}

	return fmt.Errorf("no player found with ID '%v'", playerID)
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
