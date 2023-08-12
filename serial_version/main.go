package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat"
)

func preprocessing() ([]float64, []float64, []float64, error) {
	// Use os package to the CSV file in the same folder
	file, err := os.Open("boston.csv")

	// If there is an error, print that there is an error.
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, nil, err
	}

	// Defer closing the file until the end of the program
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Skip column header on the first row.
	_, _ = reader.Read()

	// Declare xData and yData outside the loop
	var (
		x_crim []float64
		x_nox  []float64
		y_mv   []float64
	)

	// Read and process each row of the CSV file
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// These are our input variables for the regressions we run.
		crim, _ := strconv.ParseFloat(record[1], 64)
		nox, _ := strconv.ParseFloat(record[5], 64)

		// This is our target variable
		mv, _ := strconv.ParseFloat(record[12], 64)

		x_crim = append(x_crim, crim)
		x_nox = append(x_nox, nox)
		y_mv = append(y_mv, mv)
	}

	return x_crim, x_nox, y_mv, err

}

func perform_regression() {

	x_crim, x_nox, y_mv, _ := preprocessing()

	var weights []float64
	origin := false

	alpha1, beta1 := stat.LinearRegression(x_crim, y_mv, weights, origin)
	alpha2, beta2 := stat.LinearRegression(x_nox, y_mv, weights, origin)
	comment1 := fmt.Sprintf("Regression 1 Results: Alpha [%f] Beta [%f]", alpha1, beta1)
	comment2 := fmt.Sprintf("Regression 2 Results: Alpha [%f] Beta [%f]", alpha2, beta2)

	fmt.Println(comment1)
	fmt.Println(comment2)
}

func main() {
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		perform_regression()
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Total time:", elapsedTime)

	// Open a file for writing
	file, err := os.Create("elapsed_time.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write elapsed time to the file
	_, err = file.WriteString(fmt.Sprintf("Time to run with concurrency: %s\n", elapsedTime.String()))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
