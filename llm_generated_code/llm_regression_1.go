package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/regression"
)

func main() {
	// Given data
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Define the data matrix
	var data mat.Dense
	data.SetCol(0, x1)
	// Create a slice for the intercept term
	intercept := make([]float64, len(x1))
	for i := range intercept {
		intercept[i] = 1
	}
	data.SetCol(1, intercept)

	// Create a linear regression model
	var model regression.Regression
	model.SetObserved(y1)
	model.SetVars("x", "intercept")
	model.SetRounding(true)

	// Fit the model
	if ok := model.Fit(&data); !ok {
		fmt.Println("Failed to fit model")
		return
	}

	// Predict the values
	var predictions mat.Dense
	predictions.Col(nil, 0)
	model.Predict(&predictions)

	// Get the coefficients
	coefficients := model.Coefficients(nil)

	// Output the results
	fmt.Println("Intercept:", coefficients[0])
	fmt.Println("Slope:", coefficients[1])
}
