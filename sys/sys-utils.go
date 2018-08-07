package main

import "hash/crc32"

const (
	minUint32      uint32 = 0
	maxUint32             = ^minUint32
	maxUint64Float        = float64(maxUint32)
)

// return float hash [0 .. 1] from given data
func hashToFloat64(data []byte) float64 {
	return float64(crc32.ChecksumIEEE(data)) / maxUint64Float
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
