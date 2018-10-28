package pogo

// LobbyCreateMsg contains information for creating a lobby
type LobbyCreateMsg struct {
}

// LobbyMsg contains a snapshot of a server lobby for a client
type LobbyMsg struct {
	ID      string   `json:"id"`
	Players []string `json:"players"`
}

type ReadyMsg struct {
	ID          string   `json:"id"`
	Players     []string `json:"players"`
	ReadyStatus string   `json:"readyStatus"`
}
