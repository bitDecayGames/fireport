namespace AnimationEngine.Animations {
    public class DefaultAnimationBehaviour : AnimationActionBehaviour {
        protected override void InternalPlay() {}

        void Update() {
            if (IsRunning) {
                AddDeltaTimeToTimeTracker();
                if (IsTimeGreaterThanTotalTime) {
                    time = 0;
                    IsPlaying = false;
                    OnFinished.Invoke(this);
                }
            }
        }
    }
}