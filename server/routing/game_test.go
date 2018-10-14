package routing

import (
	"testing"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func LobbyCreationHelper(name string)*services.Lobby{
	return &services.Lobby{
		Name:    name,
		ID:      uuid.NewV4(),
		Players: []string{"Player1"+name, "Player2"+name, "Player3"+name},
	}
}
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

func TestCreateGame(t *testing.T) {	
	_, svcs := startTestServer()

	lobbyOne := LobbyCreationHelper("FirstLobby")

	game := svcs.Game.CreateGame(lobbyOne)

	assert.Equal(t, game.Name, lobbyOne.Name)
	assert.Equal(t, game.ID, lobbyOne.ID)
	assert.Equal(t, game.Players, lobbyOne.Players)
}
