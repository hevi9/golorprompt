package sys

import (
	"reflect"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestHashToFloat64(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashToFloat64(tt.args.data); got != tt.want {
				t.Errorf("HashToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("maxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumColorToHSV(t *testing.T) {
	type args struct {
		numColor int
	}
	tests := []struct {
		name string
		args args
		want colorful.Color
	}{
		{
			"test", args{10}, colorful.Hsv(100.0, 0.5, 0.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumColorToHSV(tt.args.numColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumColorToHSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
