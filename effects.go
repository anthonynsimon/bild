package bild

import (
	"image"
	"image/color"
	"math"
)

// Grayscale returns a copy of the image in Grayscale using the weights: 0.3R + 0.6G + 0.1B
func Grayscale(src image.Image) *image.NRGBA {
	img := cloneAsNRGBA(src)

	fn := func(c color.NRGBA) color.NRGBA {

		v := 0.3*float64(c.R) + 0.6*float64(c.G) + 0.1*float64(c.B)
		result := uint8(clamp(math.Ceil(v), 0, 255))

		return color.NRGBA{
			result,
			result,
			result,
			c.A}
	}

	apply(img, fn)

	return img
}
