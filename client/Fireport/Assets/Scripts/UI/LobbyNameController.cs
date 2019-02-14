using Model.Message;
using Network;
using TMPro;
using UnityEngine;
using Utils;

public class LobbyNameController : MonoBehaviour, IDownStreamSubscriber
{
	private void Start()
	{
		WebSocketListener.Instance().Subscribe(this);
		LobbyInfoController lobby = LobbyInfoController.Instance();
		updateLobbyName(lobby.msg.id);
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			updateLobbyName(lobbyMsg.id);
		}
	}

	private void updateLobbyName(string lobbyID)
	{
		GetComponent<TextMeshProUGUI>().text = "Lobby: " + lobbyID;
	}
}
