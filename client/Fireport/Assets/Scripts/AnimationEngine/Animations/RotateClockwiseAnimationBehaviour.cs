using UnityEngine;

namespace AnimationEngine.Animations {
    public class RotateClockwiseAnimationBehaviour : AnimationActionBehaviour {
        private float target;
        private float source;
        private float diff;
        
        public override void Play() {
            if (!IsPlaying) {
                IsPlaying = true;
                OnPlay.Invoke(this);
                
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
        }

        void Update() {
            // since this is the default animation, just immediately mark it as finished
            if (IsPlaying && !IsPaused) {
                time += Time.deltaTime;
                if (time > TotalTime) {
                    SetRotation(target);
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                } else {
                    SetRotation(diff * (time / TotalTime) + source);
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