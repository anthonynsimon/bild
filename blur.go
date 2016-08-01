package bild

import "image"

// Box Blur returns a blurred (average) version of the image
func Box(src image.Image) *image.NRGBA {
	k := Kernel{[3][3]int32{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}}

	img := cloneAsNRGBA(src)
	return convolute(img, &k)
}

// Gaussian Blur returns a blurred version of the image using
// an approximation to the Gaussian function
func Gaussian(src image.Image, iterations int) *image.NRGBA {
	k := Kernel{[3][3]int32{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}}

	result := cloneAsNRGBA(src)
	for i := 0; i < iterations; i++ {
		result = convolute(result, &k)
	}

	return result
}
