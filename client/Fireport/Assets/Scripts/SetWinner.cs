using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SetWinner : MonoBehaviour {

	// Use this for initialization
	void Start () { 
		UnityEngine.UI.Text text = GetComponent<UnityEngine.UI.Text>();
		text.text = string.Format("A Winner is: {0}", WinnerRememberer.Winner);
	}
	
	// Update is called once per frame
	void Update () {
		
	}
}
