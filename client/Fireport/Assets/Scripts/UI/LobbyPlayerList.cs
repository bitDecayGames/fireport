using System.Collections.Generic;
using Model.Message;
using Network;
using UnityEngine;
using Utils;

public class LobbyPlayerList : MonoBehaviour, IDownStreamSubscriber
{
	public PlayerRowController playerRowPrefab;
	private GoToScene Goto;

	private void Start()
	{
		WebSocketListener.Instance().Subscribe(this);
		LobbyInfoController lobbyInfo = LobbyInfoController.Instance();
		updatePlayers(lobbyInfo.msg.players.ToArray(), lobbyInfo.msg.readyStatus);
		WebSocketListener.Instance().StartListening(lobbyInfo.msg.id, lobbyInfo.playerName, () =>
		{
			Debug.Log("I'm listening now");
		});
		Goto = FindObjectOfType<GoToScene>();
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		Debug.Log("Got downstream message: \n" + message);
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			Debug.Log("Ready status: " + lobbyMsg.readyStatus);
			updatePlayers(lobbyMsg.players.ToArray(), lobbyMsg.readyStatus);
		} else if (messageType == MsgTypes.GAME_START) {
			Goto.Go("GameScene");	
		}else {
			Debug.Log("Got unhandled message: " + messageType);
		}
	}

	public void updatePlayers(string[] playerNames, Dictionary<string, bool> readyStatus)
	{
		foreach (Transform child in transform) {
			Destroy(child.gameObject);
		}
		foreach (var player in playerNames)
		{
			var rowItem = Instantiate(playerRowPrefab, transform);
			rowItem.label.text = player;
			if (readyStatus != null && readyStatus.ContainsKey(player) && readyStatus[player]) rowItem.Readied();
			// TODO: Put ready image there, too.
		}
	}

	private void OnDestroy()
	{
		WebSocketListener.Instance().CancelSubscription(this);
	}
}
