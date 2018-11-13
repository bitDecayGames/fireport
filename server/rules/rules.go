package rules

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
	"strings"
)

// GameRule is meant to validate the previous and new proposed state, it WILL NOT make any modifications to either state
type GameRule interface {
	Apply(a *pogo.GameState, b *pogo.GameState) error
}

// DefaultGameRules this is this list of default rules for a game
var DefaultGameRules = []GameRule{
	&OneTurnAtATimeRule{},
	&CreatedMustRemainConstantRule{},
	&UpdatedMustAlwaysMoveForwardRule{},
	&IDCounterMustOnlyMoveForwardRule{},
	&BoardWidthAndHeightMustRemainConstantRule{},
	&NumberOfBoardSpacesCannotChangeRule{},
	&NumberOfBoardSpacesMustEqualWidthAndHeightRule{},
	&BoardSpaceIdsCannotBeChangedRule{},
	&PlayerIdsCannotBeChangedRule{},
	&NumberOfPlayersCannotChangeRule{},
	&CardIdsCannotBeChangedRule{},
	&CardTypesCannotBeChangedRule{},
	&NumberOfPlayerCardsCannotChangeRule{},
	&PlayersCannotOccupyTheSameSpaceRule{},
	&PlayersMustOccupyAnExistingSpaceRule{},
	&PlayerHandsMustContainSpecificNumberOfCards{NumberOfCardsInHand: 5},
	&AllIdsMustBeUniqueRule{},
}

// ApplyGameRules takes two states and a list of rules and compiles a list of errors that correspond to the rules that have been violated
func ApplyGameRules(a *pogo.GameState, b *pogo.GameState, rules []GameRule) error {
	var errors []string
	for _, rule := range rules {
		err := rule.Apply(a, b)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	return nil
}

// OneTurnAtATimeRule makes sure the turn ticker only moves up by 1
type OneTurnAtATimeRule struct{}

// Apply apply this rule
func (r *OneTurnAtATimeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.Turn == b.Turn {
		return fmt.Errorf("the turn must increment by one each turn")
	} else if a.Turn > b.Turn {
		return fmt.Errorf("the turn cannot go backward")
	} else if a.Turn+1 != b.Turn {
		return fmt.Errorf("the turn cannot increment more than one each turn")
	} else {
		return nil
	}
}

// CreatedMustRemainConstantRule the created date on the game state should never change
type CreatedMustRemainConstantRule struct{}

// Apply apply this rule
func (r *CreatedMustRemainConstantRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.Created != b.Created {
		return fmt.Errorf("cannot change the created date of a game")
	}
	return nil
}

// UpdatedMustAlwaysMoveForwardRule the updated date on the state should never move backwards
type UpdatedMustAlwaysMoveForwardRule struct{}

// Apply apply this rule
func (r *UpdatedMustAlwaysMoveForwardRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.Updated >= b.Updated {
		return fmt.Errorf("updated date much change for each new game state")
	}
	return nil
}

// IDCounterMustOnlyMoveForwardRule in order to avoid duplicate ids, the counter should only go up
type IDCounterMustOnlyMoveForwardRule struct{}

// Apply apply this rule
func (r *IDCounterMustOnlyMoveForwardRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.IDCounter > b.IDCounter {
		return fmt.Errorf("id counter cannot go down in value")
	}
	return nil
}

// BoardWidthAndHeightMustRemainConstantRule you should not be allowed to change the board size after init
type BoardWidthAndHeightMustRemainConstantRule struct{}

// Apply apply this rule
func (r *BoardWidthAndHeightMustRemainConstantRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.BoardWidth != b.BoardWidth {
		return fmt.Errorf("board width is not allowed to change mid-game")
	} else if a.BoardHeight != b.BoardHeight {
		return fmt.Errorf("board height is not allowed to change mid-game")
	}
	return nil
}

// NumberOfBoardSpacesCannotChangeRule the array containing the spaces should not change in length
type NumberOfBoardSpacesCannotChangeRule struct{}

// Apply apply this rule
func (r *NumberOfBoardSpacesCannotChangeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if len(a.BoardSpaces) != len(b.BoardSpaces) {
		return fmt.Errorf("number of board spaces changed from %v to %v", len(a.BoardSpaces), len(b.BoardSpaces))
	}
	return nil
}

// NumberOfBoardSpacesMustEqualWidthAndHeightRule the length of the spaces array should match the width and height values
type NumberOfBoardSpacesMustEqualWidthAndHeightRule struct{}

// Apply apply this rule
func (r *NumberOfBoardSpacesMustEqualWidthAndHeightRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.BoardWidth*a.BoardHeight != len(b.BoardSpaces) {
		return fmt.Errorf("expected %v board spaces and instead got %v", b.BoardWidth*b.BoardHeight, len(b.BoardSpaces))
	}
	return nil
}

// BoardSpaceIdsCannotBeChangedRule the board space ids should never change
type BoardSpaceIdsCannotBeChangedRule struct{}

// Apply apply this rule
func (r *BoardSpaceIdsCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, spaceA := range a.BoardSpaces {
		var found = false
		for _, spaceB := range b.BoardSpaces {
			if spaceA.ID == spaceB.ID {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("could not find board space %v in new state", spaceA.ID)
		}
	}
	return nil
}

// PlayerIdsCannotBeChangedRule player ids should never change
type PlayerIdsCannotBeChangedRule struct{}

// Apply apply this rule
func (r *PlayerIdsCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		var found = false
		for _, playerB := range b.Players {
			if playerA.ID == playerB.ID {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("could not find player %v in new state", playerA.ID)
		}
	}
	return nil
}

// NumberOfPlayersCannotChangeRule you cannot add or remove players from a game in progress
type NumberOfPlayersCannotChangeRule struct{}

// Apply apply this rule
func (r *NumberOfPlayersCannotChangeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if len(a.Players) != len(b.Players) {
		return fmt.Errorf("number of players changed from %v to %v", len(a.Players), len(b.Players))
	}
	return nil
}

// CardIdsCannotBeChangedRule the ids of the cards in the game should never change
type CardIdsCannotBeChangedRule struct{}

// Apply apply this rule
func (r *CardIdsCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		for _, playerB := range b.Players {
			if playerA.ID == playerB.ID {
				for _, card := range playerA.Hand {
					if !doesPlayerHaveCardID(&playerB, card.ID) {
						return fmt.Errorf("missing card %v from hand", card.ID)
					}
				}
				for _, card := range playerA.Deck {
					if !doesPlayerHaveCardID(&playerB, card.ID) {
						return fmt.Errorf("missing card %v from deck", card.ID)
					}
				}
				for _, card := range playerA.Discard {
					if !doesPlayerHaveCardID(&playerB, card.ID) {
						return fmt.Errorf("missing card %v from discard", card.ID)
					}
				}
				break
			}
		}
	}
	return nil
}

// CardTypesCannotBeChangedRule the types of the cards in the game should never change
type CardTypesCannotBeChangedRule struct{}

// Apply apply this rule
func (r *CardTypesCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		var playerACards = collectAllPlayersCards(&playerA)
		for _, playerB := range b.Players {
			if playerA.ID == playerB.ID {
				var playerBCards = collectAllPlayersCards(&playerB)
				for _, aCard := range playerACards {
					for _, bCard := range playerBCards {
						if aCard.ID == bCard.ID && aCard.CardType != bCard.CardType {
							return fmt.Errorf("card %v from player %v changed type from %v to %v", aCard.ID, playerA.ID, aCard.CardType, bCard.CardType)
						}
					}
				}
				break
			}
		}
	}
	return nil
}

// NumberOfPlayerCardsCannotChangeRule the total number of a given player's card cannot change
type NumberOfPlayerCardsCannotChangeRule struct{}

// Apply apply this rule
func (r *NumberOfPlayerCardsCannotChangeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		var aTotal = len(playerA.Hand) + len(playerA.Deck) + len(playerA.Discard)
		for _, playerB := range b.Players {
			if playerA.ID == playerB.ID {
				var bTotal = len(playerB.Hand) + len(playerB.Deck) + len(playerB.Discard)
				if aTotal != bTotal {
					return fmt.Errorf("expected %v cards from player %v but instead got %v", aTotal, playerA.ID, bTotal)
				}
				break
			}
		}
	}
	return nil
}

// PlayersCannotOccupyTheSameSpaceRule no two players can occupy the same space on the board at the same time
type PlayersCannotOccupyTheSameSpaceRule struct{}

// Apply apply this rule
func (r *PlayersCannotOccupyTheSameSpaceRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerB1 := range b.Players {
		for _, playerB2 := range b.Players {
			if playerB1.ID != playerB2.ID && playerB1.Location == playerB2.Location {
				return fmt.Errorf("player %v and player %v are occupying the same space %v", playerB1.ID, playerB2.ID, playerB1.Location)
			}
		}
	}
	return nil
}

// PlayersMustOccupyAnExistingSpaceRule each player must be located at a space that actually exists on the board
type PlayersMustOccupyAnExistingSpaceRule struct{}

// Apply apply this rule
func (r *PlayersMustOccupyAnExistingSpaceRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerB := range b.Players {
		if playerB.Location < 0 || playerB.Location >= len(b.BoardSpaces) {
			return fmt.Errorf("player %v is occupying a space %v that does not exist", playerB.ID, playerB.Location)
		}
	}
	return nil
}

// PlayerHandsMustContainSpecificNumberOfCards the player's hand must always contain a specific number of cards
type PlayerHandsMustContainSpecificNumberOfCards struct {
	NumberOfCardsInHand int
}

// Apply apply this rule
func (r *PlayerHandsMustContainSpecificNumberOfCards) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerB := range b.Players {
		if len(playerB.Hand) != r.NumberOfCardsInHand {
			return fmt.Errorf("expected player %v hand to contain %v cards instead found %v", playerB.ID, r.NumberOfCardsInHand, len(playerB.Hand))
		}
	}
	return nil
}

// AllIdsMustBeUniqueRule you cannot have a card and a space have the same id in the game
type AllIdsMustBeUniqueRule struct{}

// Apply apply this rule
func (r *AllIdsMustBeUniqueRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	var ids = collectAllGameStateIds(b)
	for i, first := range ids {
		for k, second := range ids {
			if i != k && first.ID == second.ID {
				return fmt.Errorf("the id %v was the same between the %v and the %v", first.ID, first.Name, second.Name)
			}
		}
	}
	return nil
}
