
using UnityEngine;

namespace AnimationEngine.Animations
{
    public class DecrementHealthBarBehaviour : AnimationActionBehaviour
    {
        protected override void InternalPlay() {
            decrement();
        }

        void Update()
        {
            if (IsRunning)
            {
                AddDeltaTimeToTimeTracker();
                if (IsTimeGreaterThanTotalTime)
                {
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                }
            }
        }

        void decrement()
        {
            Debug.Log("DECREMENTING");
            var health = GamePiece.gameObject.GetComponentInChildren<HealthBar>();
            if (health != null)
            {
                Debug.Log("notNull");
                health.reduceByOne();
            }
        }

    }
}