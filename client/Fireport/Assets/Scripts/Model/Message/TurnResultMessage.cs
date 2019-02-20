using System.Collections.Generic;
using Model.State;

namespace Model.Message {
    [System.Serializable]
    public class TurnResultMessage {
        public string gameID;
        public GameState previousState;
        public GameState currentState;
        public List<List<AnimationAction>> animations;
    }
}