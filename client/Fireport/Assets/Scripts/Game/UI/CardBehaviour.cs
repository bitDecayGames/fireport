using System;
using Model.State;
using TMPro;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

namespace Game.UI {
    public class CardBehaviour : MonoBehaviour {

        public Image CardImage;
        public TextMeshProUGUI OrderText;
        public TextMeshProUGUI TitleText;
        public TextMeshProUGUI DetailsText;

        public CardBehaviourEvent OnSelected = new CardBehaviourEvent();
        
        public CardState Card { get; private set; }
        public bool IsSelected { get; private set; }
        
        public void SetCard(CardState card) {
            Card = card;
            
            // TODO: Pick the card image based on the card.cardType
            //CardImage.sprite = // pick a sprite here
            
            // TODO: Pick the TitleText and DetailsText based on card.cardType
            TitleText.text = "Title";
            DetailsText.text = "Details";
        }

        /// <summary>
        /// Sets the selected order of the card, with <= 0 being none
        /// </summary>
        /// <param name="orderNumber"></param>
        public void SetOrder(int orderNumber) {
            IsSelected = orderNumber > 0;
            if (orderNumber > 0) OrderText.text = "" + orderNumber;
            else OrderText.text = "";
        }

        public void ClickedTraySlot() {
            OnSelected.Invoke(this);
        }
        
        public void ClickedCard() {
            OnSelected.Invoke(this);
        }

        [Serializable]
        public class CardBehaviourEvent : UnityEvent<CardBehaviour> { }    
    }
}