package main

import (
	"image/color"

	"github.com/Arzeeq/geom"
	"github.com/Arzeeq/geom/pkg/gen"
	"github.com/pzsz/voronoi"
)

func main() {
	w, h, cage, scale := 1000, 800, 20, 1.0
	borderX := float64(w/cage/2) * scale
	borderY := float64(h/cage/2) * scale

	n := 10
	sites := make([]voronoi.Vertex, n)
	for i := range n {
		x, y := genPoint(w, h, cage, scale)
		sites[i] = voronoi.Vertex{X: x, Y: y}
	}

	bbox := voronoi.NewBBox(-borderX, borderX, -borderY, borderY)
	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	drawVoronoiDiagram(diagram)
	drawDeloneTriangulation(diagram)
}

func genPoint(w, h, cage int, scale float64) (float64, float64) {
	borderX := float64(w/cage/2) * scale
	borderY := float64(h/cage/2) * scale
	x := gen.RandFloat64Between(-borderX, borderX)
	y := gen.RandFloat64Between(-borderY, borderY)
	return x, y
}

func drawDeloneTriangulation(d *voronoi.Diagram) {
	w, h, cage, scale := 1000, 800, 20, 1.0
	canvas := geom.NewCanvas(w, h, cage, scale)

	canvas.DrawDelone(d)
	canvas.SetColor(color.RGBA{0, 255, 0, 255})
	canvas.Stroke()

	canvas.DrawVoronoiPoints(d)
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.Stroke()

	canvas.SavePNG("delone.png")
}

func drawVoronoiDiagram(d *voronoi.Diagram) {
	w, h, cage, scale := 1000, 800, 20, 1.0
	canvas := geom.NewCanvas(w, h, cage, scale)

	canvas.DrawVoronoiEdges(d)
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.Stroke()

	canvas.DrawVoronoiPoints(d)
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.Stroke()

	canvas.SavePNG("voronoi.png")
}
