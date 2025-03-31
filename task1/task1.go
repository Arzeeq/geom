package main

import (
	"arzeeq/geometry/abc"
	"math"
)

func main() {
	A := abc.NewPoint2D(-1, 3)
	B := abc.NewPoint2D(7, 4)
	C := abc.NewPoint2D(1, -3)
	triangle := abc.Polygon{A, B, C}

	// перенос на (5, 3)
	Ta := abc.NewTransferTransform(5, 3)

	// поворот на pi/4
	R := abc.NewRotateTransform(math.Pi / 4)

	// oceвая симметрия
	S := abc.NewAxialSymetryTransform()

	// гомотетия
	H := abc.NewHomothetyTransform(3, 3)

	comp := abc.Composition(H, R)

	abc.DrawPolygonTransformation(triangle, Ta, "triangleTa.png")
	abc.DrawPolygonTransformation(triangle, R, "triangleR.png")
	abc.DrawPolygonTransformation(triangle, S, "triangleS.png")
	abc.DrawPolygonTransformation(triangle, H, "triangleH.png")
	abc.DrawPolygonTransformation(triangle, comp, "triangleCOMP.png")

	M := abc.NewPoint2D(1, 1)
	N := abc.NewPoint2D(1, 4)
	P := abc.NewPoint2D(4, 4)
	Q := abc.NewPoint2D(4, 1)
	square := abc.Polygon{M, N, P, Q}
	sh := abc.NewShearTransform(4)
	sh1 := abc.NewShearTransform(-4)
	abc.DrawPolygonTransformation(square, sh, "square.png")
	abc.DrawPolygonTransformation(sh.TransformPolygon(square), sh1, "square1.png")
}
