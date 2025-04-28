package main

import (
	"image/color"
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	width := 800
	height := 600
	canvas := geom.NewCanvas(width, height, 20, 1)

	// Ellipse axis
	a := 10.0
	b := 5.0
	x := func(t float64) float64 { return a * math.Cos(t) }
	y := func(t float64) float64 { return b * math.Sin(t) }
	x1 := func(t float64) float64 { return -a * math.Sin(t) }
	y1 := func(t float64) float64 { return b * math.Cos(t) }
	x2 := func(t float64) float64 { return -a * math.Cos(t) }
	y2 := func(t float64) float64 { return -b * math.Sin(t) }
	f := geom.NewFunc(x, y, x1, y1, x2, y2)
	drawEllipseApproximation(canvas, f)
	drawEvolute(canvas, f)
	err := canvas.SavePNG("ellipse.png")
	if err != nil {
		panic(err)
	}
}

func drawEllipseApproximation(canvas *geom.Canvas, f geom.Func) {
	// approximation degree
	n := 20

	t1 := 0.0
	t2 := 2 * math.Pi
	tangents := make([]geom.Line, n+1)
	for i := range n + 1 {
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		tangent, err := geom.NewLine(
			geom.NewPoint2D(f.X(t), f.Y(t)),
			geom.NewPoint2D(f.X(t)+f.X1(t), f.Y(t)+f.Y1(t)))
		if err != nil {
			panic(err)
		}
		tangents[i] = tangent
	}

	points := make([]geom.Point2D, n)
	for i := 1; i < n+1; i++ {
		p, ok := geom.LinesIntersection(tangents[i-1], tangents[i])
		if !ok {
			panic("lines do not intersect")
		}
		points[i-1] = p
	}

	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.DrawPolygon(points)
	canvas.Stroke()
}

func drawEvolute(canvas *geom.Canvas, f geom.Func) {
	// approximation degree
	n := 100

	t1 := 0.0
	t2 := 2 * math.Pi
	for i := range n + 1 {
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		if i == 0 {
			canvas.MoveTo(f.EvoluteX(t), f.EvoluteY(t))
		} else {
			canvas.LineTo(f.EvoluteX(t), f.EvoluteY(t))
		}
	}
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.Stroke()
}
