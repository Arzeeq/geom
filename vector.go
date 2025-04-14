package geom

import (
	"math"
)

type Vector2D struct {
	X, Y float64
}

func NewVector2D(dx, dy float64) *Vector2D {
	return &Vector2D{X: dx, Y: dy}
}

func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func NewVector2DFromPoints(begin, end Point2D) *Vector2D {
	return NewVector2D(end.X()-begin.X(), end.Y()-begin.Y())
}

func SkewProduct(v1, v2 *Vector2D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func ScalarProduct(v1, v2 *Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}
