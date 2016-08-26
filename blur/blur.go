/*Package blur provides image blurring functions.*/
package blur

import (
	"image"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/convolution"
)

// BoxBlur returns a blurred (average) version of the image.
// Radius must be larger than 0.
func BoxBlur(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return clone.AsRGBA(src)
	}

	length := int(math.Ceil(2*radius + 1))
	k := convolution.NewKernel(length, length)

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			k.Matrix[y*length+x] = 1
		}
	}

	return convolution.Convolve(src, k.Normalized(), &convolution.Options{Bias: 0, Wrap: false, CarryAlpha: false})
}

// GaussianBlur returns a smoothly blurred version of the image using
// a Gaussian function. Radius must be larger than 0.
func GaussianBlur(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return clone.AsRGBA(src)
	}

	length := int(math.Ceil(2*radius + 1))
	k := convolution.NewKernel(length, length)

	gaussianFn := func(x, y, sigma float64) float64 {
		return math.Exp(-(x*x/sigma + y*y/sigma))
	}

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			k.Matrix[y*length+x] = gaussianFn(float64(x)-radius, float64(y)-radius, 4*radius)

		}
	}

	return convolution.Convolve(src, k.Normalized(), &convolution.Options{Bias: 0, Wrap: false, CarryAlpha: false})
}
