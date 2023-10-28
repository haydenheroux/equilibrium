package main

import "thermodynamics/regression"

func main() {
	result, _ := regression.LeastSquares("data.csv")

	println(result.String())
}
