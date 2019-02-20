using System;
using Model.Message;
using Network;
using UnityEngine;
using UnityEngine.UI;
using Utils;
using Random = UnityEngine.Random;

namespace MainSceneScripts {
	public class JoinLobbySceneBehaviour : MonoBehaviour {

		public InputField LobbyIDInput;
		public InputField PlayerNameInput;
		public Button JoinLobbyButton;
		public RestApi Api;
		public GoToScene sceneChanger;

		public LobbyInfoController lobbyInfo;
		
		void Start () {
			WebSocketListener.Instance();
			LobbyInfoController.Instance();
			LobbyInfoController.ClearLobbyObject();
			LobbyIDInput.onValueChanged.AddListener(UpdateInput);
			JoinLobbyButton.onClick.AddListener(JoinLobby);
			JoinLobbyButton.interactable = false;
		}

		private void UpdateInput(string input) {
			JoinLobbyButton.interactable = !string.IsNullOrEmpty(input);
		}

		private void JoinLobby() {
			if (!string.IsNullOrEmpty(LobbyIDInput.text) && !string.IsNullOrEmpty(PlayerNameInput.text)) {
				Api.JoinLobby(LobbyIDInput.text, PlayerNameInput.text, (body) => {
					var lobby = LobbyInfoController.Instance();
					lobby.name = LobbyInfoController.objectName;
					var lobbyMessage = JsonUtility.FromJson<LobbyMessage>(body);
					lobby.msg = lobbyMessage;
					sceneChanger.Go("LobbyScene");
				});
			}
		}
	}
}
