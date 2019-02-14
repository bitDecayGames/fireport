using Model.Message;
using Network;
using UnityEngine;
using Utils;

public class LobbyPlayerList : MonoBehaviour, IDownStreamSubscriber
{
	public PlayerRowController playerRowPrefab;

	private void Start()
	{
		WebSocketListener.Instance().Subscribe(this);
		LobbyInfoController lobbyInfo = LobbyInfoController.Instance();
		updatePlayers(lobbyInfo.msg.players.ToArray());
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			updatePlayers(lobbyMsg.players.ToArray());
		}
	}

	public void updatePlayers(string[] playerNames)
	{
		foreach (var player in playerNames)
		{
			var rowItem = Instantiate(playerRowPrefab, transform);
			rowItem.label.text = player;
			// TODO: Put ready image there, too.
		}
	}

	private void OnDestroy()
	{
		WebSocketListener.Instance().CancelSubscription(this);
	}
}
