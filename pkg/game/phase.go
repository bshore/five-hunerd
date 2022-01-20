package game

import (
	"fmt"

	"github.com/bshore/five-hunerd/pkg/objects"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

var revealedBlindLocation = []float64{650, 600}
var trickCardLocation = []float64{650, 600}

type PhaseType int

const (
	PhaseTypeBid PhaseType = iota
	PhaseTypeBlindExchange
	PhaseTypeTrick
	PhaseTypeGameOver
)

func (g *Game) RenderBiddingPhase() []*LabeledButton {
	var buttons []*LabeledButton

	g.imd.Color = colornames.Black
	g.imd.Push(pixel.V(500, 450), pixel.V(1100, 750))
	g.imd.Rectangle(1)

	passButton := g.NewLabeledButton([4]pixel.Vec{{X: 525, Y: 650}, {X: 600, Y: 650}, {X: 600, Y: 725}, {X: 525, Y: 725}}, "Pass", colornames.Yellow, 1, 0)
	passButton.Event = EventTypeBidPass
	bidButton := g.NewLabeledButton([4]pixel.Vec{{X: 625, Y: 650}, {X: 900, Y: 650}, {X: 900, Y: 725}, {X: 625, Y: 725}}, fmt.Sprintf("%d %s", g.Round.Bid.Take, g.Round.Bid.Suit), colornames.Whitesmoke, 1, 0)
	lockInButton := g.NewLabeledButton([4]pixel.Vec{{X: 925, Y: 650}, {X: 1075, Y: 650}, {X: 1075, Y: 725}, {X: 925, Y: 725}}, "Lock In", colornames.Green, 1, 0)
	lockInButton.Event = EventTypeBidLockIn

	bidSixButton := g.NewLabeledButton([4]pixel.Vec{{X: 525, Y: 550}, {X: 600, Y: 550}, {X: 600, Y: 600}, {X: 525, Y: 600}}, "Six", colornames.Paleturquoise, 0.5, 0)
	bidSixButton.Value = 6
	bidSixButton.Event = EventTypeBid
	bidSevenButton := g.NewLabeledButton([4]pixel.Vec{{X: 625, Y: 550}, {X: 700, Y: 550}, {X: 700, Y: 600}, {X: 625, Y: 600}}, "Seven", colornames.Paleturquoise, 0.5, 0)
	bidSevenButton.Value = 7
	bidSevenButton.Event = EventTypeBid
	bidEightButton := g.NewLabeledButton([4]pixel.Vec{{X: 725, Y: 550}, {X: 800, Y: 550}, {X: 800, Y: 600}, {X: 725, Y: 600}}, "Eight", colornames.Paleturquoise, 0.5, 0)
	bidEightButton.Value = 8
	bidEightButton.Event = EventTypeBid
	bidNineButton := g.NewLabeledButton([4]pixel.Vec{{X: 825, Y: 550}, {X: 900, Y: 550}, {X: 900, Y: 600}, {X: 825, Y: 600}}, "Nine", colornames.Paleturquoise, 0.5, 0)
	bidNineButton.Value = 9
	bidNineButton.Event = EventTypeBid
	bidTenButton := g.NewLabeledButton([4]pixel.Vec{{X: 925, Y: 550}, {X: 1000, Y: 550}, {X: 1000, Y: 600}, {X: 925, Y: 600}}, "Ten", colornames.Paleturquoise, 0.5, 0)
	bidTenButton.Value = 10
	bidTenButton.Event = EventTypeBid

	bidSpadesButton := g.NewLabeledButton([4]pixel.Vec{{X: 525, Y: 475}, {X: 600, Y: 475}, {X: 600, Y: 525}, {X: 525, Y: 525}}, "Spades", colornames.Darkgray, 0.5, 0)
	bidSpadesButton.Value = objects.SuitSpade
	bidSpadesButton.Event = EventTypeBid
	bidClubsButton := g.NewLabeledButton([4]pixel.Vec{{X: 625, Y: 475}, {X: 700, Y: 475}, {X: 700, Y: 525}, {X: 625, Y: 525}}, "Clubs", colornames.Darkgray, 0.5, 0)
	bidClubsButton.Value = objects.SuitClub
	bidClubsButton.Event = EventTypeBid
	bidDiamondsButton := g.NewLabeledButton([4]pixel.Vec{{X: 725, Y: 475}, {X: 800, Y: 475}, {X: 800, Y: 525}, {X: 725, Y: 525}}, "Diamonds", colornames.Palevioletred, 0.5, 0)
	bidDiamondsButton.Value = objects.SuitDiamond
	bidDiamondsButton.Event = EventTypeBid
	bidHeartsButton := g.NewLabeledButton([4]pixel.Vec{{X: 825, Y: 475}, {X: 900, Y: 475}, {X: 900, Y: 525}, {X: 825, Y: 525}}, "Hearts", colornames.Palevioletred, 0.5, 0)
	bidHeartsButton.Value = objects.SuitHeart
	bidHeartsButton.Event = EventTypeBid
	bidNoTrumpButton := g.NewLabeledButton([4]pixel.Vec{{X: 925, Y: 475}, {X: 1000, Y: 475}, {X: 1000, Y: 525}, {X: 925, Y: 525}}, "No Trump", colornames.Gold, 0.5, 0)
	bidNoTrumpButton.Value = objects.SuitNo
	bidNoTrumpButton.Event = EventTypeBid

	return append(buttons, []*LabeledButton{
		passButton, bidButton, lockInButton,
		bidSixButton, bidSevenButton, bidEightButton, bidNineButton, bidTenButton,
		bidSpadesButton, bidClubsButton, bidDiamondsButton, bidHeartsButton, bidNoTrumpButton,
	}...)
}

// TODO: Swap cards from player hand with the blind
func (g *Game) RenderBlindExchangePhase() []*LabeledButton {
	var buttons []*LabeledButton
	increment = 0.0
	for i, card := range g.Round.Blind {
		sprite := pixel.NewSprite(g.CardFaceGraphics.SpriteSheet, card.Frame)
		sprite.Draw(g.CardFaceGraphics.Batch, pixel.IM.Moved(pixel.V(revealedBlindLocation[0]+increment, revealedBlindLocation[1])))

		minX := revealedBlindLocation[0] + increment - (objects.CardWidth / 2)
		minY := revealedBlindLocation[1] - (objects.CardHeight / 2)
		maxX := revealedBlindLocation[0] + increment - (objects.CardWidth / 4)
		maxY := revealedBlindLocation[1] + (objects.CardHeight / 2)

		if i+1 == len(g.Round.Blind) {
			maxX = revealedBlindLocation[0] + increment + (objects.CardWidth / 2)
		}

		cardRect := pixel.R(minX, minY, maxX, maxY)

		if card.Selected {
			cardPoints := cardRect.Vertices()
			g.imd.Color = colornames.Green
			g.imd.Push(cardPoints[0:]...)
			g.imd.Polygon(2)
		}
		cardButton := &LabeledButton{
			Rect:   cardRect,
			Matrix: pixel.IM,
			Event:  EventTypeBlindCardClicked,
			Value:  i,
		}
		buttons = append(buttons, cardButton)

		increment += incrementBy
	}
	switchButton := g.NewLabeledButton([4]pixel.Vec{{X: 575, Y: 450}, {X: 625, Y: 450}, {X: 625, Y: 500}, {X: 575, Y: 500}}, "Swap", colornames.Paleturquoise, 0.5, 0)
	switchButton.Event = EventTypeBlindExchangeSwap
	doneButton := g.NewLabeledButton([4]pixel.Vec{{X: 675, Y: 450}, {X: 725, Y: 450}, {X: 725, Y: 500}, {X: 675, Y: 500}}, "Done", colornames.Green, 0.5, 0)
	doneButton.Event = EventTypeBlindExchangeDone
	buttons = append(buttons, switchButton, doneButton)
	return buttons
}

func (g *Game) RenderTrickPhase() []*LabeledButton {
	var buttons []*LabeledButton
	increment = 0.0
	g.imd.Color = colornames.Black
	g.imd.Push(pixel.V(500, 450), pixel.V(1100, 750))
	g.imd.Rectangle(1)

	tricksButton := g.NewLabeledButton([4]pixel.Vec{{X: 500, Y: 450}, {X: 1100, Y: 450}, {X: 1100, Y: 750}, {X: 500, Y: 750}}, "Play", colornames.Black, 0, 2)
	tricksButton.Event = EventTypeTrickPlaceCard
	buttons = append(buttons, tricksButton)

	if g.Round.TrickCard == trickPhaseBlockActionEvaluate && g.Round.Finished() {
		weTookButton := g.NewLabeledButton([4]pixel.Vec{{X: 600, Y: 550}, {X: 699, Y: 550}, {X: 699, Y: 600}, {X: 600, Y: 600}}, fmt.Sprintf("We Took %d", g.Round.TeamWeTricks), colornames.Paleturquoise, 0.5, 0)
		theyTookButton := g.NewLabeledButton([4]pixel.Vec{{X: 750, Y: 550}, {X: 849, Y: 550}, {X: 849, Y: 600}, {X: 750, Y: 600}}, fmt.Sprintf("They Took %d", g.Round.TeamTheyTricks), colornames.Paleturquoise, 0.5, 0)

		nextRoundButton := g.NewLabeledButton([4]pixel.Vec{{X: 700, Y: 400}, {X: 850, Y: 400}, {X: 850, Y: 448}, {X: 700, Y: 448}}, "Next Round", colornames.Green, 0, 2)
		nextRoundButton.Event = EventTypeRoundOver
		buttons = append(buttons, weTookButton, theyTookButton, nextRoundButton)
	} else {
		for i, card := range g.Round.Tricks[g.Round.OnTrick].Cards {
			sprite := pixel.NewSprite(g.CardFaceGraphics.SpriteSheet, card.Frame)
			sprite.Draw(g.CardFaceGraphics.Batch, pixel.IM.Moved(pixel.V(trickCardLocation[0]+increment, trickCardLocation[1])))
			increment += incrementBy * 2

			minX := trickCardLocation[0] + increment - (objects.CardWidth)
			minY := trickCardLocation[1] - (objects.CardHeight / 2)
			maxX := trickCardLocation[0] + increment - objects.CardWidth/2
			maxY := trickCardLocation[1] + (objects.CardHeight / 2)

			if i+1 == len(g.Round.Tricks[g.Round.OnTrick].Cards) {
				maxX = trickCardLocation[0] + increment
			}

			cardRect := pixel.R(minX, minY, maxX, maxY)
			if card.Trump {
				cardPoints := cardRect.Vertices()
				g.imd.Color = colornames.Blue
				g.imd.Push(cardPoints[0:]...)
				g.imd.Polygon(2)
			}
		}

		if g.Round.AllPlayersActed() {
			winningCard := g.Round.Tricks[g.Round.OnTrick].WinningCard()
			if winningCard.TeamWe {
				weTakeButton := g.NewLabeledButton([4]pixel.Vec{{X: 700, Y: 400}, {X: 850, Y: 400}, {X: 850, Y: 448}, {X: 700, Y: 448}}, "We Take", colornames.Green, 0, 2)
				weTakeButton.Event = EventTypeWeTakeTrick
				buttons = append(buttons, weTakeButton)
			} else {
				theyTakeButton := g.NewLabeledButton([4]pixel.Vec{{X: 700, Y: 400}, {X: 850, Y: 400}, {X: 850, Y: 448}, {X: 700, Y: 448}}, "They Take", colornames.Red, 0, 2)
				theyTakeButton.Event = EventTypeTheyTakeTrick
				buttons = append(buttons, theyTakeButton)
			}
			g.Round.TrickPlaysNext = winningCard.PlayedBy
		}
	}

	return buttons
}

func (g *Game) RenderGameOverPhase() []*LabeledButton {
	var buttons []*LabeledButton

	g.imd.Color = colornames.Black
	g.imd.Push(pixel.V(500, 450), pixel.V(1100, 750))
	g.imd.Rectangle(2)

	weTookButton := g.NewLabeledButton([4]pixel.Vec{{X: 600, Y: 650}, {X: 699, Y: 650}, {X: 699, Y: 700}, {X: 600, Y: 700}}, fmt.Sprintf("We: %d points", g.TeamWeScore), colornames.Paleturquoise, 0.5, 0)
	theyTookButton := g.NewLabeledButton([4]pixel.Vec{{X: 750, Y: 650}, {X: 849, Y: 650}, {X: 849, Y: 700}, {X: 750, Y: 700}}, fmt.Sprintf("They %d points", g.TeamTheyScore), colornames.Paleturquoise, 0.5, 0)

	newGameButton := g.NewLabeledButton([4]pixel.Vec{{X: 700, Y: 500}, {X: 850, Y: 500}, {X: 850, Y: 550}, {X: 700, Y: 550}}, "New Game", colornames.Green, 0, 2)
	newGameButton.Event = EventTypeNewGame

	quitButton := g.NewLabeledButton([4]pixel.Vec{{X: 700, Y: 452}, {X: 850, Y: 452}, {X: 850, Y: 498}, {X: 700, Y: 498}}, "Quit", colornames.Green, 0, 2)
	quitButton.Event = EventTypeQuit
	return append(buttons, weTookButton, theyTookButton, newGameButton, quitButton)
}
