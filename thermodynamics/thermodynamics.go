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

func EquilibriumTemperature(cupMass float64, cupTemp float64, waterMass float64, waterTemp float64, iceMass float64, iceTemp float64) float64 {
	cupHeatCapacity := HeatCapacity(cupMass, AluminumSpecificHeat, cupTemp)

	liquidHeatCapacity := HeatCapacity(waterMass, WaterSpecificHeat, waterTemp)

	iceToMeltHeatTransfer := HeatTransfer(iceMass, IceSpecificHeat, iceTemp, 0)

	iceMeltLatentHeat := LatentHeat(iceMass, IceLatentHeatFusion)

	totalWaterMass := waterMass + iceMass

	remainingHeatCapacity := (cupHeatCapacity + liquidHeatCapacity) - (iceMeltLatentHeat + iceToMeltHeatTransfer)

	waterDenominator := PartialHeatCapacity(totalWaterMass, WaterSpecificHeat)
	cupDenominator := PartialHeatCapacity(cupMass, AluminumSpecificHeat)

	joulesPerCelsius := waterDenominator + cupDenominator
	celsiusPerJoule := 1 / joulesPerCelsius

	return remainingHeatCapacity * celsiusPerJoule
}
