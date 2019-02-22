using System.Collections.Generic;
using Network;
using UnityEngine;

namespace Model.State {
    [System.Serializable]
    public class AnimationAction {
        public int ID;
        public string Name;
        public int Owner;
        public float Time;

        public AnimationAction() {}

        public AnimationAction(int ID, string Name, int Owner, float Time) {
            this.ID = ID;
            this.Name = Name;
            this.Owner = Owner;
            this.Time = Time;
        }

        public override string ToString() {
            return JsonUtility.ToJson(this);
        }

        public static AnimationAction FromDictionary(Dictionary<string, object> data) {
            var a = new AnimationAction();
            if (data.ContainsKey("ID")) a.ID = (int) (long) data["ID"];
            if (data.ContainsKey("Name")) a.Name = (string) data["Name"];
            if (data.ContainsKey("Owner")) a.Owner = (int) (long) data["Owner"];
            if (data.ContainsKey("Time")) a.Time = (float) (double) data["Time"];
            else a.Time = 1f;
            return a;
        }
    }
}