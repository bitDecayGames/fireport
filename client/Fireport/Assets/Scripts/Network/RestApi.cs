using System;
using Model.Message;
using UnityEngine;
using Utils;

namespace Network {
    public class RestApi : MonoBehaviour {

        public void CreateLobby(Action<string> onSuccess) {
            var req = new RESTEasyRequest();
            req.Body(" ").Url(State.HTTP_HOST + "/api/v1/lobby").OnSuccess(onSuccess).OnFailure(handleFailure);
            StartCoroutine(req.Post());
        }
        
        public void JoinLobby(string code, string playerName, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new LobbyJoinMessage();
            body.lobbyID = code;
            body.playerID = playerName;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/join").OnSuccess(onSuccess).OnFailure(handleFailure);
            StartCoroutine(req.Put());
        }

        public void ReadyUp(string code, string playerName, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new PlayerReadyMessage();
            body.playerName = playerName;
            body.ready = true;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/" + code + "/ready").OnSuccess(onSuccess).OnFailure(handleFailure);
            StartCoroutine(req.Put());
        }

        public void StartGame(string code, Action onSuccess) {
            var req = new RESTEasyRequest();
            var body = new GameStartMessage();
            body.gameID = code;
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/lobby/" + code + "/start").OnSuccess(onSuccess).OnFailure(handleFailure);
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
            req.Body(JsonUtility.ToJson(body)).Url(State.HTTP_HOST + "/api/v1/game/" + gameId + "/turn/" + turn + "/player/" + playerName).OnSuccess(onSuccess).OnFailure(handleFailure);
            StartCoroutine(req.Put());
        }

        public void GetCurrentTurn(string gameId, int turn, Action onSuccess) {
            var req = new RESTEasyRequest();
            req.Url(State.HTTP_HOST + "/api/v1/game/" + gameId + "/turn/" + turn).OnSuccess(onSuccess).OnFailure(handleFailure);
            StartCoroutine(req.Get());
        }

        public void GetGameState(string gameId, int turn, string playerName, int playerId, Action onSuccess) {
            var req = new RESTEasyRequest();
            req.Body(JsonUtility.ToJson("{}")).Url(State.HTTP_HOST + "/api/v1/game/" + gameId + "/turn/" + turn + "/player/" + playerName).OnSuccess(onSuccess).OnFailure(handleFailure);
            StartCoroutine(req.Get());
        }

        private void handleFailure(string error, int status) {
            Debug.LogError(string.Format("ApiFailure({0}): {1}", status, error));
        }
    }
}