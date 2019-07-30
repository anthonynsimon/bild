package transform

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

// benchResult is used to avoid having the compiler optimize the benchmark code calls
var benchResult interface{}

func BenchmarkTranslate(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		benchResult = Translate(img, 512, 512)
	}
}

func TestTranslate(t *testing.T) {
	cases := []struct {
		name     string
		dx       int
		dy       int
		img      *image.RGBA
		expected *image.RGBA
	}{
		{
			name: "empty with translation",
			dx:   2,
			dy:   2,
			img: &image.RGBA{
				Stride: 0,
				Rect:   image.Rect(0, 0, 0, 0),
				Pix:    []uint8{},
			},
			expected: &image.RGBA{
				Stride: 0,
				Rect:   image.Rect(0, 0, 0, 0),
				Pix:    []uint8{},
			},
		},
		{
			name: "empty no translation",
			dx:   0,
			dy:   0,
			img: &image.RGBA{
				Stride: 0,
				Rect:   image.Rect(0, 0, 0, 0),
				Pix:    []uint8{},
			},
			expected: &image.RGBA{
				Stride: 0,
				Rect:   image.Rect(0, 0, 0, 0),
				Pix:    []uint8{},
			},
		},
		{
			name: "no translation",
			dx:   0,
			dy:   0,
			img: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
			expected: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
		},
		{
			name: "dx +1",
			dx:   1,
			dy:   0,
			img: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
			expected: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF,
				},
			},
		},
		{
			name: "dy +1",
			dx:   0,
			dy:   1,
			img: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
			expected: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			name: "dx +1 dy +1",
			dx:   1,
			dy:   1,
			img: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
			expected: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			name: "dx -1",
			dx:   -1,
			dy:   0,
			img: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
			expected: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0x40, 0x40, 0x40, 0xFF, 0x00, 0x00, 0x00, 0x00,
					0x20, 0x20, 0x20, 0xFF, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			name: "dy -3",
			dx:   0,
			dy:   -3,
			img: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x20, 0x20, 0x20, 0xFF,
				},
			},
			expected: &image.RGBA{
				Stride: 2 * 4,
				Rect:   image.Rect(0, 0, 2, 2),
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}

	for _, c := range cases {
		result := Translate(c.img, c.dx, c.dy)
		if !util.RGBAImageEqual(result, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v", "Translate "+c.name, util.RGBAToString(c.expected), util.RGBAToString(result))
		}
	}
}
