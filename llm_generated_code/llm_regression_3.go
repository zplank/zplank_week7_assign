package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/regression"
)

func main() {
	// Given data
	x := [][]float64{
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // Predictor variable x1
		{8, 8, 8, 8, 8, 8, 8, 4, 12, 7, 5},     // Predictor variable x2
	}
	y := [][]float64{
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}, // Response variable y1
		{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},    // Response variable y2
	}

	// Number of predictor variables and response variables
	numPredictors := len(x)
	numResponses := len(y)

	for i := 0; i < numResponses; i++ {
		// Define the data matrix
		var data mat.Dense
		data.SetCol(0, x[0]) // Set the first predictor variable
		data.SetCol(1, x[1]) // Set the second predictor variable
		// Add more columns if you have more predictor variables

		// Create a slice for the intercept term
		intercept := make([]float64, len(x[0]))
		for j := range intercept {
			intercept[j] = 1
		}
		data.SetCol(numPredictors, intercept)

		// Create a linear regression model
		var model regression.Regression
		model.SetObserved(y[i])
		model.SetVars("x1", "x2", "intercept")
		// Add more predictor variable names if you have more predictor variables

		// Fit the model
		if ok := model.Fit(&data); !ok {
			fmt.Println("Failed to fit model for response variable", i+1)
			continue
		}

		// Get the coefficients
		coefficients := model.Coefficients(nil)

		// Output the results
		fmt.Printf("For response variable %d:\n", i+1)
		fmt.Println("Intercept:", coefficients[0])
		fmt.Println("Coefficient for x1:", coefficients[1])
		fmt.Println("Coefficient for x2:", coefficients[2])
		// Add more coefficient outputs if you have more predictor variables
	}
}
