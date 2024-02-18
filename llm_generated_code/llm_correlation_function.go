package statistics

import "gonum.org/v1/gonum/stat"

// Correlation calculates the correlation coefficient between two sets of data.
func Correlation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}
