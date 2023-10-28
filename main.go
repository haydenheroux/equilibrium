package main

import (
	"fmt"
	"thermodynamics/stats"
)

func main() {
	points, _ := stats.LoadPoints("data.csv")

	lt := stats.LeastSquares(points)

	x := 41.7454137

	y := lt.Apply(x)

	fmt.Println(y)
}
