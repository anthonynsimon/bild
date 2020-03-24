package caman

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/caman/helpers"
	"github.com/anthonynsimon/bild/caman/utils"
)

type RGB struct {
	R, G, B float64
}

// Curves applies curves effect to image
// TEST PASSED
func Curves(img image.Image, chans string, second, third, fourth, fifth [2]float64) *image.RGBA {
	bezier := helpers.Bezier(second, third, fourth, fifth, 0, 255)

	start := second
	for i := 0; i < int(start[0]); i++ {
		if start[0] > 0 {
			bezier[i] = start[1]
		}
	}

	end := fifth
	for i := int(end[0]); i <= 255; i++ {
		if end[0] < 255 {
			bezier[i] = end[1]
		}
	}

	fn := func(c color.RGBA) color.RGBA {
		if strings.Contains(chans, "r") {
			c.R = uint8(math.Round(bezier[int(c.R)]))
		}
		if strings.Contains(chans, "g") {
			c.G = uint8(math.Round(bezier[int(c.G)]))
		}
		if strings.Contains(chans, "b") {
			c.B = uint8(math.Round(bezier[int(c.B)]))
		}

		return c
	}

	return adjust.Apply(img, fn)
}

// Saturation applies Saturation to image
// adj must be negative
// TEST PASSED
func Saturation(img image.Image, adj float64) *image.RGBA {
	adj *= -0.01

	fn := func(c color.RGBA) color.RGBA {
		max := utils.MaxUint8(c.R, utils.MaxUint8(c.G, c.B))
		if c.R != max {
			c.R += uint8(math.Round(float64(max-c.R) * adj))
		}
		if c.G != max {
			c.G += uint8(math.Round(float64(max-c.G) * adj))
		}
		if c.B != max {
			c.B += uint8(math.Round(float64(max-c.B) * adj))
		}

		return c
	}

	return adjust.Apply(img, fn)
}

// Exposure applies exposure effect to image
func Exposure(img image.Image, adj float64) *image.RGBA {

	p := math.Abs(adj) / 100
	ctrl1 := [2]float64{0, 255 * p}
	ctrl2 := [2]float64{255 - (255 * p), 255}

	if adj < 0 {
		// reverse arrays
		ctrl1 = [2]float64{255 * p, 0}
		ctrl2 = [2]float64{255, 255 - (255 * p)}
	}
	return Curves(img, "rgb", [2]float64{0, 0}, ctrl1, ctrl2, [2]float64{255, 255})
}

// Gamma applies gamma effect tom image
func Gamma(img image.Image, adj float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		c.R = uint8(math.Pow(float64(c.R)/255, adj) * 255)
		c.G = uint8(math.Pow(float64(c.G)/255, adj) * 255)
		c.B = uint8(math.Pow(float64(c.B)/255, adj) * 255)

		return c
	}

	return adjust.Apply(img, fn)
}

// Channels aplies channels effect to the image
func Channels(img image.Image, options map[string]float64) *image.RGBA {
	for k := range options {
		options[k] /= 100
	}
	var maxColor float64 = 255

	fn := func(c color.RGBA) color.RGBA {
		if red, ok := options["red"]; ok {
			if red > 0 {
				c.R += uint8(math.Round((maxColor - float64(c.R)) * red))
			} else {
				c.R -= uint8(math.Round(float64(c.R) * -red))
			}
		}
		if green, ok := options["green"]; ok {
			if green > 0 {
				c.G += uint8(math.Round((maxColor - float64(c.G)) * green))
			} else {
				c.G -= uint8(math.Round(float64(c.G) * -green))
			}
		}
		if blue, ok := options["blue"]; ok {
			if blue > 0 {
				c.B += uint8(math.Round((maxColor - float64(c.B)) * blue))
			} else {
				c.B -= uint8(math.Round(float64(c.B) * -blue))
			}
		}
		return c
	}

	return adjust.Apply(img, fn)
}

// Sepia applies sepia effect to the image
func Sepia(img image.Image, adj float64) *image.RGBA {
	adj /= 100
	var maxColor float64 = 255

	fn := func(rgba color.RGBA) color.RGBA {
		floatR := float64(rgba.R)
		floatG := float64(rgba.G)
		floatB := float64(rgba.B)

		rgba.R = uint8(math.Min(maxColor, (floatR*(1-(0.607*adj)))+(floatG*(0.769*adj))+(floatB*(0.189*adj))))
		rgba.G = uint8(math.Min(maxColor, (floatR*(0.349*adj))+(floatG*(1-(0.314*adj)))+(floatB*(0.168*adj))))
		rgba.B = uint8(math.Min(maxColor, (floatR*(0.272*adj))+(floatG*(0.534*adj))+(floatB*(1-(0.869*adj)))))

		return rgba
	}

	return adjust.Apply(img, fn)
}

// Colorize applies colorize effect to image
func Colorize(img image.Image, rgb *RGB, level float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		floatR := float64(c.R)
		floatG := float64(c.G)
		floatB := float64(c.B)

		c.R -= uint8(math.Round((floatR - rgb.R) * level / 100))
		c.G -= uint8(math.Round((floatG - rgb.G) * level / 100))
		c.B -= uint8(math.Round((floatB - rgb.B) * level / 100))
		return c
	}

	return adjust.Apply(img, fn)
}
