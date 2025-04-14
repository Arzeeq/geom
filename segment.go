package geom

import "errors"

type Segment struct {
	P1, P2 Point2D
}

func NewSegment(p1, p2 Point2D) Segment {
	var segment Segment
	segment.P1 = p1
	segment.P2 = p2
	return segment
}

func (s *Segment) Contains(p Point2D) bool {
	l, err := NewLine(s.P1, s.P2)
	if errors.Is(err, ErrLineCreation) {
		return IsEqualPoints(s.P1, p)
	}

	return l.Contains(p) && IsPointInRect(p, s.P1, s.P2)
}
