using UnityEngine;

namespace Utils {
    public class ClickDragToPan : MonoBehaviour {
        private Vector3 BOTTOM_LEFT = new Vector3(-5, -5, 0);
        private Vector3 TOP_RIGHT = new Vector3(5, 5, 0);

        private Camera cam;
        private bool isDragging = false;
        private Vector3 camDown;
        private Vector3 mouseDown;

        void Start() {
            cam = FindObjectOfType<Camera>();
        }
        
        void Update() {
            if (!isDragging && Input.GetMouseButtonDown(0)) {
                isDragging = true;
                mouseDown = Input.mousePosition;
                camDown = cam.transform.position;
            }
            if (isDragging) {
                var dragDiff = Input.mousePosition - mouseDown;
                var worldDragDiff = cam.ScreenToWorldPoint(Vector3.zero) - cam.ScreenToWorldPoint(dragDiff); 
                cam.transform.position = clamp(camDown + worldDragDiff);
                if (Input.GetMouseButtonUp(0)) isDragging = false;
            }
        }

        private Vector3 clamp(Vector3 v) {
            return new Vector3(Mathf.Clamp(v.x, BOTTOM_LEFT.x, TOP_RIGHT.x), Mathf.Clamp(v.y, BOTTOM_LEFT.y, TOP_RIGHT.y), v.z);
        }
    }
}