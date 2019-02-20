using Model.Message;
using UnityEngine;

public class LobbyInfoController : MonoBehaviour
{
	public static string objectName = "LobbyInfo";
	private static LobbyInfoController instance = null;

	public static LobbyInfoController Instance() {
		if (!instance) {
			var go = new GameObject(objectName);
			instance = go.AddComponent<LobbyInfoController>();
		}

		return instance;
	}

	void Start() {
		DontDestroyOnLoad(gameObject);
	}

	public string playerName;
	public int playerId;
	
	public LobbyMessage msg;

	public static void ClearLobbyObject() {
		if (instance != null) Destroy(instance.gameObject);
	}
}
