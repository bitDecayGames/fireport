using Model.State;
using UnityEngine;

namespace Game {
    public class GameBoardBehaviour : MonoBehaviour {

        public GamePieceFactory Factory;
        public float GridSizeMultiplier = 1;

        private GameObject board;
        
        public void Populate(GameState state) {
            if (board != null) Destroy(board);
            board = new GameObject("board");
            board.transform.parent = transform;
            board.transform.localPosition = new Vector3();
            
            for (int x = 0; x < state.BoardWidth; x++) {
                for (int y = 0; y < state.BoardHeight; y++) {
                    var boardSpace = state.BoardSpaces[x + y * state.BoardWidth];
                    var boardSpaceBehaviour = Factory.Build("BoardSpace", board.transform);
                    boardSpaceBehaviour.Id = boardSpace.ID;
                    boardSpaceBehaviour.transform.localPosition = new Vector3(x * GridSizeMultiplier, -y * GridSizeMultiplier, 0);
                }
            }
            
            state.Players.ForEach(playerState => {
                var playerBehaviour = Factory.Build("Player", board.transform);
                playerBehaviour.Id = playerState.ID;
                playerBehaviour.transform.localPosition = new Vector3(playerState.Location % state.BoardWidth * GridSizeMultiplier, playerState.Location / state.BoardWidth * -GridSizeMultiplier, 0);
                playerBehaviour.transform.localRotation = facingToRotation(playerState.Facing);
            });
            
            // TODO: MW set the position of the camera to center on the middle? Or maybe on the current player?
        }

        private Quaternion facingToRotation(int facing) {
            var rot = 0f;
            switch (facing) {
                case 0:
                    rot = 0;
                    break;
                case 1:
                    rot = 90;
                    break;
                case 2:
                    rot = 180;
                    break;
                case 3:
                    rot = 270;
                    break;
            }

            return Quaternion.Euler(0, 0, rot);
        }
    }
}