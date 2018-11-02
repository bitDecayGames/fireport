﻿using System.Collections.Generic;
using Model.Message;
using Model.State;
using Network;
using UnityEngine;
using UnityEngine.UI;
using Utils;

namespace Dev {
    public class DevGameBehaviour : MonoBehaviour, IDownStreamSubscriber {
        public const int MAX_CARD_SELECTIONS = 3;

        public RestApi Api;
        public WebSocketListener Listener;

        public InputField GameCodeInputField;
        public InputField PlayerNameInputField;
        public Button CreateButton;
        public Button JoinButton;
        public Button ReadyButton;
        public Button StartButton;
        public Text ActivityText;
        public Text GameInfoText;
        public Transform CardGroup;

        private List<Button> cardButtons = new List<Button>();
        private List<Text> cards = new List<Text>();
        private List<int> cardSelections = new List<int>();
        private int currentTurn = 0;
        private int playerId = 0;
        private GameState currentState;
        private PlayerState currentPlayer;

        private void Start() {
            Listener.Subscribe(this);

            cards.AddRange(CardGroup.GetComponentsInChildren<Text>());
            cardButtons.AddRange(CardGroup.GetComponentsInChildren<Button>());

            CreateButton.onClick.AddListener(CreateLobby);
            JoinButton.onClick.AddListener(JoinLobby);
            ReadyButton.onClick.AddListener(ReadyUp);
            StartButton.onClick.AddListener(StartGame);

            PlayerNameInputField.interactable = false;
            ReadyButton.interactable = false;
            StartButton.interactable = false;
            cardButtons.ForEach(b => b.interactable = false);
        }

        public void CreateLobby() {
            addToActivityStream("Attempt to create lobby");
            Api.CreateLobby(lobbyCode => {
                GameCodeInputField.text = lobbyCode;
                GameCodeInputField.enabled = false;
                addToActivityStream("Created lobby " + lobbyCode);
                CreateButton.interactable = false;
            });
        }

        public void JoinLobby() {
            addToActivityStream("Attempt to join lobby " + GameCodeInputField.text);
            Api.JoinLobby(GameCodeInputField.text, PlayerNameInputField.text, () => {
                addToActivityStream("Joined lobby " + GameCodeInputField.text);
                // TODO: probably need to start listening on the websocket
                PlayerNameInputField.interactable = false;
                JoinButton.interactable = false;
                ReadyButton.interactable = true;
                StartButton.interactable = true;
            });
        }

        public void ReadyUp() {
            addToActivityStream("Attempt to ready up");
            Api.ReadyUp(GameCodeInputField.text, PlayerNameInputField.text,
                () => {
                    addToActivityStream(PlayerNameInputField.text + " is ready");
                    ReadyButton.interactable = false;
                });
        }

        public void StartGame() {
            addToActivityStream("Attempt to start game");
            Api.StartGame(GameCodeInputField.text, () => {
                addToActivityStream("Game started");
                cardButtons.ForEach(b => b.interactable = true);
                StartButton.interactable = false;
            });
        }

        public void SelectCard(int index) {
            if (!cardSelections.Contains(index)) {
                addToActivityStream("Select card " + index);
                cardSelections.Add(index);
                if (cardSelections.Count >= MAX_CARD_SELECTIONS) {
                    var cardIds = cardSelections.ConvertAll(i => currentPlayer.Hand[i].ID);
                    Api.SubmitTurn(GameCodeInputField.text, currentTurn, PlayerNameInputField.text, playerId,
                        cardIds.ToArray(),
                        () => { addToActivityStream("Submitted selections"); });
                    cardSelections.Clear();
                }
            }
        }

        private void OnDestroy() {
            Listener.StopListening();
        }

        public void handleDownStreamMessage(string messageType, string message) {
            Debug.Log("Got message: " + messageType + "\n" + message);
            addToActivityStream("Received message: " + messageType);

            switch (messageType) {
                case MsgTypes.TURN_RESULT:
                    var msg = JsonUtility.FromJson<TurnResultMessage>(message);
                    currentState = msg.currentState;
                    currentPlayer = currentState.Players.Find(p => p.Name == PlayerNameInputField.text);
                    msg.animationActions.ForEach(a => addToActivityStream("Action: " + a.Name));
                    gameStateToInfoText();
                    cardStatesToButtonText();
                    break;
                default:
                    addToActivityStream("Message unhandled: " + messageType);
                    break;
            }
        }

        private void addToActivityStream(string message) {
            ActivityText.text = message + "\n" + ActivityText.text;
        }

        private void gameStateToInfoText() {
            GameInfoText.text = currentState.ToString();
        }

        private void cardStatesToButtonText() {
            for (int i = 0; i < currentPlayer.Hand.Count; i++) {
                var card = currentPlayer.Hand[i];
                if (System.Enum.IsDefined(typeof(CardType), card.CardType)) {
                    var cardType = (CardType) card.CardType;
                    cards[i].text = cardType.ToString();
                }
                else {
                    cards[i].text = "Unknown(" + card.ID + "): " + card.CardType;
                }
            }
        }
    }
}