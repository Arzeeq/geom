package geom

import (
	"math"

	"github.com/Arzeeq/geom/pkg/gen"
)

type Polygon []Point2D

func (c *Canvas) DrawPolygon(p Polygon) {
	c.MoveTo(p[0].X(), p[0].Y())
	for _, point := range p {
		c.LineTo(point.X(), point.Y())
	}
	c.LineTo(p[0].X(), p[0].Y())
}

func (p Polygon) Center() Point2D {
	if len(p) == 0 {
		return NewPoint2D(0, 0)
	}
	var x, y float64
	for _, point := range p {
		x += point.X()
		y += point.Y()
	}
	return NewPoint2D(x/float64(len(p)), y/float64(len(p)))
}

func (p Polygon) Perimeter() float64 {
	if len(p) < 3 {
		return 0
	}

	v := NewVector2DFromPoints(p[0], p[len(p)-1])
	result := v.Length()
	for i := 1; i < len(p); i++ {
		v = NewVector2DFromPoints(p[i-1], p[i])
		result += v.Length()
	}

	return result
}

func (p Polygon) Area() float64 {
	if len(p) < 3 {
		return 0
	}
	if len(p) == 3 {
		v1 := NewVector2DFromPoints(p[0], p[1])
		v2 := NewVector2DFromPoints(p[0], p[2])
		return math.Abs(SkewProduct(v1, v2)) / 2
	}

	var area float64
	for i := 2; i < len(p); i++ {
		triangle := Polygon{p[0], p[i-1], p[i]}
		area += triangle.Area()
	}

	return area
}

// Polygon must be without self-intersections
func (p Polygon) ContainsRayMethod(point Point2D) bool {
	if len(p) < 3 {
		panic("polygon must consist of at least 3 points")
	}

	for i := 0; i < len(p); i++ {
		var s Segment
		if i == 0 {
			s = NewSegment(p[0], p[len(p)-1])
		} else {
			s = NewSegment(p[i-1], p[i])
		}

		if s.Contains(point) {
			return true
		}
	}

	rayCount := 3

	isInsideCnt := 0
	isOutsideCnt := 0
	for range rayCount {
		t := gen.RandFloat64n(2 * math.Pi)
		direction := NewPoint2D(point.X()+math.Cos(t), point.Y()+math.Sin(t))
		r := NewRay(point, direction)

		intersectionsCnt := 0
		if IsRayIntersectSeg(r, NewSegment(p[0], p[len(p)-1])) {
			intersectionsCnt++
		}

		for i := 0; i < len(p)-1; i++ {
			if IsRayIntersectSeg(r, NewSegment(p[i], p[i+1])) {
				intersectionsCnt++
			}
		}

		if intersectionsCnt%2 == 0 {
			isOutsideCnt++
		} else {
			isInsideCnt++
		}
	}
	return isInsideCnt > isOutsideCnt
}

// Polygon must be without self-intersections
func (p Polygon) ContainsAngleMethod(point Point2D) bool {
	n := len(p)
	if n < 3 {
		return false
	}

	angle := 0.0
	for i := 0; i < n; i++ {
		v1 := NewVector2DFromPoints(point, p[i])
		v2 := NewVector2DFromPoints(point, p[(i+1)%n])
		angle += AngleBetween(v1, v2)
	}

	return math.Abs(math.Abs(angle)-2*math.Pi) < EPS
}
