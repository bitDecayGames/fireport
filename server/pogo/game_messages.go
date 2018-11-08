package pogo

// GameStartMsg contains all information a client needs to start
// playing in a game
type GameStartMsg struct {
	TypedMsg

	GameID    string    `json:"gameID"`
	GameState GameState `json:"gameState"`
	Players   []string  `json:"players"`
}

// TurnTimeoutMsg contains information for when a player's turn has timed out
type TurnTimeoutMsg struct {
	TypedMsg

	// Currently no needed information
}

// GameTimeoutMsg contains information for when a game has had no activity
// and will be ended by the server
type GameTimeoutMsg struct {
	TypedMsg

	// Currently no needed information
}

// TurnResultMsg contains information on game state changes that occurred
// in the latest turn
type TurnResultMsg struct {
	TypedMsg

	GameID        string
	PreviousState GameState
	CurrentState  GameState
	Animations    []AnimationAction
}

// TurnSubmissionMsg is a full turn of actions submitted by a player
type TurnSubmissionMsg struct {
	TypedMsg

	GameID   string         `json:"gameID"`
	PlayerID string         `json:"playerID"`
	Inputs   []GameInputMsg `json:"inputs"`
}

// GameInputMsg is sent by the user to tell the server which card the user is selecting
type GameInputMsg struct {
	TypedMsg

	CardID int // this should maybe just be a query parameter or something simple
	Owner  int // this should come from the authentication layer
	Order  int // the order of the card for this player's selection
	Swap   int // the id of the basic movement type that this input will use, 0 means use the card as is
}

// GameResultSummaryMsg is sent out to each player when the game has been won
type GameResultSummaryMsg struct {
	TypedMsg

	Winner string
	// TODO: attach stats and other summary information
}

// CurrentTurnMsg contains information of what turn a game is on
type CurrentTurnMsg struct {
	TypedMsg

	GameID      string `json:"gameID"`
	CurrentTurn int    `json:"currentTurn"`
}
