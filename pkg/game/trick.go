package game

import (
	"github.com/bshore/five-hunerd/pkg/objects"
)

const (
	trickPhaseNoCardChosen        int = -1
	trickPhaseBlockAction         int = -2
	trickPhaseBlockActionEvaluate int = -3
)

type Trick struct {
	Cards []*objects.Card
}

func NewTrickSet() []*Trick {
	return []*Trick{
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
		{Cards: []*objects.Card{}},
	}
}

func (t *Trick) WinningCard() *objects.Card {
	highestCard := t.Cards[0]
	for i := 1; i < len(t.Cards); i++ {
		card := t.Cards[i]
		if card.Suit == highestCard.Suit && card.Rank > highestCard.Rank {
			highestCard = card
		}
		if !highestCard.Trump && card.Trump {
			highestCard = card
		}
	}
	return highestCard
}
