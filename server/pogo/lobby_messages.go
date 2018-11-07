package pogo

// LobbyCreateMsg contains information for creating a lobby
type LobbyCreateMsg struct {
	TypedMsg
}

// LobbyJoinMsg is a client message for a player to join a lobby
type LobbyJoinMsg struct {
	TypedMsg

	LobbyID  string `json:"lobbyID"`
	PlayerID string `json:"playerID"`
}

// LobbyMsg contains a snapshot of a server lobby for a client
type LobbyMsg struct {
	TypedMsg

	ID          string          `json:"id"`
	Players     []string        `json:"players"`
	ReadyStatus map[string]bool `json:"readyStatus"`
}

// PlayerReadyMsg contains a ready(true or false) message from the player
type PlayerReadyMsg struct {
	TypedMsg

	PlayerName string `json:"playerName"`
	Ready      bool   `json:"ready"`
}
