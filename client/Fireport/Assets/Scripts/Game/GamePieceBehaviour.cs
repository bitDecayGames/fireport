using TMPro;
using UnityEngine;

namespace Game {
    /// <summary>
    /// Identifies this game object as a piece of the game board, whether that be the players, a card, a mountain, etc.
    ///
    /// Anything that has an ID on the GameState should be identified in the world with this component.
    /// </summary>
    public class GamePieceBehaviour : MonoBehaviour {
        public int Id;

        private bool debug = true;

        void Start() {
            if (debug) {
                var go = new GameObject("GamePieceIdDebugLabel");
                var t = go.transform;
                t.parent = transform;
                t.localPosition = new Vector3();
                var tmp = go.AddComponent<TextMeshPro>();
                tmp.text = "" + Id;
                tmp.fontSize = 4;
                tmp.alignment = TextAlignmentOptions.Center;
                tmp.sortingOrder = GetComponentInChildren<Renderer>().sortingOrder + 1;
            }
        }
    }
}