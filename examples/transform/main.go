package main

import (
	"image/color"
	"math"

	"github.com/Arzeeq/geom"
)

func main() {
	triangleTransformations()
	squareTransformations()
}

func drawPolygonTransformation(p geom.Polygon, a geom.AffineTransform2D, filename string) {
	canvas := geom.NewCanvas(1000, 800, 20, 1)

	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	canvas.DrawPolygon(p)
	canvas.Stroke()

	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	canvas.DrawPolygon(a.TransformPolygon(p))
	canvas.Stroke()

	if err := canvas.SavePNG(filename); err != nil {
		panic(err)
	}
}

func triangleTransformations() {
	triangle := geom.Polygon{
		geom.NewPoint2D(0, 4),
		geom.NewPoint2D(8, 5),
		geom.NewPoint2D(2, -2),
	}

	// transfer to (5, 3)
	T := geom.NewTransferTransform(5, 3)

	// rotate around the triangle center
	R1 := geom.NewRotateTransform(math.Pi/4, triangle.Center())

	// rotate around point (1, 1)
	R2 := geom.NewRotateTransform(math.Pi, geom.NewPoint2D(1, 1))

	// axial symmetry along line l
	l, err := geom.NewLine(geom.NewPoint2D(0, 0), geom.NewPoint2D(0, 1))
	if err != nil {
		panic(err)
	}
	S := geom.NewAxialSymetryTransform(l)

	// homothety relative to the point (0, 0)
	H1 := geom.NewHomothetyTransform(2, 2, geom.NewPoint2D(0, 0))

	// homothety relative to the point (1, 1)
	H2 := geom.NewHomothetyTransform(2, 2, geom.NewPoint2D(1, 1))

	// create compositions of several affine transforms
	comp1 := geom.Composition(S, H1)
	comp2 := geom.Composition(H2, R2)

	drawPolygonTransformation(triangle, T, "triangleT.png")
	drawPolygonTransformation(triangle, R1, "triangleR.png")
	drawPolygonTransformation(triangle, comp1, "triangleCOMP1.png")
	drawPolygonTransformation(triangle, comp2, "triangleCOMP2.png")
}

func squareTransformations() {
	square := geom.Polygon{
		geom.NewPoint2D(0, 0),
		geom.NewPoint2D(0, 3),
		geom.NewPoint2D(3, 3),
		geom.NewPoint2D(3, 0),
	}

	// create transforms
	shear := geom.NewShearTransform(math.Sqrt(3))
	homothety := geom.NewHomothetyTransform(2, 2, geom.NewPoint2D(0, 0))
	transfer := geom.NewTransferTransform(9, 9)

	// create composition and inverse of composition
	comp := geom.Composition(transfer, shear, homothety)
	inverseComp, err := comp.Inverse()
	if err != nil {
		panic(err)
	}

	drawPolygonTransformation(square, comp, "squareTransform.png")
	drawPolygonTransformation(comp.TransformPolygon(square), inverseComp, "squareInverseTransform.png")
}
