package graphics

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type Manager struct {
	config     pixelgl.WindowConfig
	window     *pixelgl.Window
	cameraMX   pixel.Matrix
	batches    []pixel.Batch
	atlas      text.Atlas
	textList   []text.Text
	imdrawList []imdraw.IMDraw
}

func NewManager(config pixelgl.WindowConfig) (*Manager, error) {
	config = pixelgl.WindowConfig{
		Title:  "Five-Hunerd",
		Bounds: pixel.R(0, 0, 1920, 1080),
		VSync:  true,
	}
	w, err := pixelgl.NewWindow(config)
	if err != nil {
		return nil, err
	}
	return &Manager{
		window: w,
		config: config,
	}, nil
}

func (m *Manager) Run() {
	var frames int
	var second = time.Tick(time.Second)
	for !m.window.Closed() {
		m.updateCamera()
		m.drawBatches()
		m.drawText()
		m.drawIMDs()

		m.window.Update()
		// Calculate FPS and update Window Title
		frames++
		select {
		case <-second:
			m.window.SetTitle(fmt.Sprintf("%s | FPS: %d", m.config.Title, frames))
		}
	}
}
