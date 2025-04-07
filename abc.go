package geom

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

func DrawPolygon(p Polygon, canvas *gg.Context, color color.Color, scale float64) {
	canvas.SetColor(color)
	canvas.MoveTo(p[0][0]*scale, p[0][1]*scale)
	for _, p := range p {
		canvas.LineTo(p[0]*scale, p[1]*scale)
	}
	canvas.LineTo(p[0][0]*scale, p[0][1]*scale)
	canvas.Stroke()
}

func FindScale(polygons []Polygon, width, height int) float64 {
	x := AbsMaxX(polygons)
	y := AbsMaxY(polygons)
	return math.Min(0.9*float64(width/2)/x, 0.9*float64(height/2)/y)
}

func DrawPolygonTransformation(p Polygon, a AffineTransform2D, filename string) {
	width := 800
	height := 600
	pT := a.TransformPolygon(p)
	k := FindScale([]Polygon{p, pT}, width, height)

	canvas := gg.NewContext(width, height)
	canvas.SetRGB(0.9, 0.9, 0.9)
	canvas.Clear()

	canvas.Translate(float64(width/2), float64(height/2))
	canvas.Scale(1, -1)
	DrawAxis(canvas, width, height)

	DrawPolygon(p, canvas, color.RGBA{0, 0, 255, 255}, k)
	DrawPolygon(pT, canvas, color.RGBA{255, 0, 0, 255}, k)

	canvas.SavePNG(filename)
}

func DrawAxis(dc *gg.Context, w, h int) {
	w1 := float64(w)
	h1 := float64(h)

	// Длина стрелок и отступ
	arrowSize := 10.0
	padding := 20.0

	// Цвет осей
	axisColor := color.RGBA{0, 0, 0, 255} // Черный

	// Рисуем ось X
	dc.SetColor(axisColor)
	dc.DrawLine(-w1/2+padding, 0, w1/2-padding, 0)
	dc.Stroke()

	// Рисуем стрелку оси X
	dc.DrawLine(w1/2-padding, 0, w1/2-padding-arrowSize, arrowSize)
	dc.DrawLine(w1/2-padding, 0, w1/2-padding-arrowSize, -arrowSize)
	dc.Stroke()

	// Рисуем ось Y
	dc.DrawLine(0, -h1/2+padding, 0, h1/2-padding)
	dc.Stroke()

	// Рисуем стрелку оси Y
	dc.DrawLine(0, h1/2-padding, arrowSize, h1/2-padding-arrowSize)
	dc.DrawLine(0, h1/2-padding, -arrowSize, h1/2-padding-arrowSize)
	dc.Stroke()

	// Подписи осей (с учетом инверсии Y)
	dc.Scale(1, -1) // Временно возвращаем обычное направление Y для текста
	dc.SetColor(color.RGBA{0, 0, 200, 255})

	// Подпись оси X
	dc.DrawStringAnchored("X", w1/2-padding-15, 15, 0.5, 0.5)

	// Подпись оси Y
	dc.DrawStringAnchored("Y", -25, -(h1/2 - padding - 15), 0.5, 0.5)

	dc.Scale(1, -1) // Возвращаем математическое направление осей
}
