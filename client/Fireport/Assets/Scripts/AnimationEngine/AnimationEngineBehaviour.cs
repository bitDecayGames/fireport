using System;
using System.Collections.Generic;
using Model.State;
using UnityEngine;

namespace AnimationEngine {
    public class AnimationEngineBehaviour : MonoBehaviour {
        private List<List<AnimationAction>> animations = null;
        private List<List<AnimationActionBehaviour>> behaviours = null;
        private int currentGroup;
        private bool isRunningGroup;
        private bool isPaused;

        void Update() {
            if (isRunningGroup && IsCurrentGroupFinished()) {
                PlayNextGroup();
            }
        }

        public void Play(List<List<AnimationAction>> animations) {
            this.animations = animations;
            if (behaviours != null) {
                behaviours.ForEach(group => {
                    if (group != null) group.ForEach(a => {
                        if (a != null) Destroy(a);
                    });
                });
                behaviours.Clear();
            }
            
            if (this.animations != null && this.animations.Count > 0) {
                this.animations.ForEach(group => {
                    if (group != null && group.Count > 0) {
                        var behaviourGroup = new List<AnimationActionBehaviour>();
                        behaviours.Add(behaviourGroup);
                        group.ForEach(a => {
                            var animationAction = AnimationActionFactory.AddComponentByName(a.Name, GetGameObjectWithId(a.Owner)); 
                            behaviourGroup.Add(animationAction);
                        });
                    }
                });
                currentGroup = -1;
                PlayNextGroup();
            }
        }

        private void PlayNextGroup() {
            currentGroup++;
            if (currentGroup < 0) throw new Exception("Current Group index out of bounds");
            if (currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.Play());
                isRunningGroup = true;
            } else isRunningGroup = false;
        }

        private bool IsCurrentGroupFinished() {
            // check if there exists any action that is currently still playing, if so, the group is not finished yet
            return !behaviours[currentGroup].Exists(a => a.IsPlaying);
        }

        private Transform GetGameObjectWithId(int id) {
            // TODO: MW this probably needs to call some sort of game manager that is mapping the state to Unity objects
            return transform;
        }

        public void Pause() {
            if (!isPaused && currentGroup >= 0 && currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.Pause());
                isPaused = true;
            }
        }

        public void UnPause() {
            if (isPaused && currentGroup >= 0 && currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.UnPause());
                isPaused = false;
            }
        }

        public void Stop() {
            if (currentGroup >= 0 && currentGroup < behaviours.Count) {
                behaviours[currentGroup].ForEach(a => a.Stop());
                isPaused = false;
                isRunningGroup = false;
            }
        }
    }
}