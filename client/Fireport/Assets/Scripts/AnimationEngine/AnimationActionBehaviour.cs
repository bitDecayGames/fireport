using System;
using Game;
using UnityEngine;
using UnityEngine.Events;

namespace AnimationEngine {
    /// <summary>
    /// Describes a single unit of Animation.  It contains some logic to play, stop, pause and fire callbacks on those events.
    /// </summary>
    public abstract class AnimationActionBehaviour : MonoBehaviour {
        public AnimationActionEvent OnPlay = new AnimationActionEvent();
        public AnimationActionEvent OnStop = new AnimationActionEvent();
        public AnimationActionEvent OnFinished = new AnimationActionEvent();

        private GamePieceBehaviour _gamePiece;
        public GamePieceBehaviour GamePiece {
            get {
                if (_gamePiece == null) _gamePiece = GetComponent<GamePieceBehaviour>();
                return _gamePiece;
            }
        }

        private bool _isPlaying;

        public bool IsPlaying {
            get { return _isPlaying; }
            protected set { _isPlaying = value; }
        }
        
        /// <summary>
        /// IsPlaying && !IsPaused
        /// </summary>
        public bool IsRunning {
            get { return IsPlaying && !IsPaused; }
        }

        private bool _isPaused;

        public bool IsPaused {
            get { return _isPaused; }
            protected set { _isPaused = value; }
        }

        private float _totalTime;

        /// <summary>
        /// The total time this animation should run for before calling OnFinished.
        /// </summary>
        public float TotalTime {
            get { return _totalTime; }
            set { _totalTime = value; }
        }
        /// <summary>
        /// A convenient way to track how much time this animation has been playing for.
        /// </summary>
        protected float time;

        /// <summary>
        /// time += Time.deltaTime;
        /// </summary>
        protected void AddDeltaTimeToTimeTracker() {
            time += Time.deltaTime;
        }

        /// <summary>
        /// time / TotalTime
        /// </summary>
        protected float TimeRatio {
            get {
                if (_totalTime == 0) return 0;
                return time / _totalTime;
            }
        }

        /// <summary>
        /// time > TotalTime
        /// </summary>
        protected bool IsTimeGreaterThanTotalTime {
            get { return time > _totalTime; }
        }

        /// <summary>
        /// Start the animation.  It will continue to play the animation until it is finished.
        /// </summary>
        public void Play() {
            if (!IsPlaying) {
                IsPlaying = true;
                OnPlay.Invoke(this);
                time = 0;
                InternalPlay();
            }
        }

        /// <summary>
        /// This is the abstract play method so that you don't have to remember to emit the OnPlay event every time
        /// </summary>
        protected abstract void InternalPlay();
        
        /// <summary>
        /// Pause the animation.  This animation can be resumed from where it was paused by UnPausing it.
        /// </summary>
        public virtual void Pause() {
            if (!IsPaused) {
                IsPaused = true;
            }
        }

        /// <summary>
        /// UnPause the animation.  This animation will resume from where it was when paused.
        /// </summary>
        public virtual void UnPause() {
            if (IsPaused) {
                IsPaused = false;
            }
        }

        /// <summary>
        /// Stop the animation.  This animation will now start from the beginning if it is started again.
        /// </summary>
        public virtual void Stop() {
            if (IsPlaying) {
                IsPlaying = false;
                OnStop.Invoke(this);
            }
        }

        [Serializable]
        public class AnimationActionEvent : UnityEvent<AnimationActionBehaviour> {}
    }
}