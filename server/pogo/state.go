package pogo

// GameState is an all-containing set of information for the state of a
// game at a given point in time
type GameState struct {
	Turn      int   // the current turn
	Created   int64 // the epoch timestamp when this game started
	Updated   int64 // the epoch timestamp for when this specific state was created
	IDCounter int   // a counter to keep track of all of the ids issued throughout the game

	Players []PlayerState // each player state corresponds to each player in the game

	BoardWidth  int          // how many tiles wide is the game
	BoardHeight int          // how many tiles long is the game
	BoardSpaces []BoardSpace // defines each space on the board reading from top left to bottom right

	IsGameFinished bool // is the game finished
	Winner         int  // the id of the winner of the game
}

// GetNewID increments the IDCounter on this game state and returns the last IDCounter
func (s *GameState) GetNewID() int {
	s.IDCounter = s.IDCounter + 1
	return s.IDCounter - 1
}

// DeepCopy returns a deep copy of this game state
func (s *GameState) DeepCopy() *GameState {
	// TODO: MW I really hate that this is manual like this... Is there no library out there that does deep dynamic copying?
	cp := &GameState{
		Turn:        s.Turn,
		Created:     s.Created,
		Updated:     s.Updated,
		IDCounter:   s.IDCounter,
		BoardWidth:  s.BoardWidth,
		BoardHeight: s.BoardHeight,
	}
	for _, player := range s.Players {
		cpP := &PlayerState{
			ID:       player.ID,
			Name:     player.Name,
			Location: player.Location,
			Facing:   player.Facing,
			Health:   player.Health,
			Hand:     deepCopyListOfCards(player.Hand),
			Deck:     deepCopyListOfCards(player.Deck),
			Discard:  deepCopyListOfCards(player.Discard),
		}
		cp.Players = append(cp.Players, *cpP)
	}
	for _, space := range s.BoardSpaces {
		cpS := &BoardSpace{
			ID:        space.ID,
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
			ID:       card.ID,
			CardType: card.CardType,
		}
		cp = append(cp, *cpC)
	}
	return cp
}

// GetPlayer returns the player with the given id or nil if they do not exist
func (s *GameState) GetPlayer(id int) *PlayerState {
	for i, player := range s.Players {
		if player.ID == id {
			return &s.Players[i]
		}
	}
	return nil
}

// GetCardType returns the type of the card with the matching id
func (s *GameState) GetCardType(id int) CardType {
	for _, player := range s.Players {
		for _, card := range player.Hand {
			if card.ID == id {
				return card.CardType
			}
		}
		for _, card := range player.Deck {
			if card.ID == id {
				return card.CardType
			}
		}
		for _, card := range player.Discard {
			if card.ID == id {
				return card.CardType
			}
		}
	}
	return Unknown
}

// PlayerState contains all of the information about a given player, their hand, their discard, everything
type PlayerState struct {
	ID       int         // unique id for this player in this game
	Name     string      // essentially the username of the player
	Hand     []CardState // the cards currently available to the player
	Discard  []CardState // the cards that have been played or discarded
	Deck     []CardState // the cards still in the deck
	Location int         // the index of the board space the player is on
	Facing   int         // the direction the player is facing 0, 1, 2, 3 for North, East, South, West
	Health   int         // the current hitpoints of this player
}

// AnimationAction tracks the specific animations required by the client to move from state A to state B
type AnimationAction struct {
	ID     int    // id for this specific action (mostly for debugging)
	Name   string // name key for the type of animation
	Target int    // the id of the target
	Value  int    // a generic integer field that will have specific meaning for each type of animation action
}

// CardState defines a single and specific instance of a card in the game
type CardState struct {
	ID       int      // the unique id that denotes this specific card
	CardType CardType // the type of card
}

// BoardSpace defines a single and specific instance of a space on the board
type BoardSpace struct {
	ID        int // the unique id that describes this space on the board
	SpaceType int // the type of the space
	State     int // the state that the space is in (0 = default, 1 = on fire, 2 = flooded, etc)
}
