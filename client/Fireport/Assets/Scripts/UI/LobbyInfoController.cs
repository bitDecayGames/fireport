using System.Collections;
using System.Collections.Generic;
using Boo.Lang.Runtime;
using Model.Message;
using UnityEngine;
using UnityEngine.Networking;


public class LobbyInfoController : MonoBehaviour
{
	public static string objectName = "LobbyInfo";
	public LobbyMessage msg;

	public static LobbyInfoController GetLobbyObject()
	{
		GameObject l = GameObject.Find(objectName);
		if (l == null)
		{
			throw new RuntimeException("No lobby info to work with");
		}
		return l.GetComponent<LobbyInfoController>();
	}

	public static void ClearLobbyObject()
	{
		GameObject l = GameObject.Find(objectName);
		if (l != null)
		{
			Destroy(l);
		}
	}
}
