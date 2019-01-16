using System.Collections.Generic;
using Model.State;

namespace DebugScripts {
    public static class DebugHelper {
        public static GameState DebugGameState() {
            var gs = new GameState();
            gs.Turn = 3;
            gs.Winner = -1;
            gs.IsGameFinished = false;
            gs.BoardWidth = 5;
            gs.BoardHeight = 5;
            gs.BoardSpaces = new List<BoardSpace>();
            for (int i = 0; i < gs.BoardWidth * gs.BoardHeight; i++) {
                var bs = new BoardSpace();
                bs.ID = getId(gs);
                gs.BoardSpaces.Add(bs);
            }
            gs.Players = new List<PlayerState>();
            for (int i = 0; i < 4; i++) {
                var ps = new PlayerState();
                ps.ID = getId(gs);
                ps.Name = "Player" + (i + 1);
                ps.Location = (i + 1) * 2;
                ps.Facing = i;
                ps.Deck = new List<CardState>();
                ps.Discard = new List<CardState>();
                ps.Hand = new List<CardState>();
                ps.Health = 10;
                gs.Players.Add(ps);
            }

            return gs;
        }

        private static int getId(GameState gs) {
            var id = gs.IDCounter;
            gs.IDCounter++;
            return id;
        }
    }
}