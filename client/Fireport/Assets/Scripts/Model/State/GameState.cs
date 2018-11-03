using System.Collections.Generic;
using System.Text;
using UnityEditor;

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

        public override string ToString() {
            StringBuilder sb = new StringBuilder();
            sb.AppendLine("Turn: " + Turn);
            for (int h = 0; h < BoardHeight; h++) {
                for (int w = 0; w < BoardWidth; w++) {
                    sb.Append(" ");
                    int i = h * BoardWidth + w;
                    var player = Players.Find(p => p.Location == i && p.Facing == 0);
                    if (player != null) sb.Append("*");
                    else sb.Append(" ");
                    sb.Append(" ");
                }

                sb.AppendLine();
                
                for (int w = 0; w < BoardWidth; w++) {
                    int i = h * BoardWidth + w;
                    var player = Players.Find(p => p.Location == i);
                    if (player != null) {
                        if (player.Facing == 3) sb.Append("*");
                        else sb.Append(" ");
                        sb.Append(playerToCharacter(player));
                        if (player.Facing == 1) sb.Append("*");
                        else sb.Append(" ");
                    }
                    else sb.Append(" . ");
                }

                sb.AppendLine();
                
                for (int w = 0; w < BoardWidth; w++) {
                    sb.Append(" ");
                    int i = h * BoardWidth + w;
                    var player = Players.Find(p => p.Location == i && p.Facing == 2);
                    if (player != null) sb.Append("*");
                    else sb.Append(" ");
                    sb.Append(" ");
                }
                sb.AppendLine();
            }
            return sb.ToString();
        }

        private string playerToCharacter(PlayerState p) {
            if (p != null && p.Name != null)
                return p.Name[0].ToString().ToUpper();
            return "0";
        }
    }
}