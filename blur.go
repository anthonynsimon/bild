package bild

import (
	"image"
	"math"
)

// BoxBlur returns a blurred (average) version of the image
func BoxBlur(src image.Image, size int) *image.NRGBA {
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			k.Matrix[x][y] = 1
		}
	}

	img := cloneAsNRGBA(src)
	return convolute(img, k)
}

// GaussianBlur returns a smoothly blurred version of the image using
// a Gaussian function
func GaussianBlur(src image.Image, radius float64) *image.NRGBA {
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
