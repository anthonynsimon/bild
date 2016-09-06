package transform

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		description string
		angle       float64
		options     *RotationOptions
		value       image.Image
		expected    *image.RGBA
	}{
		{
			description: "angle 0.0 at center",
			angle:       0.0,
			options:     &RotationOptions{ResizeBounds: false},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			description: "angle 90.0 at center",
			angle:       90.0,
			options:     &RotationOptions{ResizeBounds: false},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xE0, 0xE0, 0xE0, 0xFF, 0x9F, 0x9F, 0x9F, 0xFF,
					0x9F, 0x9F, 0x9F, 0xFF, 0xE0, 0xE0, 0xE0, 0xFF,
				},
			},
		},
		{
			description: "angle 180.0 at center",
			angle:       180.0,
			options:     &RotationOptions{ResizeBounds: false},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xED, 0xED, 0xED, 0xFF, 0xED, 0xED, 0xED, 0xFF,
					0x92, 0x92, 0x92, 0xFF, 0x92, 0x92, 0x92, 0xFF,
				},
			},
		},
		{
			description: "angle 360.0 at center",
			angle:       360.0,
			options:     &RotationOptions{ResizeBounds: false},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			description: "angle -90.0 at center",
			angle:       -90.0,
			options:     &RotationOptions{ResizeBounds: false},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x92, 0x92, 0x92, 0xFF, 0xED, 0xED, 0xED, 0xFF,
					0x92, 0x92, 0x92, 0xFF, 0xED, 0xED, 0xED, 0xFF,
				},
			},
		},
		{
			description: "angle -90.0 at middle bottom",
			angle:       -90.0,
			options:     &RotationOptions{ResizeBounds: false, Pivot: &image.Point{1, 2}},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x1F, 0x1F, 0x1F, 0x1F, 0x5, 0x5, 0x5, 0x5,
					0xBC, 0xBC, 0xBC, 0xBC, 0x1F, 0x1F, 0x1F, 0x1F,
				},
			},
		},
		{
			description: "angle 45.0 at center, don't preserve bounds",
			angle:       45.0,
			options:     &RotationOptions{ResizeBounds: true},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 5, 5),
				Stride: 5 * 4,
				Pix: []uint8{
					0x5C, 0x5C, 0x5C, 0x87, 0x85, 0x85, 0x85, 0xF7, 0x81, 0x81, 0x81, 0xFF, 0x8C, 0x8C, 0x8C, 0xD1, 0x33, 0x33, 0x33, 0x33,
					0xF0, 0xF0, 0xF0, 0xF7, 0xD3, 0xD3, 0xD3, 0xFF, 0xAD, 0xAD, 0xAD, 0xFF, 0xEF, 0xEF, 0xEF, 0xFD, 0x95, 0x95, 0x95, 0x95,
					0xDF, 0xDE, 0xDE, 0xDE, 0xFD, 0xDA, 0xDA, 0xDB, 0xFC, 0xB3, 0xB3, 0xB7, 0xF6, 0xEC, 0xEC, 0xEC, 0x7E, 0x7E, 0x7E, 0x7E,
					0x35, 0x2B, 0x2B, 0x2B, 0xC9, 0x6F, 0x6F, 0x6F, 0xF5, 0x7C, 0x7C, 0x7C, 0x82, 0x54, 0x54, 0x54, 0xA, 0xA, 0xA, 0xA,
					0x00, 0x00, 0x00, 0x00, 0x35, 0x1B, 0x1B, 0x1B, 0x79, 0x3D, 0x3D, 0x3D, 0xA, 0x5, 0x5, 0x5, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Rotate(c.value, c.angle, c.options)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Rotate "+c.description, util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestFlipH(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
		},
	}

	for _, c := range cases {
		actual := FlipH(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "FlipH", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestFlipV(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := FlipV(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "FlipV", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func BenchmarkRotation256(b *testing.B) {
	benchRotate(256, 256, 90.0, b)
}

func BenchmarkRotation512(b *testing.B) {
	benchRotate(512, 512, 90.0, b)
}

func BenchmarkRotation1024(b *testing.B) {
	benchRotate(1024, 1024, 90.0, b)
}

func BenchmarkRotation2048(b *testing.B) {
	benchRotate(2048, 2048, 90.0, b)
}

func BenchmarkRotation4096(b *testing.B) {
	benchRotate(4096, 4096, 90.0, b)
}

func benchRotate(w, h int, rot float64, bench *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		Rotate(img, rot, nil)
	}
}
