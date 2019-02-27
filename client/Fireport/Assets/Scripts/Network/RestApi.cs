using System;
using Model.Message;
using UnityEngine;
using Utils;

namespace Network {
    public class RestApi : MonoBehaviour {

        public void CreateLobby(Action<string> onSuccess) {
            var req = new RESTEasyRequest();
            req.Body(" ").Url(State.HTTP_HOST + "/api/v1/lobby").OnSuccess(onSuccess).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Post());
        }
        
        public void JoinLobby(string code, string playerName, Action<string> onSuccess) {
            var req = new RESTEasyRequest();
            var body = new LobbyJoinMessage();
            body.lobbyID = code;
            body.playerID = playerName;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/join").OnSuccess(onSuccess).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Put());
        }
        
        public void LeaveLobby(string code, string playerName, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new LobbyLeaveMessage();
            body.playerID = playerName;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/"+code+"/leave").OnSuccess(onSuccess).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Put());
        }

        public void ReadyUp(string code, string playerName, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new PlayerReadyMessage();
            body.playerName = playerName;
            body.ready = true;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/" + code + "/ready").OnSuccess(onSuccess).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Put());
        }

        public void StartGame(string code, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new GameStartMessage();
            body.gameID = code;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/" + code + "/start").OnSuccess(onSuccess).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Put());
        }

        public void SubmitTurn(string gameId, int turn, string playerName, int playerId, int[] cardIds, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new TurnSubmissionMessage();
            body.gameID = gameId;
            body.playerID = playerName;
            for (int i = 0; i < cardIds.Length; i++) {
                var input = new GameInputMessage();
                input.order = i;
                input.owner = playerId;
                input.swap = 0; // TODO: swap cards
                input.cardID = cardIds[i];
                body.inputs.Add(input);
            }
            Debug.Log(string.Format("Card submissions: {0}", body));
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/game/" + gameId + "/turn/" + playerName).OnSuccess(onSuccess).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Put());
        }

        public void GetCurrentTurn(string gameId, Action<CurrentTurnMessage> onSuccess) {
            var req = new RESTEasyRequest();
            req.Url(State.HTTP_HOST + "/api/v1/game/" + gameId + "/turn").OnSuccess((resp, status) => {
                var msg = JsonUtility.FromJson<CurrentTurnMessage>(resp);
                onSuccess(msg);
            }).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Get());
        }

        public void GetGameState(string gameId, Action<CurrentStateMessage> onSuccess) {
            var req = new RESTEasyRequest();
            req.Url(State.HTTP_HOST + "/api/v1/game/" + gameId + "/turn/state").OnSuccess((resp, status) => {
                var msg = JsonUtility.FromJson<CurrentStateMessage>(resp);
                onSuccess(msg);
            }).OnFailure((s, i) => handleFailure(req, s, i));
            StartCoroutine(req.Get());
        }

        private void handleFailure(RESTEasyRequest req, string error, int status) {
            Debug.LogError(string.Format("{2} ({0}): {1}", status, error, req.url));
        }
    }
}