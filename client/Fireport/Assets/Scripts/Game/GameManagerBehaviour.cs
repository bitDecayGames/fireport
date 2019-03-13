using System;
using System.Collections.Generic;
using System.Text;
using AnimationEngine;
using Game;
using Game.UI;
using Model.Message;
using Model.State;
using Network;
using TMPro;
using UnityEngine;
using Utils;

public class GameManagerBehaviour : MonoBehaviour, IDownStreamSubscriber {
    public const int MAX_CARD_SELECTIONS = 3;
    
    public TextMeshProUGUI ActivityText;
    public TextMeshProUGUI PlayerInfoText;

    private AnimationEngineBehaviour AnimationEngine;
    private RestApi Api;
    private CardTrayBehaviour CardTray;
    private GameBoardBehaviour Board;
    private LobbyInfoController lobbyInfo;

    private int currentTurn = 0;
    private GameState currentState;
    private PlayerState currentPlayer;

    private Action onAnimationFinish;

    void Start() {
        Api = FindObjectOfType<RestApi>();
        
        WebSocketListener.Instance().Subscribe(this);

        CardTray = FindObjectOfType<CardTrayBehaviour>();
        CardTray.OnSelected.AddListener(OnCardSelections);
        
        Board = FindObjectOfType<GameBoardBehaviour>();

        AnimationEngine = FindObjectOfType<AnimationEngineBehaviour>();
        AnimationEngine.OnComplete.AddListener(onAnimationsComplete);
        
        lobbyInfo = LobbyInfoController.Instance();
        if (lobbyInfo != null && lobbyInfo.msg != null) WebSocketListener.Instance().StartListening(lobbyInfo.msg.id, lobbyInfo.playerName, () => {
            Debug.Log("I'm listening...");
            if (lobbyInfo.gameStartMessage != null) handleDownStreamMessage(MsgTypes.GAME_START, lobbyInfo.gameStartMessage);
        });
    }

    private void OnDestroy() {
        WebSocketListener.Instance().Unsubscribe(this);
    }

    public void OnCardSelections(List<CardBehaviour> selections) {
        if (selections.Count >= MAX_CARD_SELECTIONS) { // TODO: MW change this to be on click of a submit button
            CardTray.SetCards(new List<CardState>());
            CardTray.Hide();
            Api.SubmitTurn(
                lobbyInfo.msg.id,
                currentTurn,
                lobbyInfo.playerName,
                lobbyInfo.playerId,
                selections.ConvertAll(c => c.Card.ID).ToArray(),
                () => { addToActivityStream("Submitted selections"); });
        }
    }

    public void handleDownStreamMessage(string messageType, string message) {
        Debug.Log("Got message: " + messageType + "\n" + message);
        addToActivityStream("Received message: " + messageType);

        switch (messageType) {
            case MsgTypes.TURN_RESULT:
                var turnResultMsg = TurnResultMessage.FromJson(message);
                Debug.Log("Animations: " + turnResultMsg.currentState.Animations);
                applyAnimations(turnResultMsg.previousState, turnResultMsg.currentState, () => {
                    nextState(turnResultMsg.currentState);
                });
                break;
            case MsgTypes.GAME_START:
                var gameStartMessage = JsonUtility.FromJson<GameStartMessage>(message);
                nextState(gameStartMessage.gameState);
                break;
            default:
                addToActivityStream("Message unhandled: " + messageType);
                break;
        }
    }

    private void applyAnimations(GameState previous, GameState next, Action onAnimationFinish) {
        this.onAnimationFinish = onAnimationFinish;
        
        var gamePieces = new List<GamePieceBehaviour>();
        gamePieces.AddRange(FindObjectsOfType<GamePieceBehaviour>()); // TODO: MW this is highly inefficient
        
        AnimationEngine.Play(next.Animations, gamePieces);
    }

    private void onAnimationsComplete() {
        if (onAnimationFinish != null) {
            onAnimationFinish();
            onAnimationFinish = null;
        }
    }

    private void nextState(GameState next) {
        Debug.Log("Got key frame state: " + JsonUtility.ToJson(next));
        currentState = next;
        Board.Populate(currentState); // TODO: MW I'm guessing we will run into problems by just continually rebuilding the board each key frame.  I imagine we will need to do some smart reloading/updating instead of just replacing.
        currentPlayer = currentState.Players.Find(p => p.Name == lobbyInfo.playerName);
        lobbyInfo.playerId = currentPlayer.ID;
        Debug.Log("Got current player: " + JsonUtility.ToJson(currentPlayer));
        
        if (currentState.IsGameFinished) addToActivityStream("Game Over! A winner is: " + currentState.Winner);
        CardTray.SetCards(currentPlayer.Hand);
        CardTray.Show();
        playerStateToInfoText();
        nextState(next);
    }

    private void addToActivityStream(string message) {
        ActivityText.text = " - " + message + "\n" + ActivityText.text;
    }

    private void checkForGameEnd(GameState next)
    {
        if (next.Winner != 0)
        {
            // TODO: Does this trigger at the right time?
            SceneNavigation.LoadScene("GameOverScene");
        }
    }

    private void playerStateToInfoText() {
        StringBuilder sb = new StringBuilder();
        sb.Append("Turn: ").Append(currentState.Turn).AppendLine();
        sb.Append("Health: ").Append(currentPlayer.Health).AppendLine();
        PlayerInfoText.text = sb.ToString();
    }
}