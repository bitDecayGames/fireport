package services

// MasterList is a master struct to hold all of our services
type MasterList struct {
	Lobby LobbyService
	Game  GameService
}

// NewMasterList will return a new instance of all core services
func NewMasterList() *MasterList {
	return &MasterList{
		Lobby: NewLobbyService(),
		Game:  &GameServiceImpl{},
	}
}
