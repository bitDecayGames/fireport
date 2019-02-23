using System.Collections.Generic;
using Network;
using UnityEngine;

namespace Model.Message {
    [System.Serializable]
    public class TurnSubmissionMessage {
        public string gameID;
        public string playerID;
        public List<GameInputMessage> inputs = new List<GameInputMessage>();

        public override string ToString() {
            return JsonUtility.ToJson(this);
        }
    }
}