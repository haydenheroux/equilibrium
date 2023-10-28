package stats

import "fmt"

type LinearTransform struct {
	M float64
	B float64
}

func (lt LinearTransform) String() string {
	return fmt.Sprintf("y=%.4fx+%.4f", lt.M, lt.B)
}

func (lt LinearTransform) Apply(x float64) float64 {
	return x*lt.M + lt.B
}
