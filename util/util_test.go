package util

import (
	"image/color"
	"math"
	"testing"

	"github.com/anthonynsimon/bild/util/compare"
)

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
		SortRGBA(c.value, 0, len(c.value)-1)
		if !compare.RGBASlicesEqual(c.value, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "SortRGBA", c.expected, c.value)
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
		actual := math.Ceil(Rank(c.value))
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "rank", c.expected, actual)
		}
	}
}
