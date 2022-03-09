package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"time"

	_ "image/png"

	"github.com/bshore/five-hunerd/pkg/game"
	"github.com/bshore/five-hunerd/pkg/objects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const (
	// cards stop appearing in the spritesheet after x=700 and below y=1330
	xCutoff = 700
	yCutoff = 1330
)

var (
	frames = 0
	second = time.Tick(time.Second)
)

func isCutoff(x, y float64) bool {
	return x >= xCutoff && y <= yCutoff
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func run() {
	// monitors := pixelgl.Monitors()
	config := pixelgl.WindowConfig{
		Title:     "Five-Hunerd",
		Bounds:    pixel.R(0, 0, 1920, 1080),
		VSync:     true,
		Resizable: true,
		// Monitor: monitors[2],
	}
	window, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	window.SetSmooth(true)
	window.Clear(colornames.Lightgray)

	// Load up all the card faces
	cardFaceSS, err := loadPicture("assets/playingCards.png")
	if err != nil {
		panic(err)
	}
	cardFaceBatch := pixel.NewBatch(&pixel.TrianglesData{}, cardFaceSS)
	var cardFaceFrames []pixel.Rect

	// Since this spritesheet's sprites start at the top left, go left to right and top to bottom
	for x := cardFaceSS.Bounds().Min.X; x < cardFaceSS.Bounds().Max.X; x += objects.CardWidth {
		for y := cardFaceSS.Bounds().Max.Y; y > cardFaceSS.Bounds().Min.Y; y -= objects.CardHeight {
			if isCutoff(x, y) {
				continue
			}
			cardFaceFrames = append(cardFaceFrames, pixel.R(x, y, x+objects.CardWidth, y-objects.CardHeight))
		}
	}

	cardFaceGraphics := &game.CardGraphic{
		SpriteSheet: cardFaceSS,
		Batch:       cardFaceBatch,
		Frames:      cardFaceFrames,
	}

	cardBackSS, err := loadPicture("assets/playingCardBacks.png")
	if err != nil {
		panic(err)
	}
	cardBackBatch := pixel.NewBatch(&pixel.TrianglesData{}, cardBackSS)
	var cardBackFrames []pixel.Rect
	for x := cardBackSS.Bounds().Min.X; x < cardBackSS.Bounds().Max.X; x += objects.CardWidth {
		for y := cardBackSS.Bounds().Min.Y; y < cardBackSS.Bounds().Max.Y; y += objects.CardHeight {
			cardBackFrames = append(cardBackFrames, pixel.R(x, y, x+objects.CardWidth, y+objects.CardHeight))
		}
	}

	cardBackGraphics := &game.CardGraphic{
		SpriteSheet: cardBackSS,
		Batch:       cardBackBatch,
		Frames:      cardBackFrames,
	}

	face, err := loadTTF("./assets/FreeUniversal-Regular.ttf", 24)
	if err != nil {
		panic(err)
	}

	var basicAtlas = text.NewAtlas(face, text.ASCII)

	imd := imdraw.New(nil)

	cardPlaceOGG, err := os.Open("./assets/cardPlace1.ogg")
	if err != nil {
		panic(err)
	}
	defer cardPlaceOGG.Close()

	cardPlaceSound, format, err := vorbis.Decode(cardPlaceOGG)
	if err != nil {
		panic(err)
	}
	defer cardPlaceSound.Close()
	speaker.Init(format.SampleRate*2, format.SampleRate.N(time.Second/10))

	cardShuffleMP3, err := os.Open("./assets/cardShuffle.mp3")
	if err != nil {
		panic(err)
	}
	defer cardShuffleMP3.Close()

	cardShuffleSound, _, err := mp3.Decode(cardShuffleMP3)
	if err != nil {
		panic(err)
	}
	defer cardShuffleSound.Close()
	speaker.Play(cardShuffleSound)
	cardShuffleSound.Seek(0)

	g := game.NewGame(basicAtlas, imd, cardFaceGraphics, cardBackGraphics, cardPlaceSound, cardShuffleSound)

	buttons := []*game.LabeledButton{}

	for !window.Closed() {
		cardFaceBatch.Clear()
		cardBackBatch.Clear()
		imd.Clear()
		for _, button := range buttons {
			if button.Text != nil {
				button.Text.Clear()
			}
		}

		buttons := g.Render()

		// Check mouse position and sort hands by suit clicked on
		if window.JustPressed(pixelgl.MouseButtonLeft) {
			pos := window.MousePosition()
			for _, button := range buttons {
				if button.Rect.Contains(pos) {
					g.HandleClick(button)
				}
			}
		}

		window.Clear(colornames.Lightgray)
		g.CardFaceGraphics.Batch.Draw(window)
		g.CardBackGraphics.Batch.Draw(window)
		imd.Draw(window)

		for _, button := range buttons {
			if button.Text != nil {
				button.Text.Draw(window, button.Matrix)
			}
		}

		window.Update()

		// calculate FPS
		frames++
		select {
		case <-second:
			window.SetTitle(fmt.Sprintf("%s | FPS: %d", config.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
