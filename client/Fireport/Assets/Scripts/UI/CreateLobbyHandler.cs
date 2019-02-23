using System.Collections;
using System.Collections.Generic;
using Model.Message;
using Network;
using TMPro;
using UnityEngine;
using UnityEngine.UI;
using Utils;

public class CreateLobbyHandler : MonoBehaviour
{
	public TMP_InputField PlayerNameInputField;
	public Button CreateLobbyButton;
	public RestApi Api;
	public GoToScene sceneChanger;

	void Start () {
		WebSocketListener.Instance();
		LobbyInfoController.Instance();
		
		LobbyInfoController.ClearLobbyObject();
		PlayerNameInputField.onValueChanged.AddListener(UpdateInput);
		CreateLobbyButton.onClick.AddListener(CreateLobby);
		CreateLobbyButton.interactable = false;
	}

	private void UpdateInput(string input) {
		CreateLobbyButton.interactable = !string.IsNullOrEmpty(input);
	}

	public void CreateLobby() {
		if (!string.IsNullOrEmpty(PlayerNameInputField.text)) {
			Api.CreateLobby((createRespBody) => {
				Api.JoinLobby(createRespBody, PlayerNameInputField.text, (body) => {
					var lobbyInfo = LobbyInfoController.Instance();
					lobbyInfo.msg = JsonUtility.FromJson<LobbyMessage>(body);
					lobbyInfo.playerName = PlayerNameInputField.text;
					sceneChanger.Go("LobbyScene");
				});
			});
		}
	}
}
