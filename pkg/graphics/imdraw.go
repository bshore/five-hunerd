package graphics

import "github.com/faiface/pixel/imdraw"

func (m *Manager) AddIMDraw(imd imdraw.IMDraw) {
	m.imdrawList = append(m.imdrawList, imd)
}

func (m *Manager) drawIMDs() {
	for _, imd := range m.imdrawList {
		imd.Draw(m.window)
	}
}
