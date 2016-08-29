package transform

import (
	"image"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

func Shear(img image.Image, amountX, amountY float64) *image.RGBA {
	src := clone.AsRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()

	// Default pivot position is center of image
	pivotX, pivotY := float64(srcW)*amountX/2, float64(srcH)*amountY/2

	dstW, dstH := srcW, srcH
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	parallel.Line(srcH, func(start, end int) {

		for y := start; y < end; y++ {
			for x := 0; x < dstW; x++ {
				ix := int(float64(x) + float64(y)*amountX - pivotX)
				iy := int(float64(y) + float64(x)*amountY - pivotY)

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

	return dst
}
