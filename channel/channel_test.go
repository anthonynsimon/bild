package channel

import "testing"
import "image"
import "github.com/anthonynsimon/bild/image/compare"

func TestExtract(t *testing.T) {
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
		actual := Extract(c.img, c.channel)
		if !compare.GrayImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual %#v", "Extract "+c.description, c.expected, actual)
		}
	}
}
