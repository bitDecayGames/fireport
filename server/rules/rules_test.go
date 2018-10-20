package rules

import (
	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestStates() (*pogo.GameState, *pogo.GameState) {
	return &pogo.GameState{
		Turn:        0,
		Created:     1000,
		Updated:     2000,
		IdCounter:   300,
		BoardWidth:  2,
		BoardHeight: 2,
		BoardSpaces: []pogo.BoardSpace{
			{Id: 0, SpaceType: 0, State: 0},
			{Id: 1, SpaceType: 0, State: 0},
			{Id: 2, SpaceType: 0, State: 0},
			{Id: 3, SpaceType: 0, State: 0},
		},
		Players: []pogo.PlayerState{
			{
				Id:       100,
				Name:     "PlayerOne",
				Location: 0,
				Hand: []pogo.CardState{
					{Id: 101, CardType: 0},
					{Id: 102, CardType: 0},
					{Id: 103, CardType: 0},
					{Id: 104, CardType: 0},
					{Id: 105, CardType: 0},
				},
				Deck: []pogo.CardState{
					{Id: 106, CardType: 0},
					{Id: 107, CardType: 0},
					{Id: 108, CardType: 0},
					{Id: 109, CardType: 0},
					{Id: 110, CardType: 0},
				},
				Discard: []pogo.CardState{
					{Id: 111, CardType: 0},
					{Id: 112, CardType: 0},
					{Id: 113, CardType: 0},
					{Id: 114, CardType: 0},
					{Id: 115, CardType: 0},
				},
			},
			{
				Id:       200,
				Name:     "PlayerTwo",
				Location: 3,
				Hand: []pogo.CardState{
					{Id: 201, CardType: 0},
					{Id: 202, CardType: 0},
					{Id: 203, CardType: 0},
					{Id: 204, CardType: 0},
					{Id: 205, CardType: 0},
				},
				Deck: []pogo.CardState{
					{Id: 206, CardType: 0},
					{Id: 207, CardType: 0},
					{Id: 208, CardType: 0},
					{Id: 209, CardType: 0},
					{Id: 210, CardType: 0},
				},
				Discard: []pogo.CardState{
					{Id: 211, CardType: 0},
					{Id: 212, CardType: 0},
					{Id: 213, CardType: 0},
					{Id: 214, CardType: 0},
					{Id: 215, CardType: 0},
				},
			},
		},
	}, &pogo.GameState{
		Turn:        0,
		Created:     1000,
		Updated:     2000,
		IdCounter:   300,
		BoardWidth:  2,
		BoardHeight: 2,
		BoardSpaces: []pogo.BoardSpace{
			{Id: 0, SpaceType: 0, State: 0},
			{Id: 1, SpaceType: 0, State: 0},
			{Id: 2, SpaceType: 0, State: 0},
			{Id: 3, SpaceType: 0, State: 0},
		},
		Players: []pogo.PlayerState{
			{
				Id:       100,
				Name:     "PlayerOne",
				Location: 0,
				Hand: []pogo.CardState{
					{Id: 101, CardType: 0},
					{Id: 102, CardType: 0},
					{Id: 103, CardType: 0},
					{Id: 104, CardType: 0},
					{Id: 105, CardType: 0},
				},
				Deck: []pogo.CardState{
					{Id: 106, CardType: 0},
					{Id: 107, CardType: 0},
					{Id: 108, CardType: 0},
					{Id: 109, CardType: 0},
					{Id: 110, CardType: 0},
				},
				Discard: []pogo.CardState{
					{Id: 111, CardType: 0},
					{Id: 112, CardType: 0},
					{Id: 113, CardType: 0},
					{Id: 114, CardType: 0},
					{Id: 115, CardType: 0},
				},
			},
			{
				Id:       200,
				Name:     "PlayerTwo",
				Location: 3,
				Hand: []pogo.CardState{
					{Id: 201, CardType: 0},
					{Id: 202, CardType: 0},
					{Id: 203, CardType: 0},
					{Id: 204, CardType: 0},
					{Id: 205, CardType: 0},
				},
				Deck: []pogo.CardState{
					{Id: 206, CardType: 0},
					{Id: 207, CardType: 0},
					{Id: 208, CardType: 0},
					{Id: 209, CardType: 0},
					{Id: 210, CardType: 0},
				},
				Discard: []pogo.CardState{
					{Id: 211, CardType: 0},
					{Id: 212, CardType: 0},
					{Id: 213, CardType: 0},
					{Id: 214, CardType: 0},
					{Id: 215, CardType: 0},
				},
			},
		},
	}
}

func TestOneTurnAtATimeRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &OneTurnAtATimeRule{}

	var err = rule.Apply(a, b)
	assert.Error(t, err)

	b.Turn = b.Turn + 1
	err = rule.Apply(a, b)
	assert.NoError(t, err)
}

func TestCreatedMustRemainConstantRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &CreatedMustRemainConstantRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Created = b.Created + 1
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestUpdatedMustAlwaysMoveForwardRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &UpdatedMustAlwaysMoveForwardRule{}

	var err = rule.Apply(a, b)
	assert.Error(t, err)

	b.Updated = b.Updated - 10
	err = rule.Apply(a, b)
	assert.Error(t, err)

	b.Updated = b.Updated + 20
	err = rule.Apply(a, b)
	assert.NoError(t, err)
}

func TestIdCounterMustOnlyMoveForwardRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &IdCounterMustOnlyMoveForwardRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.IdCounter = b.IdCounter + 1
	err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.IdCounter = b.IdCounter - 10
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestBoardWidthAndHeightMustRemainConstantRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &BoardWidthAndHeightMustRemainConstantRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.BoardWidth = b.BoardWidth + 1
	err = rule.Apply(a, b)
	assert.Error(t, err)

	b.BoardWidth = a.BoardWidth
	b.BoardHeight = b.BoardHeight + 1
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestNumberOfBoardSpacesCannotChangeRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &NumberOfBoardSpacesCannotChangeRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.BoardSpaces = append(b.BoardSpaces, pogo.BoardSpace{Id: 3000})
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestNumberOfBoardSpacesMustEqualWidthAndHeightRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &NumberOfBoardSpacesMustEqualWidthAndHeightRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.BoardSpaces = append([]pogo.BoardSpace{}, b.BoardSpaces[1:]...)
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestBoardSpaceIdsCannotBeChangedRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &BoardSpaceIdsCannotBeChangedRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.BoardSpaces[0].Id = 3000
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestPlayerIdsCannotBeChangedRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &PlayerIdsCannotBeChangedRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Id = 3000
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestNumberOfPlayersCannotChangeRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &NumberOfPlayersCannotChangeRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players = append(b.Players, pogo.PlayerState{Id: 3000})
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestCardIdsCannotBeChangedRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &CardIdsCannotBeChangedRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Hand[0].Id = 3000
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestCardTypesCannotBeChangedRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &CardTypesCannotBeChangedRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Hand[0].CardType = 3000
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestNumberOfPlayerCardsCannotChangeRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &NumberOfPlayerCardsCannotChangeRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Hand = append(b.Players[0].Hand, pogo.CardState{Id: 3000})
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestPlayersCannotOccupyTheSameSpaceRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &PlayersCannotOccupyTheSameSpaceRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Location = b.Players[1].Location
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestPlayersMustOccupyAnExistingSpaceRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &PlayersMustOccupyAnExistingSpaceRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Location = 7
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestPlayerHandsMustContainSpecificNumberOfCards(t *testing.T) {
	var a, b = getTestStates()
	var rule = &PlayerHandsMustContainSpecificNumberOfCards{NumberOfCardsInHand: 5}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	var drawCard = b.Players[0].Deck[0]
	b.Players[0].Deck = append([]pogo.CardState{}, b.Players[0].Deck[1:]...)
	b.Players[0].Hand = append(b.Players[0].Hand, drawCard)
	err = rule.Apply(a, b)
	assert.Error(t, err)
}

func TestAllIdsMustBeUniqueRule(t *testing.T) {
	var a, b = getTestStates()
	var rule = &AllIdsMustBeUniqueRule{}

	var err = rule.Apply(a, b)
	assert.NoError(t, err)

	b.Players[0].Hand[0].Id = b.Players[1].Id
	err = rule.Apply(a, b)
	assert.Error(t, err)
}
