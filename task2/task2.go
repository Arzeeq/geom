package main

import (
	"arzeeq/geometry/geom"

	"github.com/fogleman/gg"
)

func main() {
	width := 800
	height := 600
	// a := 200.0
	// b := 140.0

	canvas := gg.NewContext(width, height)
	canvas.SetRGB(0.9, 0.9, 0.9)
	canvas.Clear()

	canvas.Translate(float64(width/2), float64(height/2))
	canvas.Scale(1, -1)
	geom.DrawAxis(canvas, width, height)

	//geom.DrawEllipse(a, b, canvas)
	//geom.DrawHypocycloid(3*50, 5*50, canvas)
	geom.DrawArhSpiral(10, 10, canvas)

	//canvas.SavePNG("ellipse.png")
	//canvas.SavePNG("hypocycloid.png")
	canvas.SavePNG("spiral.png")
}
