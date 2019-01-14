using UnityEngine;

namespace Game {
    /// <summary>
    /// Identifies this game object as a piece of the game board, whether that be the players, a card, a mountain, etc.
    ///
    /// Anything that has an ID on the GameState should be identified in the world with this component.
    /// </summary>
    public class GamePieceBehaviour : MonoBehaviour {
        public int Id;
    }
}