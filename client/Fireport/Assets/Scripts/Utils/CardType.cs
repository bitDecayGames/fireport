using System;

namespace Utils {
    public enum CardType {
        Unknown = -1,
        SkipTurn = 0,
        MoveForwardOne = 100,
        MoveForwardTwo = 101,
        MoveForwardThree = 102,
        MoveBackwardOne = 103,
        MoveBackwardTwo = 104,
        MoveBackwardThree = 105,
        RotateRight = 110,
        RotateLeft = 111,
        Rotate180 = 112,
        TurnRight = 120,
        TurnLeft = 121
    }
    
    public static class CardTypeUtils {
        private static Random random = new Random();
        
        public static CardType RandomCardType() {
            Array values = Enum.GetValues(typeof(CardType));
            return (CardType)values.GetValue(random.Next(values.Length));
        }
    }
}