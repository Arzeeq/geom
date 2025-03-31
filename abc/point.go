package abc

import (
	"fmt"
	"math"
)

const (
	EPS = 1e-9
)

type Point2D []float64

// Однородные координаты
func NewPoint2D(x, y float64) Point2D {
	return []float64{x, y, 1}
}

// TODO: Ввести структуру вектора
func NewVector(dx, dy float64) Point2D {
	return []float64{dx, dy, 0}
}

func (v *Point2D) Length() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y())
}

func VectorFromPoints(begin, end Point2D) Point2D {
	return NewVector(end.X()-begin.X(), end.Y()-begin.Y())
}

func (p *Point2D) X() float64 {
	return (*p)[0]
}

func (p *Point2D) Y() float64 {
	return (*p)[1]
}

func SkewProduct(v1, v2 Point2D) (float64, error) {
	if !(v1[2] < EPS) || !(v2[2] < EPS) {
		return 0, fmt.Errorf("v1 and v2 must be vectors")
	}

	return v1.X()*v2.Y() - v1.Y()*v2.X(), nil
}

func ScalarProduct(v1, v2 Point2D) (float64, error) {
	if !(v1[2] < EPS) || !(v2[2] < EPS) {
		return 0, fmt.Errorf("v1 and v2 must be vectors")
	}

	return v1.X()*v2.X() + v1.Y()*v2.Y(), nil
}

// check if point (p Point2D) lies inside rectangular with opposite vertexes (p1, p2 Point2D)
func IsPointInRect(p, p1, p2 Point2D) bool {
	return math.Min(p1.X(), p2.X()) <= p.X() && p.X() <= math.Max(p1.X(), p2.X()) &&
		math.Min(p1.Y(), p2.Y()) <= p.Y() && p.Y() <= math.Max(p1.Y(), p2.Y())
}
