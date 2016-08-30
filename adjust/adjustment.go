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

func Saturation(img image.Image, value float64) *image.RGBA {
	value = f64.Clamp(value, -1.0, 1.0)

	fn := func(c color.RGBA) color.RGBA {
		h, s, v := util.RGBToHSV(c)
		s += value
		outColor := util.HSVToRGB(h, s, v)
		outColor.A = c.A
		return outColor
	}

	return Apply(img, fn)
}

func Hue(img image.Image, change int) *image.RGBA {
	change = change % 360

	fn := func(c color.RGBA) color.RGBA {
		h, s, v := util.RGBToHSV(c)
		h += float64(change)
		outColor := util.HSVToRGB(h, s, v)
		outColor.A = c.A
		return outColor
	}

	return Apply(img, fn)
}
