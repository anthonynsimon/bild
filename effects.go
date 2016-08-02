package bild

import (
	"image"
	"image/color"
	"math"
)

// Grayscale returns a copy of the image in Grayscale using the weights: 0.3R + 0.6G + 0.1B
func Grayscale(src image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {

		v := 0.3*float64(c.R) + 0.6*float64(c.G) + 0.1*float64(c.B)
		result := uint8(clamp(math.Ceil(v), 0, 255))

		return color.RGBA{
			result,
			result,
			result,
			c.A}
	}

	img := apply(src, fn)

	return img
}

// EdgeDetection returns a copy of the image with it's edges marked
func EdgeDetection(src image.Image, radius float64) *image.RGBA {
	size := int(math.Ceil(2*radius + 1))
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			v := -1.0
			if x == size/2 && y == size/2 {
				v = float64(size*size) - 1
			}
			k.Matrix[x][y] = v

		}
	}
	return convolute(src, k, 0)
}

// Emboss returns a copy of the image with a 3D shadow effect
func Emboss(src image.Image) *image.RGBA {
	k := Kernel{[][]float64{
		{-1, -1, 0},
		{-1, 0, 1},
		{0, 1, 1},
	}}

	return convolute(src, &k, 128)
}
