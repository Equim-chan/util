package number

// drunk...

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

func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func MaxInt8N(a []int8) (int, int8) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func MaxInt16N(a []int16) (int, int16) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func MaxInt32N(a []int32) (int, int32) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MaxInt64N(a []int64) (int, int64) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func MaxUintN(a []uint) (int, uint) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func MaxUint8N(a []uint8) (int, uint8) {
	maxI := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxI] {
			maxI = i
		}
	}

	return maxI, a[maxI]
}

func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func MaxUint16N(a []uint16) (int, uint16) {
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

func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func MaxUint64N(a []uint64) (int, uint64) {
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

func MinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

func MinInt8N(a []int8) (int, int8) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MinInt32N(a []int32) (int, int32) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinInt64N(a []int64) (int, int64) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func MinUintN(a []uint) (int, uint) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

func MinUint8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

func MinUint8N(a []uint8) (int, uint8) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}

func MinUint16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func MinUint16N(a []uint16) (int, uint16) {
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

func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func MinUint64N(a []uint64) (int, uint64) {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}

	return minI, a[minI]
}
