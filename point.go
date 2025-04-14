package geom

import (
	"math"
)

const (
	EPS = 1e-9
)

type Point2D [3]float64

// Однородные координаты
func NewPoint2D(x, y float64) Point2D {
	return [3]float64{x, y, 1}
}

func (p Point2D) X() float64 {
	return p[0]
}

func (p Point2D) Y() float64 {
	return p[1]
}

// check if point (p Point2D) lies inside rectangular with opposite vertexes (p1, p2 Point2D)
func IsPointInRect(p, p1, p2 Point2D) bool {
	return math.Min(p1.X(), p2.X()) <= p.X() && p.X() <= math.Max(p1.X(), p2.X()) &&
		math.Min(p1.Y(), p2.Y()) <= p.Y() && p.Y() <= math.Max(p1.Y(), p2.Y())
}

func IsEqualPoints(p1, p2 Point2D) bool {
	return math.Abs(p1.X()-p2.X()) < EPS && math.Abs(p1.Y()-p2.Y()) < EPS
}
