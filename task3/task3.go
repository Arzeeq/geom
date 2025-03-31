package main

import (
	"arzeeq/geometry/abc"
	"image/color"
)

func main() {
	// p1 := abc.NewPoint2D(-10, -7)
	// p2 := abc.NewPoint2D(15, 23)

	// canvas1 := abc.NewCanvas(100, 100, true)
	// canvas1.SetColor(color.RGBA{0, 0, 255, 255})
	// abc.SimpleSegmentDrawer(p1, p2, canvas1)
	// canvas1.SavePNG("simple-segment.png")

	// canvas2 := abc.NewCanvas(100, 100, true)
	// canvas2.SetColor(color.RGBA{255, 0, 0, 255})
	// abc.DrawSegmentBresenham(p1, p2, canvas2)
	// canvas2.SavePNG("bresenham-segment.png")

	// canvas3 := abc.NewCanvas(400, 400, true)
	// canvas3.SetColor(color.RGBA{0, 0, 255, 255})
	// p0 := abc.NewPoint2D(50, 50)
	// p1 := abc.NewPoint2D(100, 100)
	// p2 := abc.NewPoint2D(150, 50)
	// abc.BezierCurve([]abc.Point2D{p0, p1, p2}, 1000, canvas3)
	// canvas3.SavePNG("BezierCurve2.png")

	// canvas3 := abc.NewCanvas(1000, 1000, true)
	// canvas3.SetColor(color.RGBA{0, 0, 255, 255})
	// p0 := abc.NewPoint2D(0, 0)
	// p1 := abc.NewPoint2D(100, 200)
	// p2 := abc.NewPoint2D(200, 0)
	// p3 := abc.NewPoint2D(300, 200)
	// abc.BezierCurve([]abc.Point2D{p0, p1, p2, p3}, 1000, canvas3)
	// canvas3.SavePNG("BezierCurveTest6.png")

	canvas3 := abc.NewCanvas(1000, 1000, true)
	canvas3.SetColor(color.RGBA{0, 0, 255, 255})
	p0 := abc.NewPoint2D(-100, -400)
	p1 := abc.NewPoint2D(-200, -150)
	p2 := abc.NewPoint2D(-100, 100)
	p3 := abc.NewPoint2D(-50, 225)
	p4 := abc.NewPoint2D(-100, 350)
	p5 := abc.NewPoint2D(0, 350)
	p6 := abc.NewPoint2D(100, 350)
	p7 := abc.NewPoint2D(50, 225)
	p8 := abc.NewPoint2D(100, 100)
	p9 := abc.NewPoint2D(200, -150)
	p10 := abc.NewPoint2D(100, -400)
	abc.BezierCurve([]abc.Point2D{p0, p1, p2}, 1000, canvas3)
	abc.BezierCurve([]abc.Point2D{p2, p3, p4}, 1000, canvas3)
	abc.BezierCurve([]abc.Point2D{p4, p5, p6}, 1000, canvas3)
	abc.BezierCurve([]abc.Point2D{p6, p7, p8}, 1000, canvas3)
	abc.BezierCurve([]abc.Point2D{p8, p9, p10}, 1000, canvas3)
	abc.BezierCurve([]abc.Point2D{p10, p0}, 1000, canvas3)
	canvas3.SavePNG("BezierCurveTest7.png")
}
