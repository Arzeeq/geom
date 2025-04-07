package geom

type Segment struct {
	P1, P2 Point2D
}

func NewSegment(p1, p2 Point2D) Segment {
	var segment Segment
	segment.P1 = p1
	segment.P2 = p2
	return segment
}
