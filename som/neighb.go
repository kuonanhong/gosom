package som

import "math"

// Gaussian calculates gaussian neghbourhood
func Gaussian(distance float64, radius float64) float64 {
	return math.Exp(-(distance * distance) / (2 * radius * radius))
}

// Bubble calculates bubble neghbourhood
func Bubble(distance float64, radius float64) float64 {
	return 0.0
}

// Mexican calculates mexican hat neghbourhood
func Mexican(distance float64, radius float64) float64 {
	return 0.0
}
