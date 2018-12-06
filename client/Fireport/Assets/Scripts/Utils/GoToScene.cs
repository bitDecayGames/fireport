using UnityEngine;

namespace Utils {
    public class GoToScene : MonoBehaviour {
        public void Go(string sceneName) {
            SceneNavigation.LoadScene(sceneName);
        }

        public void Back() {
            SceneNavigation.LoadPreviousScene();
        }
    }
}