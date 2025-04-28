package geom

import (
	"errors"
)

// Ray is a geometric shape that represents a set of points lying on a straiht line
// from the beginning (begin Point2D), and to the infinity in the
// direction of the point (direction Point2D)
type Ray struct {
	begin     Point2D
	direction Point2D
}

func NewRay(begin, direction Point2D) Ray {
	return Ray{begin: begin, direction: direction}
}

func (r *Ray) Contains(p Point2D) bool {
	l, err := NewLine(r.begin, r.direction)

	// r.begin == r.end
	if errors.Is(err, ErrLineCreation) {
		return IsEqualPoints(r.begin, p)
	}

	v1 := NewVector2DFromPoints(r.begin, r.direction)
	v2 := NewVector2DFromPoints(r.begin, p)

	return l.Contains(p) && ScalarProduct(v1, v2) > 0
}
