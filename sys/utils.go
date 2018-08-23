package sys

import (
	"hash/crc32"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	minUint32          uint32 = 0
	maxUint32                 = ^minUint32
	maxUint32AsFloat64        = float64(maxUint32)
)

// HashToFloat64 Return float hash [0 .. 1] from given data
func HashToFloat64(data []byte) float64 {
	return float64(crc32.ChecksumIEEE(data)) / maxUint32AsFloat64
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// color 045050080 => H=45 S=50 V=80
func NumColorToHSV(numColor int) colorful.Color {
	return colorful.Hsv(180, 0.5, 0.5)
}
