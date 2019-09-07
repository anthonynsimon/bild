package noise

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/anthonynsimon/bild/parallel"
	"github.com/roz3x/bild/perlin"
)

// Fn is a noise function that generates values between 0 and 255.
type Fn func() uint8

var (
	// Uniform distribution noise function.
	Uniform Fn
	// Binary distribution noise function.
	Binary Fn
	// Gaussian distribution noise function.
	Gaussian Fn
)

func init() {
	Uniform = func() uint8 {
		return uint8(rand.Intn(256))
	}
	Binary = func() uint8 {
		return 0xFF * uint8(rand.Intn(2))
	}
	Gaussian = func() uint8 {
		return uint8(rand.NormFloat64()*32.0 + 128.0)
	}
}

// Options to configure the noise generation.
type Options struct {
	// NoiseFn is a noise function that will be called for each pixel
	// on the image being generated.
	NoiseFn Fn
	// Monochrome sets if the resulting image is grayscale or colored,
	// the latter meaning that each RGB channel was filled with different values.
	Monochrome bool
}

// PerlinGenerate outputs the perlin image of given height and width
// and freqency , freq from 0.1 to 2 is a good range
func PerlinGenerate(height, width int, freq float64) *image.RGBA {
	//keep these values as such
	alpha, beta, n := 2., 2., 3

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	p := perlin.NewPerlin(alpha, beta, n, rand.Int63())

	// serial implimentation
	// works well
	// see output folder
	for x := 0.; x < float64(height); x++ {
		for y := 0.; y < float64(width); y++ {
			t := p.Noise2D((x/10)*freq, (y/10)*freq)
			img.Set(int(x), int(y), color.NRGBA{
				R: uint8((t + 1) * 126),
				G: uint8((t + 1) * 126),
				B: uint8((t + 1) * 126),
				A: 255,
			})
		}
	}
	// parrel implimentation but doesnt work quite well
	// which means output image has pixellated effect
	// could be  useful for somecases
	// see output folder

	// parallel.Line(height, func(start, end int) {
	// 	for y := start; y < end; y++ {
	// 		for x := 0; x < width; x++ {
	// 			t := p.Noise2D(float64(x/10)*freq, float64(y/10)*freq)
	// 			img.Set(int(x), int(y), color.NRGBA{
	// 				R: uint8((t + 1) * 126),
	// 				G: uint8((t + 1) * 126),
	// 				B: uint8((t + 1) * 126),
	// 				A: 255,
	// 			})
	// 		}
	// 	}
	// })

	return img
}

// Generate returns an image of the parameter width and height filled
// with the values from a noise function.
// If no options are provided, defaults will be used.
func Generate(width, height int, o *Options) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// Get options or defaults
	noiseFn := Uniform
	monochrome := false
	if o != nil {
		if o.NoiseFn != nil {
			noiseFn = o.NoiseFn
		}
		monochrome = o.Monochrome
	}

	rand.Seed(time.Now().UTC().UnixNano())

	if monochrome {
		fillMonochrome(dst, noiseFn)
	} else {
		fillColored(dst, noiseFn)
	}

	return dst
}

func fillMonochrome(img *image.RGBA, noiseFn Fn) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	parallel.Line(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				pos := y*img.Stride + x*4
				v := noiseFn()

				img.Pix[pos+0] = v
				img.Pix[pos+1] = v
				img.Pix[pos+2] = v
				img.Pix[pos+3] = 0xFF
			}
		}
	})
}

func fillColored(img *image.RGBA, noiseFn Fn) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	parallel.Line(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				pos := y*img.Stride + x*4

				img.Pix[pos+0] = noiseFn()
				img.Pix[pos+1] = noiseFn()
				img.Pix[pos+2] = noiseFn()
				img.Pix[pos+3] = 0xFF
			}
		}
	})
}
