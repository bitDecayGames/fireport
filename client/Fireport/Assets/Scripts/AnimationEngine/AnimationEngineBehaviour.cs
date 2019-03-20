using System;
using System.Collections.Generic;
using Game;
using Model.State;
using UnityEngine;
using UnityEngine.Events;

namespace AnimationEngine {
    /// <summary>
    /// This is the manager class for all of a game's animation sequences.  There should only be one of these
    /// on the scene at a time.
    /// </summary>
    public class AnimationEngineBehaviour : MonoBehaviour {
        private List<List<AnimationAction>> animations = null;
        private List<List<AnimationActionBehaviour>> behaviours = null;
        private List<GamePieceBehaviour> gamePieces;
        private int currentGroup;
        private bool isRunningGroup;
        private bool isPaused;

        public UnityEvent OnComplete = new UnityEvent();

        void Update() {
            if (isRunningGroup && IsCurrentGroupFinished()) {
                PlayNextGroup();
            }
        }

        /// <summary>
        /// Given a list of lists of animation actions with a list of all of the game pieces create the animation behaviours
        /// and run through each animation sequence.
        /// </summary>
        /// <param name="animations">the animation metadata</param>
        /// <param name="gamePieces">all of the game pieces relevant to these animation actions</param>
        public void Play(List<List<AnimationAction>> animations, List<GamePieceBehaviour> gamePieces) {
            
            this.gamePieces = gamePieces;
            this.animations = animations;
            if (behaviours != null) {
                behaviours.ForEach(group => {
                    if (group != null) group.ForEach(a => {
                        if (a != null) Destroy(a);
                    });
                });
                behaviours.Clear();
            } else behaviours = new List<List<AnimationActionBehaviour>>();
            
            if (this.animations != null && this.animations.Count > 0) {
                this.animations.ForEach(group => {
                    if (group != null && group.Count > 0) {
                        var behaviourGroup = new List<AnimationActionBehaviour>();
                        behaviours.Add(behaviourGroup);
                        group.ForEach(a => {
                            var animationAction = AnimationActionFactory.AddComponentByAnimationAction(a, GetGameObjectWithId(a.Owner)); 
                            if (animationAction != null) behaviourGroup.AddRange(animationAction);
                        });
                    }
                });
                currentGroup = -1;
                PlayNextGroup();
            } else OnComplete.Invoke();
        }

        private void PlayNextGroup() {
            currentGroup++;
            if (currentGroup < 0) throw new Exception("Current Group index cannot be less than 0: " + currentGroup);
            if (currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.Play());
                isRunningGroup = true;
            } else {
                isRunningGroup = false;
                OnComplete.Invoke();
            }
        }

        private bool IsCurrentGroupFinished() {
            // check if there exists any action that is currently still playing, if so, the group is not finished yet
            return !behaviours[currentGroup].Exists(a => a.IsPlaying);
        }

        private GamePieceBehaviour GetGameObjectWithId(int id) {
            GamePieceBehaviour obj = null;
            if (gamePieces != null) obj = gamePieces.Find(g => g.Id == id);
            return obj;
        }

        /// <summary>
        /// Pause the currently running animation group
        /// </summary>
        public void Pause() {
            if (!isPaused && currentGroup >= 0 && currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.Pause());
                isPaused = true;
            }
        }

        /// <summary>
        /// UnPause the currently running animation group
        /// </summary>
        public void UnPause() {
            if (isPaused && currentGroup >= 0 && currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.UnPause());
                isPaused = false;
            }
        }

        /// <summary>
        /// Stops not just the current animation group, but the entire animation sequence.  There is no way
        /// to resume from a stopped animation sequence.  If you intend to start the sequence back up again,
        /// you should use the Pause and UnPause methods.
        /// </summary>
        public void Stop() {
            if (currentGroup >= 0 && currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.Stop());
                isPaused = false;
                isRunningGroup = false;
                OnComplete.Invoke();
            }
        }
    }
}