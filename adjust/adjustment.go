/*Package adjust provides basic color correction functions.*/
package adjust

import (
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/math/f64"
	"github.com/anthonynsimon/bild/util"
)

// Brightness returns a copy of the image with the adjusted brightness.
// Change is the normalized amount of change to be applied (range -1.0 to 1.0).
func Brightness(src image.Image, change float64) *image.RGBA {
	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		lookup[i] = uint8(f64.Clamp(float64(i)*(1+change), 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{lookup[c.R], lookup[c.G], lookup[c.B], c.A}
	}

	img := Apply(src, fn)

	return img
}

// Gamma returns a gamma corrected copy of the image. Provided gamma param must be larger than 0.
func Gamma(src image.Image, gamma float64) *image.RGBA {
	gamma = math.Max(0.00001, gamma)

	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		lookup[i] = uint8(f64.Clamp(math.Pow(float64(i)/255, 1.0/gamma)*255, 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{lookup[c.R], lookup[c.G], lookup[c.B], c.A}
	}

	img := Apply(src, fn)

	return img
}

// Contrast returns a copy of the image with its difference in high and low values adjusted by the change param.
// Change is the normalized amount of change to be applied, in the range of -1.0 to 1.0.
// If Change is set to 0.0, then the values remain the same, if it's set to 0.5, then all values will be moved 50% away from the middle value.
func Contrast(src image.Image, change float64) *image.RGBA {
	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		lookup[i] = uint8(f64.Clamp(((((float64(i)/255)-0.5)*(1+change))+0.5)*255, 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{lookup[c.R], lookup[c.G], lookup[c.B], c.A}
	}

	img := Apply(src, fn)

	return img
}

// Hue adjusts the overall hue of the provided image and returns the result.
// Parameter change is the amount of change to be applied and is of the range
// -360 to 360. It corresponds to the hue angle in the HSL color model.
func Hue(img image.Image, change int) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		h, s, l := util.RGBToHSL(c)
		h = float64((int(h) + change) % 360)
		outColor := util.HSLToRGB(h, s, l)
		outColor.A = c.A
		return outColor
	}

	return Apply(img, fn)
}

// Saturation adjusts the saturation of the image and returns the result.
// Parameter change is the amount of change to be applied and is of the range
// -1.0 to 1.0. It's applied as relative change. For example if the current color
// saturation is 1.0 and the saturation change is set to -0.5, a change of -50%
// will be applied so that the resulting saturation is 0.5 in the HSL color model.
func Saturation(img image.Image, change float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		h, s, l := util.RGBToHSL(c)
		s = f64.Clamp(s*(1+change), 0.0, 1.0)
		outColor := util.HSLToRGB(h, s, l)
		outColor.A = c.A
		return outColor
	}

	return Apply(img, fn)
}
