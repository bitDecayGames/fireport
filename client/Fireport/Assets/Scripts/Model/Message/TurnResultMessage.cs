using System;
using System.Collections.Generic;
using Model.State;
using Network;
using UnityEngine;

namespace Model.Message {
    [Serializable]
    public class TurnResultMessage {
        public string gameID;
        public GameState previousState;
        public GameState currentState;

        public static TurnResultMessage FromJson(string json) {
            var msg = JsonUtility.FromJson<TurnResultMessage>(json);
            var data = Json.Deserialize(json) as Dictionary<string, object>;
            // TODO: MW holy shit this is ugly... but JsonUtility can't handle List<List<T>> so... ya...
            msg.previousState.Animations = GetAnimationsFromState(data, "previousState");
            msg.currentState.Animations = GetAnimationsFromState(data, "currentState");
            return msg;
        }

        private static List<List<AnimationAction>> GetAnimationsFromState(Dictionary<string, object> data, string key) {
            var animations = new List<List<AnimationAction>>();
            var state = data[key] as Dictionary<string, object>;
            if (state != null) {
                var outerList = state["Animations"] as List<object>;
                if (outerList != null) {
                    animations.AddRange(outerList.ConvertAll(o => {
                        var innerList = o as List<object>;
                        if (innerList != null) return innerList.ConvertAll(a => a as Dictionary<string, object>).ConvertAll(d => AnimationAction.FromDictionary(d));
                        return new List<AnimationAction>();
                    }));
                }
            }
            return animations;
        }
    }
}