package objects

import "github.com/faiface/pixel"

const (
	CardWidth  = 140
	CardHeight = 190
)

type Suit int

const (
	SuitUnassigned Suit = iota
	SuitSpade
	SuitClub
	SuitDiamond
	SuitHeart
	SuitNo
)

func (s Suit) Int() int {
	return int(s)
}

func (s Suit) String() string {
	return []string{
		"Unassigned",
		"Spades",
		"Clubs",
		"Diamonds",
		"Hearts",
		"No Trump",
	}[s]
}

func (s Suit) Points() int {
	return []int{0, 40, 60, 80, 100, 120}[s]
}

type Rank int

const (
	RankTwo Rank = iota + 2
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
	RankLeftJack
	RankRightJack
	RankJoker
)

func (r Rank) Int() int {
	return int(r)
}

func (r Rank) String() string {
	return []string{
		"Unknown 0",
		"Unknown 1",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
		"LeftJack",
		"RightJack",
		"Joker",
	}[r]
}

type Card struct {
	Rank         Rank
	Suit         Suit
	OriginalSuit Suit
	Frame        pixel.Rect
	PlayedBy     int
	Trump        bool
	TeamWe       bool
	Selected     bool
}

// SpriteSheetCardOrder returns a slice of Cards in the order they are scanned from the SpriteSheet
func SpriteSheetCardOrder() []*Card {
	return []*Card{
		{RankQueen, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankKing, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankJack, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankAce, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTen, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankNine, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankEight, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSeven, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSix, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFive, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFour, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankThree, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTwo, SuitSpade, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankJoker, SuitNo, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankQueen, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankKing, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankJack, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankAce, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTen, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankNine, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankEight, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSeven, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSix, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFive, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFour, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankThree, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTwo, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankQueen, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankKing, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankJack, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankAce, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTen, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankNine, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankEight, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSeven, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSix, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFive, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFour, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankThree, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTwo, SuitDiamond, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankQueen, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankKing, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankJack, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankAce, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTen, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankNine, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankEight, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSeven, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankSix, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFive, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankFour, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankThree, SuitClub, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
		{RankTwo, SuitHeart, SuitUnassigned, pixel.Rect{}, 0, false, false, false},
	}
}
