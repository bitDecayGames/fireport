using System.Collections.Generic;
using UnityEngine.SceneManagement;

namespace Utils {
    public static class SceneNavigation {
        private static List<string> history = new List<string>();
        

        public static void LoadScene(string sceneName, bool trackHistory = true) {
            if (!ResetHistoryLoops(sceneName) && trackHistory) AddToHistory(SceneManager.GetActiveScene().name);
            SceneManager.LoadScene(sceneName);
        }

        public static void LoadSceneClearHistory(string sceneName) {
            history.Clear();
            LoadScene(sceneName, false);
        }

        public static void RefreshScene() {
            SceneManager.LoadScene(SceneManager.GetActiveScene().name);
        }

        public static void LoadPreviousScene(string orElse = null) {
            var last = PopPreviousSceneName();
            if (last != null) SceneManager.LoadScene(last);
            else if (orElse != null) LoadScene(orElse);
        }

        public static string PreviousSceneName() {
            if (history != null && history.Count > 0) return history[0];
            return null;
        }

        private static bool ResetHistoryLoops(string sceneName) {
            var found = false;
            var index = history.FindIndex(s => sceneName == s);
            for (int i = 0; i <= index; i++) {
                PopPreviousSceneName();
                found = true;
            }

            return found;
        }

        private static void AddToHistory(string sceneName) {
            var last = PreviousSceneName();
            if (sceneName != null && sceneName != last) history.Insert(0, sceneName);
        }

        private static string PopPreviousSceneName() {
            var last = PreviousSceneName();
            if (last != null) history.RemoveAt(0);
            return last;
        }
    }
}