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
		public Button JoinLobbyButton;
		public RestApi Api;
		public GoToScene sceneChanger;

		public LobbyInfoController lobbyInfo;
		
		void Start () {
			LobbyIDInput.onValueChanged.AddListener(UpdateInput);
			JoinLobbyButton.onClick.AddListener(JoinLobby);
			JoinLobbyButton.interactable = false;
		}

		private void UpdateInput(string input) {
			JoinLobbyButton.interactable = !string.IsNullOrEmpty(input);
		}

		private void JoinLobby() {
			if (!string.IsNullOrEmpty(LobbyIDInput.text)) {
				var pNum = Random.Range(1, 1000);
				Api.JoinLobby(LobbyIDInput.text, "Player " + pNum.ToString(), (body) => {
					// TODO: parse resp data and pass it to the scene
					var lobby = Instantiate(lobbyInfo);
					DontDestroyOnLoad(lobby.transform.gameObject);
					lobby.name = LobbyInfoController.objectName;
					var lobbyMessage = JsonUtility.FromJson<LobbyMessage>(body);
					lobby.msg = lobbyMessage;
					sceneChanger.Go("LobbyScene");
				});
			}
		}
	}
}
