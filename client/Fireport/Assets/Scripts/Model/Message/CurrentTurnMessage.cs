using System;

namespace Model.Message {
    [Serializable]
    public class CurrentTurnMessage {
        public string gameID;
        public int currentTurn;
    }
}