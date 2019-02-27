using AI;
using Network;
using UnityEngine;
using Utils;

namespace MainSceneScripts {
    public class LobbySceneBehaviour : MonoBehaviour {
        private RestApi Api;
        private GoToScene Goto;
        
        private void Start(){
            Api = FindObjectOfType<RestApi>();
            Goto = FindObjectOfType<GoToScene>();
        }

        public void StartGame() {
            Api.StartGame(LobbyInfoController.Instance().msg.id, () => {
                Goto.Go("GameScene");
            });
        }

        public void ReadyUp() {
            var lobbyInfo = LobbyInfoController.Instance();
            Api.ReadyUp(lobbyInfo.msg.id, lobbyInfo.playerName, () => {
                Debug.Log("I'm ready!");
            });
        }

        public void Leave()
        {
            WebSocketListener.Instance().StopListening();
            var lobbyInfo = LobbyInfoController.Instance();
            Api.LeaveLobby(lobbyInfo.msg.id, lobbyInfo.playerName, () => {
                Debug.Log("I left the lobby");
                Goto.Go("MainMenuScene");
            });
        }

        public void AddDumbAi() {
            GameObject go = new GameObject();
            var dumbAi = go.AddComponent<DumbAI>();
            dumbAi.Initialize();
        }
    }
}