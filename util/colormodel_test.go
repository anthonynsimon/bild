package util

import (
	"image/color"
	"math"
	"testing"
)

func TestRGBToHSV(t *testing.T) {
	cases := []struct {
		input    color.RGBA
		expected [3]float64
	}{
		{
			input:    color.RGBA{45, 166, 115, 255},
			expected: [3]float64{155, 0.73, 0.65},
		},
		{
			input:    color.RGBA{0, 255, 0, 255},
			expected: [3]float64{120, 1, 1},
		},
		{
			input:    color.RGBA{242, 220, 97, 255},
			expected: [3]float64{51, 0.6, 0.95},
		},
		{
			input:    color.RGBA{10, 10, 10, 255},
			expected: [3]float64{0, 0.0, 0.04},
		},
		{
			input:    color.RGBA{255, 255, 255, 255},
			expected: [3]float64{0, 0.0, 1.0},
		},
		{
			input:    color.RGBA{0, 0, 0, 255},
			expected: [3]float64{0, 0.0, 0.0},
		},
		{
			input:    color.RGBA{255, 0, 0, 255},
			expected: [3]float64{0, 1.0, 1.0},
		},
		{
			input:    color.RGBA{255, 0, 255, 255},
			expected: [3]float64{300, 1.0, 1.0},
		},
	}

	for _, c := range cases {
		h, s, v := RGBToHSV(c.input)
		h = math.Floor(h + 0.5)
		s = math.Floor((s*100)+0.5) / 100
		v = math.Floor((v*100)+0.5) / 100
		if h != c.expected[0] || s != c.expected[1] || v != c.expected[2] {
			t.Errorf("RGBToHSV failed: expected: %#v, actual: %#v, %#v, %#v", c.expected, h, s, v)
		}
	}
}

func TestHSVToRGB(t *testing.T) {
	cases := []struct {
		input    [3]float64
		expected color.RGBA
	}{
		{
			input:    [3]float64{155, 0.73, 0.65},
			expected: color.RGBA{45, 166, 115, 255},
		},
		{
			input:    [3]float64{120, 1, 1},
			expected: color.RGBA{0, 255, 0, 255},
		},
		{
			input:    [3]float64{51, 0.6, 0.95},
			expected: color.RGBA{242, 220, 97, 255},
		},
		{
			input:    [3]float64{0, 0.0, 0.04},
			expected: color.RGBA{10, 10, 10, 255},
		},
		{
			input:    [3]float64{0, 0.0, 1.0},
			expected: color.RGBA{255, 255, 255, 255},
		},
		{
			input:    [3]float64{0, 0.0, 0.0},
			expected: color.RGBA{0, 0, 0, 255},
		},
		{
			input:    [3]float64{0, 1.0, 1.0},
			expected: color.RGBA{255, 0, 0, 255},
		},
		{
			input:    [3]float64{300, 1.0, 1.0},
			expected: color.RGBA{255, 0, 255, 255},
		},
	}

	for _, c := range cases {
		actual := HSVToRGB(c.input[0], c.input[1], c.input[2])
		if actual != c.expected {
			t.Errorf("HSVToRGB failed: expected: %#v, actual: %#v", c.expected, actual)
		}
	}
}

func TestRGBToHSL(t *testing.T) {
	cases := []struct {
		input    color.RGBA
		expected [3]float64
	}{
		{
			input:    color.RGBA{45, 166, 115, 255},
			expected: [3]float64{155, 0.57, 0.41},
		},
		{
			input:    color.RGBA{0, 255, 0, 255},
			expected: [3]float64{120, 1, 0.5},
		},
		{
			input:    color.RGBA{242, 220, 97, 255},
			expected: [3]float64{51, 0.85, 0.66},
		},
		{
			input:    color.RGBA{10, 10, 10, 255},
			expected: [3]float64{0, 0.0, 0.04},
		},
		{
			input:    color.RGBA{255, 255, 255, 255},
			expected: [3]float64{0, 0.0, 1.0},
		},
		{
			input:    color.RGBA{0, 0, 0, 255},
			expected: [3]float64{0, 0.0, 0.0},
		},
		{
			input:    color.RGBA{255, 0, 0, 255},
			expected: [3]float64{0, 1.0, 0.5},
		},
		{
			input:    color.RGBA{0, 0, 255, 255},
			expected: [3]float64{240, 1.0, 0.5},
		},
		{
			input:    color.RGBA{255, 0, 255, 255},
			expected: [3]float64{300, 1.0, 0.5},
		},
	}

	for _, c := range cases {
		h, s, l := RGBToHSL(c.input)
		h = math.Floor(h + 0.5)
		s = math.Floor((s*100)+0.5) / 100
		l = math.Floor((l*100)+0.5) / 100
		if h != c.expected[0] || s != c.expected[1] || l != c.expected[2] {
			t.Errorf("RGBToHSL failed: expected: %#v, actual: %#v, %#v, %#v", c.expected, h, s, l)
		}
	}
}

func TestHSLToRGB(t *testing.T) {
	cases := []struct {
		input    [3]float64
		expected color.RGBA
	}{
		{
			input:    [3]float64{155, 0.57, 0.41},
			expected: color.RGBA{0x2d, 0xa4, 0x72, 0xff},
		},
		{
			input:    [3]float64{120, 1, 0.5},
			expected: color.RGBA{0, 255, 0, 255},
		},
		{
			input:    [3]float64{51, 0.85, 0.66},
			expected: color.RGBA{0xf2, 0xdc, 0x5f, 0xff},
		},
		{
			input:    [3]float64{0, 0.0, 0.04},
			expected: color.RGBA{10, 10, 10, 255},
		},
		{
			input:    [3]float64{0, 0.0, 1.0},
			expected: color.RGBA{255, 255, 255, 255},
		},
		{
			input:    [3]float64{0, 0.0, 0.0},
			expected: color.RGBA{0, 0, 0, 255},
		},
		{
			input:    [3]float64{0, 1.0, 0.5},
			expected: color.RGBA{255, 0, 0, 255},
		},
		{
			input:    [3]float64{240, 1.0, 0.5},
			expected: color.RGBA{0, 0, 255, 255},
		},
		{
			input:    [3]float64{300, 1.0, 0.5},
			expected: color.RGBA{255, 0, 255, 255},
		},
	}

	for _, c := range cases {
		actual := HSLToRGB(c.input[0], c.input[1], c.input[2])
		if actual != c.expected {
			t.Errorf("HSLToRGB failed: expected: %#v, actual: %#v", c.expected, actual)
		}
	}
}
