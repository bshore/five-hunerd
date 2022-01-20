package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bshore/five-hunerd/pkg/objects"
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	// card spread
	increment   float64 = 0.0
	incrementBy float64 = objects.CardWidth / 4

	playerHandLocation = [][]float64{
		{300, 300},
		{300, 900},
		{900, 300},
		{900, 900},
	}

	blindLocation = []float64{100, 600}
)

type CardGraphic struct {
	SpriteSheet pixel.Picture
	Batch       *pixel.Batch
	Frames      []pixel.Rect
}

type Game struct {
	Deck             []*objects.Card
	Round            *Round
	PreviousRounds   []*Round
	TeamWeScore      int
	TeamTheyScore    int
	CardFaceGraphics *CardGraphic
	CardBackGraphics *CardGraphic
	CardPlaceSound   beep.StreamSeekCloser
	DeckShuffleSound beep.StreamSeekCloser

	atlas *text.Atlas
	imd   *imdraw.IMDraw
}

//
func NewGame(atlas *text.Atlas, imd *imdraw.IMDraw, face, back *CardGraphic, place, shuffle beep.StreamSeekCloser) *Game {
	cards := objects.SpriteSheetCardOrder()
	for i, frame := range face.Frames {
		cards[i].Frame = frame
	}
	cards = objects.RemoveTwosAndThrees(cards)
	game := &Game{
		Deck:             cards,
		CardFaceGraphics: face,
		CardBackGraphics: back,
		atlas:            atlas,
		imd:              imd,
		CardPlaceSound:   place,
		DeckShuffleSound: shuffle,
	}
	game.NewRound()
	return game
}

//
func (g *Game) NewRound() {
	if g.Round != nil {
		g.PreviousRounds = append(g.PreviousRounds, g.Round)
	}
	rand.Seed(time.Now().UnixNano())
	for i := range g.Deck {
		g.Deck[i].Trump = false
		g.Deck[i].Selected = false
	}
	g.Deck = objects.Shuffle(g.Deck)
	// shuffle once, then again anywhere between 0 and 10 more times
	for a := 0; a < rand.Intn(10); a++ {
		g.Deck = objects.Shuffle(g.Deck)
	}
	p1 := g.Deck[:10]
	p2 := g.Deck[10:20]
	p3 := g.Deck[20:30]
	p4 := g.Deck[30:40]
	blind := g.Deck[40:]
	g.Round = newRound(p1, p2, p3, p4, blind)
}

func (g *Game) GameFinished() bool {
	return (g.TeamWeScore >= 500 || g.TeamWeScore <= -500) || (g.TeamTheyScore >= 500 || g.TeamTheyScore <= -500)
}

func (g *Game) Render() []*LabeledButton {
	var buttons []*LabeledButton

	if g.GameFinished() {
		g.Round.Phase = PhaseTypeGameOver
	}

	switch g.Round.Phase {
	case PhaseTypeBid:
		buttons = append(buttons, g.RenderBiddingPhase()...)
	case PhaseTypeBlindExchange:
		buttons = append(buttons, g.RenderBlindExchangePhase()...)
	case PhaseTypeTrick:
		buttons = append(buttons, g.RenderTrickPhase()...)
		if g.Round.AllPlayersActed() {
			g.Round.TrickCard = trickPhaseBlockActionEvaluate
		} else if g.Round.TrickCard == trickPhaseBlockAction {
			thisPlayer := g.Round.TrickPlaysNext
			idx, playedCard := g.Round.Players[thisPlayer].DecideCard(g.Round.Tricks[g.Round.OnTrick], g.Round.Bid.Suit)
			playedCard.PlayedBy = thisPlayer
			g.Round.Players[thisPlayer].Hand = popCard(g.Round.Players[thisPlayer].Hand, idx)
			g.Round.Tricks[g.Round.OnTrick].Cards = append(g.Round.Tricks[g.Round.OnTrick].Cards, playedCard)
			if thisPlayer == 2 {
				g.Round.TrickPlaysNext = -1
				g.Round.TrickCard = trickPhaseNoCardChosen
			} else {
				g.Round.TrickPlaysNext++
				g.Round.TrickCard = trickPhaseBlockAction
			}
		}
	case PhaseTypeGameOver:
		return append(buttons, g.RenderGameOverPhase()...)
	}
	buttons = append(buttons, g.RenderScoreboard()...)

	// Sort Hand Buttons
	sortSpadesButton := g.NewLabeledButton([4]pixel.Vec{{X: 200, Y: 100}, {X: 290, Y: 100}, {X: 290, Y: 150}, {X: 200, Y: 150}}, "Sort Spades", colornames.Paleturquoise, 0.5, 0)
	sortSpadesButton.Value = objects.SuitSpade
	sortSpadesButton.Event = EventTypeSort
	sortClubsButton := g.NewLabeledButton([4]pixel.Vec{{X: 300, Y: 100}, {X: 390, Y: 100}, {X: 390, Y: 150}, {X: 300, Y: 150}}, "Sort Clubs", colornames.Paleturquoise, 0.5, 0)
	sortClubsButton.Value = objects.SuitClub
	sortClubsButton.Event = EventTypeSort
	sortDiamondsButton := g.NewLabeledButton([4]pixel.Vec{{X: 400, Y: 100}, {X: 490, Y: 100}, {X: 490, Y: 150}, {X: 400, Y: 150}}, "Sort Diamonds", colornames.Paleturquoise, 0.5, 0)
	sortDiamondsButton.Value = objects.SuitDiamond
	sortDiamondsButton.Event = EventTypeSort
	sortHeartsButton := g.NewLabeledButton([4]pixel.Vec{{X: 500, Y: 100}, {X: 590, Y: 100}, {X: 590, Y: 150}, {X: 500, Y: 150}}, "Sort Hearts", colornames.Paleturquoise, 0.5, 0)
	sortHeartsButton.Value = objects.SuitHeart
	sortHeartsButton.Event = EventTypeSort
	sortNoTrumpButton := g.NewLabeledButton([4]pixel.Vec{{X: 600, Y: 100}, {X: 690, Y: 100}, {X: 690, Y: 150}, {X: 600, Y: 150}}, "Sort No Trump", colornames.Paleturquoise, 0.5, 0)
	sortNoTrumpButton.Value = objects.SuitNo
	sortNoTrumpButton.Event = EventTypeSort

	buttons = append(buttons, sortSpadesButton, sortClubsButton, sortDiamondsButton, sortHeartsButton, sortNoTrumpButton)

	// rand.Seed(time.Now().UnixNano())
	cardBack := g.CardBackGraphics.Frames[6] //[rand.Intn(len(g.CardBackGraphics.Frames))]
	backSprite := pixel.NewSprite(g.CardBackGraphics.SpriteSheet, cardBack)

	increment = 0.0
	for i, card := range g.Round.Hand {
		coords := playerHandLocation[0]
		cardMatrix := pixel.IM.Moved(pixel.V(coords[0]+increment, coords[1]))
		cardSprite := pixel.NewSprite(g.CardFaceGraphics.SpriteSheet, card.Frame)
		cardSprite.Draw(g.CardFaceGraphics.Batch, cardMatrix)

		minX := coords[0] + increment - (objects.CardWidth / 2)
		minY := coords[1] - (objects.CardHeight / 2)
		maxX := coords[0] + increment - (objects.CardWidth / 4)
		maxY := coords[1] + (objects.CardHeight / 2)

		if i+1 == len(g.Round.Hand) {
			maxX = coords[0] + increment + (objects.CardWidth / 2)
		}

		cardRect := pixel.R(minX, minY, maxX, maxY)
		if card.Selected {
			cardPoints := cardRect.Vertices()
			g.imd.Color = colornames.Red
			g.imd.Push(cardPoints[0:]...)
			g.imd.Polygon(2)
		}
		cardButton := &LabeledButton{
			Rect:   cardRect,
			Matrix: pixel.IM,
			Event:  EventTypeHandCardClicked,
			Value:  i,
		}
		buttons = append(buttons, cardButton)

		increment += incrementBy
	}

	for playerNum, player := range g.Round.Players {
		coords := playerHandLocation[playerNum+1]
		increment = 0.0
		for i := range player.Hand {
			_ = i
			cardMatrix := pixel.IM.Moved(pixel.V(coords[0]+increment, coords[1]))
			backSprite.Draw(g.CardBackGraphics.Batch, cardMatrix)
			increment += incrementBy
		}
	}

	if g.Round.Phase != PhaseTypeBlindExchange {
		increment = 0.0
		for i := range g.Round.Blind {
			_ = i
			backSprite.Draw(g.CardBackGraphics.Batch, pixel.IM.Moved(pixel.V(blindLocation[0]+increment, blindLocation[1])))
			increment += incrementBy
		}
	}
	return buttons
}

func (g *Game) RenderScoreboard() []*LabeledButton {
	var buttons []*LabeledButton
	g.imd.Color = colornames.Black
	g.imd.Push(pixel.V(1400, 0), pixel.V(1400, 1080))
	g.imd.Line(2)

	scorecardHeaderButton := g.NewLabeledButton([4]pixel.Vec{{X: 1550, Y: 1010}, {X: 1755, Y: 1010}, {X: 1755, Y: 1079}, {X: 1550, Y: 1079}}, "Scoreboard", colornames.Lightgray, 1, 0)
	if g.Round.Bid.Locked {
		currentBidButton := g.NewLabeledButton([4]pixel.Vec{{X: 1550, Y: 980}, {X: 1755, Y: 980}, {X: 1755, Y: 1010}, {X: 1550, Y: 1010}}, fmt.Sprintf("Current Bid: %s", g.Round.Bid.String()), colornames.Lightgray, 1, 0)
		buttons = append(buttons, currentBidButton)
	}

	g.imd.Color = colornames.Black
	g.imd.Push(pixel.V(1400, 990), pixel.V(1920, 990))
	g.imd.Line(2)

	g.imd.Push(pixel.V(1660, 0), pixel.V(1660, 990))
	g.imd.Line(2)

	weButtonText := g.NewLabeledButton([4]pixel.Vec{{X: 1425, Y: 925}, {X: 1625, Y: 925}, {X: 1625, Y: 975}, {X: 1425, Y: 975}}, "We", colornames.Lightgray, 1, 0)
	theyButtonText := g.NewLabeledButton([4]pixel.Vec{{X: 1685, Y: 925}, {X: 1825, Y: 925}, {X: 1825, Y: 975}, {X: 1685, Y: 975}}, "They", colornames.Lightgray, 1, 0)

	yMax := 900.0
	for _, round := range g.PreviousRounds {
		wePoints := g.NewLabeledButton([4]pixel.Vec{{X: 1425, Y: yMax - 50}, {X: 1625, Y: yMax - 50}, {X: 1625, Y: yMax}, {X: 1425, Y: yMax}}, fmt.Sprintf("%d", round.TeamWeScore), colornames.Lightgray, 1, 0)
		bidTxt := g.NewLabeledButton([4]pixel.Vec{{X: 1500, Y: yMax - 50}, {X: 1675, Y: yMax - 50}, {X: 1675, Y: yMax}, {X: 1500, Y: yMax}}, fmt.Sprint(round.Bid.String()), colornames.Lightgray, 1, 0)
		theyPoints := g.NewLabeledButton([4]pixel.Vec{{X: 1685, Y: yMax - 50}, {X: 1825, Y: yMax - 50}, {X: 1825, Y: yMax}, {X: 1685, Y: yMax}}, fmt.Sprintf("%d", round.TeamTheyScore), colornames.Lightgray, 1, 0)
		yMax -= 50.0
		buttons = append(buttons, wePoints, bidTxt, theyPoints)
	}

	g.imd.Color = colornames.Black
	g.imd.Push(pixel.V(1400, yMax-10), pixel.V(1920, yMax-10))
	g.imd.Line(2)

	weCurrentScore := g.NewLabeledButton([4]pixel.Vec{{X: 1425, Y: yMax - 50}, {X: 1625, Y: yMax - 50}, {X: 1625, Y: yMax}, {X: 1425, Y: yMax}}, fmt.Sprintf("Total: %d", g.TeamWeScore), colornames.Lightgray, 1, 0)
	theyCurrentScore := g.NewLabeledButton([4]pixel.Vec{{X: 1685, Y: yMax - 50}, {X: 1825, Y: yMax - 50}, {X: 1825, Y: yMax}, {X: 1685, Y: yMax}}, fmt.Sprintf("Total: %d", g.TeamTheyScore), colornames.Lightgray, 1, 0)
	buttons = append(buttons, weCurrentScore, theyCurrentScore)

	return append(buttons, scorecardHeaderButton, weButtonText, theyButtonText)
}

func (g *Game) NextPhase() {
	switch g.Round.Phase {
	case PhaseTypeBid:
		g.Round.Phase = PhaseTypeBlindExchange
	case PhaseTypeBlindExchange:
		g.Round.Phase = PhaseTypeTrick
	case PhaseTypeTrick:
		g.Round.Phase = PhaseTypeBid
	}
}
