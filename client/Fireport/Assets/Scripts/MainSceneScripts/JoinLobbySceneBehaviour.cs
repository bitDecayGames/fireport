using UnityEngine;
using UnityEngine.UI;

namespace MainSceneScripts {
	public class JoinLobbySceneBehaviour : MonoBehaviour {

		public InputField LobbyIDInput;
		public Button JoinLobbyButton;
		
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
				// TODO: join the lobby
//				Api.JoinLobby(GameCodeInputField.text, PlayerNameInputField.text, () => {
//					addToActivityStream("Joined lobby " + GameCodeInputField.text);
//					Listener.StartListening(GameCodeInputField.text, PlayerNameInputField.text,
//						() => { addToActivityStream("Made websocket connection"); });
//					GameCodeInputField.interactable = false;
//					PlayerNameInputField.interactable = false;
//					CreateButton.interactable = false;
//					JoinButton.interactable = false;
//					ReadyButton.interactable = true;
//				});
			}
		}
	}
}
