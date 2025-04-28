package main

import (
	"image/color"
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	quadraticBezier()
	qubicBezier()
	spline()
}

func quadraticBezier() {
	T := geom.NewTransferTransform(-3, -3)
	Sh := geom.NewShearTransform(2)
	R := geom.NewRotateTransform(math.Pi/4, geom.NewPoint2D(0, 0))
	comp := geom.Composition(T, Sh, R)

	canvas := geom.NewCanvas(1000, 800, 20, 1)
	p0 := geom.NewPoint2D(0, 0)
	p1 := geom.NewPoint2D(5, 5)
	p2 := geom.NewPoint2D(10, 0)

	canvas.BezierCurve([]geom.Point2D{p0, p1, p2}, 50)
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.Stroke()

	canvas.BezierCurve(comp.TransformPolygon([]geom.Point2D{p0, p1, p2}), 50)
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.Stroke()

	canvas.SavePNG("bezier2.png")
}

func qubicBezier() {
	canvas := geom.NewCanvas(1000, 800, 20, 1)
	p0 := geom.NewPoint2D(0, 0)
	p1 := geom.NewPoint2D(5, 10)
	p2 := geom.NewPoint2D(10, 0)
	p3 := geom.NewPoint2D(15, 10)

	canvas.BezierCurve([]geom.Point2D{p0, p1, p2, p3}, 50)
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.Stroke()

	canvas.SavePNG("bezier3.png")
}

func spline() {
	canvas := geom.NewCanvas(1000, 800, 20, 25)
	p0 := geom.NewPoint2D(-100, -400)
	p1 := geom.NewPoint2D(-200, -150)
	p2 := geom.NewPoint2D(-100, 100)
	p3 := geom.NewPoint2D(-50, 225)
	p4 := geom.NewPoint2D(-100, 350)
	p5 := geom.NewPoint2D(0, 350)
	p6 := geom.NewPoint2D(100, 350)
	p7 := geom.NewPoint2D(50, 225)
	p8 := geom.NewPoint2D(100, 100)
	p9 := geom.NewPoint2D(200, -150)
	p10 := geom.NewPoint2D(100, -400)
	canvas.BezierCurve([]geom.Point2D{p0, p1, p2}, 50)
	canvas.BezierCurve([]geom.Point2D{p2, p3, p4}, 50)
	canvas.BezierCurve([]geom.Point2D{p4, p5, p6}, 50)
	canvas.BezierCurve([]geom.Point2D{p6, p7, p8}, 50)
	canvas.BezierCurve([]geom.Point2D{p8, p9, p10}, 50)
	canvas.BezierCurve([]geom.Point2D{p10, p0}, 50)

	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.Stroke()
	err := canvas.SavePNG("spline.png")
	if err != nil {
		panic(err)
	}
}
