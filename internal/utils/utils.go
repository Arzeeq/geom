package utils

import "math/rand"

func RandFloat64n(n float64) float64 {
	return n * rand.Float64()
}
