package abc

import "github.com/fogleman/gg"

func NewCanvas(w, h int, axis bool) *gg.Context {
	canvas := gg.NewContext(w, h)
	canvas.SetRGB(0.9, 0.9, 0.9)
	canvas.Clear()

	canvas.Translate(float64(w/2), float64(h/2))
	canvas.Scale(1, -1)
	if axis {
		DrawAxis(canvas, w, h)
	}
	return canvas
}
