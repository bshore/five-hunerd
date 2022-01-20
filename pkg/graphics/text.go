package graphics

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

func (m *Manager) AddText(t text.Text) {
	m.textList = append(m.textList, t)
}

func (m *Manager) drawText() {
	for _, text := range m.textList {
		text.Draw(m.window, pixel.IM)
	}
}
