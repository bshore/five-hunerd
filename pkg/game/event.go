package game

import (
	"fmt"
	"os"

	"github.com/bshore/five-hunerd/pkg/objects"
	"github.com/faiface/beep/speaker"
)

type EventType int

const (
	EventTypeNoAction EventType = iota
	EventTypeSort
	EventTypeBid
	EventTypeBidPass
	EventTypeBidLockIn
	EventTypeHandCardClicked
	EventTypeBlindCardClicked
	EventTypeBlindExchangeSwap
	EventTypeBlindExchangeDone
	EventTypeTrickPlaceCard
	EventTypeWeTakeTrick
	EventTypeTheyTakeTrick
	EventTypeRoundOver
	EventTypeNewGame
	EventTypeQuit
)

func (g *Game) HandleClick(button *LabeledButton) {
	switch button.Event {
	case EventTypeSort:
		if button.Value == objects.SuitUnassigned {
			return
		}
		g.Round.Hand = objects.SortHand(g.Round.Hand, button.Value.(objects.Suit), true)
	case EventTypeBid:
		switch v := button.Value.(type) {
		case objects.Suit:
			g.Round.Bid.Suit = v
		case int:
			g.Round.Bid.Take = v
		}
	case EventTypeBidPass:
		g.Round.Bid.Take = 0
		g.Round.Bid.Suit = objects.SuitUnassigned
		g.NewRound()
	case EventTypeBidLockIn:
		if g.Round.Bid.Suit == objects.SuitUnassigned || g.Round.Bid.Take < 6 {
			return
		}
		g.Round.Bid.Locked = true
		g.Round.Hand = objects.SetTrumps(g.Round.Hand, g.Round.Bid.Suit)
		for i := range g.Round.Players {
			teamWe := false
			if i == 1 {
				teamWe = true
			}
			g.Round.Players[i].Hand = objects.SortHand(g.Round.Players[i].Hand, g.Round.Bid.Suit, teamWe)
			g.Round.Players[i].Hand = objects.SetTrumps(g.Round.Players[i].Hand, g.Round.Bid.Suit)
			g.Round.Blind = objects.SetTrumps(g.Round.Blind, g.Round.Bid.Suit)
		}
		g.NextPhase()
	case EventTypeHandCardClicked:
		i := button.Value.(int)
		cardClicked := g.Round.Hand[i]
		if g.Round.Phase == PhaseTypeBlindExchange {
			if g.Round.HandExchangeCards[i] == nil {
				g.Round.HandExchangeCards[i] = cardClicked
			} else {
				g.Round.HandExchangeCards[i] = nil
			}
			g.Round.Hand[i].Selected = !cardClicked.Selected
		} else if g.Round.Phase == PhaseTypeTrick && g.Round.TrickCard != trickPhaseBlockAction && g.Round.TrickCard != trickPhaseBlockActionEvaluate {
			if len(g.Round.Tricks[g.Round.OnTrick].Cards) != 0 {
				botLeadWithCard := g.Round.Tricks[g.Round.OnTrick].Cards[0]
				if cardClicked.Suit != botLeadWithCard.Suit && hasSuit(g.Round.Hand, botLeadWithCard.Suit) {
					return
				}
			}
			g.Round.TrickCard = i
			for j := range g.Round.Hand {
				if j != i {
					g.Round.Hand[j].Selected = false
				}
			}
			g.Round.Hand[i].Selected = true
		}
	case EventTypeBlindCardClicked:
		i := button.Value.(int)
		card := g.Round.Blind[i]

		if g.Round.BlindExchangeCards[i] == nil {
			g.Round.BlindExchangeCards[i] = card
		} else {
			g.Round.BlindExchangeCards[i] = nil
		}
		g.Round.Blind[i].Selected = !g.Round.Blind[i].Selected
	case EventTypeBlindExchangeSwap:
		if sumOfNotNil(g.Round.HandExchangeCards) != sumOfNotNil(g.Round.BlindExchangeCards) {
			return
		}
		if sumOfNotNil(g.Round.HandExchangeCards) == 0 && sumOfNotNil(g.Round.BlindExchangeCards) == 0 {
			return
		}
		outOfHandPlaceholder := []*objects.Card{}
		outOfBlindPlaceholder := []*objects.Card{}
		for i, card := range g.Round.HandExchangeCards {
			if card != nil {
				outOfHandPlaceholder = append(outOfHandPlaceholder, card)
				g.Round.Hand[i] = nil
			}
		}

		for i, card := range g.Round.BlindExchangeCards {
			if card != nil {
				outOfBlindPlaceholder = append(outOfBlindPlaceholder, card)
				g.Round.Blind[i] = nil
			}
		}

		for _, toHand := range outOfBlindPlaceholder {
			for i, card := range g.Round.Hand {
				if card == nil {
					toHand.Selected = false
					g.Round.Hand[i] = toHand
					break
				}
			}
		}

		for _, toBlind := range outOfHandPlaceholder {
			for i, card := range g.Round.Blind {
				if card == nil {
					toBlind.Selected = false
					g.Round.Blind[i] = toBlind
					break
				}
			}
		}
		g.Round.HandExchangeCards = make([]*objects.Card, 10)
		g.Round.BlindExchangeCards = make([]*objects.Card, 5)

	case EventTypeBlindExchangeDone:
		g.NextPhase()
	case EventTypeTrickPlaceCard:
		if g.Round.TrickCard == trickPhaseNoCardChosen || g.Round.TrickCard == trickPhaseBlockAction || g.Round.TrickCard == trickPhaseBlockActionEvaluate {
			return
		}
		playedCard := g.Round.Hand[g.Round.TrickCard]
		playedCard.PlayedBy = -1
		g.Round.Tricks[g.Round.OnTrick].Cards = append(g.Round.Tricks[g.Round.OnTrick].Cards, playedCard)
		g.Round.Hand = popCard(g.Round.Hand, g.Round.TrickCard)
		g.Round.TrickCard = trickPhaseBlockAction
		g.Round.TrickPlaysNext = 0
		speaker.Play(g.CardPlaceSound)
		g.CardPlaceSound.Seek(0)
	case EventTypeWeTakeTrick:
		g.Round.TeamWeTricks++
		if g.Round.OnTrick < 9 {
			g.Round.OnTrick++
		}
		if g.Round.TrickPlaysNext != -1 {
			g.Round.TrickCard = trickPhaseBlockAction
		} else {
			g.Round.TrickCard = trickPhaseNoCardChosen
		}
	case EventTypeTheyTakeTrick:
		g.Round.TeamTheyTricks++
		if g.Round.OnTrick < 9 {
			g.Round.OnTrick++
		}
		if g.Round.TrickPlaysNext != -1 {
			g.Round.TrickCard = trickPhaseBlockAction
		} else {
			g.Round.TrickCard = trickPhaseNoCardChosen
		}
	case EventTypeRoundOver:
		score := g.Round.Bid.Score()
		if g.Round.TeamWeTricks >= g.Round.Bid.Take {
			g.TeamWeScore += score
			g.Round.TeamWeScore = score
		} else {
			g.TeamWeScore -= score
			g.Round.TeamWeScore = 0 - score
		}
		g.TeamTheyScore += 10 * g.Round.TeamTheyTricks
		g.Round.TeamTheyScore = 10 * g.Round.TeamTheyTricks
		speaker.Play(g.DeckShuffleSound)
		g.DeckShuffleSound.Seek(0)
		g.NewRound()
	case EventTypeNewGame:
		g.TeamWeScore = 0
		g.TeamTheyScore = 0
		g.Round = nil
		g.PreviousRounds = []*Round{}
		g.NewRound()
	case EventTypeQuit:
		os.Exit(0)
	default:
		fmt.Printf("\nUnknown Button Event: %+v\n", button)
	}
}

func sumOfNotNil(cards []*objects.Card) int {
	var sum int
	for _, card := range cards {
		if card != nil {
			sum++
		}
	}
	return sum
}

func popCard(hand []*objects.Card, i int) []*objects.Card {
	return append(hand[:i], hand[i+1:]...)
}

func hasSuit(hand []*objects.Card, suit objects.Suit) bool {
	for _, card := range hand {
		if card.Suit == suit {
			return true
		}
	}
	return false
}
