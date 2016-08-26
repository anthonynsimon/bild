package adjust

import (
	"image"
	"image/color"
	"testing"

	"github.com/anthonynsimon/bild/math/f64"
	"github.com/anthonynsimon/bild/util/compare"
)

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
		actual := Apply(c.value, c.fn)
		if !compare.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "apply "+c.desc, c.expected, actual)
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
		actual := f64.Clamp(c.value, 0.0, 1.0)
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "clampFloat64", c.expected, actual)
		}
	}
}
