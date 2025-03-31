package abc

import (
	"fmt"
	"math"
)

func LinesIntersection(l1, l2 Line) (Point2D, bool) {
	if IsParallel(l1, l2) {
		return nil, false
	}

	var x, y float64
	if math.Abs(l1.A) < EPS {
		x = (l2.B*l1.C - l1.B*l2.C) / (l2.A*l1.B - l1.A*l2.B)
		y = (-l1.A*x - l1.C) / l1.B
	} else {
		y = (l2.A*l1.C - l1.A*l2.C) / (l1.A*l2.B - l2.A*l1.B)
		x = (-l1.B*y - l1.C) / l1.A
	}
	fmt.Printf("x: %f y: %f\n", x, y)
	fmt.Printf("l1: %f, %f, %f l2: %f, %f, %f\n", l1.A, l1.B, l1.C, l2.A, l2.B, l2.C)

	return NewPoint2D(x, y), true
}

func IsLinesIntersect(l1, l2 Line) bool {
	return !IsParallel(l1, l2)
}

func SegmentsIntersection(s1, s2 Segment) (Point2D, bool) {
	l1 := NewLine(s1.P1, s1.P2)
	l2 := NewLine(s2.P1, s2.P2)

	m, isLinesIntersect := LinesIntersection(l1, l2)
	var isSegmentsIntersect bool = IsPointInRect(m, s1.P1, s1.P2) && IsPointInRect(m, s2.P1, s2.P2)

	return m, isLinesIntersect && isSegmentsIntersect
}

func IsSegmentsIntersect(s1, s2 Segment) bool {
	f := func(seg1, seg2 Segment) bool {
		v := VectorFromPoints(seg1.P1, seg1.P2)
		r1 := VectorFromPoints(seg1.P1, seg2.P1)
		r2 := VectorFromPoints(seg1.P1, seg2.P2)
		skew1, _ := SkewProduct(r1, v)
		skew2, _ := SkewProduct(r2, v)
		return math.Signbit(skew1) != math.Signbit(skew2)
	}

	return f(s1, s2) && f(s2, s1)
}

func IsRayIntersectSeg(r Ray, s Segment) bool {
	fmt.Println("IsRayIntersectSeg", r, s)
	l1, l2 := NewLine(r.begin, r.direction), NewLine(s.P1, s.P2)
	p, isLinesIntersect := LinesIntersection(l1, l2)
	if !isLinesIntersect {
		return false
	}
	fmt.Println("intersection: ", p)
	scalarProduct, _ := ScalarProduct(VectorFromPoints(r.begin, r.direction), VectorFromPoints(r.begin, p))
	isOnRay := scalarProduct > 0
	isOnSegment := IsPointInRect(p, s.P1, s.P2)
	fmt.Printf("isOnRay: %t, isOnSegment: %t\n", isOnRay, isOnSegment)
	return isOnRay && isOnSegment
}
