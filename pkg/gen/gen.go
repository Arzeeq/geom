package gen

import "math/rand"

// generate random float64 number [0, n)
func RandFloat64n(n float64) float64 {
	return n * rand.Float64()
}

// generate random float64 number [x1, x2)
func RandFloat64Between(x1, x2 float64) float64 {
	return x1 + RandFloat64n(x2-x1)
}
