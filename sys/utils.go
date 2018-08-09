package sys

import "hash/crc32"

const (
	MinUint32 uint32 = 0

	MaxUint32 = ^MinUint32

	MaxUint64Float = float64(MaxUint32)
)

// HashToFloat64 Return float hash [0 .. 1] from given data
func HashToFloat64(data []byte) float64 {
	return float64(crc32.ChecksumIEEE(data)) / MaxUint64Float
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
