package game

import (
	"fmt"

	"github.com/bshore/five-hunerd/pkg/objects"
)

type Bid struct {
	Take   int
	Suit   objects.Suit
	Locked bool
}

func (b *Bid) String() string {
	return fmt.Sprintf("%d%v", b.Take, b.Suit.String()[0:1])
}

func (b *Bid) Score() int {
	return (100 * (b.Take - 6)) + b.Suit.Points()
}

type Round struct {
	Hand               []*objects.Card
	Players            []*Player
	Blind              []*objects.Card
	HandExchangeCards  []*objects.Card
	BlindExchangeCards []*objects.Card
	Bids               []*Bid
	Bid                *Bid
	TrickPlaysNext     int
	TrickCard          int
	OnTrick            int
	Tricks             []*Trick
	Phase              PhaseType
	TeamWeTricks       int
	TeamWeScore        int
	TeamTheyTricks     int
	TeamTheyScore      int
}

func newRound(p1, p2, p3, p4, blind []*objects.Card) *Round {
	return &Round{
		Hand: p1,
		Players: []*Player{
			{Hand: p2},
			{Hand: p3, TeamWe: true},
			{Hand: p4}},
		Blind:              blind,
		HandExchangeCards:  make([]*objects.Card, 10),
		BlindExchangeCards: make([]*objects.Card, 5),
		Bids:               []*Bid{},
		Bid:                &Bid{},
		TrickPlaysNext:     -1,
		TrickCard:          -1,
		OnTrick:            0,
		Tricks:             NewTrickSet(),
		Phase:              PhaseTypeBid,
		TeamWeTricks:       0,
		TeamTheyTricks:     0,
	}
}

func (r *Round) AllPlayersActed() bool {
	return len(r.Tricks[r.OnTrick].Cards) == 4
}

// ?
func (r *Round) Finished() bool {
	playerHand := len(r.Hand)
	botHands := 0
	for i := range r.Players {
		botHands += len(r.Players[i].Hand)
	}
	return playerHand+botHands == 0 && r.TeamWeTricks+r.TeamTheyTricks == 10
}
