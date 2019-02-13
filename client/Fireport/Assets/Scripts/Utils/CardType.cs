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

    public enum CardTypeGroup {
        Unknown = -1,
        Movement = 1,
        Utility = 2,
        Attack = 3
    }
    
    public static class CardTypeUtils {
        private static Random random = new Random();
        
        public static CardType RandomCardType() {
            Array values = Enum.GetValues(typeof(CardType));
            return (CardType)values.GetValue(random.Next(values.Length));
        }

        public static string CardTypeName(int value) {
            var name = ((CardType)value).ToString();
            int outI;
            if (Int32.TryParse(name, out outI)) return "UnknownCardType(" + name + ")";
            return name;
        }
        
        public static CardTypeGroup GroupFromCardType(int cardType) {
            if (cardType < 100) return CardTypeGroup.Unknown;
            if (cardType < 200) return CardTypeGroup.Movement;
            if (cardType < 300) return CardTypeGroup.Utility;
            if (cardType < 400) return CardTypeGroup.Attack;
            return CardTypeGroup.Unknown;
        }
    }
}