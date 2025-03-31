package main

import (
	"arzeeq/geometry/abc"
	"arzeeq/geometry/internal/utils"
)

func main() {
	for j := range 10 {
		points := make([]abc.Point2D, 5)
		for i := 0; i < len(points); i++ {
			points[i] = abc.NewPoint2D(utils.RandFloat64n(100), utils.RandFloat64n(100))
		}
		abc.BuildConvexHull(points, j)
	}
}
