using System;
using System.Collections.Generic;
using Model.State;
using UnityEngine;
using UnityEngine.Events;
using Utils;

namespace Game.UI {
	public class CardTrayBehaviour : MonoBehaviour {

		public CardBehaviour CardTemplate;
		public SpriteFactory Skinner;
		public CardTrayBehaviourEvent OnSelected = new CardTrayBehaviourEvent();

		private List<CardBehaviour> cards = new List<CardBehaviour>();
		private List<CardBehaviour> selectedCards = new List<CardBehaviour>();
	
		void Start () {
			CardTemplate.gameObject.SetActive(false); // we just use this object as a template for future cards
		}

		public void SetCards(List<CardState> cardStates) {
			ClearCards();
			cardStates.ForEach(cs => {
				var card = Instantiate(CardTemplate, CardTemplate.transform.parent);
				card.gameObject.SetActive(true);
				cards.Add(card);
				card.SetCard(cs);
			});
		}

		public void ClearCards() {
			cards.ForEach(c => Destroy(c.gameObject));
			cards.Clear();
		}

		public void CardSelected(CardBehaviour card) {
			if (selectedCards.Contains(card)) {
				selectedCards.Remove(card);
				card.SetOrder(0, cards.Count);
			} else {
				selectedCards.Add(card);
			}

			for (int i = 0; i < selectedCards.Count; i++) {
				selectedCards[i].SetOrder(i + 1, cards.Count);
			}
			OnSelected.Invoke(selectedCards);
		}

		public void ClearSelectedCards() {
			selectedCards.ForEach(c => c.SetOrder(0, cards.Count));
			selectedCards.Clear();
		}
	
		[Serializable]
		public class CardTrayBehaviourEvent : UnityEvent<List<CardBehaviour>> { }    
	}
}
