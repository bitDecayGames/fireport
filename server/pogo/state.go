package pogo

// GameState is an all-containing set of information for the state of a
// game at a given point in time
type GameState struct {
	Turn      int   // the current turn
	Created   int64 // the epoch timestamp when this game started
	Updated   int64 // the epoch timestamp for when this specific state was created
	IdCounter int   // a counter to keep track of all of the ids issued throughout the game

	Players []PlayerState // each player state corresponds to each player in the game

	BoardWidth  int          // how many tiles wide is the game
	BoardHeight int          // how many tiles long is the game
	BoardSpaces []BoardSpace // defines each space on the board reading from top left to bottom right
}

// GetNewId increments the IdCounter on this game state and returns the last IdCounter
func (s *GameState) GetNewId() int {
	s.IdCounter = s.IdCounter + 1
	return s.IdCounter - 1
}

// DeepCopy returns a deep copy of this game state
func (s *GameState) DeepCopy() *GameState {
	// TODO: MW I really hate that this is manual like this... Is there no library out there that does deep dynamic copying?
	cp := &GameState{
		Turn:        s.Turn,
		Created:     s.Created,
		Updated:     s.Updated,
		IdCounter:   s.IdCounter,
		BoardWidth:  s.BoardWidth,
		BoardHeight: s.BoardHeight,
	}
	for _, player := range s.Players {
		cpP := &PlayerState{
			Id:       player.Id,
			Name:     player.Name,
			Location: player.Location,
			Facing:   player.Facing,
			Hand:     deepCopyListOfCards(player.Hand),
			Deck:     deepCopyListOfCards(player.Deck),
			Discard:  deepCopyListOfCards(player.Discard),
		}
		cp.Players = append(cp.Players, *cpP)
	}
	for _, space := range s.BoardSpaces {
		cpS := &BoardSpace{
			Id:        space.Id,
			SpaceType: space.SpaceType,
			State:     space.State,
		}
		cp.BoardSpaces = append(cp.BoardSpaces, *cpS)
	}
	return cp
}

func deepCopyListOfCards(cards []CardState) []CardState {
	var cp []CardState
	for _, card := range cards {
		cpC := &CardState{
			Id:       card.Id,
			CardType: card.CardType,
		}
		cp = append(cp, *cpC)
	}
	return cp
}

// GetPlayer returns the player with the given id or nil if they do not exist
func (s *GameState) GetPlayer(id int) *PlayerState {
	for i, player := range s.Players {
		if player.Id == id {
			return &s.Players[i]
		}
	}
	return nil
}

// PlayerState contains all of the information about a given player, their hand, their discard, everything
type PlayerState struct {
	Id       int         // unique id for this player in this game
	Name     string      // essentially the username of the player // TODO: MW is this necessary?
	Hand     []CardState // the cards currently available to the player
	Discard  []CardState // the cards that have been played or discarded
	Deck     []CardState // the cards still in the deck
	Location int         // the id of the board space this player is occupying
	Facing   int         // the direction the player is facing 0, 1, 2, 3 for Up, Right, Down, Left
	// TODO: MW there could be more here like how much health the player has, if that is something we want
}

// AnimationAction tracks the specific animations required by the client to move from state A to state B
type AnimationAction struct {
	Id     int    // id for this specific action (mostly for debugging)
	Name   string // name key for the type of animation
	Target int    // the id of the target
	Value  int    // a generic integer field that will have specific meaning for each type of animation action
}

// CardState defines a single and specific instance of a card in the game
type CardState struct {
	Id       int // the unique id that denotes this specific card
	CardType int // the type of card
}

// BoardSpace defines a single and specific instance of a space on the board
type BoardSpace struct {
	Id        int // the unique id that describes this space on the board
	SpaceType int // the type of the space
	State     int // the state that the space is in (0 = default, 1 = on fire, 2 = flooded, etc)
}
