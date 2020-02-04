package adjust

import (
	"image"
	"image/color"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

// Apply returns a copy of the provided image after applying the provided color function to each pixel.
func Apply(img image.Image, fn func(color.RGBA) color.RGBA) *image.RGBA {
	bounds := img.Bounds()
	dst := clone.AsRGBA(img)
	w, h := bounds.Dx(), bounds.Dy()

	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				dstPos := y*dst.Stride + x*4

				c := color.RGBA{}

				dr := &dst.Pix[dstPos+0]
				dg := &dst.Pix[dstPos+1]
				db := &dst.Pix[dstPos+2]
				da := &dst.Pix[dstPos+3]

				c.R = *dr
				c.G = *dg
				c.B = *db
				c.A = *da

				c = fn(c)

				*dr = c.R
				*dg = c.G
				*db = c.B
				*da = c.A
			}
		}
	})

	return dst
}
