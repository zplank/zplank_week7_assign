package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	// Given data
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Calculate correlation coefficient
	correlation := stat.Correlation(x1, y1, nil)

	fmt.Printf("Correlation coefficient: %.2f\n", correlation)
}
