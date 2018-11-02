using System.Collections.Generic;

namespace Model.Message {
    [System.Serializable]
    public class GameStartMessage {
        public string gameID;
        public List<string> players;
        public string message;
    }
}