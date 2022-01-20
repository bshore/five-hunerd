package objects

import (
	"math/rand"
	"sort"
	"time"
)

func RemoveTwosAndThrees(fullDeck []*Card) []*Card {
	var newDeck []*Card
	for _, card := range fullDeck {
		if card.Rank == RankTwo || card.Rank == RankThree {
			continue
		}
		newDeck = append(newDeck, card)
	}
	return newDeck
}

func Shuffle(deck []*Card) []*Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func SetTrumps(hand []*Card, suit Suit) []*Card {
	var sortedHand []*Card
	for _, card := range hand {
		if card.Suit == suit {
			card.Trump = true
		}
		sortedHand = append(sortedHand, card)
	}
	return sortedHand
}

func SortHand(hand []*Card, suit Suit, teamWe bool) []*Card {
	var sortedHand []*Card
	// Break out the hand into suits first
	var spades, clubs, diamonds, hearts, no []*Card
	for _, card := range hand {
		card.TeamWe = teamWe
		if card.Rank == RankJoker {
			card.Suit = suit
			card.Trump = true
		}
		// if previously marked as Left Jack, revert to Jack & orignal suit
		if card.Rank == RankLeftJack {
			card.Suit = card.OriginalSuit
			card.Rank = RankJack
			card.Trump = false
		}
		// if previously marked as Right Jack, revert back to Jack
		if card.Rank == RankRightJack {
			card.Rank = RankJack
			card.Trump = false
		}
		switch card.Suit {
		case SuitSpade:
			if card.Rank == RankJack && suit == SuitSpade {
				card.Rank = RankRightJack
				card.Trump = true
			} else if card.Rank == RankJack && suit == SuitClub {
				card.OriginalSuit = SuitSpade
				card.Suit = SuitClub
				card.Rank = RankLeftJack
				card.Trump = true
				clubs = append(clubs, card)
				continue
			}
			spades = append(spades, card)
		case SuitClub:
			if card.Rank == RankJack && suit == SuitClub {
				card.Rank = RankRightJack
				card.Trump = true
			} else if card.Rank == RankJack && suit == SuitSpade {
				card.OriginalSuit = SuitClub
				card.Suit = SuitSpade
				card.Rank = RankLeftJack
				card.Trump = true
				spades = append(spades, card)
				continue
			}
			clubs = append(clubs, card)
		case SuitDiamond:
			if card.Rank == RankJack && suit == SuitDiamond {
				card.Rank = RankRightJack
				card.Trump = true
			} else if card.Rank == RankJack && suit == SuitHeart {
				card.OriginalSuit = SuitDiamond
				card.Suit = SuitHeart
				card.Rank = RankLeftJack
				card.Trump = true
				hearts = append(hearts, card)
				continue
			}
			diamonds = append(diamonds, card)
		case SuitHeart:
			if card.Rank == RankJack && suit == SuitHeart {
				card.Rank = RankRightJack
				card.Trump = true
			} else if card.Rank == RankJack && suit == SuitDiamond {
				card.OriginalSuit = SuitHeart
				card.Suit = SuitDiamond
				card.Rank = RankLeftJack
				card.Trump = true
				diamonds = append(diamonds, card)
				continue
			}
			hearts = append(hearts, card)
		case SuitNo:
			switch suit {
			case SuitSpade:
				spades = append(spades, card)
			case SuitClub:
				clubs = append(clubs, card)
			case SuitDiamond:
				diamonds = append(diamonds, card)
			case SuitHeart:
				hearts = append(hearts, card)
			case SuitNo:
				no = append(no, card)
			default:
			}
		default:
		}
	}
	spades = sortCards(spades)
	sortedHand = append(sortedHand, spades...)
	diamonds = sortCards(diamonds)
	sortedHand = append(sortedHand, diamonds...)
	clubs = sortCards(clubs)
	sortedHand = append(sortedHand, clubs...)
	hearts = sortCards(hearts)
	sortedHand = append(sortedHand, hearts...)
	sortedHand = append(sortedHand, no...)
	return sortedHand
}

func sortCards(cards []*Card) []*Card {
	if len(cards) <= 1 {
		return cards
	}
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	return cards
}
