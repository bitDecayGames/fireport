using System;
using UnityEngine;

namespace Model.Message {
    [Serializable]
    public class GameInputMessage {
        public int cardID;
        public int owner;
        public int order;
        public int swap;

        public override string ToString() {
            return JsonUtility.ToJson(this);
        }
    }
}