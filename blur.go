package bild

import (
	"image"
	"math"
)

// BoxBlur returns a blurred (average) version of the image.
// Radius must be larger than 0.
func BoxBlur(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return CloneAsRGBA(src)
	}

	size := int(math.Ceil(2*radius + 1))
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			k.Matrix[x][y] = 1
		}
	}

	return Convolute(src, k.Normalized(), &ConvolutionOptions{Bias: 0, Wrap: true, CarryAlpha: false})
}

// GaussianBlur returns a smoothly blurred version of the image using
// a Gaussian function. Radius must be larger than 0.
func GaussianBlur(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return CloneAsRGBA(src)
	}

	size := int(math.Ceil(2*radius + 1))
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			k.Matrix[x][y] = gaussianFunc(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}

	return Convolute(src, k.Normalized(), &ConvolutionOptions{Bias: 0, Wrap: true, CarryAlpha: false})
}
