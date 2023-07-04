package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func BenchmarkRegressAll(b *testing.B) {
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
	Regress(data1)
	Regress(data2)
	Regress(data3)
	Regress(data4)

}

func TestRegress(t *testing.T) {
	gradient_all := 0.5
	intercept_all := 3.0

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
	gradient1, intercept1 := Regress(data1)
	gradient2, intercept2 := Regress(data2)
	gradient3, intercept3 := Regress(data3)
	gradient4, intercept4 := Regress(data4)

	if gradient1 != gradient_all {
		t.Error("The gradient in the 1st dataset is incorrect")
	}
	if gradient2 != gradient_all {
		t.Error("The gradient in the 2nd dataset is incorrect")
	}
	if gradient3 != gradient_all {
		t.Error("The gradient in the 3rd dataset is incorrect")
	}
	if gradient4 != gradient_all {
		t.Error("The gradient in the 4th dataset is incorrect")
	}
	if intercept1 != intercept_all {
		t.Error("The intercept in the 1st dataset is incorrect")
	}
	if intercept2 != intercept_all {
		t.Error("The intercept in the 2nd dataset is incorrect")
	}
	if intercept3 != intercept_all {
		t.Error("The intercept in the 3rd dataset is incorrect")
	}
	if intercept4 != intercept_all {
		t.Error("The intercept in the 4th dataset is incorrect")
	}
}
