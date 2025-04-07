package main

import (
	"arzeeq/geometry/geom"
	"fmt"
)

func main() {
	p0 := geom.NewPoint2D(3, 2)
	p1 := geom.NewPoint2D(2, 5)
	p2 := geom.NewPoint2D(3, 8)
	p3 := geom.NewPoint2D(7, 9)
	p4 := geom.NewPoint2D(10, 7)
	p5 := geom.NewPoint2D(9, 3)
	p6 := geom.NewPoint2D(8, 1)

	q := geom.NewPoint2D(1, 2)
	//rand.Seed(1)
	fmt.Println(geom.IsPointInsidePoligon(q, geom.Polygon{p0, p1, p2, p3, p4, p5, p6}))
}
