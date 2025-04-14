package geom

import "math"

type AffineTransform2D [3][3]float64

func NewAffineTransform2D(matrix [3][3]float64) AffineTransform2D {
	return matrix
}

func (a AffineTransform2D) Transform(p Point2D) Point2D {
	n := 3
	result := Point2D{}
	for i := range n {
		for j := range n {
			result[i] += a[i][j] * p[j]
		}
	}
	return result
}

func (a AffineTransform2D) TransformPolygon(p Polygon) Polygon {
	newPoints := make(Polygon, len(p))
	for i, point := range p {
		newPoints[i] = a.Transform(point)
	}
	return newPoints
}

func (a AffineTransform2D) Composition(b AffineTransform2D) AffineTransform2D {
	result := AffineTransform2D{}
	for i1 := range len(a) {
		for j2 := range len(a) {
			for j1 := range len(a) {
				result[i1][j2] += a[i1][j1] * b[j1][j2]
			}
		}
	}
	return result
}

func Composition(an ...AffineTransform2D) AffineTransform2D {
	if len(an) == 0 {
		return NewIdentityTransform()
	} else if len(an) == 1 {
		return an[0]
	}

	a1 := an[0]
	a2 := Composition(an[1:]...)
	return a1.Composition(a2)
}

func NewTransferTransform(x, y float64) AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	})
}

func NewRotateTransform(angle float64) AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{math.Cos(angle), -math.Sin(angle), 0},
		{math.Sin(angle), math.Cos(angle), 0},
		{0, 0, 1},
	})
}

func NewAxialSymetryTransform() AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 1},
	})
}

func NewHomothetyTransform(x, y float64) AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, 1},
	})
}

func NewShearTransform(sh float64) AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{1, sh, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
}

func NewIdentityTransform() AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
}
