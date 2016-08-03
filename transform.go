package bild

import "image"

// FlipH returns a horizontally flipped version of the image
func FlipH(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Max.X, dst.Bounds().Max.Y

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

// FlipV returns a vertically flipped version of the image
func FlipV(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Max.X, dst.Bounds().Max.Y

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
