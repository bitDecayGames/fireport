using System;
using Game;
using UnityEngine;

namespace DebugScripts {
    public class DebugGameBoardBehaviour : MonoBehaviour {

        private GameBoardBehaviour gameBoardBehaviour;
        
        void Start() {
            gameBoardBehaviour = GetComponent<GameBoardBehaviour>();
            if (gameBoardBehaviour == null) throw new Exception("Could not find GameBoardBehaviour on DebugGameBoardBehaviour object");
            
            gameBoardBehaviour.Populate(DebugHelper.DebugGameState());
        }
    }
}