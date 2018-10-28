package pogo

// LobbyCreateMsg contains information for creating a lobby
type LobbyCreateMsg struct {
}

// LobbyMsg contains a snapshot of a server lobby for a client
type LobbyMsg struct {
	ID          string          `json:"id"`
	Players     []string        `json:"players"`
	ReadyStatus map[string]bool `json:"readyStatus"`
}

type PlayerReadyMsg struct {
	PlayerName string `json:"playerName"`
	Ready      bool   `json:"ready"`
}
