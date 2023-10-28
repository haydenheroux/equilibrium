package main

import (
	"equilibrium/stats"
	"equilibrium/thermodynamics"
	"fmt"
)

func main() {
	points, _ := stats.LoadPoints("data.csv")

	lt := stats.LeastSquares(points)

	var waterTemp, waterMass, iceTemp, iceMass float64

	fmt.Print("water temp: ")
	fmt.Scanf("%f", &waterTemp)
	fmt.Print("water mass: ")
	fmt.Scanf("%f", &waterMass)
	fmt.Print("ice temp: ")
	fmt.Scanf("%f", &iceTemp)
	fmt.Print("ice mass: ")
	fmt.Scanf("%f", &iceMass)

	x := thermodynamics.EquilibriumTemperature(waterMass, waterTemp, iceMass, iceTemp)

	estimatedTemp := lt.Apply(x)

	fmt.Printf("estimated temp: %.4f\n", estimatedTemp)

	var actualTemp float64

	fmt.Print("actual temp: ")
	fmt.Scanf("%f", &actualTemp)

	error := stats.ApproximationError(estimatedTemp, actualTemp)

	fmt.Printf("percent error: %.4f\n", error*100)
}
