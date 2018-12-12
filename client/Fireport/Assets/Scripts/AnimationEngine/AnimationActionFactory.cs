using AnimationEngine.Animations;
using Game;
using Model.State;

namespace AnimationEngine {
    public static class AnimationActionFactory {

        public static AnimationActionBehaviour AddComponentByAnimationAction(AnimationAction animationAction, GamePieceBehaviour gamePiece) {
            AnimationActionBehaviour behaviour = null;
            if (gamePiece != null && animationAction != null) {
                switch (animationAction.Name) {
                    // TODO: MW these names are just placeholders, use what ever is in the actual server code
                    case "MoveForward":
                        behaviour = gamePiece.gameObject.AddComponent<MoveForwardAnimationBehaviour>();
                        break;
                    case "MoveBackward":
                        behaviour = gamePiece.gameObject.AddComponent<MoveBackwardAnimationBehaviour>();
                        break;
                    case "RotateClockwise":
                        behaviour = gamePiece.gameObject.AddComponent<RotateClockwiseAnimationBehaviour>();
                        break;
                    case "RotateCounterClockwise":
                        behaviour = gamePiece.gameObject.AddComponent<RotateCounterClockwiseAnimationBehaviour>();
                        break;
                    case "Default":
                        behaviour = gamePiece.gameObject.AddComponent<DefaultAnimationBehaviour>();
                        break;
                    default:
                        // TODO: MW maybe we want to throw here? but I suspect not for stability's sake
                        behaviour = gamePiece.gameObject.AddComponent<DefaultAnimationBehaviour>();
                        break;
                }

                if (behaviour != null) behaviour.TotalTime = animationAction.Time;
            }

            return behaviour;
        }
    }
}