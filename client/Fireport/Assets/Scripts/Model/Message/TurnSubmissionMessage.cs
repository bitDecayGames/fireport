using System.Collections.Generic;

namespace Model.Message {
    [System.Serializable]
    public class TurnSubmissionMessage {
        public string gameID;
        public string playerID;
        public List<GameInputMessage> inputs = new List<GameInputMessage>();
    }
}