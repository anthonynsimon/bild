package clone

import (
	"image"
	"image/color"
	"testing"

	"github.com/anthonynsimon/bild/image/compare"
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
		actual := AsRGBA(c.value)
		if !compare.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "CloneAsRGBA from "+c.desc, c.expected, actual)
		}
	}
}
