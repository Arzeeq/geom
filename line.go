package geom

import "math"

// Ax + By + C = 0
type Line struct {
	A, B, C float64
}

func NewLine(p1, p2 Point2D) Line {
	dx := p1.X() - p2.X()
	dy := p1.Y() - p2.Y()
	if math.Abs(dx) < EPS && math.Abs(dy) < EPS {
		panic("p1 and p2 must be different points")
	}

	if math.Abs(dy) < EPS {
		return Line{A: 0, B: 1, C: -p1.Y()}
	}
	var A float64 = 1
	var B float64 = -A * dx / dy
	var C float64 = -A*p1.X() - B*p1.Y()
	return Line{A: A, B: B, C: C}
}

func IsParallel(l1, l2 Line) bool {
	return math.Abs(l1.A*l2.B-l2.A*l1.B) < EPS
}
