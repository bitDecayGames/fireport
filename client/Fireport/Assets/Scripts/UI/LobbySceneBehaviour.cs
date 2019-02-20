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
    }
}