using UnityEngine;

namespace AnimationEngine.Animations {
    public class RotateClockwiseAnimationBehaviour : AnimationActionBehaviour {
        private float target;
        private float source;
        private float diff;
        
        protected override void InternalPlay() {
            if (GamePiece != null) {
                // TODO: MW start some sort of sprite rotating animation...
                source = transform.localRotation.eulerAngles.z;
                target = source - 90;
                diff = target - source;
                time = 0;
            } else {
                Stop();
            }
        }

        void Update() {
            if (IsRunning) {
                AddDeltaTimeToTimeTracker();
                if (IsTimeGreaterThanTotalTime) {
                    SetRotation(target);
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                } else {
                    SetRotation(diff * TimeRatio + source);
                }
            }
        }

        private void SetRotation(float zRotation) {
            var rot = transform.localRotation.eulerAngles;
            rot.z = zRotation;
            transform.localRotation = Quaternion.Euler(rot);
        }
    }
}