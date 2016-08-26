package fcolor

import (
	"math"
	"testing"
)

func TestNewRGBA(t *testing.T) {
	cases := []struct {
		value    [4]uint8
		expected RGBAF64
	}{
		{
			value:    [4]uint8{0, 0, 0, 0},
			expected: RGBAF64{0, 0, 0, 0},
		},
		{
			value:    [4]uint8{255, 128, 64, 32},
			expected: RGBAF64{1.0, 0.5, 0.25, 0.125},
		},
		{
			value:    [4]uint8{10, 20, 30, 40},
			expected: RGBAF64{0.04, 0.078, 0.12, 0.16},
		},
		{
			value:    [4]uint8{255, 255, 255, 255},
			expected: RGBAF64{1.0, 1.0, 1.0, 1.0},
		},
	}

	for _, c := range cases {
		actual := NewRGBAF64(c.value[0], c.value[1], c.value[2], c.value[3])
		if !rgbaf64Equal(actual, c.expected, 0.01) {
			t.Errorf("%s: expected: %#v, actual: %#v", "NewRGBAF6", c.expected, actual)
		}
	}
}

func TestClamp(t *testing.T) {
	cases := []struct {
		value    RGBAF64
		expected RGBAF64
	}{
		{
			value:    RGBAF64{0, 0, 0, 0},
			expected: RGBAF64{0, 0, 0, 0},
		},
		{
			value:    RGBAF64{10.0, 0.55, -0.25, 1.125},
			expected: RGBAF64{1.0, 0.55, 0.0, 1.0},
		},
		{
			value:    RGBAF64{1.04, 0.078, -0.12, 1.01},
			expected: RGBAF64{1.0, 0.078, 0.0, 1.0},
		},
		{
			value:    RGBAF64{1.0, 1.0, 1.0, 1.0},
			expected: RGBAF64{1.0, 1.0, 1.0, 1.0},
		},
	}

	for _, c := range cases {
		c.value.Clamp()
		if !rgbaf64Equal(c.value, c.expected, 0.01) {
			t.Errorf("%s: expected: %#v, actual: %#v", "NewRGBAF6", c.expected, c.value)
		}
	}
}

func rgbaf64Equal(a, b RGBAF64, maxDiff float64) bool {
	if math.Abs(a.R-b.R) > maxDiff {
		return false
	}
	if math.Abs(a.G-b.G) > maxDiff {
		return false
	}
	if math.Abs(a.B-b.B) > maxDiff {
		return false
	}
	if math.Abs(a.A-b.A) > maxDiff {
		return false
	}
	return true
}
