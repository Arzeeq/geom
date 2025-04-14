package main

import (
	"image/color"

	"github.com/Arzeeq/geom"
)

func main() {
	w, h := 1000, 800
	canvas := geom.NewCanvas(w, h, 20, 1)
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.DrawLine(0, 0, 5, 4)
	canvas.Stroke()
	canvas.SavePNG("canvas.png")
}
