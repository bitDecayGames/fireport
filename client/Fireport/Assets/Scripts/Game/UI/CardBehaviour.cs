using System;
using Model.State;
using TMPro;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;
using Utils;

namespace Game.UI {
    public class CardBehaviour : MonoBehaviour {

        public Transform CardTransform;
        public Image CardImage;
        public TextMeshProUGUI OrderText;
        public TextMeshProUGUI TitleText;
        public TextMeshProUGUI DetailsText;
        public SpriteFactory Skinner;

        public CardBehaviourEvent OnSelected = new CardBehaviourEvent();
        public CardState Card { get; private set; }
        public bool IsSelected { get; private set; }

        private const float HEIGHT_SCALAR = 10.0f;

        void Start() {
            SetOrder(0, 0);
        }
        
        public void SetCard(CardState card) {
            Card = card;
            
            // TODO: MW Pick the card image based on the card.cardType group
            Sprite sprite = null;
            switch (CardTypeUtils.GroupFromCardType(card.CardType)) {
                case CardTypeGroup.Unknown:
                    sprite = Skinner.Get("Unknown").sprite;
                    break;
                case CardTypeGroup.Movement:
                    sprite = Skinner.Get("Movement").sprite;
                    break;
                case CardTypeGroup.Utility:
                    sprite = Skinner.Get("Utility").sprite;
                    break;
                case CardTypeGroup.Attack:
                    sprite = Skinner.Get("Attack").sprite;
                    break;
                default:
                    sprite = Skinner.Get("Unknown").sprite;
                    break;
            }

            CardImage.sprite = sprite;

            TitleText.text = CardTypeUtils.CardTypeName(card.CardType);
            // TODO: MW Pick the DetailsText based on card.cardType
            DetailsText.text = "Details";
        }

        /// <summary>
        /// Sets the selected order of the card, with <= 0 being none
        /// </summary>
        /// <param name="orderNumber"></param>
        public void SetOrder(int orderNumber, int totalNumber) {
            if (orderNumber < 0) orderNumber = 0;
            IsSelected = orderNumber > 0;
            if (orderNumber > 0) {
                // TODO: MW convert this from 1, 2, 3 to 1st, 2nd, 3rd...
                OrderText.text = "" + orderNumber;
                SetHeight((totalNumber - orderNumber) * HEIGHT_SCALAR);
            } else {
                OrderText.text = "";
                SetHeight(0);
            }
            
        }

        private void SetHeight(float amount) {
            var pos = CardTransform.localPosition;
            pos.y = amount;
            CardTransform.localPosition = pos;
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