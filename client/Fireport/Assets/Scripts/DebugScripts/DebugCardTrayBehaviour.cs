using System.Collections.Generic;
using Game.UI;
using Model.State;
using UnityEngine;
using Utils;

namespace DebugScripts {
    public class DebugCardTrayBehaviour : MonoBehaviour {
        private CardTrayBehaviour cardTray;
        
        void Start() {
            cardTray = FindObjectOfType<CardTrayBehaviour>();
            
            List<CardState> cardStates = new List<CardState>();
            cardStates.Add(new CardState(0, (int) CardType.TurnRight));
            cardStates.Add(new CardState(1, (int) CardType.TurnLeft));
            cardStates.Add(new CardState(2, (int) CardType.Rotate180));
            cardStates.Add(new CardState(3, (int) CardType.MoveForwardTwo));
            cardStates.Add(new CardState(4, (int) CardType.MoveForwardThree));
            cardStates.Add(new CardState(5, (int) CardType.MoveBackwardOne));
            cardTray.SetCards(cardStates);
        }
    }
}