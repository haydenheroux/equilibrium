package thermodynamics

func HeatCapacity(mass float64, specificHeat float64, temp float64) float64 {
	return mass * specificHeat * temp
}

func PartialHeatCapacity(mass float64, specificHeat float64) float64 {
	return mass * specificHeat
}

func HeatTransfer(mass float64, specificHeat float64, initTemp float64, finalTemp float64) float64 {
	dt := finalTemp - initTemp

	return mass * specificHeat * dt
}

func LatentHeat(mass float64, latentHeat float64) float64 {
	return mass * latentHeat
}

func EquilibriumTemperature(waterMass float64, waterTemp float64, iceMass float64, iceTemp float64) float64 {
	liquidHeatCapacity := HeatCapacity(waterMass, WaterSpecificHeat, waterTemp)

	iceToMeltHeatTransfer := HeatTransfer(iceMass, IceSpecificHeat, iceTemp, 0)

	iceMeltLatentHeat := LatentHeat(iceMass, IceLatentHeatFusion)

	totalWaterMass := waterMass + iceMass

	remainingHeatCapacity := (liquidHeatCapacity) - (iceMeltLatentHeat + iceToMeltHeatTransfer)

	waterDenominator := PartialHeatCapacity(totalWaterMass, WaterSpecificHeat)

	joulesPerCelsius := waterDenominator
	celsiusPerJoule := 1 / joulesPerCelsius

	return remainingHeatCapacity * celsiusPerJoule
}
