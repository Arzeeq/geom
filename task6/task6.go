package main

import (
	"image/color"
	"strconv"

	"github.com/Arzeeq/geom"
	"github.com/Arzeeq/geom/pkg/utils"
)

func main() {
	canvasSize := 200
	randomPoints := make([]geom.Point2D, 100)
	for i := 0; i < len(randomPoints); i++ {
		x := utils.RandFloat64Between(-float64(canvasSize)/2, float64(canvasSize)/2)
		y := utils.RandFloat64Between(-float64(canvasSize)/2, float64(canvasSize)/2)
		randomPoints[i] = geom.NewPoint2D(x, y)
	}
	for j := range 10 {
		points := make([]geom.Point2D, 50)
		for i := 0; i < len(points); i++ {
			x := utils.RandFloat64Between(-float64(canvasSize)/2, float64(canvasSize)/2)
			y := utils.RandFloat64Between(-float64(canvasSize)/2, float64(canvasSize)/2)
			points[i] = geom.NewPoint2D(x, y)
		}
		convexHull := geom.BuildConvexHull(points)

		canvas := geom.NewCanvas(canvasSize, canvasSize, false)
		geom.DrawPolygon(convexHull, canvas, color.RGBA{0, 0, 0, 255}, 1)
		for _, p := range points {
			if geom.IsPointInsidePoligon(p, convexHull) {
				canvas.SetColor(color.RGBA{0, 255, 0, 255})
			} else {
				canvas.SetColor(color.RGBA{255, 0, 0, 255})
			}

			canvas.DrawPoint(p.X(), p.Y(), 1)
			canvas.Stroke()
		}
		for _, p := range randomPoints {
			if geom.IsPointInsidePoligon(p, convexHull) {
				canvas.SetColor(color.RGBA{0, 255, 0, 255})
			} else {
				canvas.SetColor(color.RGBA{255, 0, 0, 255})
			}

			canvas.DrawPoint(p.X(), p.Y(), 1)
			canvas.Stroke()
		}
		canvas.Stroke()
		canvas.SavePNG(strconv.Itoa(j) + ".png")
	}

}
