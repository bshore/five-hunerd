package graphics

import "github.com/faiface/pixel"

func (m *Manager) AddBatch(batch pixel.Batch) {
	m.batches = append(m.batches, batch)
}

func (m *Manager) drawBatches() {
	for _, batch := range m.batches {
		batch.Draw(m.window)
	}
}
