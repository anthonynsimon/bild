/*Package bild provides a collection of common image processing functions.
The input images must implement the image.Image interface and the functions return an *image.RGBA.

The aim of this project is simplicity in use and development over high performance, but most algorithms are designed to be efficient and make use of parallelism when available.
It is based on standard Go packages to reduce dependecy use and development abstractions.*/
package bild

import (
	"image"
	"image/color"
	"math"
)

type normColor struct {
	r, g, b, a float64
}

// Add combines the foreground and background images by adding their values and
// returns the resulting image.
func Add(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(clampFloat64(r0+r1, 0, 255))
		g2 := uint8(clampFloat64(g0+g1, 0, 255))
		b2 := uint8(clampFloat64(b0+b1, 0, 255))
		a2 := uint8(clampFloat64(a0+a1, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Multiply combines the foreground and background images by multiplying their
// normalized values and returns the resulting image.
func Multiply(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(r0 * r1 / 255)
		g2 := uint8(g0 * g1 / 255)
		b2 := uint8(b0 * b1 / 255)
		a2 := uint8(a0 * a1 / 255)

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Overlay combines the foreground and background images by using Multiply when channel values < 0.5
// or using Screen otherwise and returns the resulting image.
func Overlay(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		var r2, g2, b2, a2 uint8
		if r1 < 0.5 {
			r2 = uint8(clampFloat64(r0*r1*2*255, 0, 255))
		} else {
			r2 = uint8(clampFloat64((1-2*(1-r0)*(1-r1))*255, 0, 255))
		}
		if g1 < 0.5 {
			g2 = uint8(clampFloat64(g0*g1*2*255, 0, 255))
		} else {
			g2 = uint8(clampFloat64((1-2*(1-g0)*(1-g1))*255, 0, 255))
		}
		if b1 < 0.5 {
			b2 = uint8(clampFloat64(b0*b1*2*255, 0, 255))
		} else {
			b2 = uint8(clampFloat64((1-2*(1-b0)*(1-b1))*255, 0, 255))
		}
		if a1 < 0.5 {
			a2 = uint8(clampFloat64(a0*a1*2*255, 0, 255))
		} else {
			a2 = uint8(clampFloat64((1-2*(1-a0)*(1-a1))*255, 0, 255))
		}

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// SoftLight combines the foreground and background images by using Pegtop's Soft Light formula and
// returns the resulting image.
func SoftLight(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64(((1-2*r1)*r0*r0+2*r0*r1)*255, 0, 255))
		g2 := uint8(clampFloat64(((1-2*g1)*g0*g0+2*g0*g1)*255, 0, 255))
		b2 := uint8(clampFloat64(((1-2*b1)*b0*b0+2*b0*b1)*255, 0, 255))
		a2 := uint8(clampFloat64(((1-2*a1)*a0*a0+2*a0*a1)*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Screen combines the foreground and background images by inverting, multiplying and inverting the output.
// The result is a brighter image which is then returned.
func Screen(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64((1-(1-r0)*(1-r1))*255, 0, 255))
		g2 := uint8(clampFloat64((1-(1-g0)*(1-g1))*255, 0, 255))
		b2 := uint8(clampFloat64((1-(1-b0)*(1-b1))*255, 0, 255))
		a2 := uint8(clampFloat64((1-(1-a0)*(1-a1))*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Difference calculates the absolute difference between the foreground and background images and
// returns the resulting image.
func Difference(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64(math.Abs(r0-r1)*255, 0, 255))
		g2 := uint8(clampFloat64(math.Abs(g0-g1)*255, 0, 255))
		b2 := uint8(clampFloat64(math.Abs(b0-b1)*255, 0, 255))
		a2 := uint8(clampFloat64(math.Abs(a0-a1)*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Divide combines the foreground and background images by diving the values from the background
// by the foreground and returns the resulting image.
func Divide(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64((r0/r1)*255, 0, 255))
		g2 := uint8(clampFloat64((g0/g1)*255, 0, 255))
		b2 := uint8(clampFloat64((b0/b1)*255, 0, 255))
		a2 := uint8(clampFloat64((a0/a1)*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// ColorBurn combines the foreground and background images by dividing the inverted
// background by the foreground image and then inverting the result which is then returned.
func ColorBurn(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64((1-(1-r0)/r1)*255, 0, 255))
		g2 := uint8(clampFloat64((1-(1-g0)/g1)*255, 0, 255))
		b2 := uint8(clampFloat64((1-(1-b0)/b1)*255, 0, 255))
		a2 := uint8(clampFloat64((1-(1-a0)/a1)*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Exclusion combines the foreground and background images applying the Exclusion blend mode and
// returns the resulting image.
func Exclusion(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64((r0+r1-2*r0*r1)*255, 0, 255))
		g2 := uint8(clampFloat64((g0+g1-2*g0*g1)*255, 0, 255))
		b2 := uint8(clampFloat64((b0+b1-2*b0*b1)*255, 0, 255))
		a2 := uint8(clampFloat64((a0+a1-2*a0*a1)*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// ColorDodge combines the foreground and background images by dividing background by the
// inverted foreground image and returns the result.
func ColorDodge(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64((r0/(1-r1))*255, 0, 255))
		g2 := uint8(clampFloat64((g0/(1-g1))*255, 0, 255))
		b2 := uint8(clampFloat64((b0/(1-b1))*255, 0, 255))
		a2 := uint8(clampFloat64((a0/(1-a1))*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// LinearBurn combines the foreground and background images by adding them and
// then subtracting 255 (1.0 in normalized scale). The resulting image is then returned.
func LinearBurn(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(clampFloat64(r1+r0-255, 0, 255))
		g2 := uint8(clampFloat64(g1+g0-255, 0, 255))
		b2 := uint8(clampFloat64(b1+b0-255, 0, 255))
		a2 := uint8(clampFloat64(a1+a0-255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// LinearLight combines the foreground and background images by a mix of a Linear Dodge and
// Linear Burn operation. The resulting image is then returned.
func LinearLight(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(clampFloat64(r1+2*r0-255, 0, 255))
		g2 := uint8(clampFloat64(g1+2*g0-255, 0, 255))
		b2 := uint8(clampFloat64(b1+2*b0-255, 0, 255))
		a2 := uint8(clampFloat64(a1+2*a0-255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Subtract combines the foreground and background images by Subtracting the background from the
// foreground. The result is then returned.
func Subtract(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(clampFloat64(r0-r1, 0, 255))
		g2 := uint8(clampFloat64(g0-g1, 0, 255))
		b2 := uint8(clampFloat64(b0-b1, 0, 255))
		a2 := uint8(clampFloat64(a0-a1, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Opacity returns an image which blends the two input images by the percentage provided.
// Percent must be of range 0 <= percent <= 1.0
func Opacity(bg image.Image, fg image.Image, percent float64) *image.RGBA {
	percent = clampFloat64(percent, 0, 1.0)

	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		r2 := uint8(clampFloat64((percent*r1+(1-percent)*r0)*255, 0, 255))
		g2 := uint8(clampFloat64((percent*g1+(1-percent)*g0)*255, 0, 255))
		b2 := uint8(clampFloat64((percent*b1+(1-percent)*b0)*255, 0, 255))
		a2 := uint8(clampFloat64((percent*a1+(1-percent)*a0)*255, 0, 255))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Darken combines the foreground and background images by picking the darkest value per channel
// for each pixel. The result is then returned.
func Darken(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {
		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(math.Min(r0, r1))
		g2 := uint8(math.Min(g0, g1))
		b2 := uint8(math.Min(b0, b1))
		a2 := uint8(math.Min(a0, a1))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// Lighten combines the foreground and background images by picking the brightest value per channel
// for each pixel. The result is then returned.
func Lighten(bg image.Image, fg image.Image) *image.RGBA {
	dst := blendOperation(bg, fg, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {
		r0 := float64(c0.R)
		g0 := float64(c0.G)
		b0 := float64(c0.B)
		a0 := float64(c0.A)

		r1 := float64(c1.R)
		g1 := float64(c1.G)
		b1 := float64(c1.B)
		a1 := float64(c1.A)

		r2 := uint8(math.Max(r0, r1))
		g2 := uint8(math.Max(g0, g1))
		b2 := uint8(math.Max(b0, b1))
		a2 := uint8(math.Max(a0, a1))

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

func blendOperation(bg image.Image, fg image.Image, fn func(color.RGBA, color.RGBA) color.RGBA) *image.RGBA {
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
				result := fn(color.RGBA{srcA.Pix[pos+0], srcA.Pix[pos+1], srcA.Pix[pos+2], srcA.Pix[pos+3]},
					color.RGBA{srcB.Pix[pos+0], srcB.Pix[pos+1], srcB.Pix[pos+2], srcB.Pix[pos+3]})

				dst.Pix[pos+0] = result.R
				dst.Pix[pos+1] = result.G
				dst.Pix[pos+2] = result.B
				dst.Pix[pos+3] = result.A
			}

		}
	})

	return dst
}
