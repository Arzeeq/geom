package main

import (
	"arzeeq/geometry/abc"
	"fmt"
)

func main() {
	p0 := abc.NewPoint2D(3, 2)
	p1 := abc.NewPoint2D(2, 5)
	p2 := abc.NewPoint2D(3, 8)
	p3 := abc.NewPoint2D(7, 9)
	p4 := abc.NewPoint2D(10, 7)
	p5 := abc.NewPoint2D(9, 3)
	p6 := abc.NewPoint2D(8, 1)

	q := abc.NewPoint2D(1, 2)
	//rand.Seed(1)
	fmt.Println(abc.IsPointInsidePoligon(q, abc.Polygon{p0, p1, p2, p3, p4, p5, p6}))
}
