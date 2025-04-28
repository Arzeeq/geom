package main

import (
	"fmt"
	"image/color"

	"github.com/Arzeeq/geom"
	"github.com/Arzeeq/geom/pkg/gen"
)

func main() {
	graham()
	jarvis()
}

func genPoint(w, h, cage int, scale float64) geom.Point2D {
	borderX := float64(w/cage/2) * scale
	borderY := float64(h/cage/2) * scale
	x := gen.RandFloat64Between(-borderX, borderX)
	y := gen.RandFloat64Between(-borderY, borderY)
	return geom.NewPoint2D(x, y)
}

func graham() {
	w, h, cage, scale := 1000, 800, 20, 1.0
	canvas := geom.NewCanvas(w, h, cage, scale)

	n := 25
	points := make([]geom.Point2D, n)
	for i := range n {
		points[i] = genPoint(400, 400, cage, scale)
		canvas.DrawPoint(points[i].X(), points[i].Y(), 1)
	}

	hull := geom.ConvexHullGraham(points)
	fmt.Printf("area: %f, perimeter: %f\n", hull.Area(), hull.Perimeter())
	canvas.DrawPolygon(hull)
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.Stroke()

	if err := canvas.SavePNG("graham.png"); err != nil {
		panic(err)
	}
}

func jarvis() {
	w, h, cage, scale := 1000, 800, 20, 1.0
	canvas := geom.NewCanvas(w, h, cage, scale)

	n := 25
	points := make([]geom.Point2D, n)
	for i := range n {
		points[i] = genPoint(400, 400, cage, scale)
		canvas.DrawPoint(points[i].X(), points[i].Y(), 1)
	}

	hull := geom.ConvexHullJarvis(points)
	fmt.Printf("area: %f, perimeter: %f\n", hull.Area(), hull.Perimeter())
	canvas.DrawPolygon(hull)
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.Stroke()
	if err := canvas.SavePNG("jarvis.png"); err != nil {
		panic(err)
	}

}
