using UnityEngine;

namespace AnimationEngine.Animations {
    public class DefaultAnimationBehaviour : AnimationActionBehaviour {
        public override void Play() {
            if (!IsPlaying) {
                IsPlaying = true;
                OnPlay.Invoke(this);
                time = 0;
            }
        }

        void Update() {
            if (IsPlaying && !IsPaused) {
                time += Time.deltaTime;
                if (time > TotalTime) {
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                }
            }
        }
    }
}