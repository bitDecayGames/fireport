using System;
using Model.State;

namespace Model.Message {
    [Serializable]
    public class CurrentStateMessage {
        public string gameID;
        public GameState currentState;
    }
}