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

	parallel.Parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				dstPos := y*dst.Stride + x*4

				c := color.RGBA{}

				c.R = dst.Pix[dstPos+0]
				c.G = dst.Pix[dstPos+1]
				c.B = dst.Pix[dstPos+2]
				c.A = dst.Pix[dstPos+3]

				c = fn(c)

				dst.Pix[dstPos+0] = c.R
				dst.Pix[dstPos+1] = c.G
				dst.Pix[dstPos+2] = c.B
				dst.Pix[dstPos+3] = c.A
			}
		}
	})

	return dst
}
