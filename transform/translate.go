package transform

import (
	"image"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

// Translate repositions a copy of the provided image by dx on the x-axis and
// by dy on the y-axis and returns the result. The bounds from the provided image
// will be kept.
// A positive dx value moves the image towards the right and a positive dy value
// moves the image upwards.
func Translate(img image.Image, dx, dy int) *image.RGBA {
	src := clone.AsRGBA(img)

	if dx == 0 && dy == 0 {
		return src
	}

	w, h := src.Bounds().Dx(), src.Bounds().Dy()
	dst := image.NewRGBA(src.Bounds())

	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				ix, iy := x-dx, y+dy

				if ix < 0 || ix >= w || iy < 0 || iy >= h {
					continue
				}

				srcPos := iy*src.Stride + ix*4
				dstPos := y*src.Stride + x*4

				copy(dst.Pix[dstPos:dstPos+4], src.Pix[srcPos:srcPos+4])
			}
		}
	})

	return dst
}
