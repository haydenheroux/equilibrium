package thermodynamics

func HeatCapacity(mass float64, specificHeat float64, temp float64) float64 {
	return mass * specificHeat * temp
}

func HeatTransfer(mass float64, specificHeat float64, initTemp float64, finalTemp float64) float64 {
	dt := finalTemp - initTemp

	return mass * specificHeat * dt
}

func LatentHeat(mass float64, latentHeat float64) float64 {
	return mass * latentHeat
}

func EquilibriumTemperature(waterMass float64, waterTemp float64, iceMass float64, iceTemp float64) float64 {
	liquidHeatCapacity := HeatCapacity(waterMass, float64(WaterSpecificHeat), waterTemp)

	iceToMeltHeatTransfer := HeatTransfer(iceMass, float64(IceSpecificHeat), iceTemp, 0)

	iceMeltLatentHeat := LatentHeat(iceMass, float64(IceLatentHeatFusion))

	remainingHeatCapacity := liquidHeatCapacity - (iceToMeltHeatTransfer + iceMeltLatentHeat)

	joulesPerCelsius := float64(WaterSpecificHeat) * (iceMass + waterMass)
	celsiusPerJoule := 1 / joulesPerCelsius

	return remainingHeatCapacity * celsiusPerJoule
}
