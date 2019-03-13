using UnityEngine;
using System.Collections;

namespace AnimationEngine.Animations
{
    public class FireCannonAnimationBehavior : AnimationActionBehaviour
    {
        private Color original;
        private Color target = Color.blue;

        protected override void InternalPlay()
        {
            if (GamePiece != null)
            {
                var spriteRenderer = GamePiece.gameObject.GetComponentInChildren<SpriteRenderer>();
                original = spriteRenderer.color;
                spriteRenderer.color = target;
                time = 0;
            }
            else
            {
                Stop();
            }
        }

        void Update()
        {
            if (IsRunning)
            {
                AddDeltaTimeToTimeTracker();
                if (IsTimeGreaterThanTotalTime)
                {
                    var spriteRenderer = GamePiece.gameObject.GetComponentInChildren<SpriteRenderer>();
                    spriteRenderer.color = original;
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                }
            }
        }
    }
}
