using Model.Message;
using Network;
using TMPro;
using UnityEngine;
using Utils;

public class GameInfoController : MonoBehaviour, IDownStreamSubscriber
{

	public TextMeshProUGUI gameTypeLabel;

	private void Start()
	{
		WebSocketListener.Instance().Subscribe(this);
		
	}

	public void handleDownStreamMessage(string messageType, string message)
	{
		if (messageType == MsgTypes.LOBBY)
		{
			var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
			// populate game type once we have that
			//lobbyMsg.gametype
		}
	}
}
