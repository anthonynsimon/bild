package bild

import (
	"image"
	"image/color"
	"math"
)

// Brightness returns a copy of the image with the adjusted brightness.
// Change is the normalized amount of change to be applied (range -1.0 to 1.0).
func Brightness(src image.Image, change float64) *image.RGBA {
	cache := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		cache[i] = uint8(clampFloat64(math.Ceil(float64(i)*(1+change)), 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{cache[c.R], cache[c.G], cache[c.B], c.A}
	}

	img := apply(src, fn)

	return img
}

// Gamma returns a gamma corrected copy of the image. Provided gamma param must be larger than 0.
func Gamma(src image.Image, gamma float64) *image.RGBA {
	gamma = math.Max(0.00001, gamma)

	cache := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		cache[i] = uint8(clampFloat64(math.Ceil(math.Pow(float64(i)/255, 1.0/gamma)*255), 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{cache[c.R], cache[c.G], cache[c.B], c.A}
	}

	img := apply(src, fn)

	return img
}
