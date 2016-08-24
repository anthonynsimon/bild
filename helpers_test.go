package bild

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"testing"
)

func TestCloneAsRGBA(t *testing.T) {
	cases := []struct {
		desc     string
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc: "RGBA",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "RGBA64",
			value: &image.RGBA64{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "NRGBA",
			value: &image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0x80,
					0xFF, 0xFF, 0xFF, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "NRGBA64",
			value: &image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x80,
					0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "Gray",
			value: &image.Gray{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 2,
				Pix: []uint8{
					0x80, 0x80,
					0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0xFF,
				},
			},
		},
		{
			desc: "Gray16",
			value: &image.Gray16{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 2,
				Pix: []uint8{
					0x80, 0x80,
					0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0xFF,
				},
			},
		},
		{
			desc: "Alpha",
			value: &image.Alpha{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1,
				Pix: []uint8{
					0x80,
					0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "Alpha16",
			value: &image.Alpha16{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1,
				Pix: []uint8{
					0x80, 0x80,
					0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "Paletted",
			value: &image.Paletted{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1,
				Palette: color.Palette{
					color.RGBA{0x00, 0x00, 0x00, 0x00},
					color.RGBA{0x80, 0x80, 0x80, 0x80},
					color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
				},
				Pix: []uint8{
					0x1, 0x2,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := CloneAsRGBA(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("CloneAsRGBA from "+c.desc, c.expected, actual))
		}
	}
}

func TestApply(t *testing.T) {
	cases := []struct {
		desc     string
		fn       func(color.RGBA) color.RGBA
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc: "no change",
			fn: func(c color.RGBA) color.RGBA {
				return c
			},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			desc: "plus 128",
			fn: func(c color.RGBA) color.RGBA {
				return color.RGBA{c.R + 128, c.G + 128, c.B + 128, c.A + 128}
			},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "minus 64",
			fn: func(c color.RGBA) color.RGBA {
				return color.RGBA{c.R - 64, c.G - 64, c.B - 64, c.A - 64}
			},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xBF, 0xBF, 0xBF, 0xBF, 0xBF, 0xBF, 0xBF, 0xBF,
					0xBF, 0xBF, 0xBF, 0xBF, 0xBF, 0xBF, 0xBF, 0xBF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := apply(c.value, c.fn)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("apply "+c.desc, c.expected, actual))
		}
	}
}

func TestClampFloat64(t *testing.T) {
	cases := []struct {
		value    float64
		expected float64
	}{
		{-1.0, 0.0},
		{1.0, 1.0},
		{0.5, 0.5},
		{1.01, 1.0},
		{255.0, 1.0},
	}

	for _, c := range cases {
		actual := clampFloat64(c.value, 0.0, 1.0)
		if actual != c.expected {
			t.Error(testFailMessage("clampFloat64", c.expected, actual))
		}
	}
}

func TestQuickSortRGBA(t *testing.T) {
	cases := []struct {
		value    []color.RGBA
		expected []color.RGBA
	}{
		{
			value:    []color.RGBA{{0, 0, 0, 0}},
			expected: []color.RGBA{{0, 0, 0, 0}},
		},
		{
			value:    []color.RGBA{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
			expected: []color.RGBA{{0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}},
		},
		{
			value:    []color.RGBA{{255, 0, 128, 0}, {255, 255, 0, 0}, {0, 0, 0, 0}},
			expected: []color.RGBA{{0, 0, 0, 0}, {255, 0, 128, 0}, {255, 255, 0, 0}},
		},
		{
			value:    []color.RGBA{{255, 255, 128, 0}, {255, 255, 0, 0}, {0, 0, 0, 0}},
			expected: []color.RGBA{{0, 0, 0, 0}, {255, 255, 0, 0}, {255, 255, 128, 0}},
		},
	}

	for _, c := range cases {
		quicksortRGBA(c.value, 0, len(c.value)-1)
		if !rgbaSlicesEqual(c.value, c.expected) {
			t.Error(testFailMessage("quicksortRGBA", c.expected, c.value))
		}
	}
}

func TestRank(t *testing.T) {
	cases := []struct {
		value    color.RGBA
		expected float64
	}{
		{
			value:    color.RGBA{0, 0, 0, 0},
			expected: 0,
		},
		{
			value:    color.RGBA{255, 255, 255, 255},
			expected: 255,
		},
		{
			value:    color.RGBA{128, 128, 128, 255},
			expected: 128,
		},
		{
			value:    color.RGBA{128, 64, 32, 255},
			expected: 80,
		},
	}

	for _, c := range cases {
		actual := math.Ceil(rank(c.value))
		if actual != c.expected {
			t.Error(testFailMessage("rank", c.expected, actual))
		}
	}
}

func rgbaSlicesEqual(a, b []color.RGBA) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func grayscaleImageEqual(a, b *image.Gray) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	if a.Stride != b.Stride {
		return false
	}

	if len(a.Pix) != len(b.Pix) {
		return false
	}

	for i := 0; i < len(a.Pix); i++ {
		if a.Pix[i] != b.Pix[i] {
			return false
		}
	}
	return true
}

func rgbaImageEqual(a, b *image.RGBA) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	if a.Stride != b.Stride {
		return false
	}

	if len(a.Pix) != len(b.Pix) {
		return false
	}

	for i := 0; i < len(a.Pix); i++ {
		if a.Pix[i] != b.Pix[i] {
			return false
		}
	}
	return true
}

func testFailMessage(name string, expected, actual interface{}) string {
	return fmt.Sprintf("Test %s failed.\nExpected: %v\nActual: %v\n", name, expected, actual)
}
