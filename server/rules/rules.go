package rules

import (
	"fmt"
	"github.com/bitdecaygames/fireport/server/pogo"
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
	&IdCounterMustOnlyMoveForwardRule{},
	&BoardWidthAndHeightMustRemainConstantRule{},
	&NumberOfBoardSpacesCannotChangeRule{},
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
func ApplyGameRules(a *pogo.GameState, b *pogo.GameState, rules []GameRule) []error {
	var errors []error = nil
	for _, rule := range rules {
		err := rule.Apply(a, b)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

type OneTurnAtATimeRule struct{}

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

type CreatedMustRemainConstantRule struct{}

func (r *CreatedMustRemainConstantRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.Created != b.Created {
		return fmt.Errorf("cannot change the created date of a game")
	} else {
		return nil
	}
}

type UpdatedMustAlwaysMoveForwardRule struct{}

func (r *UpdatedMustAlwaysMoveForwardRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.Updated >= b.Updated {
		return fmt.Errorf("updated date much change for each new game state")
	} else {
		return nil
	}
}

type IdCounterMustOnlyMoveForwardRule struct{}

func (r *IdCounterMustOnlyMoveForwardRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.IdCounter > b.IdCounter {
		return fmt.Errorf("id counter cannot go down in value")
	} else {
		return nil
	}
}

type BoardWidthAndHeightMustRemainConstantRule struct{}

func (r *BoardWidthAndHeightMustRemainConstantRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if a.BoardWidth != b.BoardWidth {
		return fmt.Errorf("board width is not allowed to change mid-game")
	} else if a.BoardHeight != b.BoardHeight {
		return fmt.Errorf("board height is not allowed to change mid-game")
	} else {
		return nil
	}
}

type NumberOfBoardSpacesCannotChangeRule struct{}

func (r *NumberOfBoardSpacesCannotChangeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if len(a.BoardSpaces) != len(b.BoardSpaces) {
		return fmt.Errorf("number of board spaces changed from %v to %v", len(a.BoardSpaces), len(b.BoardSpaces))
	} else {
		return nil
	}
}

type BoardSpaceIdsCannotBeChangedRule struct{}

func (r *BoardSpaceIdsCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, spaceA := range a.BoardSpaces {
		var found = false
		for _, spaceB := range b.BoardSpaces {
			if spaceA.Id == spaceB.Id {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("could not find board space %v in new state", spaceA.Id)
		}
	}
	return nil
}

type PlayerIdsCannotBeChangedRule struct{}

func (r *PlayerIdsCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		var found = false
		for _, playerB := range b.Players {
			if playerA.Id == playerB.Id {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("could not find player %v in new state", playerA.Id)
		}
	}
	return nil
}

type NumberOfPlayersCannotChangeRule struct{}

func (r *NumberOfPlayersCannotChangeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	if len(a.Players) != len(b.Players) {
		return fmt.Errorf("number of players changed from %v to %v", len(a.Players), len(b.Players))
	} else {
		return nil
	}
}

type CardIdsCannotBeChangedRule struct{}

func (r *CardIdsCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		for _, playerB := range b.Players {
			if playerA.Id == playerB.Id {
				for _, card := range playerA.Hand {
					if !doesPlayerHaveCardId(&playerB, card.Id) {
						return fmt.Errorf("missing card %v from hand", card.Id)
					}
				}
				for _, card := range playerA.Deck {
					if !doesPlayerHaveCardId(&playerB, card.Id) {
						return fmt.Errorf("missing card %v from deck", card.Id)
					}
				}
				for _, card := range playerA.Discard {
					if !doesPlayerHaveCardId(&playerB, card.Id) {
						return fmt.Errorf("missing card %v from discard", card.Id)
					}
				}
				break
			}
		}
	}
	return nil
}

type CardTypesCannotBeChangedRule struct{}

func (r *CardTypesCannotBeChangedRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		var playerACards = collectAllPlayersCards(&playerA)
		for _, playerB := range b.Players {
			if playerA.Id == playerB.Id {
				var playerBCards = collectAllPlayersCards(&playerB)
				for _, aCard := range playerACards {
					for _, bCard := range playerBCards {
						if aCard.Id == bCard.Id && aCard.CardType != bCard.CardType {
							return fmt.Errorf("card %v from player %v changed type from %v to %v", aCard.Id, playerA.Id, aCard.CardType, bCard.CardType)
						}
					}
				}
				break
			}
		}
	}
	return nil
}

type NumberOfPlayerCardsCannotChangeRule struct{}

func (r *NumberOfPlayerCardsCannotChangeRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerA := range a.Players {
		var aTotal = len(playerA.Hand) + len(playerA.Deck) + len(playerA.Discard)
		for _, playerB := range b.Players {
			if playerA.Id == playerB.Id {
				var bTotal = len(playerB.Hand) + len(playerB.Deck) + len(playerB.Discard)
				if aTotal != bTotal {
					return fmt.Errorf("expected %v cards from player %v but instead got %v", aTotal, playerA.Id, bTotal)
				}
				break
			}
		}
	}
	return nil
}

type PlayersCannotOccupyTheSameSpaceRule struct{}

func (r *PlayersCannotOccupyTheSameSpaceRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerB1 := range b.Players {
		for _, playerB2 := range b.Players {
			if playerB1.Id != playerB2.Id && playerB1.Location == playerB2.Location {
				return fmt.Errorf("player %v and player %v are occupying the same space %v", playerB1.Id, playerB2.Id, playerB1.Location)
			}
		}
	}
	return nil
}

type PlayersMustOccupyAnExistingSpaceRule struct{}

func (r *PlayersMustOccupyAnExistingSpaceRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerB := range b.Players {
		var found = false
		for _, space := range b.BoardSpaces {
			if playerB.Location == space.Id {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("player %v is occupying a space %v that does not exist", playerB.Id, playerB.Location)
		}
	}
	return nil
}

type PlayerHandsMustContainSpecificNumberOfCards struct {
	NumberOfCardsInHand int
}

func (r *PlayerHandsMustContainSpecificNumberOfCards) Apply(a *pogo.GameState, b *pogo.GameState) error {
	for _, playerB := range b.Players {
		if len(playerB.Hand) != r.NumberOfCardsInHand {
			return fmt.Errorf("expected player %v hand to contain %v cards instead found %v", playerB.Id, r.NumberOfCardsInHand, len(playerB.Hand))
		}
	}
	return nil
}

type AllIdsMustBeUniqueRule struct{}

func (r *AllIdsMustBeUniqueRule) Apply(a *pogo.GameState, b *pogo.GameState) error {
	var ids = collectAllGameStateIds(b)
	for i, first := range ids {
		for k, second := range ids {
			if i != k && first.Id == second.Id {
				return fmt.Errorf("the id %v was the same between the %v and the %v", first.Id, first.Name, second.Name)
			}
		}
	}
	return nil
}
