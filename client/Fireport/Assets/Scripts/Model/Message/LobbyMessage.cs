using System;
using System.Collections.Generic;

namespace Model.Message {
    [Serializable]
    public class LobbyMessage {
        public string id;
        public List<string> players;
        public Dictionary<string, bool> readyStatus;

        public override string ToString() {
            return string.Format("LobbyMessage[id: {0} players: {1} readyStatus: {2}]", id, players, readyStatus);
        }
    }
}