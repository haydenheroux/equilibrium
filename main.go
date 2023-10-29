package main

import (
	"equilibrium/stats"
	"equilibrium/thermodynamics"
	"fmt"
)

func main() {
	points, _ := stats.LoadPoints("data.csv")

	lt := stats.LeastSquares(points)

	var initialMass, waterTemp, postWaterMass, iceTemp, postIceMass float64

	fmt.Print("initial mass: ")
	fmt.Scanf("%f", &initialMass)
	fmt.Print("water temp: ")
	fmt.Scanf("%f", &waterTemp)
	fmt.Print("post-water mass: ")
	fmt.Scanf("%f", &postWaterMass)
	fmt.Print("ice temp: ")
	fmt.Scanf("%f", &iceTemp)
	fmt.Print("post-ice mass: ")
	fmt.Scanf("%f", &postIceMass)

	waterMass := postWaterMass - initialMass
	iceMass := postIceMass - postWaterMass

	x := thermodynamics.EquilibriumTemperature(waterMass, waterTemp, iceMass, iceTemp)

	estimatedTemp := lt.Apply(x)

	fmt.Printf("estimated temp: %.4f\n", estimatedTemp)

	fmt.Printf("\nThe equilibrium temperature will be %.2f degress Celsius.\n\n", estimatedTemp)

	var actualTemp float64

	fmt.Print("actual temp: ")
	fmt.Scanf("%f", &actualTemp)

	error := stats.ApproximationError(estimatedTemp, actualTemp)

	fmt.Printf("error: %.4f\n", error)

	percentError := error * 100
	if percentError < 0 {
		fmt.Printf("\nThe actual value is %.1f%% below the estimated value.\n\n", -1*percentError)
	} else {
		fmt.Printf("\nThe actual value is %.1f%% above the estimated value.\n\n", percentError)
	}
}
