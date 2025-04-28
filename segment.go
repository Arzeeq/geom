package geom

import (
	"errors"
	"math"
)

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

// Draw segment based on Bresenham algorythm
func (c *Canvas) DrawSegmentBresenham(x1, y1, x2, y2 int) {
	// inverseMode swaps x and y, because bresenham need x2 - x1 > y2 - y1
	var inverseMode bool
	deltaX := abs(x2 - x1)
	deltaY := abs(y2 - y1)
	if deltaX < deltaY {
		inverseMode = true
		deltaX, deltaY = deltaY, deltaX
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	// swap point1 and point2, because x1 < x2 needed
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	// bresenham
	e := 0
	deltaerr := (deltaY + 1)
	y := y1
	var diry int
	if y2-y1 > 0 {
		diry = 1
	} else {
		diry = -1
	}
	for x := x1; x <= x2; x++ {
		if inverseMode {
			c.SetPixel(y, x)
		} else {
			c.SetPixel(x, y)
		}
		e += deltaerr
		if e >= (deltaX + 1) {
			y += diry
			e -= deltaX + 1
		}
	}
}

func (c *Canvas) DrawSegmentSimpleAlg(x1, y1, x2, y2 int) {
	// inverseMode swaps x and y to be able to draw verical lines
	inverseMode := abs(x1-x2) < abs(y1-y2)
	if inverseMode {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	// swap point1 and point2, because x1 < x2 needed
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	// algorithm based on equation y = kx + b
	k := float64(y2-y1) / float64(x2-x1)
	b := float64(y1) - k*float64(x1)
	for x := x1; x <= x2; x++ {
		curX := float64(x)
		curY := math.Round(curX*k + b)

		if inverseMode {
			curX, curY = curY, curX
		}
		c.SetPixel(int(curX), int(curY))
	}
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
