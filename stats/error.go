package stats

func ApproximationError(estimate float64, actual float64) float64 {
	delta := actual - estimate

	return delta / estimate
}
