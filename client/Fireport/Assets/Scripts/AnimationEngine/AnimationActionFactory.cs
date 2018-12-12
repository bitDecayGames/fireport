using AnimationEngine.Animations;
using UnityEngine;

namespace AnimationEngine {
    public static class AnimationActionFactory {

        public static AnimationActionBehaviour AddComponentByName(string animationName, Transform obj) {
            switch (animationName) {
                case "Default":
                    return obj.gameObject.AddComponent<DefaultAnimationBehaviour>();
                default:
                    // TODO: MW maybe we want to throw here? but I suspect not for stability's sake
                    return obj.gameObject.AddComponent<DefaultAnimationBehaviour>();
            }
        }
    }
}