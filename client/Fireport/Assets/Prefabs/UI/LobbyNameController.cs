using System.Collections;
using System.Collections.Generic;
using Boo.Lang.Runtime;
using Model.Message;
using Network;
using TMPro;
using UnityEngine;
using Utils;

public class LobbyNameController : MonoBehaviour, IDownStreamSubscriber
{
	public WebSocketListener Listener;

	private void Start()
	{
		Listener.Subscribe(this);
		LobbyInfoController lobby = LobbyInfoController.GetLobbyObject();
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
