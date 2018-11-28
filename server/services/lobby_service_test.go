package services

import (
	"github.com/bitdecaygames/fireport/server/pogo"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLobbyReadyStatus will test when to start the game based on player readiness
func TestLobbyReadyStatus(t *testing.T) {
	lobbySvc := NewLobbyService()
	lobby := lobbySvc.CreateLobby()

	res, found := lobbySvc.IsReady(lobby.ID)
	assert.True(t, found)
	assert.False(t, res)

	_, err := lobbySvc.JoinLobby(pogo.LobbyJoinMsg{LobbyID: lobby.ID, PlayerID: "Player1"})
	assert.NoError(t, err)

	res, found = lobbySvc.IsReady(lobby.ID)
	assert.True(t, found)
	assert.False(t, res)

	_, err = lobbySvc.ReadyPlayer(lobby.ID, pogo.PlayerReadyMsg{PlayerName: "Player1", Ready: true})
	assert.NoError(t, err)

	res, found = lobbySvc.IsReady(lobby.ID)
	assert.True(t, found)
	assert.False(t, res)

	_, err = lobbySvc.JoinLobby(pogo.LobbyJoinMsg{LobbyID: lobby.ID, PlayerID: "Player2"})
	assert.NoError(t, err)

	_, err = lobbySvc.JoinLobby(pogo.LobbyJoinMsg{LobbyID: lobby.ID, PlayerID: "Player3"})
	assert.NoError(t, err)

	res, found = lobbySvc.IsReady(lobby.ID)
	assert.True(t, found)
	assert.False(t, res)

	_, err = lobbySvc.ReadyPlayer(lobby.ID, pogo.PlayerReadyMsg{PlayerName: "Player2", Ready: true})
	assert.NoError(t, err)

	res, found = lobbySvc.IsReady(lobby.ID)
	assert.True(t, found)
	assert.False(t, res)

	_, err = lobbySvc.ReadyPlayer(lobby.ID, pogo.PlayerReadyMsg{PlayerName: "Player3", Ready: true})
	assert.NoError(t, err)

	res, found = lobbySvc.IsReady(lobby.ID)
	assert.True(t, found)
	assert.True(t, res)
}
