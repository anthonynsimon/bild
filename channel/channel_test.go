package channel

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

func TestExtractMultiple(t *testing.T) {
	cases := []struct {
		description string
		channels    []Channel
		img         image.Image
		expected    *image.RGBA
	}{
		{
			description: "red empty image",
			channels:    []Channel{Red},
			img:         &image.RGBA{},
			expected:    &image.RGBA{},
		},
		{
			description: "red empty pix",
			channels:    []Channel{Red},
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 0, 0),
				Stride: 0 * 4,
				Pix:    []uint8{},
			},
			expected: &image.RGBA{},
		},
		{
			description: "red single pixel",
			channels:    []Channel{Red},
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x20, 0x00, 0x00, 0x00,
				}},
		},
		{
			description: "green single pixel",
			channels:    []Channel{Green},
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x00, 0x60, 0x00, 0x00,
				}},
		},
		{
			description: "blue single pixel",
			channels:    []Channel{Blue},
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x00, 0x00, 0x90, 0x00,
				}},
		},
		{
			description: "alpha single pixel",
			channels:    []Channel{Alpha},
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF,
				}},
		},
		{
			description: "multiple single pixel",
			channels:    []Channel{Red, Alpha},
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x20, 0x60, 0x90, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1,
				Pix: []uint8{
					0x20, 0x00, 0x00, 0xFF,
				}},
		},
	}

	for _, c := range cases {
		actual := ExtractMultiple(c.img, c.channels...)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual %#v", "Extract "+c.description, c.expected, actual)
		}
	}
}

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
		if !util.GrayImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual %#v", "Extract "+c.description, c.expected, actual)
		}
	}
}
