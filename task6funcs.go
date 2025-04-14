package geom

import (
	"sort"
)

func BuildConvexHull(points []Point2D) Polygon {
	// define the index p0Idx of the lowest point
	p := make([]Point2D, len(points))
	copy(p, points)
	p0Idx := 0
	for i := range p {
		if p[p0Idx].Y() == p[i].Y() {
			if p[p0Idx].X() > p[i].X() {
				p0Idx = i
			}
		} else if p[p0Idx].Y() > p[i].Y() {
			p0Idx = i
		}
	}

	// create append the lowest point to the convexHull
	convexHull := make([]Point2D, 0)
	convexHull = append(convexHull, p[p0Idx])
	p = append(p[:p0Idx], p[p0Idx+1:]...)

	// sort points by their polar angle
	sort.Slice(p, func(i, j int) bool {
		vi := NewVector2DFromPoints(convexHull[0], p[i])
		vj := NewVector2DFromPoints(convexHull[0], p[j])
		return vi.X/vi.Y < vj.X/vj.Y
	})

	// build convexHull
	for i := 0; i < len(p); i++ {
		for {
			if len(convexHull) <= 1 {
				convexHull = append(convexHull, p[i])
				break
			}
			p2 := convexHull[len(convexHull)-1]
			p1 := convexHull[len(convexHull)-2]
			sk := SkewProduct(NewVector2DFromPoints(p1, p2), NewVector2DFromPoints(p2, p[i]))

			if sk <= 0 {
				convexHull = append(convexHull, p[i])
				break
			} else {
				convexHull = convexHull[:len(convexHull)-1]
			}
		}
	}

	return convexHull
}
