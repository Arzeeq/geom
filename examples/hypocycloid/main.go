package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	width := 800
	height := 600
	canvas := geom.NewCanvas(width, height, 20, 1)

	//R - outter circle radius, r - inner circle radius
	r := 6.0
	R := 10.0
	k := R / r

	x := func(t float64) float64 { return r * (k - 1) * (math.Cos(t) + math.Cos((k-1)*t)/(k-1)) }
	y := func(t float64) float64 { return r * (k - 1) * (math.Sin(t) - math.Sin((k-1)*t)/(k-1)) }
	x1 := func(t float64) float64 { return r * (k - 1) * (-math.Sin(t) - math.Sin(t*(k-1))) }
	y1 := func(t float64) float64 { return r * (k - 1) * (math.Cos(t) - math.Cos(t*(k-1))) }
	x2 := func(t float64) float64 { return r * (k - 1) * (-math.Cos(t) - (k-1)*math.Cos(t*(k-1))) }
	y2 := func(t float64) float64 { return r * (k - 1) * (-math.Sin(t) + (k-1)*math.Sin(t*(k-1))) }
	f := geom.NewFunc(x, y, x1, y1, x2, y2)

	n := 100
	t1 := 0.0
	t2 := 6 * math.Pi
	for i := range n {
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		if i == 0 {
			canvas.MoveTo(f.X(t), f.Y(t))
		} else {
			canvas.LineTo(f.X(t), f.Y(t))
		}
	}

	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.Stroke()

	// draw tangent and norm
	t := 5.5 * math.Pi
	tangent := geom.NewVector2D(f.X1(t), f.Y1(t))
	norm := geom.NewVector2D(-f.Y1(t), f.X1(t))
	tangent.ScaleToLen(2)
	norm.ScaleToLen(2)
	canvas.DrawVector(geom.NewPoint2D(f.X(t), f.Y(t)), tangent)
	canvas.DrawVector(geom.NewPoint2D(f.X(t), f.Y(t)), norm)
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.Stroke()

	// radius of curvature
	fmt.Println(f.CurvatureRadius(t))

	canvas.SavePNG("hypocycloid.png")
}
