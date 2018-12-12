using System;
using UnityEngine;
using UnityEngine.Events;

namespace AnimationEngine {
    public abstract class AnimationActionBehaviour : MonoBehaviour {
        public AnimationActionEvent OnPlay = new AnimationActionEvent();
        public AnimationActionEvent OnStop = new AnimationActionEvent();
        public AnimationActionEvent OnFinished = new AnimationActionEvent();

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

        // TODO: MW maybe make these virtual so you can do things like set IsPlaying or IsPaused with super...
        public abstract void Play();
        public abstract void Pause();
        public abstract void UnPause();
        public abstract void Stop();

        [Serializable]
        public class AnimationActionEvent : UnityEvent<AnimationActionBehaviour> {}
    }
}