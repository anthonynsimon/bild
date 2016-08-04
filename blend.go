package bild

import (
	"image"
	"image/color"
)

type normColor struct {
	r, g, b, a float64
}

// Add returns an image with the added color values of two images
func Add(a image.Image, b image.Image) *image.RGBA {
	dst := blendOperation(a, b, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

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

// Multiply returns an image with the normalized color values of two images multiplied
func Multiply(a image.Image, b image.Image) *image.RGBA {
	dst := blendOperation(a, b, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

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

// Overlay returns an image that combines Multiply and Screen blend modes
func Overlay(a image.Image, b image.Image) *image.RGBA {
	dst := blendOperation(a, b, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

		r0 := float64(c0.R) / 255
		g0 := float64(c0.G) / 255
		b0 := float64(c0.B) / 255
		a0 := float64(c0.A) / 255

		r1 := float64(c1.R) / 255
		g1 := float64(c1.G) / 255
		b1 := float64(c1.B) / 255
		a1 := float64(c1.A) / 255

		var r2, g2, b2, a2 uint8
		if 0.3*r1+0.6*g1+0.1*b1 < 0.5 {
			r2 = uint8(clampFloat64(r0*r1*2*255, 0, 255))
			g2 = uint8(clampFloat64(g0*g1*2*255, 0, 255))
			b2 = uint8(clampFloat64(b0*b1*2*255, 0, 255))
			a2 = uint8(clampFloat64(a0*a1*2*255, 0, 255))
		} else {
			r2 = uint8(clampFloat64((1-2*(1-r0)*(1-r1))*255, 0, 255))
			g2 = uint8(clampFloat64((1-2*(1-g0)*(1-g1))*255, 0, 255))
			b2 = uint8(clampFloat64((1-2*(1-b0)*(1-b1))*255, 0, 255))
			a2 = uint8(clampFloat64((1-2*(1-a0)*(1-a1))*255, 0, 255))
		}

		return color.RGBA{r2, g2, b2, a2}
	})

	return dst
}

// SoftLight returns an image has the Soft Light blend mode applied
func SoftLight(a image.Image, b image.Image) *image.RGBA {
	dst := blendOperation(a, b, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

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

// Screen returns an image that has the screen blend mode applied
func Screen(a image.Image, b image.Image) *image.RGBA {
	dst := blendOperation(a, b, func(c0 color.RGBA, c1 color.RGBA) color.RGBA {

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

func blendOperation(a image.Image, b image.Image, fn func(color.RGBA, color.RGBA) color.RGBA) *image.RGBA {
	// Currently only equal size images are supported
	if a.Bounds() != b.Bounds() {
		panic("blend operation: only equal size images are supported")
	}

	bounds := a.Bounds()
	srcA := CloneAsRGBA(a)
	srcB := CloneAsRGBA(b)

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
