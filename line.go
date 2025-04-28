package geom

import (
	"errors"
	"math"
)

var ErrLineCreation = errors.New("p1 and p2 must be different points")

// Line is a geometric shape that represents a set of points
// where every point satisfies equation Ax + By + C = 0
type Line struct {
	A, B, C float64
}

func NewLine(p1, p2 Point2D) (Line, error) {
	if IsEqualPoints(p1, p2) {
		return Line{}, ErrLineCreation
	}

	dx := p1.X() - p2.X()
	dy := p1.Y() - p2.Y()

	if math.Abs(dy) < EPS {
		return Line{A: 0, B: 1, C: -p1.Y()}, nil
	}
	var A, B, C float64
	A = 1
	B = -A * dx / dy
	C = -A*p1.X() - B*p1.Y()

	return Line{A: A, B: B, C: C}, nil
}

func (l *Line) Norm() {
	scale := math.Sqrt(l.A*l.A + l.B*l.B)
	l.A /= scale
	l.B /= scale
	l.C /= scale
}

func IsParallel(l1, l2 Line) bool {
	return math.Abs(l1.A*l2.B-l2.A*l1.B) < EPS
}

func (l *Line) Contains(p Point2D) bool {
	return math.Abs(l.A*p.X()+l.B*p.Y()+l.C) < EPS
}
