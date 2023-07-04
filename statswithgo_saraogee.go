package main

//Importing packages including stats
import (
	"fmt"
	"math"

	"github.com/montanaflynn/stats"
)

// Rounding function from gosamples.dev/round-float/ to specify 2 significant digits
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// stats package does not output gradient and intercept and must be recalculated with this function
func Regress(data []stats.Coordinate) (gradient, intercept float64) {
	r, _ := stats.LinearRegression(data)
	// Gradient and intercepts use 7th and 8th points because others result in zero values with the fourth dataset
	gradient = roundFloat((r[7].Y-r[6].Y)/(r[7].X-r[6].X), 2)
	intercept = roundFloat((r[7].Y - (r[7].X * gradient)), 2)
	return gradient, intercept
}

func main() {
	//Inputing data as Series types for the Anscombe Quartet
	// Anscombe, F. J. 1973, February. Graphs in statistical analysis.
	//  The American Statistician 27: 17–21.
	data1 := []stats.Coordinate{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}
	data2 := []stats.Coordinate{
		{X: 10, Y: 9.14},
		{X: 8, Y: 8.14},
		{X: 13, Y: 8.74},
		{X: 9, Y: 8.77},
		{X: 11, Y: 9.26},
		{X: 14, Y: 8.1},
		{X: 6, Y: 6.13},
		{X: 4, Y: 3.1},
		{X: 12, Y: 9.13},
		{X: 7, Y: 7.26},
		{X: 5, Y: 4.74},
	}
	data3 := []stats.Coordinate{
		{X: 10, Y: 7.46},
		{X: 8, Y: 6.77},
		{X: 13, Y: 12.74},
		{X: 9, Y: 7.11},
		{X: 11, Y: 7.81},
		{X: 14, Y: 8.84},
		{X: 6, Y: 6.08},
		{X: 4, Y: 5.39},
		{X: 12, Y: 8.15},
		{X: 7, Y: 6.42},
		{X: 5, Y: 5.73},
	}
	data4 := []stats.Coordinate{
		{X: 8, Y: 6.58},
		{X: 8, Y: 5.76},
		{X: 8, Y: 7.71},
		{X: 8, Y: 8.84},
		{X: 8, Y: 8.47},
		{X: 8, Y: 7.04},
		{X: 8, Y: 5.25},
		{X: 19, Y: 12.5},
		{X: 8, Y: 5.56},
		{X: 8, Y: 7.91},
		{X: 8, Y: 6.89},
	}

	//Performing the least squares regression using Go's stats package
	gradient1, intercept1 := Regress(data1)
	gradient2, intercept2 := Regress(data2)
	gradient3, intercept3 := Regress(data3)
	gradient4, intercept4 := Regress(data4)

	//Printing out results for each model with gradients and intercepts
	fmt.Println("Results for x1 vs y1")
	fmt.Println("Gradient: ", gradient1)
	fmt.Println("Intercept: ", intercept1)
	fmt.Println(" ")

	fmt.Println("Results for x2 vs y2")
	fmt.Println("Gradient: ", gradient2)
	fmt.Println("Intercept: ", intercept2)
	fmt.Println(" ")

	fmt.Println("Results for x3 vs y3")
	fmt.Println("Gradient: ", gradient3)
	fmt.Println("Intercept: ", intercept3)
	fmt.Println(" ")

	fmt.Println("Results for x4 vs y4")
	fmt.Println("Gradient: ", gradient4)
	fmt.Println("Intercept: ", intercept4)
	fmt.Println(" ")
}
