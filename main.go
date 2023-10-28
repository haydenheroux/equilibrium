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

	y := lt.Apply(x)

	fmt.Println(y)
}
