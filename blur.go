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

	length := int(math.Ceil(2*radius + 1))
	k := NewKernel(length)

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			k.Matrix[y*length+x] = 1
		}
	}

	return Convolve(src, k.Normalized(), &ConvolutionOptions{Bias: 0, Wrap: true, CarryAlpha: false})
}

// GaussianBlur returns a smoothly blurred version of the image using
// a Gaussian function. Radius must be larger than 0.
func GaussianBlur(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return CloneAsRGBA(src)
	}

	length := int(math.Ceil(2*radius + 1))
	k := NewKernel(length)

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			k.Matrix[y*length+x] = gaussianFunc(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}

	return Convolve(src, k.Normalized(), &ConvolutionOptions{Bias: 0, Wrap: true, CarryAlpha: false})
}
