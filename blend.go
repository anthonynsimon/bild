package bild

import "image"

// Add returns an image with the added color values of two images
func Add(a image.Image, b image.Image) *image.RGBA {
	// Currently only equal size images are supported
	if a.Bounds() != b.Bounds() {
		panic("only equal size images are supported")
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

				r0 := float64(srcA.Pix[pos+0])
				g0 := float64(srcA.Pix[pos+1])
				b0 := float64(srcA.Pix[pos+2])
				a0 := float64(srcA.Pix[pos+3])

				r1 := float64(srcB.Pix[pos+0])
				g1 := float64(srcB.Pix[pos+1])
				b1 := float64(srcB.Pix[pos+2])
				a1 := float64(srcB.Pix[pos+3])

				dst.Pix[pos+0] = uint8(clampFloat64(r0+r1, 0, 255))
				dst.Pix[pos+1] = uint8(clampFloat64(g0+g1, 0, 255))
				dst.Pix[pos+2] = uint8(clampFloat64(b0+b1, 0, 255))
				dst.Pix[pos+3] = uint8(clampFloat64(a0+a1, 0, 255))
			}
		}
	})

	return dst
}

// Multiply returns an image with the normalized color values of two images multiplied
func Multiply(a image.Image, b image.Image) *image.RGBA {
	// Currently only equal size images are supported
	if a.Bounds() != b.Bounds() {
		panic("only equal size images are supported")
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

				r0 := float64(srcA.Pix[pos+0]) / 255
				g0 := float64(srcA.Pix[pos+1]) / 255
				b0 := float64(srcA.Pix[pos+2]) / 255
				a0 := float64(srcA.Pix[pos+3]) / 255

				r1 := float64(srcB.Pix[pos+0]) / 255
				g1 := float64(srcB.Pix[pos+1]) / 255
				b1 := float64(srcB.Pix[pos+2]) / 255
				a1 := float64(srcB.Pix[pos+3]) / 255

				dst.Pix[pos+0] = uint8(r0 * r1 * 255)
				dst.Pix[pos+1] = uint8(g0 * g1 * 255)
				dst.Pix[pos+2] = uint8(b0 * b1 * 255)
				dst.Pix[pos+3] = uint8(a0 * a1 * 255)
			}
		}
	})

	return dst
}

// Overlay returns an image that combines Multiply and Screen blend modes
func Overlay(a image.Image, b image.Image) *image.RGBA {
	// Currently only equal size images are supported
	if a.Bounds() != b.Bounds() {
		panic("only equal size images are supported")
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

				r0 := float64(srcA.Pix[pos+0]) / 255
				g0 := float64(srcA.Pix[pos+1]) / 255
				b0 := float64(srcA.Pix[pos+2]) / 255
				a0 := float64(srcA.Pix[pos+3]) / 255

				r1 := float64(srcB.Pix[pos+0]) / 255
				g1 := float64(srcB.Pix[pos+1]) / 255
				b1 := float64(srcB.Pix[pos+2]) / 255
				a1 := float64(srcB.Pix[pos+3]) / 255

				if 0.3*r0+0.6*g0+0.1*b0 < 0.5 {
					dst.Pix[pos+0] = uint8(clampFloat64(r0*r1*2*255, 0, 255))
					dst.Pix[pos+1] = uint8(clampFloat64(g0*g1*2*255, 0, 255))
					dst.Pix[pos+2] = uint8(clampFloat64(b0*b1*2*255, 0, 255))
					dst.Pix[pos+3] = uint8(clampFloat64(a0*a1*2*255, 0, 255))
				} else {
					dst.Pix[pos+0] = uint8(clampFloat64((1-(2*(1-r0)*(1-r1)))*255, 0, 255))
					dst.Pix[pos+1] = uint8(clampFloat64((1-(2*(1-g0)*(1-g1)))*255, 0, 255))
					dst.Pix[pos+2] = uint8(clampFloat64((1-(2*(1-b0)*(1-b1)))*255, 0, 255))
					dst.Pix[pos+3] = uint8(clampFloat64((1-(2*(1-a0)*(1-a1)))*255, 0, 255))
				}
			}
		}
	})

	return dst
}
