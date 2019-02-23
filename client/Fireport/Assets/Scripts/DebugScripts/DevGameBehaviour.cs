using System.Collections.Generic;
using System.Text;
using Model.Message;
using Model.State;
using Network;
using UnityEngine;
using UnityEngine.UI;
using Utils;

namespace DebugScripts {
    public class DevGameBehaviour : MonoBehaviour, IDownStreamSubscriber {
        public const int MAX_CARD_SELECTIONS = 3;
        public const int CARDS_IN_HAND = 5;

        public RestApi Api;

        public InputField GameCodeInputField;
        public InputField PlayerNameInputField;
        public Button CreateButton;
        public Button JoinButton;
        public Button ReadyButton;
        public Button StartButton;
        public Text ActivityText;
        public Text GameInfoText;
        public Text PlayerInfoText;
        public Transform CardGroup;

        private List<Button> cardButtons = new List<Button>();
        private List<Text> cards = new List<Text>();
        private List<int> cardSelections = new List<int>();
        private int currentTurn = 0;
        private int playerId = 0;
        private GameState currentState;
        private PlayerState currentPlayer;

        private void Start() {
            WebSocketListener.Instance().Subscribe(this);

            cards.AddRange(CardGroup.GetComponentsInChildren<Text>());
            cardButtons.AddRange(CardGroup.GetComponentsInChildren<Button>());

            CreateButton.onClick.AddListener(CreateLobby);
            JoinButton.onClick.AddListener(JoinLobby);
            ReadyButton.onClick.AddListener(ReadyUp);
            StartButton.onClick.AddListener(StartGame);

            ReadyButton.interactable = false;
            StartButton.interactable = false;
            cardButtons.ForEach(b => b.interactable = false);
            
            PlayerNameInputField.Select();
            
            debugPrintGameState();
        }

        public void CreateLobby() {
            addToActivityStream("Attempt to create lobby");
            Api.CreateLobby(lobbyCode => {
                Debug.Log(lobbyCode);
                GameCodeInputField.text = lobbyCode;
                GameCodeInputField.interactable = false;
                addToActivityStream("Created lobby " + lobbyCode);
                CreateButton.interactable = false;
                PlayerNameInputField.Select();
            });
        }

        public void JoinLobby() {
            if (PlayerNameInputField.text.Length > 0) {
                if (GameCodeInputField.text.Length <= 0) GameCodeInputField.text = "GAME";
                addToActivityStream("Attempt to join lobby " + GameCodeInputField.text);
                Api.JoinLobby(GameCodeInputField.text, PlayerNameInputField.text, (resp) => {
                    addToActivityStream("Joined lobby " + GameCodeInputField.text);
                    WebSocketListener.Instance().StartListening(GameCodeInputField.text, PlayerNameInputField.text,
                        () => { addToActivityStream("Made websocket connection"); });
                    GameCodeInputField.interactable = false;
                    PlayerNameInputField.interactable = false;
                    CreateButton.interactable = false;
                    JoinButton.interactable = false;
                    ReadyButton.interactable = true;
                });
            } else addToActivityStream("Failed to join lobby, must have a player name");
        }

        public void ReadyUp() {
            addToActivityStream("Attempt to ready up");
            Api.ReadyUp(GameCodeInputField.text, PlayerNameInputField.text,
                () => {
                    addToActivityStream(PlayerNameInputField.text + " is ready");
                    ReadyButton.interactable = false;
                    StartButton.interactable = true;
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
            if (!cardSelections.Contains(index) && index < cardButtons.Count && index >= 0) {
                addToActivityStream("Select card " + index);
                cardSelections.Add(index);
                cardButtons[index].interactable = false;
                if (cardSelections.Count >= MAX_CARD_SELECTIONS) {
                    Debug.Log("Card selections: " + string.Join(", ", cardSelections.ConvertAll(i => i.ToString()).ToArray()));
                    Debug.Log("Player hand: " + string.Join(", ", currentPlayer.Hand.ConvertAll(c => c.ID.ToString()).ToArray()));
                    var cardIds = cardSelections.ConvertAll(i => currentPlayer.Hand[i].ID);
                    Api.SubmitTurn(GameCodeInputField.text, currentTurn, PlayerNameInputField.text, playerId,
                        cardIds.ToArray(),
                        () => { addToActivityStream("Submitted selections"); });
                    cardSelections.Clear();
                    cardButtons.ForEach(b => b.interactable = false);
                }
            }
        }

        private void OnDestroy() {
            WebSocketListener.Instance().StopListening();
        }

        public void handleDownStreamMessage(string messageType, string message) {
            Debug.Log("Got message: " + messageType + "\n" + message);
            addToActivityStream("Received message: " + messageType);

            switch (messageType) {
                case MsgTypes.LOBBY:
                    var lobbyMsg = JsonUtility.FromJson<LobbyMessage>(message);
                    addToActivityStream("Lobby status: " + lobbyMsg.readyStatus);
                    break;
                case MsgTypes.GAME_START:
                    var gameStartMsg = JsonUtility.FromJson<GameStartMessage>(message);
                    addToActivityStream("Game started");
                    nextState(gameStartMsg.gameState);
                    break;
                case MsgTypes.TURN_RESULT:
                    var turnResultMsg = JsonUtility.FromJson<TurnResultMessage>(message);
                    turnResultMsg.currentState.Animations.ForEach(aL => aL.ForEach(a => addToActivityStream("Action: " + a.Name)));
                    nextState(turnResultMsg.currentState);
                    break;
                default:
                    addToActivityStream("Message unhandled: " + messageType);
                    break;
            }
        }

        private void nextState(GameState next) {
            Debug.Log("Got current state: " + JsonUtility.ToJson(next));
            currentState = next;
            currentPlayer = currentState.Players.Find(p => p.Name == PlayerNameInputField.text);
            playerId = currentPlayer.ID;
            Debug.Log(string.Format("Cards in Hand: {0} Deck: {1} Discard: {2}", currentPlayer.Hand.Count, currentPlayer.Deck.Count, currentPlayer.Discard.Count));
            Debug.Log("Got current player: " + JsonUtility.ToJson(currentPlayer));
            gameStateToInfoText();
            playerStateToInfoText();
            cardStatesToButtonText();
        }

        private void addToActivityStream(string message) {
            ActivityText.text = " - " + message + "\n" + ActivityText.text;
        }

        private void gameStateToInfoText() {
            GameInfoText.text = currentState.ToString();
        }

        private void playerStateToInfoText() {
            StringBuilder sb = new StringBuilder();
            sb.Append("Turn: ").Append(currentState.Turn).AppendLine();
            sb.Append("Health: ").Append(currentPlayer.Health).AppendLine();
            if (currentState.IsGameFinished) sb.Append("Game Over").AppendLine().Append(currentState.Winner);
            PlayerInfoText.text = sb.ToString();
        }

        private void cardStatesToButtonText() {
            for (int i = 0; i < CARDS_IN_HAND; i++) {
                if (i < currentPlayer.Hand.Count) {
                    var card = currentPlayer.Hand[i];
                    if (System.Enum.IsDefined(typeof(CardType), card.CardType)) {
                        var cardType = (CardType) card.CardType;
                        cards[i].text = cardType.ToString();
                    }
                    else {
                        cards[i].text = "???(" + card.ID + "): " + card.CardType;
                    }

                    cardButtons[i].interactable = true;
                }
                else {
                    cards[i].text = "<empty>";
                    cardButtons[i].interactable = false;
                }
            }
        }

        private void debugPrintGameState() {
            var a = new GameState();
            a.Turn = 7;
            a.BoardWidth = 7;
            a.BoardHeight = 7;
            a.Players = new List<PlayerState>();
            a.Players.Add(new PlayerState());
            a.Players[0].Name = "Mike";
            a.Players[0].Location = 7;
            a.Players[0].Hand = new List<CardState>();
            a.Players[0].Hand.Add(randomCard(101));
            a.Players[0].Hand.Add(randomCard(102));
            a.Players[0].Hand.Add(randomCard(103));
            a.Players[0].Hand.Add(randomCard(104));
            a.Players[0].Hand.Add(randomCard(105));
            a.Players.Add(new PlayerState());
            a.Players[1].Name = "Bob";
            a.Players[1].Location = 2;
            a.Players[1].Facing = 2;
            a.Players[1].Hand = new List<CardState>();
            a.Players[1].Hand.Add(randomCard(201));
            a.Players[1].Hand.Add(randomCard(202));
            a.Players[1].Hand.Add(randomCard(203));
            a.Players[1].Hand.Add(randomCard(204));
            a.Players[1].Hand.Add(randomCard(205));
            a.Players.Add(new PlayerState());
            a.Players[2].Name = "Steve";
            a.Players[2].Location = 18;
            a.Players[2].Facing = 3;
            a.Players[2].Hand = new List<CardState>();
            a.Players[2].Hand.Add(randomCard(301));
            a.Players[2].Hand.Add(randomCard(302));
            a.Players[2].Hand.Add(randomCard(303));
            a.Players[2].Hand.Add(randomCard(304));
            a.Players[2].Hand.Add(randomCard(305));
            currentState = a;
            currentPlayer = a.Players[0];
            gameStateToInfoText();
            playerStateToInfoText();
            cardStatesToButtonText();
        }

        private CardState randomCard(int id) {
            var card = new CardState();
            card.ID = id;
            card.CardType = (int) CardTypeUtils.RandomCardType();
            return card;
        }
    }
}