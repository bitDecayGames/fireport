using System.Collections.Generic;

namespace Model.Message {
    [System.Serializable]
    public class LobbyMessage {
        public string id;
        public List<string> players;
        public Dictionary<string, bool> readyStatus;
    }
}