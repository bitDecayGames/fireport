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
		WebSocketListener.Instance().StartListening(lobbyInfo.msg.id, lobbyInfo.playerName, () => {});
		Goto = FindObjectOfType<GoToScene>();
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		Debug.Log("Got downstream message: \n" + message);
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			lobbyMsg.readyStatus = parserReadyStatus(message);
			updatePlayers(lobbyMsg.players.ToArray(), lobbyMsg.readyStatus);
		} else if (messageType == MsgTypes.GAME_START) {
			LobbyInfoController.Instance().gameStartMessage = message;
			Goto.Go("GameScene");	
		}else {
			Debug.Log("Got unhandled message: " + messageType);
		}
	}

	private Dictionary<string, bool> parserReadyStatus(string message) {
		var msg = Json.Deserialize(message) as Dictionary<string, object>;
		if (msg != null) {
			var readyStatus = msg["readyStatus"] as Dictionary<string, object>;
			var ready = new Dictionary<string, bool>();
			if (readyStatus != null) {
				foreach (var key in readyStatus.Keys) {
					var value = readyStatus[key];
					if (value is bool) ready[key] = (bool) readyStatus[key];
				}

				return ready;
			}
		}
		return null;
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
