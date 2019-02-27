using System.Collections.Generic;
using AnimationEngine;
using Game;
using Model.State;
using UnityEngine;

namespace DebugScripts {
    public class DebugAnimationEngineBehaviour : MonoBehaviour {

        private AnimationEngineBehaviour engine;
        
        void Start() {
            engine = gameObject.AddComponent<AnimationEngineBehaviour>();
            
            List<List<AnimationAction>> animations = new List<List<AnimationAction>>();
            animations.Add(new List<AnimationAction>{
                new AnimationAction(0, "MoveForward", 0, 1),
                new AnimationAction(0, "MoveForward", 1, 2),
                new AnimationAction(0, "MoveForward", 2, 3),
            });
            animations.Add(new List<AnimationAction>{
                new AnimationAction(0, "Default", 1, 1)
            });
            animations.Add(new List<AnimationAction>{
                new AnimationAction(0, "MoveForward", 1, 1),
                new AnimationAction(0, "TurnClockwise90", 0, 1),
                new AnimationAction(0, "MoveBackward", 2, 1),
            });
            animations.Add(new List<AnimationAction>{
                new AnimationAction(0, "Default", 1, 1)
            });
            animations.Add(new List<AnimationAction>{
                new AnimationAction(0, "MoveForward", 2, 2),
                new AnimationAction(0, "TurnCounterClockwise90", 1, 1),
                new AnimationAction(0, "MoveForward", 0, 2),
            });

            var pieces = new List<GamePieceBehaviour>();
            pieces.AddRange(FindObjectsOfType<GamePieceBehaviour>());
            
            engine.Play(animations, pieces);
        }
    }
}