package abc

import (
	"image/color"
	"sort"
	"strconv"
)

func BuildConvexHull(points []Point2D, name int) {
	p0 := points[0]
	for _, p := range points {
		if p0.Y() == p.Y() {
			if p0.X() > p.X() {
				p0 = p
			}
		}
		if p0.Y() > p.Y() {
			p0 = p
		}
	}

	sort.Slice(points, func(i, j int) bool {
		o := NewVector(1, 0)
		v1 := VectorFromPoints(p0, points[i])
		v2 := VectorFromPoints(p0, points[j])
		sk1, _ := SkewProduct(o, v1)
		sk2, _ := SkewProduct(o, v2)
		sk1 /= v1.Length()
		sk2 /= v2.Length()
		return sk1 < sk2
	})

	stack := make([]Point2D, 0)

	for i := 0; i < len(points); i++ {
		for {
			if len(stack) <= 1 {
				stack = append(stack, points[i])
				break
			}
			p2 := stack[len(stack)-1]
			p1 := stack[len(stack)-2]
			sk, _ := SkewProduct(VectorFromPoints(p1, p2), VectorFromPoints(p2, points[i]))

			if sk >= 0 {
				stack = append(stack, points[i])
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}

	}

	canvas1 := NewCanvas(200, 200, false)
	DrawPolygon(points, canvas1, color.RGBA{0, 0, 0, 255}, 1)
	canvas1.SetColor(color.RGBA{255, 0, 0, 255})
	for _, p := range points {
		canvas1.DrawPoint(p.X(), p.Y(), 1)
	}
	canvas1.Stroke()
	canvas1.SavePNG(strconv.Itoa(name) + "_.png")

	canvas := NewCanvas(200, 200, false)
	DrawPolygon(stack, canvas, color.RGBA{0, 0, 0, 255}, 1)
	canvas.SetColor(color.RGBA{255, 0, 0, 255})
	for _, p := range points {
		canvas.DrawPoint(p.X(), p.Y(), 1)
	}
	canvas.Stroke()
	canvas.SavePNG(strconv.Itoa(name) + ".png")
}
