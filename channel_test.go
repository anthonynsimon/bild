package bild

import "testing"
import "image"

func TestExtractChannel(t *testing.T) {
	cases := []struct {
		description string
		channel     Channel
		img         image.Image
		expected    *image.Gray
	}{
		{
			description: "red empty image",
			channel:     Red,
			img:         &image.RGBA{},
			expected:    &image.Gray{},
		},
		{
			description: "red empty pix",
			channel:     Red,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 0, 0),
				Stride: 0 * 4,
				Pix:    []uint8{},
			},
			expected: &image.Gray{},
		},
		{
			description: "red single pixel",
			channel:     Red,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x20,
				}},
		},
		{
			description: "green single pixel",
			channel:     Green,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x60,
				}},
		},
		{
			description: "blue single pixel",
			channel:     Blue,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x90,
				}},
		},
		{
			description: "alpha single pixel",
			channel:     Alpha,
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0xFF,
				}},
		},
	}

	for _, c := range cases {
		actual := ExtractChannel(c.img, c.channel)
		if !grayscaleImageEqual(actual, c.expected) {
			t.Error(testFailMessage("ExtractChannel "+c.description, c.expected, actual))
		}
	}
}
