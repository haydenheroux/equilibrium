package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func LeastSquares(name string) (LeastSquaresResult, error) {
	file, err := os.Open(name)

	if err != nil {
		return LeastSquaresResult{}, nil
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	ds := make([]LeastSquaresData, len(records))

	for i, record := range records {
		x, _ := strconv.ParseFloat(record[0], 64)
		y, _ := strconv.ParseFloat(record[1], 64)

		ds[i] = LeastSquaresData{X: x, Y: y}
	}

	return doLeastSquares(ds), nil
}

type LeastSquaresData struct {
	X float64
	Y float64
}

func (d LeastSquaresData) String() string {
	return fmt.Sprintf("(%.4f, %.4f)", d.X, d.Y)
}

type LeastSquaresResult struct {
	M float64
	B float64
}

func (r LeastSquaresResult) String() string {
	return fmt.Sprintf("y=%.4fx+%.4fx", r.M, r.B)
}

func doLeastSquares(ds []LeastSquaresData) LeastSquaresResult {
	sumX := 0.0
	sumY := 0.0
	sumXX := 0.0
	sumXY := 0.0

	n := float64(len(ds))

	for _, d := range ds {
		sumX += d.X
		sumY += d.Y
		sumXX += d.X * d.X
		sumXY += d.X * d.Y
	}

	m := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - m*sumX) / n

	return LeastSquaresResult{
		M: m,
		B: b,
	}
}
