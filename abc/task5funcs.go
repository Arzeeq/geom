package abc

import (
	"arzeeq/geometry/internal/utils"
	"fmt"
	"math"
)

// Polygon without self-intersections
func IsPointInsidePoligon(point Point2D, polygon Polygon) bool {
	if len(polygon) < 3 {
		panic("polygon must consist of at least 3 points")
	}

	rayCount := 3

	isInsideCnt := 0
	isOutsideCnt := 0
	for range rayCount {
		t := utils.RandFloat64n(2 * math.Pi)
		direction := NewPoint2D(point.X()+math.Cos(t), point.Y()+math.Sin(t))
		r := NewRay(point, direction)

		intersectionsCnt := 0
		if IsRayIntersectSeg(r, NewSegment(polygon[0], polygon[len(polygon)-1])) {
			intersectionsCnt++
		}

		for i := 0; i < len(polygon)-1; i++ {
			if IsRayIntersectSeg(r, NewSegment(polygon[i], polygon[i+1])) {
				intersectionsCnt++
			}
		}

		if intersectionsCnt%2 == 0 {
			isOutsideCnt++
		} else {
			isInsideCnt++
		}
	}
	// debugging
	fmt.Printf("inside: %d, outside: %d\n", isInsideCnt, isOutsideCnt)
	return isInsideCnt > isOutsideCnt
}
