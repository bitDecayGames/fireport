using AnimationEngine.Animations;
using Game;
using Model.State;
using System.Collections.Generic;
using UnityEngine;

namespace AnimationEngine {
    /// <summary>
    /// Builds every animation action in the game
    /// </summary>
    public static class AnimationActionFactory {

        /// <summary>
        /// Create a new animation action component on the given game piece with the given animation action object.
        /// This can also add special logic to the creation of a component based on the given game piece.
        /// </summary>
        /// <param name="animationAction">the animation action metadata</param>
        /// <param name="gamePiece">the game piece to apply this new animation action behaviour to</param>
        /// <returns>the newly created animation action behaviour component</returns>
        public static List<AnimationActionBehaviour> AddComponentByAnimationAction(AnimationAction animationAction, GamePieceBehaviour gamePiece) {
            List<AnimationActionBehaviour> behaviours = new List<AnimationActionBehaviour>();
            if (gamePiece != null && animationAction != null) {
                switch (animationAction.Name) {
                    // TODO: MW these names are just placeholders, use what ever is in the actual server code
                    case "MoveForward":
                        behaviours.Add(gamePiece.gameObject.AddComponent<MoveForwardAnimationBehaviour>());
                        break;
                    case "MoveBackward":
                        behaviours.Add(gamePiece.gameObject.AddComponent<MoveBackwardAnimationBehaviour>());
                        break;
                    case "TurnClockwise90":
                        behaviours.Add(gamePiece.gameObject.AddComponent<RotateClockwiseAnimationBehaviour>());
                        break;
                    case "TurnCounterClockwise90":
                        behaviours.Add(gamePiece.gameObject.AddComponent<RotateCounterClockwiseAnimationBehaviour>());
                        break;
                    case "DoBumpInto":
                        behaviours.Add(gamePiece.gameObject.AddComponent<TurnRedAnimationBehaviour>());
                        behaviours.Add(gamePiece.gameObject.AddComponent<DecrementHealthBarBehaviour>());
                        break;
                    case "HitByCanon":
                        behaviours.Add(gamePiece.gameObject.AddComponent<TurnRedAnimationBehaviour>());
                        behaviours.Add(gamePiece.gameObject.AddComponent<DecrementHealthBarBehaviour>());
                        break;
                    case "FireCanon":
                        behaviours.Add(gamePiece.gameObject.AddComponent<FireCannonAnimationBehavior>());
                        break;
                    case "Default":
                        behaviours.Add(gamePiece.gameObject.AddComponent<DefaultAnimationBehaviour>());
                        break;
                    default:
                        // TODO: MW maybe we want to throw here? but I suspect not for stability's sake
                        Debug.LogWarning("Failed to find an animation behaviour factory method for animation: " + animationAction.Name);
                        behaviours.Add(gamePiece.gameObject.AddComponent<DefaultAnimationBehaviour>());
                        break;
                }

                foreach (AnimationActionBehaviour behaviour in behaviours)
                {
                    behaviour.TotalTime = animationAction.Time;
                }
            }

            return behaviours;
        }
    }
}