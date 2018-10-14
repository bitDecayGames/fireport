package routing

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// LobbyCreationHelper creates a lobby with the passed in name and 3 players in the lobby's Players slice
func LobbyCreationHelper(name string)*services.Lobby{
	return &services.Lobby{
		Name:    name,
		ID:      uuid.NewV4(),
		Players: []string{"Player1"+name, "Player2"+name, "Player3"+name},
	}
}

// TestGetActiveGame built to test the GetActiveGame function from the game
func TestGetActiveGame(t *testing.T) {	
	_, svcs := startTestServer()

	lobbyOne := LobbyCreationHelper("FirstLobby")
	lobbyTwo := LobbyCreationHelper("SecondLobby")
	lobbyThree := LobbyCreationHelper("ThirdLobby")

	svcs.Game.CreateGame(lobbyOne)
	svcs.Game.CreateGame(lobbyTwo)
	svcs.Game.CreateGame(lobbyThree)

	retrievedGame1,_:= svcs.Game.GetActiveGame(lobbyOne.ID)
	retrievedGame2,_:= svcs.Game.GetActiveGame(lobbyTwo.ID)
	retrievedGame3,_:= svcs.Game.GetActiveGame(lobbyThree.ID)

	assert.Equal(t, retrievedGame1.ID, lobbyOne.ID)
	assert.Equal(t, retrievedGame2.ID, lobbyTwo.ID)
	assert.Equal(t, retrievedGame3.ID, lobbyThree.ID)
}

// TestCreateGame built to test the CreateGame function from the game
func TestCreateGame(t *testing.T) {	
	_, svcs := startTestServer()

	testLobby := LobbyCreationHelper("testLobby")

	game := svcs.Game.CreateGame(testLobby)

	assert.Equal(t, game.Name, testLobby.Name)
	assert.Equal(t, game.ID, testLobby.ID)
	assert.Equal(t, game.Players, testLobby.Players)
}
