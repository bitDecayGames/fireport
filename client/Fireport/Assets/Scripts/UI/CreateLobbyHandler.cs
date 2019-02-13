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

	public LobbyInfoController lobbyInfo;

	void Start () {
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
					var lobby = Instantiate(lobbyInfo);
					DontDestroyOnLoad(lobby.transform.gameObject);
					lobby.name = LobbyInfoController.objectName;
					var lobbyMessage = JsonUtility.FromJson<LobbyMessage>(body);
					lobby.msg = lobbyMessage;
					sceneChanger.Go("LobbyScene");
				});
			});
		}
	}
}
