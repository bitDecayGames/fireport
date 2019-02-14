using Network;
using UnityEngine;
using UnityEngine.UI;

namespace DebugScripts {
    public class DebugWebsockets : MonoBehaviour, IDownStreamSubscriber {
        
        [HideInInspector] public string GAME_ID = "GAME";
        [HideInInspector] public string PLAYER_NAME = "Mike";
        
        public InputField GameCodeInputField;
        public InputField PlayerNameInputField;
        public InputField MessageInputField;

        void Start() {
            GameCodeInputField.text = GAME_ID;
            PlayerNameInputField.text = PLAYER_NAME;
        }

        public void Connect() {
            WebSocketListener.Instance().StartListening(GAME_ID, PLAYER_NAME, () => { UnityEngine.Debug.Log("Websocket connection successful"); });
        }

        public void Disconnect() {
            WebSocketListener.Instance().StopListening();
        }

        public void Send() {
            WebSocketListener.Instance().Send(MessageInputField.text);
        }
        
        
        public void handleDownStreamMessage(string messageType, string message) {
            Debug.Log("(" + messageType + ") " + message);
        }
    }
}