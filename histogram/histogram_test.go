package histogram

import (
	"fmt"
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util/compare"
)

func TestHistogramMin(t *testing.T) {
	cases := []struct {
		hist     Histogram
		expected int
	}{
		{
			hist: Histogram{
				Bins: []int{},
			},
			expected: 0,
		},
		{
			hist: Histogram{
				Bins: []int{0, 10, 50, 60, 70},
			},
			expected: 0,
		},
		{
			hist: Histogram{
				Bins: []int{10, 50, 60, 70},
			},
			expected: 10,
		},
		{
			hist: Histogram{
				Bins: []int{55, 40, 30, 7, 1},
			},
			expected: 1,
		},
		{
			hist: Histogram{
				Bins: []int{55, -40, 30, 7, 1},
			},
			expected: -40,
		},
	}

	for _, c := range cases {
		actual := c.hist.Min()
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "HistogramMin", c.expected, actual)
		}
	}
}

func TestHistogramMax(t *testing.T) {
	cases := []struct {
		hist     Histogram
		expected int
	}{
		{
			hist: Histogram{
				Bins: []int{},
			},
			expected: 0,
		},
		{
			hist: Histogram{
				Bins: []int{0, 10, 50, 60, 70},
			},
			expected: 70,
		},
		{
			hist: Histogram{
				Bins: []int{10, 50, 60, 70},
			},
			expected: 70,
		},
		{
			hist: Histogram{
				Bins: []int{55, 40, 30, 7, 1},
			},
			expected: 55,
		},
		{
			hist: Histogram{
				Bins: []int{55, -40, 30, 7, 1},
			},
			expected: 55,
		},
	}

	for _, c := range cases {
		actual := c.hist.Max()
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "HistogramMax", c.expected, actual)
		}
	}
}

func TestHistogramCumulative(t *testing.T) {
	cases := []struct {
		hist     Histogram
		expected Histogram
	}{
		{
			hist: Histogram{
				Bins: []int{},
			},
			expected: Histogram{
				Bins: []int{},
			}},
		{
			hist: Histogram{
				Bins: []int{0, 10, 50, 60, 70},
			},
			expected: Histogram{
				Bins: []int{0, 10, 60, 120, 190},
			},
		},
		{
			hist: Histogram{
				Bins: []int{10, 50, 60, 70},
			},
			expected: Histogram{
				Bins: []int{10, 60, 120, 190},
			},
		},
		{
			hist: Histogram{
				Bins: []int{55, 40, 30, 7, 1},
			},
			expected: Histogram{
				Bins: []int{55, 95, 125, 132, 133},
			},
		},
		{
			hist: Histogram{
				Bins: []int{55, -40, 30, 7, 1},
			},
			expected: Histogram{
				Bins: []int{55, 15, 45, 52, 53},
			},
		},
	}

	for _, c := range cases {
		actual := c.hist.Cumulative()
		if !histogramBinsEqual(actual.Bins, c.expected.Bins) {
			t.Errorf("%s: expected: %#v, actual: %#v", "HistogramCumulative", c.expected, actual)
		}
	}
}

func TestHistogramImage(t *testing.T) {
	cases := []struct {
		hist     Histogram
		expected *image.Gray
	}{
		{
			hist: Histogram{
				Bins: []int{},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 0, 0),
				Stride: 0,
				Pix:    []uint8{},
			},
		},
		{
			hist: Histogram{
				Bins: []int{0, 10, 50, 60, 70},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 5, 5),
				Stride: 5,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0xFF,
					0x00, 0x00, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			hist: Histogram{
				Bins: []int{0, 0, 0, 0, 0},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 5, 5),
				Stride: 5,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			hist: Histogram{
				Bins: []int{10, 50, 20, 5, 3, 9, 40},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 7, 7),
				Stride: 7,
				Pix: []uint8{
					0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0xFF, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := c.hist.Image()
		if !compare.GrayImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "HistogramImage", formatImageGrayString(c.expected), formatImageGrayString(actual))
		}
	}
}

func TestNewRGBAHistogram(t *testing.T) {
	cases := []struct {
		img      image.Image
		expected RGBAHistogram
	}{
		{
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 0, 0),
				Stride: 0,
				Pix:    []uint8{},
			},
			expected: RGBAHistogram{
				R: Histogram{Bins: make([]int, 256)},
				G: Histogram{Bins: make([]int, 256)},
				B: Histogram{Bins: make([]int, 256)},
				A: Histogram{Bins: make([]int, 256)},
			},
		},
		{
			img: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x80,
					0x00, 0x80, 0xF, 0x20, 0xff, 0x00, 0x00, 0x00,
				},
			},
			expected: RGBAHistogram{
				R: Histogram{Bins: make([]int, 256)},
				G: Histogram{Bins: make([]int, 256)},
				B: Histogram{Bins: make([]int, 256)},
				A: Histogram{Bins: make([]int, 256)},
			},
		},
	}

	// Manually set expected index counts
	cases[1].expected.R.Bins[0] = 3
	cases[1].expected.R.Bins[255] = 1

	cases[1].expected.G.Bins[0] = 2
	cases[1].expected.G.Bins[128] = 1
	cases[1].expected.G.Bins[255] = 1

	cases[1].expected.B.Bins[0] = 3
	cases[1].expected.B.Bins[15] = 1

	cases[1].expected.A.Bins[0] = 1
	cases[1].expected.A.Bins[32] = 1
	cases[1].expected.A.Bins[128] = 1
	cases[1].expected.A.Bins[255] = 1

	for _, c := range cases {
		actual := NewRGBAHistogram(c.img)
		if !rgbaHistogramEqual(actual, &c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "NewRGBAHistogram", c.expected, actual)
		}
	}
}

func TestRGBAHistogramCumulative(t *testing.T) {
	cases := []struct {
		hist     RGBAHistogram
		expected RGBAHistogram
	}{
		{
			hist: RGBAHistogram{
				R: Histogram{Bins: []int{}},
				G: Histogram{Bins: []int{}},
				B: Histogram{Bins: []int{}},
				A: Histogram{Bins: []int{}},
			},
			expected: RGBAHistogram{
				R: Histogram{Bins: []int{}},
				G: Histogram{Bins: []int{}},
				B: Histogram{Bins: []int{}},
				A: Histogram{Bins: []int{}},
			},
		},
		{
			hist: RGBAHistogram{
				R: Histogram{Bins: []int{40, 60, 70}},
				G: Histogram{Bins: []int{70, 15, 10}},
				B: Histogram{Bins: []int{10, 255, 255}},
				A: Histogram{Bins: []int{5, -5, 67}},
			},
			expected: RGBAHistogram{
				R: Histogram{Bins: []int{40, 100, 170}},
				G: Histogram{Bins: []int{70, 85, 95}},
				B: Histogram{Bins: []int{10, 265, 520}},
				A: Histogram{Bins: []int{5, 0, 67}},
			},
		},
	}

	for _, c := range cases {
		actual := c.hist.Cumulative()
		if !rgbaHistogramEqual(actual, &c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "RGBAHistogramCumulative", c.expected, actual)
		}
	}
}

func TestRGBAHistogramImage(t *testing.T) {

	cases := []struct {
		hist     *RGBAHistogram
		expected *image.RGBA
	}{
		{
			hist: NewRGBAHistogram(&image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xC0, 0x80, 0x40, 0xFF, 0xC0, 0x80, 0x40, 0xFF,
					0xC0, 0x80, 0x40, 0xFF, 0xC0, 0x80, 0x40, 0xFF,
				},
			}),
			expected: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 256, 256))

				rx := 0xC0
				for y := 0; y < 256; y++ {
					pos := y*img.Stride + rx*4
					img.Pix[pos+0] = 0xFF
				}

				gx := 0x80
				for y := 0; y < 256; y++ {
					pos := y*img.Stride + gx*4
					img.Pix[pos+1] = 0xFF
				}

				bx := 0x40
				for y := 0; y < 256; y++ {
					pos := y*img.Stride + bx*4
					img.Pix[pos+2] = 0xFF
				}

				for y := 0; y < 256; y++ {
					for x := 0; x < 256; x++ {
						pos := y*img.Stride + x*4
						img.Pix[pos+3] = 0xFF
					}
				}

				return img
			}(),
		},
		{
			hist: NewRGBAHistogram(&image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			}),
			expected: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 256, 256))

				for y := 0; y < 256; y++ {
					pos := y * img.Stride
					img.Pix[pos+0] = 0xFF
					img.Pix[pos+1] = 0xFF
					img.Pix[pos+2] = 0xFF
				}

				for y := 0; y < 256; y++ {
					for x := 0; x < 256; x++ {
						pos := y*img.Stride + x*4
						img.Pix[pos+3] = 0xFF
					}
				}

				return img
			}(),
		},
		{
			hist: NewRGBAHistogram(&image.RGBA{}),
			expected: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 256, 256))

				// for y := 0; y < 256; y++ {
				// 	pos := y * img.Stride
				// 	img.Pix[pos+0] = 0xFF
				// 	img.Pix[pos+1] = 0xFF
				// 	img.Pix[pos+2] = 0xFF
				// }

				for y := 0; y < 256; y++ {
					for x := 0; x < 256; x++ {
						pos := y*img.Stride + x*4
						img.Pix[pos+3] = 0xFF
					}
				}

				return img
			}(),
		},
	}

	for _, c := range cases {
		actual := c.hist.Image()
		if !compare.RGBAImageEqual(actual, c.expected) {
			// Actual and expected values too large to display in log
			t.Error("RGBAHistogramImage failed")
		}
	}
}

func rgbaHistogramEqual(a, b *RGBAHistogram) bool {
	if len(a.R.Bins) != len(b.R.Bins) ||
		len(a.G.Bins) != len(b.G.Bins) ||
		len(a.B.Bins) != len(b.B.Bins) ||
		len(a.A.Bins) != len(b.A.Bins) {
		return false
	}

	for i := range a.R.Bins {
		if a.R.Bins[i] != b.R.Bins[i] ||
			a.G.Bins[i] != b.G.Bins[i] ||
			a.B.Bins[i] != b.B.Bins[i] ||
			a.A.Bins[i] != b.A.Bins[i] {
			return false
		}
	}

	return true
}

func histogramBinsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func formatImageGrayString(img *image.Gray) string {
	var result string
	for y := 0; y < img.Bounds().Dy(); y++ {
		result += "\n"
		for x := 0; x < img.Bounds().Dx(); x++ {
			pos := y*img.Stride + x
			result += fmt.Sprintf("%#X, ", img.Pix[pos+0])
		}
	}
	result += "\n"
	return result
}
