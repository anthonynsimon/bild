/*Package blur provides image blurring functions.*/
package blur

import (
	"image"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/convolution"
)

// Box returns a blurred (average) version of the image.
// Radius must be larger than 0.
func Box(src image.Image, radius float64) *image.RGBA {
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

	return convolution.Convolve(src, k.Normalized(), &convolution.Options{Bias: 0, Wrap: false, KeepAlpha: false})
}

// Gaussian returns a smoothly blurred version of the image using
// a Gaussian function. Radius must be larger than 0.
func Gaussian(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return clone.AsRGBA(src)
	}

	// Create the 1-d gaussian kernel
	length := int(math.Ceil(2*radius + 1))
	k := convolution.NewKernel(length, 1)
	for i, x := 0, -radius; i < length; i, x = i+1, x+1 {
		k.Matrix[i] = math.Exp(-(x * x / 4 / radius))
	}
	normK := k.Normalized()

	// Perform separable convolution
	options := convolution.Options{Bias: 0, Wrap: false, KeepAlpha: false}
	result := convolution.Convolve(src, normK, &options)
	result = convolution.Convolve(result, normK.Transposed(), &options)

	return result
}
