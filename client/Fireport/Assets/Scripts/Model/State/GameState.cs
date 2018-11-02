using System.Collections.Generic;

namespace Model.State {
    [System.Serializable]
    public class GameState {
        public int Turn;
        public long Created;
        public long Updated;
        public int IDCounter;
        public List<PlayerState> Players;
        public int BoardWidth;
        public int BoardHeight;
        public List<BoardSpace> BoardSpaces;
        public bool IsGameFinished;
        public int Winner;
    }
}