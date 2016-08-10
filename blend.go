/*Package bild provides a collection of common image processing functions.
The input images must implement the image.Image interface and the functions return an *image.RGBA.

The aim of this project is simplicity in use and development over high performance, but most algorithms are designed to be efficient and make use of parallelism when available.
It is based on standard Go packages to reduce dependecy use and development abstractions.*/
package bild

import (
	"image"
	"math"
)

// Normal combines the foreground and background images by placing the foreground over the
// background using alpha compositing. The resulting image is then returned.
func Normal(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		return alphaComp(c0, c1)
	})

	return dst
}

// Add combines the foreground and background images by adding their values and
// returns the resulting image.
func Add(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := c0.R + c1.R
		g := c0.G + c1.G
		b := c0.B + c1.B

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Multiply combines the foreground and background images by multiplying their
// normalized values and returns the resulting image.
func Multiply(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := c0.R * c1.R
		g := c0.G * c1.G
		b := c0.B * c1.B

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Overlay combines the foreground and background images by using Multiply when channel values < 0.5
// or using Screen otherwise and returns the resulting image.
func Overlay(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		var r, g, b float64
		if c0.R > 0.5 {
			r = 1 - (1-2*(c0.R-0.5))*(1-c1.R)
		} else {
			r = 2 * c0.R * c1.R
		}
		if c0.G > 0.5 {
			g = 1 - (1-2*(c0.G-0.5))*(1-c1.G)
		} else {
			g = 2 * c0.G * c1.G
		}
		if c0.B > 0.5 {
			b = 1 - (1-2*(c0.B-0.5))*(1-c1.B)
		} else {
			b = 2 * c0.B * c1.B
		}

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// SoftLight combines the foreground and background images by using Pegtop's Soft Light formula and
// returns the resulting image.
func SoftLight(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := (1-2*c1.R)*c0.R*c0.R + 2*c0.R*c1.R
		g := (1-2*c1.G)*c0.G*c0.G + 2*c0.G*c1.G
		b := (1-2*c1.B)*c0.B*c0.B + 2*c0.B*c1.B

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})
	return dst
}

// Screen combines the foreground and background images by inverting, multiplying and inverting the output.
// The result is a brighter image which is then returned.
func Screen(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := 1 - (1-c0.R)*(1-c1.R)
		g := 1 - (1-c0.G)*(1-c1.G)
		b := 1 - (1-c0.B)*(1-c1.B)

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Difference calculates the absolute difference between the foreground and background images and
// returns the resulting image.
func Difference(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := math.Abs(c0.R - c1.R)
		g := math.Abs(c0.G - c1.G)
		b := math.Abs(c0.B - c1.B)

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Divide combines the foreground and background images by diving the values from the background
// by the foreground and returns the resulting image.
func Divide(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		var r, g, b float64
		if c1.R == 0 {
			r = 1
		} else {
			r = c0.R / c1.R
		}
		if c1.G == 0 {
			g = 1
		} else {
			g = c0.G / c1.G
		}
		if c1.B == 0 {
			b = 1
		} else {
			b = c0.B / c1.B
		}

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// ColorBurn combines the foreground and background images by dividing the inverted
// background by the foreground image and then inverting the result which is then returned.
func ColorBurn(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		var r, g, b float64
		if c1.R == 0 {
			r = 0
		} else {
			r = 1 - (1-c0.R)/c1.R
		}
		if c1.G == 0 {
			g = 0
		} else {
			g = 1 - (1-c0.G)/c1.G
		}
		if c1.B == 0 {
			b = 0
		} else {
			b = 1 - (1-c0.B)/c1.B
		}

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Exclusion combines the foreground and background images applying the Exclusion blend mode and
// returns the resulting image.
func Exclusion(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := 0.5 - 2*(c0.R-0.5)*(c1.R-0.5)
		g := 0.5 - 2*(c0.G-0.5)*(c1.G-0.5)
		b := 0.5 - 2*(c0.B-0.5)*(c1.B-0.5)

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst

}

// ColorDodge combines the foreground and background images by dividing background by the
// inverted foreground image and returns the result.
func ColorDodge(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		var r, g, b float64
		if c1.R == 1 {
			r = 1
		} else {
			r = c0.R / (1 - c1.R)
		}
		if c1.G == 1 {
			g = 1
		} else {
			g = c0.G / (1 - c1.G)
		}
		if c1.B == 1 {
			b = 1
		} else {
			b = c0.B / (1 - c1.B)
		}

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// LinearBurn combines the foreground and background images by adding them and
// then subtracting 255 (1.0 in normalized scale). The resulting image is then returned.
func LinearBurn(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := c0.R + c1.R - 1
		g := c0.G + c1.G - 1
		b := c0.B + c1.B - 1

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// LinearLight combines the foreground and background images by a mix of a Linear Dodge and
// Linear Burn operation. The resulting image is then returned.
func LinearLight(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		var r, g, b float64
		if c1.R > 0.5 {
			r = c0.R + 2*c1.R - 0.5
		} else {
			r = c0.R + 2*c1.R - 1
		}
		if c1.G > 0.5 {
			g = c0.G + 2*c1.G - 0.5
		} else {
			g = c0.G + 2*c1.G - 1
		}
		if c1.B > 0.5 {
			b = c0.B + 2*c1.B - 0.5
		} else {
			b = c0.B + 2*c1.B - 1
		}

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Subtract combines the foreground and background images by Subtracting the background from the
// foreground. The result is then returned.
func Subtract(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := c1.R - c0.R
		g := c1.G - c0.G
		b := c1.B - c0.B

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Opacity returns an image which blends the two input images by the percentage provided.
// Percent must be of range 0 <= percent <= 1.0
func Opacity(bg image.Image, fg image.Image, percent float64) *image.RGBA {
	percent = clampFloat64(percent, 0, 1.0)

	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := c1.R*percent + (1-percent)*c0.R
		g := c1.G*percent + (1-percent)*c0.G
		b := c1.B*percent + (1-percent)*c0.B

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Darken combines the foreground and background images by picking the darkest value per channel
// for each pixel. The result is then returned.
func Darken(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := math.Min(c0.R, c1.R)
		g := math.Min(c0.G, c1.G)
		b := math.Min(c0.B, c1.B)

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// Lighten combines the foreground and background images by picking the brightest value per channel
// for each pixel. The result is then returned.
func Lighten(bg image.Image, fg image.Image) *image.RGBA {
	dst := blend(bg, fg, func(c0, c1 RGBAF64) RGBAF64 {
		r := math.Max(c0.R, c1.R)
		g := math.Max(c0.G, c1.G)
		b := math.Max(c0.B, c1.B)

		c2 := RGBAF64{r, g, b, c1.A}
		return alphaComp(c0, c2)
	})

	return dst
}

// blend two images together by applying the provided function for each pixel.
func blend(bg image.Image, fg image.Image, fn func(RGBAF64, RGBAF64) RGBAF64) *image.RGBA {
	// Currently only equal size images are supported
	if bg.Bounds() != fg.Bounds() {
		panic("blend operation: only equal size images are supported")
	}

	bounds := bg.Bounds()
	srcA := CloneAsRGBA(bg)
	srcB := CloneAsRGBA(fg)

	dst := image.NewRGBA(bounds)

	w, h := bounds.Max.X, bounds.Max.Y

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {
				pos := y*dst.Stride + x*4
				result := fn(
					NewRGBAF64(srcA.Pix[pos+0], srcA.Pix[pos+1], srcA.Pix[pos+2], srcA.Pix[pos+3]),
					NewRGBAF64(srcB.Pix[pos+0], srcB.Pix[pos+1], srcB.Pix[pos+2], srcB.Pix[pos+3]))

				result.Clamp()
				dst.Pix[pos+0] = uint8(result.R * 255)
				dst.Pix[pos+1] = uint8(result.G * 255)
				dst.Pix[pos+2] = uint8(result.B * 255)
				dst.Pix[pos+3] = uint8(result.A * 255)
			}

		}
	})

	return dst
}
