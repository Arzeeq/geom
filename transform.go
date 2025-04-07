package geom

import "math"

type AffineTransform2D struct {
	matrix [][]float64
	n, m   int
}

func NewAffineTransform2D(matrix [][]float64) AffineTransform2D {
	if len(matrix) != 3 {
		panic("Invalid transform matrix")
	}
	for i := range matrix {
		if len(matrix[i]) != 3 {
			panic("Invalid transform matrix")
		}
	}
	var result AffineTransform2D
	result.matrix = matrix
	result.n = 3
	result.m = 3
	return result
}

func (a *AffineTransform2D) Transform(p Point2D) Point2D {
	result := make(Point2D, 3)
	for i := range (*a).n {
		for j := range (*a).m {
			result[i] += (*a).matrix[i][j] * p[j]
		}
	}
	return result
}

func (a *AffineTransform2D) TransformPolygon(p Polygon) Polygon {
	newPoints := make(Polygon, len(p))
	for i, point := range p {
		newPoints[i] = a.Transform(point)
	}
	return newPoints
}

func Composition(an ...AffineTransform2D) AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	})
	if len(an) == 0 {
		return NewIdentityTransform()
	} else if len(an) == 1 {
		return an[0]
	}
	a1 := an[0]
	a2 := Composition(an[1:]...)
	for i1 := range a1.n {
		for j2 := range a2.m {
			for j1 := range a1.m {
				a.matrix[i1][j2] += a1.matrix[i1][j1] * a2.matrix[j1][j2]
			}
		}
	}
	return a
}

func NewTransferTransform(x, y float64) AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	})
	return a
}

func NewRotateTransform(angle float64) AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{math.Cos(angle), -math.Sin(angle), 0},
		{math.Sin(angle), math.Cos(angle), 0},
		{0, 0, 1},
	})
	return a
}

func NewAxialSymetryTransform() AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 1},
	})
	return a
}

func NewHomothetyTransform(x, y float64) AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, 1},
	})
	return a
}

func NewShearTransform(sh float64) AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{1, sh, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	return a
}

func NewIdentityTransform() AffineTransform2D {
	var a AffineTransform2D = NewAffineTransform2D([][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	return a
}
