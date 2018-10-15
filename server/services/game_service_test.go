package services

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// LobbyCreationHelper creates a lobby with the passed in name and 3 players in the lobby's Players slice
func LobbyCreationHelper(name string) *Lobby {
	return &Lobby{
		Name:    name,
		ID:      uuid.NewV4(),
		Players: []string{"Player1" + name, "Player2" + name, "Player3" + name},
	}
}

// TestGetActiveGame built to test the GetActiveGame function from the game service
func TestGetActiveGame(t *testing.T) {
	gameSvc := &GameServiceImpl{}

	lobbyOne := LobbyCreationHelper("lobbyOne")
	lobbyTwo := LobbyCreationHelper("lobbyTwo")
	lobbyThree := LobbyCreationHelper("lobbyThree")

	gameSvc.CreateGame(lobbyOne)
	gameSvc.CreateGame(lobbyTwo)
	gameSvc.CreateGame(lobbyThree)

	retrievedGame1, _ := gameSvc.GetActiveGame(lobbyOne.ID)
	retrievedGame2, _ := gameSvc.GetActiveGame(lobbyTwo.ID)
	retrievedGame3, _ := gameSvc.GetActiveGame(lobbyThree.ID)

	assert.Equal(t, retrievedGame1.ID, lobbyOne.ID)
	assert.Equal(t, retrievedGame2.ID, lobbyTwo.ID)
	assert.Equal(t, retrievedGame3.ID, lobbyThree.ID)
}

// TestCreateGame built to test the CreateGame function from the game service
func TestCreateGame(t *testing.T) {
	gameSvc := &GameServiceImpl{}

	testLobby := LobbyCreationHelper("testLobby")

	game := gameSvc.CreateGame(testLobby)

	assert.Equal(t, game.Name, testLobby.Name)
	assert.Equal(t, game.ID, testLobby.ID)
	assert.Equal(t, game.Players, testLobby.Players)
}
