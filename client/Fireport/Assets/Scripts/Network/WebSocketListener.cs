using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Utils;

namespace Network {
    public class WebSocketListener : MonoBehaviour {

        private static WebSocketListener instance = null;
        public static WebSocketListener Instance() {
            if (instance == null) {
                var go = new GameObject("WebSocketListener");
                instance = go.AddComponent<WebSocketListener>();
            }

            return instance;
        }
        private static WebSocket webSocket;
        private static bool started;
        private List<IDownStreamSubscriber> subscribers = new List<IDownStreamSubscriber>();
        
        private static bool created;

        void Awake()
        {
            if (!created)
            {
                DontDestroyOnLoad(gameObject); // keep this game object around between scenes
                created = true;
            }
        }
        
        /// <summary>
        /// Requires you to wait for a bit before you can actually send messages
        /// </summary>
        public void StartListening(string lobbyId, string playerId, Action onSuccess) {
            if (!started) {
                StartCoroutine(startWebsocket(lobbyId, playerId, onSuccess));
            } else onSuccess.Invoke();
        }

        public void StopListening() {
            if (started) {
                started = false;
                webSocket.Close();
            }
        }

        private void OnDestroy() {
            StopListening(); // tries to clean up the websocket connection before destroying this behaviour
        }

        public void Subscribe(IDownStreamSubscriber subscriber) {
            subscribers.Add(subscriber);
        }
        
        /// <summary>
        /// Alias for CancelSubscription
        /// </summary>
        /// <param name="subscriber"></param>
        public void Unsubscribe(IDownStreamSubscriber subscriber) {
            CancelSubscription(subscriber);
        }

        public void CancelSubscription(IDownStreamSubscriber subscriber) {
            if (subscribers.Contains(subscriber)) subscribers.Remove(subscriber);
        }

        public void Send(string msg) {
            if (started) webSocket.SendString(msg);
            else Debug.LogError("Failed to send message because Websocket was still initializing");
        }

        private IEnumerator startWebsocket(string lobbyId, string playerId, Action onSuccess) {
            var url = State.WEBSOCKET_HOST + "/api/v1/pubsub/" + lobbyId + "/" + playerId;
            Debug.Log("Attempt to connect to websocket: " + url);
            var ws = new WebSocket(new Uri(url));
            yield return StartCoroutine(ws.Connect());
            started = true;
            webSocket = ws;
            onSuccess.Invoke();
            Debug.Log("WebSocket now listening");
            while (started) {
                string msg = webSocket.RecvString();
                if (msg != null) {
                    var json = JsonUtility.FromJson<MessageType>(msg); // grab the messageType string for future classification
                    subscribers.ForEach(s => {
                        try {
                            s.handleDownStreamMessage(json.msgType, msg);
                        } catch (Exception e) {
                            Debug.LogError("Error while sending message to subscriber: " + e + "\n" + e.StackTrace);
                        }
                    });
                }

                if (webSocket.error != null) {
                    Debug.LogError("WebSocketError: " + webSocket.error);
                    break;
                }

                yield return 0;
            }

            StopListening();
        }
    
        public class MessageType {
            public string msgType;
        }
    }

    public interface IDownStreamSubscriber {
        void handleDownStreamMessage(string messageType, string message);
    }
}