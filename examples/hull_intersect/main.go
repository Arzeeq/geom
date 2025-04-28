package main

import (
	"image/color"

	"github.com/Arzeeq/geom"
	"github.com/Arzeeq/geom/pkg/gen"
)

func main() {
	w, h, cage, scale := 1000, 800, 20, 1.0
	canvas := geom.NewCanvas(w, h, cage, scale)

	n, m := 30, 30
	e1 := make([]geom.Point2D, n)
	e2 := make([]geom.Point2D, m)

	for i := range e1 {
		e1[i] = genPoint(-15, 5, -15, 15)
	}

	for i := range e2 {
		e2[i] = genPoint(-5, 15, -15, 15)
	}

	hull1 := geom.ConvexHullGraham(e1)
	hull2 := geom.ConvexHullGraham(e2)
	intersectHull := geom.PolygonIntersection(hull1, hull2)

	canvas.DrawPolygon(hull1)
	canvas.DrawPolygon(hull2)
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.Stroke()

	canvas.DrawPolygon(intersectHull)
	canvas.SetColor(color.RGBA{255, 0, 255, 255})
	canvas.Stroke()

	for i := range e1 {
		if intersectHull.ContainsRayMethod(e1[i]) {
			canvas.SetColor(color.RGBA{0, 255, 0, 255})
		} else {
			canvas.SetColor(color.RGBA{0, 0, 255, 255})
		}
		canvas.DrawPoint(e1[i].X(), e1[i].Y(), 1)
		canvas.Stroke()
	}

	for i := range e2 {
		if intersectHull.ContainsRayMethod(e2[i]) {
			canvas.SetColor(color.RGBA{0, 255, 0, 255})
		} else {
			canvas.SetColor(color.RGBA{0, 0, 255, 255})
		}
		canvas.DrawPoint(e2[i].X(), e2[i].Y(), 1)
		canvas.Stroke()
	}

	canvas.SavePNG("hull_intersect.png")
}

func genPoint(x1, x2, y1, y2 float64) geom.Point2D {
	x := gen.RandFloat64Between(x1, x2)
	y := gen.RandFloat64Between(y1, y2)
	return geom.NewPoint2D(x, y)
}
