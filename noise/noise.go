package bild

import (
	"image"
	"math/rand"
)

// NoiseFn is a noise function that generates values between 0 and 255.
type NoiseFn func() uint8

var (
	// Uniform distribution noise function.
	Uniform NoiseFn
	// Binary distribution noise function.
	Binary NoiseFn
	// Gaussian distribution noise function.
	Gaussian NoiseFn
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
	// Fn is a noise function that will be called for each pixel
	// on the image being generated.
	Fn NoiseFn
	// Monochrome sets if the resulting image is grayscale or colored,
	// the latter meaning that each RGB channel was filled with different values.
	Monochrome bool
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
		if o.Fn != nil {
			noiseFn = o.Fn
		}
		monochrome = o.Monochrome
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*dst.Stride + x*4
			if monochrome {
				dst.Pix[pos+0] = noiseFn()
				dst.Pix[pos+1] = noiseFn()
				dst.Pix[pos+2] = noiseFn()
				dst.Pix[pos+3] = 0xFF
			} else {
				v := noiseFn()
				dst.Pix[pos+0] = v
				dst.Pix[pos+1] = v
				dst.Pix[pos+2] = v
				dst.Pix[pos+3] = 0xFF
			}
		}
	}

	return dst
}
