using System.Collections.Generic;

namespace Model.State {
    [System.Serializable]
    public class PlayerState {
        public int ID;
        public string Name;
        public List<CardState> Hand;
        public List<CardState> Discard;
        public List<CardState> Deck;
        public int Location;
        public int Facing;
        public int Health;
    }
}