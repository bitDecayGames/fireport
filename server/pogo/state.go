package pogo

import (
	"math/rand"

	"github.com/bitdecaygames/fireport/server/animations"
)

// GameState is an all-containing set of information for the state of a
// game at a given point in time
type GameState struct {
	Turn      int        // the current turn
	RNG       *rand.Rand // random number generator for this state
	Created   int64      // the epoch timestamp when this game started
	Updated   int64      // the epoch timestamp for when this specific state was created
	IDCounter int        // a counter to keep track of all of the ids issued throughout the game

	Players []PlayerState // each player state corresponds to each player in the game

	BoardWidth  int          // how many tiles wide is the game
	BoardHeight int          // how many tiles long is the game
	BoardSpaces []BoardSpace // defines each space on the board reading from top left to bottom right

	IsGameFinished bool // is the game finished
	Winner         int  // the id of the winner of the game

	Animations [][]animations.AnimationAction // Animations used to get the current GameState
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
		RNG:         s.RNG,
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
	for x := range s.Animations {
		animationGroup := []animations.AnimationAction{}
		for _, animation := range s.Animations[x] {
			cpA := &animations.AnimationAction{
				ID:    animation.ID,
				Name:  animation.Name,
				Owner: animation.Owner,
			}
			animationGroup = append(animationGroup, *cpA)
		}
		cp.Animations = append(cp.Animations, animationGroup)
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

// AppendAnimation appends an animation to the len(Animations) -1 nested slice
func (s *GameState) AppendAnimation(animation animations.AnimationAction) {
	if len(s.Animations) == 0 {
		s.AddEmptyAnimationSlice()
	}
	s.Animations[len(s.Animations)-1] = append(s.Animations[len(s.Animations)-1], animation)
}

// AddEmptyAnimationSlice appends an empty AnimationAction slice to Animations.
// Used to designate the next AnimationAction group.
func (s *GameState) AddEmptyAnimationSlice() {
	s.Animations = append(s.Animations, []animations.AnimationAction{})
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
