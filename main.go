package main

import (
	"equilibrium/replay"
	"equilibrium/stats"
	"equilibrium/thermodynamics"
	"fmt"
)

func main() {
	replayed, _ := replay.Replay("replay.csv")

	for _, point := range replayed {
		fmt.Println(point)
	}

	fmt.Println(stats.LeastSquares(replayed))

	var setupMass, waterTemp, postWaterMass, iceTemp, postIceMass float64

	fmt.Print("setup mass (g): ")
	fmt.Scanf("%f", &setupMass)
	fmt.Print("water temp (C): ")
	fmt.Scanf("%f", &waterTemp)
	fmt.Print("post-water mass (g): ")
	fmt.Scanf("%f", &postWaterMass)
	fmt.Print("ice temp (C): ")
	fmt.Scanf("%f", &iceTemp)
	fmt.Print("post-ice mass (g): ")
	fmt.Scanf("%f", &postIceMass)

	waterMass := postWaterMass - setupMass
	iceMass := postIceMass - postWaterMass

	x := thermodynamics.EquilibriumTemperature(waterMass, waterTemp, iceMass, iceTemp)

	estimatedTemp := stats.LeastSquares(replayed).Apply(x)

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
