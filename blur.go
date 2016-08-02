package bild

import (
	"image"
	"math"
)

// Box Blur returns a blurred (average) version of the image
func Box(src image.Image, size int) *image.NRGBA {
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			k.Matrix[x][y] = 1
		}
	}

	img := cloneAsNRGBA(src)
	return convolute(img, k)
}

// Gaussian Blur returns a blurred version of the image using
// an approximation to the Gaussian function
func Gaussian(src image.Image, radius float64) *image.NRGBA {
	size := int(math.Ceil(radius) * 2)
	k := NewKernel(size + 1)

	for x := 0; x <= size; x++ {
		for y := 0; y <= size; y++ {
			k.Matrix[x][y] = gaussianFunc(float64(x)-radius, float64(y)-radius, radius)
		}
	}

	img := cloneAsNRGBA(src)
	return convolute(img, k)
}
