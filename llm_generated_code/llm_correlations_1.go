package main

import (
	"fmt"
	"math"
)

func correlationCoefficient(x, y []float64) float64 {
	n := len(x)
	if n != len(y) || n == 0 {
		return math.NaN() // return NaN if the slices are not of equal length or empty
	}

	var sumX, sumY, sumXY, sumXSquare, sumYSquare float64
	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
		sumYSquare += y[i] * y[i]
	}

	numerator := float64(n)*sumXY - sumX*sumY
	denominator := math.Sqrt((float64(n)*sumXSquare - sumX*sumX) * (float64(n)*sumYSquare - sumY*sumY))

	if denominator == 0 {
		return math.NaN() // return NaN if denominator is zero to avoid division by zero
	}

	return numerator / denominator
}

func main() {
	// Example input data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 3, 4, 5, 6}

	// Calculate correlation coefficient
	correlation := correlationCoefficient(x, y)

	fmt.Printf("Correlation coefficient: %.2f\n", correlation)
}
