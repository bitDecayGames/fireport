using System.Collections.Generic;

namespace Model.State {
    [System.Serializable]
    public class PlayerState {
        public int ID;
        public string Name;
        public List<CardState> Hand;
        public List<CardState> Discard;
        public List<CardState> Deck;
        /// <summary>
        /// The index of the board space the player is standing on
        /// </summary>
        public int Location;
        /// <summary>
        /// 0 -> North, 1 -> East, 2 -> South, 3 -> West
        /// </summary>
        public int Facing;
        public int Health;
    }
}