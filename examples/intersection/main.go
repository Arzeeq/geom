package main

import (
	"image/color"

	"github.com/Arzeeq/geom"
	"github.com/Arzeeq/geom/pkg/gen"
)

func main() {
	width, height := 1000, 800
	cage, scale := 20, 1.0
	canvas := geom.NewCanvas(width, height, cage, scale)

	n := 10
	segments := make([]geom.Segment, n)
	for i := range n {
		p1 := genPoint(width, height, cage, scale)
		p2 := genPoint(width, height, cage, scale)
		canvas.DrawLine(p1.X(), p1.Y(), p2.X(), p2.Y())
		segments[i] = geom.NewSegment(p1, p2)
	}
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.Stroke()

	for i := range n {
		for j := range n {
			p, ok := geom.SegmentsIntersection(segments[i], segments[j])
			if ok {
				canvas.DrawPoint(p.X(), p.Y(), 1)
			}
		}
	}
	canvas.SetColor(color.RGBA{0, 255, 0, 255})
	canvas.Stroke()

	canvas.SavePNG("intersection.png")
}

func genPoint(w, h, cage int, scale float64) geom.Point2D {
	borderX := float64(w/cage/2) * scale
	borderY := float64(h/cage/2) * scale
	x := gen.RandFloat64Between(-borderX, borderX)
	y := gen.RandFloat64Between(-borderY, borderY)
	return geom.NewPoint2D(x, y)
}
