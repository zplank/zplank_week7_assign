package statistics

import "testing"

func TestCorrelation(t *testing.T) {
	// Given data
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Calculate correlation coefficient
	correlation := Correlation(x, y)

	expectedCorrelation := 0.81642051634484
	if correlation != expectedCorrelation {
		t.Errorf("Correlation coefficient mismatch. Expected: %.10f, Got: %.10f", expectedCorrelation, correlation)
	}
}
