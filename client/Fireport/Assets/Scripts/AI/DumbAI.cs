using System;
using System.Collections;
using System.Collections.Generic;
using Model.Message;
using Model.State;
using Network;
using UnityEngine;
using Utils;

namespace AI {
    /// <summary>
    /// This DumbAI doesn't listen to its own websocket, therefore, when we start filtering
    /// the data that comes down the pipe to not include other players, this script will
    /// stop working.
    /// </summary>
    public class DumbAI : MonoBehaviour, IDownStreamSubscriber {
        public static List<DumbAI> Ais = new List<DumbAI>();
        
        public string gameCode;
        public int playerId = -1;
        public string playerName;

        private GameState gameState;

        private RestApi _api;
        public RestApi Api {
            get {
                if (!_api) RefreshApi();
                return _api;
            }
        }

        void Start() {
            DontDestroyOnLoad(gameObject);
            WebSocketListener.Instance().Subscribe(this);
        }
        
        private void RefreshApi() {
            _api = FindObjectOfType<RestApi>();
            if (!_api) throw new Exception("A RestApi object is required in the scene for a DumpAI");
        }

        public void Initialize() {
            gameCode = LobbyInfoController.Instance().msg.id;
            playerName = string.Format("DumbAI_{0}", Ais.Count);
            gameObject.name = playerName;
            Ais.Add(this);
            StartCoroutine(WaitThenDo(1, () => JoinLobby()));
        }

        public void JoinLobby() {
            Api.JoinLobby(gameCode, playerName, s => {
                Debug.Log(string.Format("{0} just joined the lobby: {1}", playerName, s));
                Api.ReadyUp(gameCode, playerName, () => {
                    Debug.Log(string.Format("{0} just readied up", playerName));
                });
            });
        }

        public void PlayCards() {
            Api.SubmitTurn(gameCode, gameState.Turn, playerName, playerId, GetCardsFromHand(3).ConvertAll(c => c.ID).ToArray(), () => Debug.Log(string.Format("{0} just played their cards", playerName)));
        }

        public List<CardState> GetCardsFromHand(int count) {
            var player = gameState.Players.Find(p => p.ID == playerId);
            List<CardState> cards = new List<CardState>();
            for (int i = 0; i < player.Hand.Count && i < count; i++) {
                cards.Add(player.Hand[i]);
            }

            return cards;
        }

        public void handleDownStreamMessage(string messageType, string message) {
            switch (messageType) {
                case MsgTypes.GAME_START:
                    var gameStartMessage = JsonUtility.FromJson<GameStartMessage>(message);
                    gameState = gameStartMessage.gameState;
                    playerId = gameState.Players.Find(p => p.Name == playerName).ID;
                    PlayCards();
                    break;
                case MsgTypes.TURN_RESULT:
                    var turnResultMessage = JsonUtility.FromJson<TurnResultMessage>(message);
                    gameState = turnResultMessage.currentState;
                    playerId = gameState.Players.Find(p => p.Name == playerName).ID;
                    PlayCards();
                    break;
                default:
                    Debug.Log(string.Format("{0} failed to handle message type({1}): {2}", playerName, messageType, message));
                    break;
            }
        }

        private void OnDestroy() {
            WebSocketListener.Instance().Unsubscribe(this);
        }

        private IEnumerator WaitThenDo(float secondsToWait, Action thenDo) {
            yield return new WaitForSeconds(secondsToWait);
            thenDo();
        }
    }
}