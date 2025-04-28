package main

import (
	"image"
	"image/color"
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	bresenhamGif()
	simpleGif()
}

func bresenhamGif() {
	n := 100
	t1 := 0.0
	t2 := 2 * math.Pi
	images := make([]image.Image, 0)
	for t := t1; t <= t2; t += (t2 - t1) / float64(n) {
		canvas := geom.NewCanvas(200, 200, 20, 1)
		canvas.SetColor(color.Black)
		canvas.DrawSegmentBresenham(20, 20, int(20+50*math.Cos(t)), int(20+50*math.Sin(t)))
		images = append(images, canvas.Image())
	}

	err := geom.CreateGif(images, "bresenham.gif")
	if err != nil {
		panic(err)
	}
}

func simpleGif() {
	n := 100
	t1 := 0.0
	t2 := 2 * math.Pi
	images := make([]image.Image, 0)
	for t := t1; t <= t2; t += (t2 - t1) / float64(n) {
		canvas := geom.NewCanvas(200, 200, 20, 1)
		canvas.SetColor(color.Black)
		canvas.DrawSegmentSimpleAlg(20, 20, int(20+50*math.Cos(t)), int(20+50*math.Sin(t)))
		images = append(images, canvas.Image())
	}

	err := geom.CreateGif(images, "simple.gif")
	if err != nil {
		panic(err)
	}
}
