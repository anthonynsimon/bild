package paint

import (
	"image"
	"image/color"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

func TestFloodFill(t *testing.T) {
	cases := []struct {
		startPoint image.Point
		fillColor  color.Color
		fuzz       float64
		value      image.Image
		expected   *image.RGBA
	}{
		{
			fuzz:       1.0,
			fillColor:  color.RGBA{0XAA, 0XAA, 0XAA, 0XAA},
			startPoint: image.Point{0, 0},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0XAA, 0XAA, 0XAA, 0XAA, 0x00, 0x00, 0x00, 0x00, 0XAA, 0XAA, 0XAA, 0XAA,
					0XAA, 0XAA, 0XAA, 0XAA, 0x00, 0x00, 0x00, 0x00, 0XAA, 0XAA, 0XAA, 0XAA,
					0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA,
				},
			},
		},
		{
			fuzz:       25.0,
			fillColor:  color.RGBA{0XAA, 0XAA, 0XAA, 0XAA},
			startPoint: image.Point{0, 0},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x88, 0x88, 0x88, 0x88,
					0xDD, 0xDD, 0xDD, 0xDD, 0x00, 0x00, 0x00, 0x00, 0x99, 0x99, 0x99, 0x99,
					0xCC, 0xCC, 0xCC, 0xCC, 0xBB, 0xBB, 0xBB, 0xBB, 0xAA, 0xAA, 0xAA, 0xAA,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0XAA, 0XAA, 0XAA, 0XAA, 0x00, 0x00, 0x00, 0x00, 0XAA, 0XAA, 0XAA, 0XAA,
					0XAA, 0XAA, 0XAA, 0XAA, 0x00, 0x00, 0x00, 0x00, 0XAA, 0XAA, 0XAA, 0XAA,
					0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA, 0XAA,
				},
			},
		},
	}

	for _, c := range cases {
		actual := FloodFill(c.value, c.startPoint, c.fillColor, c.fuzz)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v\n", "Flood Fill", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func BenchmarkFloodFill(b *testing.B) {

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	for n := 0; n < b.N; n++ {
		FloodFill(img, image.Point{0, 0}, color.RGBA{128, 0, 128, 128}, 10)
	}
}
