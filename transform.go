package geom

import (
	"errors"
	"math"
)

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

func NewRotateTransform(angle float64, o Point2D) AffineTransform2D {
	return Composition(
		NewTransferTransform(o.X(), o.Y()),
		[3][3]float64{
			{math.Cos(angle), -math.Sin(angle), 0},
			{math.Sin(angle), math.Cos(angle), 0},
			{0, 0, 1},
		},
		NewTransferTransform(-o.X(), -o.Y()),
	)
}

func NewAxialSymetryTransform(l Line) AffineTransform2D {
	l.Norm()
	return [3][3]float64{
		{1 - 2*l.A*l.A, -2 * l.A * l.B, -2 * l.A * l.C},
		{-2 * l.A * l.B, 1 - 2*l.B*l.B, -2 * l.B * l.C},
		{0, 0, 1},
	}
}

func NewHomothetyTransform(kx, ky float64, p Point2D) AffineTransform2D {
	return NewAffineTransform2D([3][3]float64{
		{kx, 0, p.X() * (1 - kx)},
		{0, ky, p.Y() * (1 - ky)},
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

func determinant(m [3][3]float64) float64 {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

func (m AffineTransform2D) Inverse() (AffineTransform2D, error) {
	det := determinant(m)
	if math.Abs(det) < EPS {
		return AffineTransform2D{}, errors.New("Inverse matrix does not exist")
	}

	invDet := 1.0 / det

	cofactorMatrix := [3][3]float64{
		{
			(m[1][1]*m[2][2] - m[1][2]*m[2][1]),
			-(m[1][0]*m[2][2] - m[1][2]*m[2][0]),
			(m[1][0]*m[2][1] - m[1][1]*m[2][0]),
		},
		{
			-(m[0][1]*m[2][2] - m[0][2]*m[2][1]),
			(m[0][0]*m[2][2] - m[0][2]*m[2][0]),
			-(m[0][0]*m[2][1] - m[0][1]*m[2][0]),
		},
		{
			(m[0][1]*m[1][2] - m[0][2]*m[1][1]),
			-(m[0][0]*m[1][2] - m[0][2]*m[1][0]),
			(m[0][0]*m[1][1] - m[0][1]*m[1][0]),
		},
	}

	inverse := AffineTransform2D{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			inverse[i][j] = cofactorMatrix[j][i] * invDet
		}
	}

	return inverse, nil
}
