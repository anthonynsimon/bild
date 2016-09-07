package noise

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/histogram"
)

func TestMonochromeNoise(t *testing.T) {
	cases := []struct {
		w, h int
		o    *Options
	}{
		{
			w: 200,
			h: 200,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
		{
			w: 512,
			h: 512,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
		{
			w: 900,
			h: 200,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
		{
			w: 5,
			h: 1000,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
	}

	for _, c := range cases {
		result := Generate(c.w, c.h, c.o)
		checkPixels(result, isMonochrome, true, "UniformNoiseMonochrome: color not monochrome.", 0, t)
	}
}

func TestColorNoise(t *testing.T) {
	cases := []struct {
		w, h int
		o    *Options
	}{
		{
			w: 200,
			h: 200,
			o: &Options{NoiseFn: Uniform, Monochrome: false},
		},
		{
			w: 512,
			h: 512,
			o: &Options{NoiseFn: Uniform, Monochrome: false},
		},
		{
			w: 900,
			h: 200,
			o: &Options{NoiseFn: Uniform, Monochrome: false},
		},
		{
			w: 5,
			h: 1000,
			o: &Options{NoiseFn: Uniform, Monochrome: false},
		},
	}

	for _, c := range cases {
		result := Generate(c.w, c.h, c.o)
		checkPixels(result, isMonochrome, false, "ColorNoise: color is monochrome.", c.w*c.h/10, t)
	}
}

func TestUniformNoise(t *testing.T) {
	cases := []struct {
		w, h int
		o    *Options
	}{
		{
			w: 200,
			h: 200,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
		{
			w: 512,
			h: 512,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
		{
			w: 900,
			h: 200,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
		{
			w: 5,
			h: 1000,
			o: &Options{NoiseFn: Uniform, Monochrome: true},
		},
	}

	for _, c := range cases {
		result := Generate(c.w, c.h, c.o)

		hist := histogram.NewRGBAHistogram(result).Cumulative()

		for i := 1; i < len(hist.R.Bins); i++ {
			// Fail if cumulative histogram does not follow a positive linear slope
			if hist.R.Bins[i] <= hist.R.Bins[i-1] || hist.G.Bins[i] <= hist.G.Bins[i-1] || hist.B.Bins[i] <= hist.B.Bins[i-1] {
				t.Errorf("UniformNoise: non uniform distribution.")
				break
			}
		}
	}
}

func TestBinaryNoise(t *testing.T) {
	cases := []struct {
		w, h int
		o    *Options
	}{
		{
			w: 200,
			h: 200,
			o: &Options{NoiseFn: Binary, Monochrome: true},
		},
		{
			w: 512,
			h: 512,
			o: &Options{NoiseFn: Binary, Monochrome: true},
		},
		{
			w: 900,
			h: 200,
			o: &Options{NoiseFn: Binary, Monochrome: true},
		},
		{
			w: 5,
			h: 1000,
			o: &Options{NoiseFn: Binary, Monochrome: true},
		},
	}

	for _, c := range cases {
		result := Generate(c.w, c.h, c.o)

		hist := histogram.NewRGBAHistogram(result)

		binCount := 0
		for i := 0; i < len(hist.R.Bins); i++ {
			if hist.R.Bins[i] != 0x00 && hist.G.Bins[i] != 0x00 && hist.B.Bins[i] != 0x00 {
				binCount++
			}
		}

		if binCount != 2 {
			t.Errorf("BinaryNoise: non binary distribution.")
			break
		}
	}
}

func BenchmarkUniformMonochrome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate(4096, 4096, &Options{NoiseFn: Uniform, Monochrome: true})
	}
}

func BenchmarkUniformColored(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate(4096, 4096, &Options{NoiseFn: Uniform, Monochrome: false})
	}
}

// checkPixels goes through each pixel in the image, extracting the RGBA channels and passing it through the
// provided test function. If the result of the function and the expected bool don't match, then fail the test
// with the provided message. Tolerance is the error count permitted.
func checkPixels(img *image.RGBA, fn func(r, g, b, a uint8) bool, expected bool, failMsg string, tolerance int, t *testing.T) {
	errorCount := 0
checkLoop:
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			pos := y*img.Stride + x*4
			if fn(img.Pix[pos+0], img.Pix[pos+1], img.Pix[pos+2], img.Pix[pos+3]) != expected {
				errorCount++
				if errorCount > tolerance {
					t.Errorf(failMsg)
					break checkLoop
				}
			}
		}
	}
}

func isMonochrome(r, g, b, a uint8) bool {
	if r == g && g == b {
		return true
	}
	return false
}
