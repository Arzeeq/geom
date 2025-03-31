package abc

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

func drawTangentSection(x0, y0, p, q float64, canvas *gg.Context) {
	l := (p*p + q*q) / 1500
	p /= math.Sqrt(l)
	q /= math.Sqrt(l)
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.DrawLine(x0-p, y0-q, x0+p, y0+q)
	canvas.Stroke()
}

func DrawEllipse(a, b float64, canvas *gg.Context) {
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	n := 100
	t1 := 0.0
	t2 := 2 * math.Pi
	for i := range n {
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		//Ellipse coordinates
		x := a * math.Cos(t)
		y := b * math.Sin(t)
		//canvas.DrawPoint(x, y, 5)
		// first derivative
		x1 := -a * math.Sin(t)
		y1 := b * math.Cos(t)
		// second derivative
		x2 := -a * math.Cos(t)
		y2 := -b * math.Sin(t)
		canvas.DrawPoint(x-y1*(x1*x1+y1*y1)/(x1*y2-x2*y1), y+x1*(x1*x1+y1*y1)/(x1*y2-x2*y1), 2)
		if i%5 == 0 {
			drawTangentSection(x, y, x1, y1, canvas)
		}
	}
	canvas.Stroke()
}

/*
R - outter circle radius
r - inner circle radius
*/
func DrawHypocycloid(r, R float64, canvas *gg.Context) {
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	k := R / r
	n := 1000
	t1 := 0.0
	t2 := 6 * math.Pi
	for i := range n {
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		//Ellipse coordinates
		x := r * (k - 1) * (math.Cos(t) + math.Cos((k-1)*t)/(k-1))
		y := r * (k - 1) * (math.Sin(t) - math.Sin((k-1)*t)/(k-1))
		canvas.DrawPoint(x, y, 5)
	}
	canvas.Stroke()
}

func DrawArhSpiral(a, b float64, canvas *gg.Context) {
	canvas.SetColor(color.RGBA{0, 0, 255, 255})
	n := 1000
	t1 := 0.0
	t2 := 6 * math.Pi
	for i := range n {
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		//Ellipse coordinates
		x := (a + b*t) * math.Cos(t)
		y := (a + b*t) * math.Sin(t)
		canvas.DrawPoint(x, y, 3)
	}
	canvas.Stroke()
}

/*
Draw vector (p, q) which starts at (x0, y0)
*/
func DrawVector(x0, y0, p, q float64, canvas *gg.Context) {
	canvas.SetColor(color.RGBA{0, 0, 0, 255})
	canvas.DrawLine(x0, y0, x0+p, y0+q)
	canvas.Stroke()
}
