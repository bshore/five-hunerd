package game

import (
	"github.com/bshore/five-hunerd/pkg/objects"
)

type Player struct {
	Hand   []*objects.Card
	TeamWe bool
}

func (p *Player) DecideCard(trick *Trick, trump objects.Suit) (int, *objects.Card) {
	if len(trick.Cards) == 0 {
		return p.Highest()
	}
	lead := trick.Cards[0]
	currentWinningCard := trick.WinningCard()
	if p.TeamWe && currentWinningCard.TeamWe {
		if p.IHave(lead.Suit) {
			idx, card := p.SluffSuit(lead.Suit)
			return idx, card
		}
		idx, card := p.Sluff()
		return idx, card
	}
	if p.IHave(lead.Suit) {
		if p.ICanTake(trick) {
			idx, card := p.HighestSuit(lead.Suit)
			return idx, card
		}
		idx, card := p.SluffSuit(lead.Suit)
		return idx, card
	} else if p.IHave(trump) {
		idx, card := p.SluffSuit(trump)
		return idx, card
	}
	idx, card := p.Sluff()
	return idx, card
}

func (p *Player) IHave(suit objects.Suit) bool {
	for _, card := range p.Hand {
		if card.Suit == suit {
			return true
		}
	}
	return false
}

func (p *Player) ICanTake(t *Trick) bool {
	leadingCard := t.Cards[0]
	for i := 1; i < len(t.Cards); i++ {
		card := t.Cards[i]
		if card.Suit == leadingCard.Suit && card.Rank > leadingCard.Rank {
			leadingCard = card
		} else if card.Trump && leadingCard.Trump && card.Rank > leadingCard.Rank {
			leadingCard = card
		}
	}
	for _, card := range p.Hand {
		if card.Suit == leadingCard.Suit && card.Rank > leadingCard.Rank {
			return true
		}
	}
	return false
}

func (p *Player) Highest() (int, *objects.Card) {
	highestCard := &objects.Card{
		Rank: objects.RankTwo,
	}
	highestTrump := &objects.Card{
		Rank: objects.RankTwo,
	}

	h := 0
	ht := 0

	for i, card := range p.Hand {
		if card.Rank > highestCard.Rank && !card.Trump {
			h = i
			highestCard = card
		} else if card.Rank > highestTrump.Rank {
			ht = i
			highestTrump = card
		}
		if highestCard.Rank == objects.RankTwo {
			return ht, highestTrump
		}
	}
	return h, highestCard
}

func (p *Player) HighestSuit(suit objects.Suit) (int, *objects.Card) {
	highestCard := &objects.Card{
		Rank: objects.RankTwo,
	}

	h := 0

	for i, card := range p.Hand {
		if card.Suit == suit && card.Rank > highestCard.Rank {
			h = i
			highestCard = card
		}
	}
	return h, highestCard
}

func (p *Player) Sluff() (int, *objects.Card) {
	lowestCard := &objects.Card{
		Rank: objects.RankJoker,
	}

	lowestTrump := &objects.Card{
		Rank: objects.RankJoker,
	}

	l := 0
	lt := 0

	for i, card := range p.Hand {
		if card.Rank < lowestCard.Rank && !card.Trump {
			l = i
			lowestCard = card
		} else if card.Rank < lowestTrump.Rank {
			lt = i
			lowestTrump = card
		}
	}
	if lowestCard.Rank == objects.RankJoker {
		return lt, lowestTrump
	}
	return l, lowestCard
}

func (p *Player) SluffSuit(suit objects.Suit) (int, *objects.Card) {
	lowestCard := &objects.Card{
		Rank: objects.RankJoker,
	}

	l := 0

	for i, card := range p.Hand {
		if card.Suit == suit && card.Rank < lowestCard.Rank {
			l = i
			lowestCard = card
		}
	}
	return l, lowestCard
}
