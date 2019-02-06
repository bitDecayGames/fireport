using System.Collections.Generic;
using Game.UI;
using UnityEngine;

public class CardTrayBehaviour : MonoBehaviour {

	public CardBehaviour CardTemplate;

	private List<CardBehaviour> cards;
	
	void Start () {
		CardTemplate.gameObject.SetActive(false); // we just use this object as a template for future cards
	}

	public void CardSelected(CardBehaviour card) {
		// TODO: do something here
	}
}
