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
    }
}