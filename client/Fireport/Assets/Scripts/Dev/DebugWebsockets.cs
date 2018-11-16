using Network;
using UnityEngine;
using UnityEngine.UI;

namespace Dev {
    public class DebugWebsockets : MonoBehaviour, IDownStreamSubscriber {
        
        [HideInInspector] public string GAME_ID = "GAME";
        [HideInInspector] public string PLAYER_NAME = "Mike";
        
        public InputField GameCodeInputField;
        public InputField PlayerNameInputField;
        public InputField MessageInputField;
        public WebSocketListener Listener;

        void Start() {
            GameCodeInputField.text = GAME_ID;
            PlayerNameInputField.text = PLAYER_NAME;
        }

        public void Connect() {
            Listener.StartListening(GAME_ID, PLAYER_NAME, () => { Debug.Log("Websocket connection successful"); });
        }

        public void Disconnect() {
            Listener.StopListening();
        }

        public void Send() {
            Listener.Send(MessageInputField.text);
        }
        
        
        public void handleDownStreamMessage(string messageType, string message) {
            Debug.Log("(" + messageType + ") " + message);
        }
    }
}