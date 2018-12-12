namespace AnimationEngine.Animations {
    public class DefaultAnimationBehaviour : AnimationActionBehaviour {
        public override void Play() {
            if (!IsPlaying) {
                IsPlaying = true;
                OnPlay.Invoke(this);
            }
        }

        public override void Pause() {
            if (!IsPaused) {
                IsPaused = true;
            }
        }

        public override void UnPause() {
            if (IsPaused) {
                IsPaused = false;
            }
        }

        public override void Stop() {
            if (IsPlaying) {
                IsPlaying = false;
                OnStop.Invoke(this);
            }
        }

        void Update() {
            // since this is the default animation, just immediately mark it as finished
            if (IsPlaying) {
                IsPlaying = false;
                OnFinished.Invoke(this);
            }
        }
    }
}