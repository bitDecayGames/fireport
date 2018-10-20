package pogo

// LobbyCreateMsg contains information for creating a lobby
type LobbyCreateMsg struct {
}

// LobbyUpdateMsg contains any important information for a client to update the
// lobby
type LobbyUpdateMsg struct {
	Players []string
}
