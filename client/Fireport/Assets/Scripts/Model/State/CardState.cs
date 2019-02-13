namespace Model.State {
    [System.Serializable]
    public class CardState {
        public int ID;
        public int CardType;

        public CardState() {
            
        }

        public CardState(int id, int cardType) {
            ID = id;
            CardType = cardType;
        }
    }
}