using System;
using System.Collections.Generic;
using Model.State;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;
using Utils;

namespace Game.UI {
	public class CardTrayBehaviour : MonoBehaviour {

		public CardBehaviour CardTemplate;
		public Button Toggle;
		public SpriteFactory Skinner;
		public CardTrayBehaviourEvent OnSelected = new CardTrayBehaviourEvent();

		private List<CardBehaviour> cards = new List<CardBehaviour>();
		private List<CardBehaviour> selectedCards = new List<CardBehaviour>();

		private SlideUpFromBottom Slider;
		
		public bool Enabled {
			get { return Slider.IsShown; }
		}
	
		void Start () {
			CardTemplate.gameObject.SetActive(false); // we just use this object as a template for future cards
			Slider = GetComponent<SlideUpFromBottom>();
			Slider.Hide(false);
			Toggle.onClick.AddListener(() => ToggleShow());
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
			selectedCards.Clear();
			SendSelectedCards();
		}

		public void CardSelected(CardBehaviour card) {
			if (Enabled) {
				if (selectedCards.Contains(card)) {
					selectedCards.Remove(card);
					card.SetOrder(0, cards.Count);
				} else {
					selectedCards.Add(card);
				}

				for (int i = 0; i < selectedCards.Count; i++) {
					selectedCards[i].SetOrder(i + 1, cards.Count);
				}

				SendSelectedCards();
			}
		}

		public void ClearSelectedCards() {
			selectedCards.ForEach(c => c.SetOrder(0, cards.Count));
			selectedCards.Clear();
			SendSelectedCards();
		}

		public void ToggleShow() {
			// TODO: MW do something visually with the button here
			if (Slider.IsShown) {
				Hide();
			} else Show();
		}

		public void Hide() {
			if (Slider.IsShown) {
				Slider.Hide();
				ClearSelectedCards();
			}
		}

		public void Show() {
			if (cards != null && cards.Count > 0) Slider.Show();
		}

		private void SendSelectedCards() {
			var tmp = new List<CardBehaviour>();
			tmp.AddRange(selectedCards);
			OnSelected.Invoke(tmp);
		}
	
		[Serializable]
		public class CardTrayBehaviourEvent : UnityEvent<List<CardBehaviour>> { }    
	}
}
