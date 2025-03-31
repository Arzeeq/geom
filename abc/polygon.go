package abc

import (
	"math"
)

type Polygon []Point2D

func (p *Polygon) AbsMaxX() float64 {
	result := 0.0
	for _, point := range *p {
		result = math.Max(result, math.Abs(point[0]))
	}
	return result
}

func (p *Polygon) AbsMaxY() float64 {
	result := 0.0
	for _, point := range *p {
		result = math.Max(result, math.Abs(point[1]))
	}
	return result
}

func AbsMaxX(p []Polygon) float64 {
	result := 0.0
	for _, polygon := range p {
		result = math.Max(result, polygon.AbsMaxX())
	}
	return result
}

func AbsMaxY(p []Polygon) float64 {
	result := 0.0
	for _, polygon := range p {
		result = math.Max(result, polygon.AbsMaxY())
	}
	return result
}
