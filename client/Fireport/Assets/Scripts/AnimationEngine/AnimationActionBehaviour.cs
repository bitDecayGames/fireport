using System;
using Game;
using UnityEngine;
using UnityEngine.Events;

namespace AnimationEngine {
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

        public abstract void Play(); // play is abstract because we almost always want to override its functionality
        public virtual void Pause() {
            if (!IsPaused) {
                IsPaused = true;
            }
        }

        public virtual void UnPause() {
            if (IsPaused) {
                IsPaused = false;
            }
        }

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