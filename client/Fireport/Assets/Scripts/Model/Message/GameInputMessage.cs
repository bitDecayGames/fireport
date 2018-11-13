using System;

namespace Model.Message {
    [Serializable]
    public class GameInputMessage {
        public int cardID;
        public int owner;
        public int order;
        public int swap;
    }
}