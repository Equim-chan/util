package number

import (
	"math"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxN(a []int) (int, int) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MaxUint32N(a []uint32) (int, uint32) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinN(a []int) (int, int) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

func MinUint32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func MinUint32N(a []uint32) (int, uint32) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

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
