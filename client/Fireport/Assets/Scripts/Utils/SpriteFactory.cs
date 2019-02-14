using System;
using System.Collections.Generic;
using UnityEngine;

namespace Utils {
    [CreateAssetMenu(menuName = "ScriptableObjects/SpriteFactory")]
    public class SpriteFactory : ScriptableObject {
        public List<SpriteSpec> sprites = new List<SpriteSpec>();

        public SpriteSpec Get(string specName) {
            if (specName == null) return null;
            return sprites.Find(p => p.name.ToLower() == specName.ToLower());
        }
        
        [Serializable]
        public class SpriteSpec {
            public string name;
            public Sprite sprite;
        }
    }
}