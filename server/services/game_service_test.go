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
		Id:      uuid.NewV4(),
		Players: []string{"Player1" + name, "Player2" + name, "Player3" + name},
	}
}

// TestGetActiveGame built to test the GetActiveGame function from the game
func TestGetActiveGame(t *testing.T) {
	gameSvc := &GameServiceImpl{}

	lobbyOne := LobbyCreationHelper("FirstLobby")
	lobbyTwo := LobbyCreationHelper("SecondLobby")
	lobbyThree := LobbyCreationHelper("ThirdLobby")

	gameSvc.CreateGame(lobbyOne)
	gameSvc.CreateGame(lobbyTwo)
	gameSvc.CreateGame(lobbyThree)

	retrievedGame1, _ := gameSvc.GetActiveGame(lobbyOne.Id)
	retrievedGame2, _ := gameSvc.GetActiveGame(lobbyTwo.Id)
	retrievedGame3, _ := gameSvc.GetActiveGame(lobbyThree.Id)

	assert.Equal(t, retrievedGame1.ID, lobbyOne.Id)
	assert.Equal(t, retrievedGame2.ID, lobbyTwo.Id)
	assert.Equal(t, retrievedGame3.ID, lobbyThree.Id)
}

// TestCreateGame built to test the CreateGame function from the game
func TestCreateGame(t *testing.T) {
	gameSvc := &GameServiceImpl{}

	testLobby := LobbyCreationHelper("testLobby")

	game := gameSvc.CreateGame(testLobby)

	assert.Equal(t, game.Name, testLobby.Name)
	assert.Equal(t, game.ID, testLobby.Id)
	assert.Equal(t, game.Players, testLobby.Players)
}
