using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Model;
using Model.Message;
using Network;
using TMPro;
using Utils;
using WebSocketSharp;

public class LobbyPlayerList : MonoBehaviour, IDownStreamSubscriber
{
	public PlayerRowController playerRowPrefab;
	public WebSocketListener Listener;

	private void Start()
	{
		Listener.Subscribe(this);
		
		updatePlayers(new string[]{"one", "two", "three"});
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
		Listener.CancelSubscription(this);
	}
}
