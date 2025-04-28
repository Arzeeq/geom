package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	canvas := geom.NewCanvas(1000, 800, 20, 1)

	a := 1.0
	b := 1.0
	x := func(t float64) float64 { return (a + b*t) * math.Cos(t) }
	y := func(t float64) float64 { return (a + b*t) * math.Sin(t) }
	x1 := func(t float64) float64 { return b*math.Cos(t) - (a+b*t)*math.Sin(t) }
	y1 := func(t float64) float64 { return b*math.Sin(t) + (a+b*t)*math.Cos(t) }
	x2 := func(t float64) float64 { return -2*b*math.Sin(t) - (a+b*t)*math.Cos(t) }
	y2 := func(t float64) float64 { return 2*b*math.Cos(t) - (a+b*t)*math.Sin(t) }
	f := geom.NewFunc(x, y, x1, y1, x2, y2)

	n := 200
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
	t := 2 * math.Pi
	tangent := geom.NewVector2D(f.X1(t), y1(t))
	norm := geom.NewVector2D(-f.Y1(t), f.X1(t))
	tangent.ScaleToLen(2)
	norm.ScaleToLen(2)
	canvas.DrawVector(geom.NewPoint2D(f.X(t), f.Y(t)), tangent)
	canvas.DrawVector(geom.NewPoint2D(f.X(t), f.Y(t)), norm)
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.Stroke()

	// radius of curvature
	fmt.Println(f.CurvatureRadius(t))

	if err := canvas.SavePNG("spiral.png"); err != nil {
		panic(err)
	}
}
