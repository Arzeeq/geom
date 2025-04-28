package main

import (
	"image/color"

	"github.com/Arzeeq/geom"
	"github.com/Arzeeq/geom/pkg/gen"
)

func main() {
	w, h, cage, scale := 1000, 800, 20, 1.0
	canvas := geom.NewCanvas(w, h, cage, scale)

	n := 100
	polygon := geom.Polygon{
		geom.NewPoint2D(15, 0),
		geom.NewPoint2D(11, 11),
		geom.NewPoint2D(0, 15),
		geom.NewPoint2D(-11, 11),
		geom.NewPoint2D(-15, 0),
		geom.NewPoint2D(-11, -11),
		geom.NewPoint2D(0, -15),
		geom.NewPoint2D(11, -11),
	}

	points := make([]geom.Point2D, n)
	for i := range n {
		points[i] = genPoint(w, h, cage, scale)
	}

	inside := make([]geom.Point2D, 0)
	outside := make([]geom.Point2D, 0)

	for _, p := range points {
		if polygon.ContainsRayMethod(p) {
			inside = append(inside, p)
		} else {
			outside = append(outside, p)
		}
	}

	canvas.DrawPolygon(polygon)
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.Stroke()

	for _, p := range inside {
		canvas.DrawPoint(p.X(), p.Y(), 1)
	}
	canvas.SetColor(color.RGBA{0, 255, 0, 255})
	canvas.Stroke()

	for _, p := range outside {
		canvas.DrawPoint(p.X(), p.Y(), 1)
	}
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.Stroke()

	canvas.SavePNG("point_in_rect.png")
}

func genPoint(w, h, cage int, scale float64) geom.Point2D {
	borderX := float64(w/cage/2) * scale
	borderY := float64(h/cage/2) * scale
	x := gen.RandFloat64Between(-borderX, borderX)
	y := gen.RandFloat64Between(-borderY, borderY)
	return geom.NewPoint2D(x, y)
}
