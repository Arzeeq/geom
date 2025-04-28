package geom

import (
	"math"
)

type Vector2D struct {
	X, Y float64
}

func NewVector2D(dx, dy float64) Vector2D {
	return Vector2D{X: dx, Y: dy}
}

func (v *Vector2D) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

// change length of vector to l
func (v *Vector2D) ScaleToLen(l float64) {
	k := l / v.Length()
	v.X *= k
	v.Y *= k
}

func NewVector2DFromPoints(begin, end Point2D) Vector2D {
	return NewVector2D(end.X()-begin.X(), end.Y()-begin.Y())
}

func SkewProduct(v1, v2 Vector2D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func ScalarProduct(v1, v2 Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (c *Canvas) DrawVector(begin Point2D, v Vector2D) {
	c.DrawLine(begin.X(), begin.Y(), begin.X()+v.X, begin.Y()+v.Y)
}

func AngleBetween(v1, v2 Vector2D) float64 {
	return math.Atan2(SkewProduct(v1, v2), ScalarProduct(v1, v2))
}
