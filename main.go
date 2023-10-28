package main

import (
	"equilibrium/stats"
	"equilibrium/thermodynamics"
	"fmt"
)

func main() {
	points, _ := stats.LoadPoints("data.csv")

	lt := stats.LeastSquares(points)

	waterMass := 0.16842
	waterTemp := 51.0

	iceMass := 0.01285
	iceTemp := 0.0

	x := thermodynamics.EquilibriumTemperature(waterMass, waterTemp, iceMass, iceTemp)

	estimate := lt.Apply(x)

	actual := 42.5

	error := stats.ApproximationError(estimate, actual)

	fmt.Println(error)
}
