using UnityEngine;

namespace Utils {
    public class ScrollWheelToZoom : MonoBehaviour {
        private const float MIN_CAMERA_SIZE = 1f;
        private const float MAX_CAMERA_SIZE = 10f;
        private const float SENSITIVITY = 0.5f;

        private Camera cam;

        void Start() {
            cam = FindObjectOfType<Camera>();
        }
        
        void Update() {
            var scroll = Input.mouseScrollDelta.y;
            cam.orthographicSize = clamp(cam.orthographicSize + scroll * SENSITIVITY);
        }

        private float clamp(float f) {
            return Mathf.Clamp(f, MIN_CAMERA_SIZE, MAX_CAMERA_SIZE);
        }
    }
}