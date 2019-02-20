using Model.Message;
using Network;
using UnityEngine;
using Utils;

public class LobbyPlayerList : MonoBehaviour, IDownStreamSubscriber
{
	public PlayerRowController playerRowPrefab;
	private RestApi Api;
	private GoToScene Goto;

	private void Start()
	{
		WebSocketListener.Instance().Subscribe(this);
		LobbyInfoController lobbyInfo = LobbyInfoController.Instance();
		updatePlayers(lobbyInfo.msg.players.ToArray());
		WebSocketListener.Instance().StartListening(lobbyInfo.msg.id, lobbyInfo.playerName, () =>
		{
			Debug.Log("I'm listening now");
		});
		Api = FindObjectOfType<RestApi>();
		Goto = FindObjectOfType<GoToScene>();
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		Debug.Log("Got downstream message");
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			updatePlayers(lobbyMsg.players.ToArray());
		} else {
			Debug.Log("Got unhandled message: " + messageType);
		}
	}

	public void updatePlayers(string[] playerNames)
	{
		foreach (Transform child in transform) {
			Destroy(child.gameObject);
		}
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
