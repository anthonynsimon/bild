package bild

import (
	"image"
	"math"
)

// BoxBlur returns a blurred (average) version of the image
func BoxBlur(src image.Image, size int) *image.RGBA {
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			k.Matrix[x][y] = 1
		}
	}

	return convolute(src, k.Normalized(), 0)
}

// GaussianBlur returns a smoothly blurred version of the image using
// a Gaussian function
func GaussianBlur(src image.Image, radius float64) *image.RGBA {
	size := int(math.Ceil(2*radius + 1))
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			k.Matrix[x][y] = gaussianFunc(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}

	return convolute(src, k.Normalized(), 0)
}
