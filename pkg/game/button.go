package game

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type LabeledButton struct {
	pixel.Rect
	*text.Text
	pixel.Matrix
	Value interface{}
	Event EventType
}

func (g *Game) NewLabeledButton(coords [4]pixel.Vec, msg string, color color.RGBA, scale, thickness float64) *LabeledButton {
	button := &LabeledButton{
		Matrix: pixel.IM,
	}
	g.imd.Color = color
	g.imd.Push(coords[0:]...)
	g.imd.Polygon(thickness)

	buttonText := text.New(pixel.V(coords[0].X+10, coords[0].Y+25), g.atlas)
	buttonText.Color = colornames.Black
	if scale != 0 {
		button.Matrix = pixel.IM.Scaled(buttonText.Orig, scale)
	}
	fmt.Fprintln(buttonText, msg)

	buttonRect := pixel.R(coords[0].X, coords[0].Y, coords[2].X, coords[2].Y)
	button.Rect = buttonRect
	button.Text = buttonText

	return button
}
