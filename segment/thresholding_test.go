package segment

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

func TestThreshold(t *testing.T) {
	cases := []struct {
		level    uint8
		img      image.Image
		expected *image.Gray
	}{
		{
			level: 0,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
		},
		{
			level: 128,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0x00, 0xFF,
					0xFF, 0x00,
				},
			},
		},
		{
			level: 255,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0x00, 0xFF,
					0xFF, 0x00,
				},
			},
		},
		{
			level: 127,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xC0, 0xC0, 0xC0, 0xFF,
					0x40, 0x40, 0x40, 0x40, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0x00, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Threshold(c.img, c.level)
		if !util.GrayImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %v actual: %v", "Threshold", c.expected, actual)
		}
	}
}
