package services

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// LobbyCreationHelper creates a lobby with the passed in name and 3 players in the lobby's Players slice
func LobbyCreationHelper(name string) Lobby {
	return Lobby{
		Name:    name,
		ID:      uuid.NewV4().String(),
		Players: []string{"Player1" + name, "Player2" + name, "Player3" + name},
	}
}

// TestGetActiveGame built to test the GetActiveGame function from the game service
func TestGetActiveGame(t *testing.T) {
	gameSvc := NewGameService()

	lobbyOne := LobbyCreationHelper("lobbyOne")
	lobbyTwo := LobbyCreationHelper("lobbyTwo")
	lobbyThree := LobbyCreationHelper("lobbyThree")

	gameOne := gameSvc.CreateGame(lobbyOne)
	gameTwo := gameSvc.CreateGame(lobbyTwo)
	gameThree := gameSvc.CreateGame(lobbyThree)

	assert.Equal(t, gameOne.ID, lobbyOne.ID)
	assert.Equal(t, gameTwo.ID, lobbyTwo.ID)
	assert.Equal(t, gameThree.ID, lobbyThree.ID)
}

// TestCreateGame built to test the CreateGame function from the game service
func TestCreateGame(t *testing.T) {
	gameSvc := NewGameService()

	testLobby := LobbyCreationHelper("testLobby")

	game := gameSvc.CreateGame(testLobby)

	assert.Equal(t, game.Name, testLobby.Name)
	assert.Equal(t, game.ID, testLobby.ID)
	assert.Equal(t, game.Players, testLobby.Players)
}

// TestGameHangsAfterInput tests #65: The game sometimes gets stuck after a few turns
func TestGameHangsAfterInput(t *testing.T) {
	gameSvc := NewGameService()
	testLobby := LobbyCreationHelper("testLobby")
	game := gameSvc.CreateGame(testLobby)
	testGameState := `{"Turn":58,"Created":1542165699,"Updated":1542166192,"IDCounter":81,"Players":[{"ID":0,"Name":"Zogan","Hand":[{"ID":14,"CardType":120},{"ID":13,"CardType":120},{"ID":11,"CardType":105},{"ID":3,"CardType":101},{"ID":8,"CardType":103}],"Discard":[],"Deck":[{"ID":9,"CardType":104},{"ID":2,"CardType":100},{"ID":5,"CardType":102},{"ID":12,"CardType":105},{"ID":1,"CardType":100},{"ID":10,"CardType":104},{"ID":7,"CardType":103},{"ID":6,"CardType":102},{"ID":4,"CardType":101}],"Location":35,"Facing":2,"Health":-128},{"ID":15,"Name":"Mike","Hand":[{"ID":29,"CardType":120},{"ID":28,"CardType":120},{"ID":19,"CardType":101},{"ID":25,"CardType":104},{"ID":23,"CardType":103}],"Discard":[],"Deck":[{"ID":18,"CardType":101},{"ID":26,"CardType":105},{"ID":22,"CardType":103},{"ID":21,"CardType":102},{"ID":17,"CardType":100},{"ID":24,"CardType":104},{"ID":20,"CardType":102},{"ID":27,"CardType":105},{"ID":16,"CardType":100}],"Location":29,"Facing":2,"Health":-79},{"ID":30,"Name":"L","Hand":[{"ID":34,"CardType":101},{"ID":36,"CardType":102},{"ID":42,"CardType":105},{"ID":31,"CardType":100},{"ID":37,"CardType":103}],"Discard":[],"Deck":[{"ID":41,"CardType":105},{"ID":35,"CardType":102},{"ID":44,"CardType":120},{"ID":43,"CardType":120},{"ID":33,"CardType":101},{"ID":38,"CardType":103},{"ID":39,"CardType":104},{"ID":40,"CardType":104},{"ID":32,"CardType":100}],"Location":26,"Facing":0,"Health":-93}],"BoardWidth":6,"BoardHeight":6,"BoardSpaces":[{"ID":45,"SpaceType":0,"State":0},{"ID":46,"SpaceType":0,"State":0},{"ID":47,"SpaceType":0,"State":0},{"ID":48,"SpaceType":0,"State":0},{"ID":49,"SpaceType":0,"State":0},{"ID":50,"SpaceType":0,"State":0},{"ID":51,"SpaceType":0,"State":0},{"ID":52,"SpaceType":0,"State":0},{"ID":53,"SpaceType":0,"State":0},{"ID":54,"SpaceType":0,"State":0},{"ID":55,"SpaceType":0,"State":0},{"ID":56,"SpaceType":0,"State":0},{"ID":57,"SpaceType":0,"State":0},{"ID":58,"SpaceType":0,"State":0},{"ID":59,"SpaceType":0,"State":0},{"ID":60,"SpaceType":0,"State":0},{"ID":61,"SpaceType":0,"State":0},{"ID":62,"SpaceType":0,"State":0},{"ID":63,"SpaceType":0,"State":0},{"ID":64,"SpaceType":0,"State":0},{"ID":65,"SpaceType":0,"State":0},{"ID":66,"SpaceType":0,"State":0},{"ID":67,"SpaceType":0,"State":0},{"ID":68,"SpaceType":0,"State":0},{"ID":69,"SpaceType":0,"State":0},{"ID":70,"SpaceType":0,"State":0},{"ID":71,"SpaceType":0,"State":0},{"ID":72,"SpaceType":0,"State":0},{"ID":73,"SpaceType":0,"State":0},{"ID":74,"SpaceType":0,"State":0},{"ID":75,"SpaceType":0,"State":0},{"ID":76,"SpaceType":0,"State":0},{"ID":77,"SpaceType":0,"State":0},{"ID":78,"SpaceType":0,"State":0},{"ID":79,"SpaceType":0,"State":0},{"ID":80,"SpaceType":0,"State":0}],"IsGameFinished":true,"Winner":-1}`
	err := game.SetTestGameState(testGameState)
	assert.NoError(t, err)

	err = gameSvc.SubmitSimpleTestTurn(game.ID, game.State.Players[0].Name, game.State.Players[0].ID, []int{3, 11, 8})
	assert.NoError(t, err)

	err = gameSvc.SubmitSimpleTestTurn(game.ID, game.State.Players[1].Name, game.State.Players[1].ID, []int{19, 25, 23})
	assert.NoError(t, err)

	err = gameSvc.SubmitSimpleTestTurn(game.ID, game.State.Players[2].Name, game.State.Players[2].ID, []int{42, 37, 31})
	assert.NoError(t, err)
}
