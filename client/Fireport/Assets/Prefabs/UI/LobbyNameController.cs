using System.Collections;
using System.Collections.Generic;
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
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			GetComponent<TextMeshProUGUI>().text = "Lobby: " + lobbyMsg.id;
		}
	}
}
