package regression

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func LeastSquares(name string) (Result, error) {
	file, err := os.Open(name)

	if err != nil {
		return Result{}, nil
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	points := make([]Point, len(records))

	for i, record := range records {
		x, _ := strconv.ParseFloat(record[0], 64)
		y, _ := strconv.ParseFloat(record[1], 64)

		points[i] = Point{X: x, Y: y}
	}

	return calculate(points), nil
}

type Point struct {
	X float64
	Y float64
}

func (point Point) String() string {
	return fmt.Sprintf("(%.4f, %.4f)", point.X, point.Y)
}

type Result struct {
	M float64
	B float64
}

func (r Result) String() string {
	return fmt.Sprintf("y=%.4fx+%.4f", r.M, r.B)
}

func calculate(points []Point) Result {
	sumX := 0.0
	sumY := 0.0
	sumXX := 0.0
	sumXY := 0.0

	n := float64(len(points))

	for _, point := range points {
		sumX += point.X
		sumY += point.Y
		sumXX += point.X * point.X
		sumXY += point.X * point.Y
	}

	m := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - m*sumX) / n

	return Result{
		M: m,
		B: b,
	}
}
