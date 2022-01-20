package graphics

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	camPos       = pixel.ZV
	camSpeed     = 500.0
	camZoom      = 1.0
	camZoomSpeed = 1.2
)

func (m *Manager) updateCamera() {
	last := time.Now()
	dt := time.Since(last).Seconds()
	last = time.Now()

	mx := pixel.IM.Scaled(camPos, camZoom).Moved(m.window.Bounds().Center().Sub(camPos))
	m.window.SetMatrix(mx)

	// Moves camera around
	if m.window.Pressed(pixelgl.KeyA) || m.window.Pressed(pixelgl.KeyLeft) {
		camPos.X -= camSpeed * dt
	}
	if m.window.Pressed(pixelgl.KeyD) || m.window.Pressed(pixelgl.KeyRight) {
		camPos.X += camSpeed * dt
	}
	if m.window.Pressed(pixelgl.KeyS) || m.window.Pressed(pixelgl.KeyDown) {
		camPos.Y -= camSpeed * dt
	}
	if m.window.Pressed(pixelgl.KeyW) || m.window.Pressed(pixelgl.KeyUp) {
		camPos.Y += camSpeed * dt
	}
	// MouseScroll zooms in/out
	camZoom *= math.Pow(camZoomSpeed, m.window.MouseScroll().Y)
}
