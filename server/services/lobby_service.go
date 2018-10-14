package services

import uuid "github.com/satori/go.uuid"

// LobbyService is responsible for managing our lobby list
type LobbyService interface {
	CreateLobby() Lobby
	GetLobbies() []Lobby
}

// Lobby is a group of players waiting to start a game
type Lobby struct {
	Name    string
	ID      uuid.UUID
	Players []string
}
