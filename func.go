package geom

import "math"

// parameter function
type Func struct {
	x, y   func(float64) float64 // function values
	x1, y1 func(float64) float64 // first derivative
	x2, y2 func(float64) float64 // second derivative
}

func NewFunc(x, y, x1, y1, x2, y2 func(float64) float64) Func {
	return Func{
		x:  x,
		y:  y,
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

// calculate X coordinate
func (f *Func) X(t float64) float64 {
	return f.x(t)
}

// calculate Y coordinate
func (f *Func) Y(t float64) float64 {
	return f.y(t)
}

// calculate first derivative of X in point (x(t), y(t))
func (f *Func) X1(t float64) float64 {
	return f.x1(t)
}

// calculate second derivative of X in point (x(t), y(t))
func (f *Func) X2(t float64) float64 {
	return f.x2(t)
}

// calculate first derivative of Y in point (x(t), y(t))
func (f *Func) Y1(t float64) float64 {
	return f.y1(t)
}

// calculate second derivative of Y in point (x(t), y(t))
func (f *Func) Y2(t float64) float64 {
	return f.y2(t)
}

// calculate X coordinate of evolute in point (x(t), y(t))
func (f *Func) EvoluteX(t float64) float64 {
	return f.x(t) - f.y1(t)*(f.x1(t)*f.x1(t)+f.y1(t)*f.y1(t))/(f.x1(t)*f.y2(t)-f.x2(t)*f.y1(t))
}

// calculate Y coordinate of evolute in point (x(t), y(t))
func (f *Func) EvoluteY(t float64) float64 {
	return f.y(t) + f.x1(t)*(f.x1(t)*f.x1(t)+f.y1(t)*f.y1(t))/(f.x1(t)*f.y2(t)-f.x2(t)*f.y1(t))
}

// calculate curvature in point (x(t), y(t))
func (f *Func) Curvature(t float64) float64 {
	n := f.x2(t)*f.y1(t) - f.x1(t)*f.y2(t)
	d := math.Pow(f.x1(t)*f.x1(t)+f.y1(t)*f.y1(t), 1.5)
	return n / d
}

// calculate radius of curvature in point (x(t), y(t))
func (f *Func) CurvatureRadius(t float64) float64 {
	return 1 / f.Curvature(t)
}
