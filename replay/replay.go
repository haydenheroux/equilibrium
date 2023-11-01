package replay

import (
	"encoding/csv"
	"equilibrium/stats"
	"equilibrium/thermodynamics"
	"os"
	"strconv"
)

func Replay(name string) ([]stats.Point, error) {
	file, err := os.Open(name)

	if err != nil {
		return []stats.Point{}, err
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	n := len(records)

	ps := make([]stats.Point, n)

	for i, record := range records {
		cupMass, _ := strconv.ParseFloat(record[0], 64)
		cupTemp, _ := strconv.ParseFloat(record[1], 64)
		waterMass, _ := strconv.ParseFloat(record[2], 64)
		waterTemp, _ := strconv.ParseFloat(record[3], 64)
		iceMass, _ := strconv.ParseFloat(record[4], 64)
		iceTemp, _ := strconv.ParseFloat(record[5], 64)

		estimate := thermodynamics.EquilibriumTemperatureCup(cupMass, cupTemp, waterMass, waterTemp, iceMass, iceTemp)

		actual, _ := strconv.ParseFloat(record[6], 64)

		ps[i] = stats.Point{X: estimate, Y: actual}
	}

	return ps, nil
}
