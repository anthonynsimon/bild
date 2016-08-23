package bild

import (
	"image"
	"math"
)

func Rotate(img image.Image, angle float64, pivot image.Point) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()
	pivotX, pivotY := float64(pivot.X), float64(pivot.Y)
	radians := angle * (math.Pi / 180)

	parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				dx := float64(x) - pivotX
				dy := float64(y) - pivotY

				ix := int(math.Cos(radians)*dx - math.Sin(radians)*dy + pivotX - 0.5)
				iy := int(math.Sin(radians)*dx + math.Cos(radians)*dy + pivotY - 0.5)

				if ix < 0 || ix >= w || iy < 0 || iy >= h {
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

// FlipH returns a horizontally flipped version of the image.
func FlipH(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()

	parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				iy := y * dst.Stride
				pos := iy + (x * 4)
				flippedX := w - x - 1
				flippedPos := iy + (flippedX * 4)

				dst.Pix[pos+0] = src.Pix[flippedPos+0]
				dst.Pix[pos+1] = src.Pix[flippedPos+1]
				dst.Pix[pos+2] = src.Pix[flippedPos+2]
				dst.Pix[pos+3] = src.Pix[flippedPos+3]
			}
		}
	})

	return dst
}

// FlipV returns a vertically flipped version of the image.
func FlipV(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()

	parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				pos := y*dst.Stride + (x * 4)
				flippedY := h - y - 1
				flippedPos := flippedY*dst.Stride + (x * 4)

				dst.Pix[pos+0] = src.Pix[flippedPos+0]
				dst.Pix[pos+1] = src.Pix[flippedPos+1]
				dst.Pix[pos+2] = src.Pix[flippedPos+2]
				dst.Pix[pos+3] = src.Pix[flippedPos+3]
			}
		}
	})

	return dst
}
