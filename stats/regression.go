package stats

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (point Point) String() string {
	return fmt.Sprintf("(%.4f, %.4f)", point.X, point.Y)
}

func LeastSquares(points []Point) LinearTransform {
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

	return LinearTransform{
		M: m,
		B: b,
	}
}
