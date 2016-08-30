package transform

import (
	"image"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

func Shear(img image.Image, horizontal, vertical float64) *image.RGBA {
	src := clone.AsRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()

	// Supersample, currently hard set to 2x
	srcW, srcH = srcW*2, srcH*2
	src = Resize(src, srcW, srcH, NearestNeighbor)

	dstW, dstH := srcW, srcH
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	pivotX, pivotY := float64(dstW)/2, float64(dstH)/2

	// // Calculate pixels as percent of pivot shift
	// // Scale by 2 for supersampling
	horizontal = horizontal * 2 / pivotX
	vertical = vertical * 2 / pivotY

	parallel.Line(dstH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < dstW; x++ {
				// Move positions to revolve around pivot
				ix := x - int(pivotX)
				iy := y - int(pivotY)

				// Apply linear transformation
				ix = ix + int(float64(iy)*horizontal)
				iy = iy + int(float64(ix)*vertical)

				// Move positions back to image coordinates
				ix += int(pivotX)
				iy += int(pivotY)

				if ix < 0 || ix >= srcW || iy < 0 || iy >= srcH {
					continue
				}

				srcPos := iy*src.Stride + ix*4
				dstPos := y*dst.Stride + x*4

				dst.Pix[dstPos+0] = src.Pix[srcPos+0]
				dst.Pix[dstPos+1] = src.Pix[srcPos+1]
				dst.Pix[dstPos+2] = src.Pix[srcPos+2]
				dst.Pix[dstPos+3] = src.Pix[srcPos+3]
			}
		}
	})

	// Downsample to original bounds as part of the Supersampling
	dst = Resize(dst, dstW/2, dstH/2, Linear)

	return dst
}
