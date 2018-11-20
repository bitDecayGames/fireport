package triggers

import (
	"github.com/bitdecaygames/fireport/server/actions"
	"github.com/bitdecaygames/fireport/server/rules"
	"testing"

	"github.com/bitdecaygames/fireport/server/pogo"
	"github.com/stretchr/testify/assert"
)

func TestWinTrigger(t *testing.T) {
	var trigger = &WinTrigger{}
	var a = pogo.GetTestState()

	assert.Equal(t, 0, a.Winner)
	assert.Equal(t, false, a.IsGameFinished)

	var b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, 0, b.Winner)
	assert.Equal(t, false, b.IsGameFinished)

	a.Players[1].Health = 0

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, 100, b.Winner)
	assert.Equal(t, true, b.IsGameFinished)

	a.Players[0].Health = 0

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	assert.Equal(t, -1, b.Winner)
	assert.Equal(t, true, b.IsGameFinished)
}

func TestNotEnoughCardsInHandTrigger(t *testing.T) {
	var trigger = &NotEnoughCardsInHandTrigger{RequiredCardsInHand: 5}
	var rule = &rules.CardIdsCannotBeChangedRule{}
	var a = pogo.GetTestState()

	var b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand))
	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))

	var action = &actions.DiscardCardAction{Owner: 100, CardID: 101}
	a, err = action.Apply(a)
	assert.NoError(t, err)

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand)-1)
	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck)+1)
}

func TestRefreshDeckTrigger(t *testing.T) {
	var trigger = &RefreshDeckTrigger{}
	var rule = &rules.CardIdsCannotBeChangedRule{}
	var a = pogo.GetTestState()

	var b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))

	a, err = actions.ApplyActions(a, []actions.Action{
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
	})
	assert.NoError(t, err)

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck)-5)

	a, err = actions.ApplyActions(b, []actions.Action{
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
		&actions.DrawCardAction{Owner: 100},
	})
	assert.NoError(t, err)

	b, err = ApplyTriggers(a, []Trigger{trigger})
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))
}

func TestDrawAndDeckRefreshTriggersTogether(t *testing.T) {
	var triggers = []Trigger{&NotEnoughCardsInHandTrigger{RequiredCardsInHand: 5}, &RefreshDeckTrigger{}}
	var rule = &rules.CardIdsCannotBeChangedRule{}
	var a = pogo.GetTestState()

	var b, err = ApplyTriggers(a, triggers)
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand))
	assert.Equal(t, len(a.Players[0].Deck), len(b.Players[0].Deck))
	assert.Equal(t, len(a.Players[0].Discard), len(b.Players[0].Discard))

	a, err = actions.ApplyActions(a, []actions.Action{
		&actions.DiscardCardAction{Owner: 100, CardID: 101},
		&actions.DiscardCardAction{Owner: 100, CardID: 102},
		&actions.DiscardCardAction{Owner: 100, CardID: 103},
		&actions.DiscardCardAction{Owner: 100, CardID: 104},
		&actions.DiscardCardAction{Owner: 100, CardID: 105},
		&actions.DrawCardAction{Owner: 100},
		&actions.DiscardCardAction{Owner: 100, CardID: 110},
		&actions.DrawCardAction{Owner: 100},
	})
	assert.NoError(t, err)

	assert.Equal(t, 1, len(a.Players[0].Hand))
	assert.Equal(t, 3, len(a.Players[0].Deck))
	assert.Equal(t, 11, len(a.Players[0].Discard))

	b, err = ApplyTriggers(a, triggers)
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, 5, len(b.Players[0].Hand))
	assert.Equal(t, 10, len(b.Players[0].Deck))
	assert.Equal(t, 0, len(b.Players[0].Discard))
}

func TestDiscardUsedCardsTrigger(t *testing.T) {
	var a = pogo.GetTestState()

	var b, err = ApplyTriggers(a, []Trigger{&DiscardUsedCardsTrigger{Cards: []int{101, 201, 202}}})
	var rule = &rules.CardIdsCannotBeChangedRule{}
	assert.NoError(t, err)

	assert.Equal(t, len(a.Players[0].Hand), len(b.Players[0].Hand)+1)
	assert.Equal(t, len(a.Players[1].Hand), len(b.Players[1].Hand)+2)

	c, err := ApplyTriggers(b, []Trigger{&DiscardUsedCardsTrigger{Cards: []int{101, 201, 202}}})
	assert.NoError(t, err)

	err = rule.Apply(a, b)
	assert.NoError(t, err)

	assert.Equal(t, len(b.Players[0].Hand), len(c.Players[0].Hand))
	assert.Equal(t, len(b.Players[1].Hand), len(c.Players[1].Hand))
}
