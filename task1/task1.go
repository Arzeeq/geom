package main

import (
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	A := geom.NewPoint2D(-1, 3)
	B := geom.NewPoint2D(7, 4)
	C := geom.NewPoint2D(1, -3)
	triangle := geom.Polygon{A, B, C}

	// перенос на (5, 3)
	Ta := geom.NewTransferTransform(5, 3)

	// поворот на pi/4
	R := geom.NewRotateTransform(math.Pi / 4)

	// oceвая симметрия
	S := geom.NewAxialSymetryTransform()

	// гомотетия
	H := geom.NewHomothetyTransform(3, 3)

	comp := geom.Composition(H, R)

	geom.DrawPolygonTransformation(triangle, Ta, "triangleTa.png")
	geom.DrawPolygonTransformation(triangle, R, "triangleR.png")
	geom.DrawPolygonTransformation(triangle, S, "triangleS.png")
	geom.DrawPolygonTransformation(triangle, H, "triangleH.png")
	geom.DrawPolygonTransformation(triangle, comp, "triangleCOMP.png")

	M := geom.NewPoint2D(1, 1)
	N := geom.NewPoint2D(1, 4)
	P := geom.NewPoint2D(4, 4)
	Q := geom.NewPoint2D(4, 1)
	square := geom.Polygon{M, N, P, Q}
	sh := geom.NewShearTransform(4)
	sh1 := geom.NewShearTransform(-4)
	geom.DrawPolygonTransformation(square, sh, "square.png")
	geom.DrawPolygonTransformation(sh.TransformPolygon(square), sh1, "square1.png")
}
