package number

import (
	"math"
)

// From github.com/gonum/gonum/floats.Round
func Round(v float64, prec int) float64 {
	if v == 0 {
		return 0
	}

	if prec >= 0 && v == math.Trunc(v) {
		return v
	}

	pow := math.Pow10(prec)
	intermed := v * pow
	if math.IsInf(intermed, 0) {
		return v
	}
	if v < 0 {
		v = math.Ceil(intermed - 0.5)
	} else {
		v = math.Floor(intermed + 0.5)
	}

	if v == 0 {
		return 0
	}

	return v / pow
}
