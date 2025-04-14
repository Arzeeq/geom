package geom

import (
	"errors"
	"math"
)

func LinesIntersection(l1, l2 *Line) (Point2D, bool) {
	if IsParallel(l1, l2) {
		return Point2D{}, false
	}

	var x, y float64
	if math.Abs(l1.A) < EPS {
		x = (l2.B*l1.C - l1.B*l2.C) / (l2.A*l1.B - l1.A*l2.B)
		y = (-l1.A*x - l1.C) / l1.B
	} else {
		y = (l2.A*l1.C - l1.A*l2.C) / (l1.A*l2.B - l2.A*l1.B)
		x = (-l1.B*y - l1.C) / l1.A
	}

	return NewPoint2D(x, y), true
}

func IsLinesIntersect(l1, l2 *Line) bool {
	return !IsParallel(l1, l2)
}

func SegmentsIntersection(s1, s2 Segment) (Point2D, bool) {
	l1, err1 := NewLine(s1.P1, s1.P2)
	l2, err2 := NewLine(s2.P1, s2.P2)
	if errors.Is(err1, ErrLineCreation) {
		return Point2D{}, s2.Contains(s1.P1)
	} else if errors.Is(err2, ErrLineCreation) {
		return Point2D{}, s1.Contains(s2.P1)
	}

	m, isLinesIntersect := LinesIntersection(l1, l2)
	var isSegmentsIntersect bool = IsPointInRect(m, s1.P1, s1.P2) && IsPointInRect(m, s2.P1, s2.P2)

	return m, isLinesIntersect && isSegmentsIntersect
}

func IsSegmentsIntersect(s1, s2 Segment) bool {
	f := func(seg1, seg2 Segment) bool {
		v := NewVector2DFromPoints(seg1.P1, seg1.P2)
		r1 := NewVector2DFromPoints(seg1.P1, seg2.P1)
		r2 := NewVector2DFromPoints(seg1.P1, seg2.P2)
		skew1 := SkewProduct(r1, v)
		skew2 := SkewProduct(r2, v)
		return math.Signbit(skew1) != math.Signbit(skew2)
	}

	return f(s1, s2) && f(s2, s1)
}

func IsRayIntersectSeg(r Ray, s Segment) bool {
	l1, err1 := NewLine(r.begin, r.direction)
	l2, err2 := NewLine(s.P1, s.P2)

	// r.begin == r.end
	if errors.Is(err1, ErrLineCreation) {
		return s.Contains(r.begin)
	}

	// s.P1 == s.P2
	if errors.Is(err2, ErrLineCreation) {
		return r.Contains(s.P1)
	}

	p, isLinesIntersect := LinesIntersection(l1, l2)
	if !isLinesIntersect {
		return false
	}
	scalarProduct := ScalarProduct(NewVector2DFromPoints(r.begin, r.direction), NewVector2DFromPoints(r.begin, p))
	isOnRay := scalarProduct > 0
	isOnSegment := IsPointInRect(p, s.P1, s.P2)
	return isOnRay && isOnSegment
}
