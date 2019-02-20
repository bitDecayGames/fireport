using System;
using Model.State;

namespace Model.Message {
    [Serializable]
    public class TurnResultMessage {
        public string gameID;
        public GameState previousState;
        public GameState currentState;
    }
}