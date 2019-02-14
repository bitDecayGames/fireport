using UnityEngine;

namespace Utils {
    public class SlideUpFromBottom : MonoBehaviour {
        public bool IsShown { get; private set; }
        private Vector3 originalPos;
        private bool isAnimating = false;
        private Vector3 target;
        private const float speedRatio = 0.05f;
        private const float fuzzyDistance = 0.06f;
        private bool isInitialized = false;
        
        void Start() {
            Init();
        }

        private void Init() {
            if (!isInitialized) {
                originalPos = transform.localPosition;
                isInitialized = true;
            }
        }
        
        void Update() {
            if (isAnimating) {
                var pos = transform.localPosition;
                if (Vector3.Distance(pos, target) <= fuzzyDistance) isAnimating = false;
                else transform.localPosition = (target - pos) * speedRatio + pos;
            }
        }

        public void Hide(bool animate = true) {
            Init();
            var height = ((RectTransform) transform).rect.height;
            IsShown = false;
            target = originalPos + new Vector3(0, -height, 0);
            if (animate) isAnimating = true;
            else transform.localPosition = target;
        }

        public void Show(bool animate = true) {
            Init();
            IsShown = true;
            target = originalPos;
            if (animate) isAnimating = true;
            else transform.localPosition = target;
        }
    }
}