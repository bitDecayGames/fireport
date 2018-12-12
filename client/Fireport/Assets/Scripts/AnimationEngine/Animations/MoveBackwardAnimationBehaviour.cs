using UnityEngine;

namespace AnimationEngine.Animations {
    public class MoveBackwardAnimationBehaviour : AnimationActionBehaviour {
        private Vector3 target;
        private Vector3 source;
        private Vector3 diff;
        
        public override void Play() {
            if (!IsPlaying) {
                IsPlaying = true;
                OnPlay.Invoke(this);
                
                if (GamePiece != null) {
                    // TODO: MW start some sort of sprite movement animation...
                    source = transform.localPosition;
                    // TODO: MW if the grid size is greater than 1x1 unit, this line will need to change
                    target = source + transform.localRotation * (new Vector3(0, 1) * -1); // up vector is used as starting point
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
                    transform.localPosition = target;
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                } else {
                    transform.localPosition = diff * (time / TotalTime) + source;
                }
            }
        }
    }
}