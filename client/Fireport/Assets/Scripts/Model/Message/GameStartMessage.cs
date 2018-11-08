using System.Collections.Generic;
using Model.State;

namespace Model.Message {
    [System.Serializable]
    public class GameStartMessage {
        public string gameID;
        public List<string> players;
        public GameState gameState;
    }
}