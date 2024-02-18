//go:generate go run generate_data.go

package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

var x2 []float64
var y2 []float64

func correlation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}

func main() {
	// calculate correlation
	corr := correlation(x2, y2)

	// print correlation coefficient
	fmt.Printf("Correlation coefficient between x2 and y2: %v\n", corr)
}
