using System;
using System.Collections.Generic;
using UnityEngine;

namespace Game {
    [CreateAssetMenu(menuName = "ScriptableObjects/GamePieceFactory")]
    public class GamePieceFactory : ScriptableObject {
        public List<GamePieceFactorySpec> pieces = new List<GamePieceFactorySpec>();

        public GamePieceBehaviour Build(string specName, Transform parent = null) {
            var piece = pieces.Find(p => p.name == specName);
            if (piece != null) {
                if (parent == null) return Instantiate(piece.prefab);
                var gp = Instantiate(piece.prefab, parent);
                var t = gp.transform;
                t.localPosition = new Vector3();
                t.localScale = new Vector3(1, 1, 1);
                t.localRotation = Quaternion.Euler(0, 0, 0);
                return gp;
            }
            Debug.LogError("Failed to find game piece factory spec that matched the name: " + specName);
            return null;
        }
        
        [Serializable]
        public class GamePieceFactorySpec {
            public string name;
            public GamePieceBehaviour prefab;
        }
    }
}