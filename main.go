package main

import "thermodynamics/leastsquares"

func main() {
	result, _ := leastsquares.LeastSquares("data.csv")

	println(result.String())
}
