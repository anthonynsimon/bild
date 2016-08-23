package bild

import (
	"image"
	"math"
)

// RotationOptions are used to config the Rotate function.
// PreserveBounds will increase the output size image if necessary to avoid
// cutting the corners when rotating.
// Pivot is the point of anchor the rotation. If not provided, a default value
// of image center will be used.
type RotationOptions struct {
	PreserveBounds bool
	Pivot          *image.Point
}

// Rotate returns a rotated image by the provided angle using the pivot as an anchor.
// Param angle is in degrees and is applied clockwise.
// Param pivot is a point which will be used as an anchor for the rotation.
// Coordinates 0, 0 represent the top left corner of the image.
func Rotate(img image.Image, angle float64, options *RotationOptions) *image.RGBA {
	src := CloneAsRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()

	if angle == 0.0 {
		return src
	}

	// Default pivot position is center of image
	pivotX, pivotY := float64(srcW)/2, float64(srcH)/2
	radians := -angle * (math.Pi / 180)
	dstW, dstH := srcW, srcH

	if options != nil {
		// Reserve larger size in destination image for full image bounds rotation
		if options.PreserveBounds {
			// Pythagorean theorem to get Hypotenuse of bounds which are the circle's
			// diameter which encapsulates it
			targetScale := math.Sqrt((float64(srcW))*(float64(srcW)) + (float64(srcH))*(float64(srcH)))
			percent := math.Abs(math.Sin(radians * 2))
			targetScale = float64(srcW) + ((targetScale - float64(srcW)) * percent)
			dstW, dstH = int(targetScale+0.5), int(targetScale+0.5)
		} else if options.Pivot != nil {
			// A custom pivot only makes sense if PreserveBounds is set to false
			pivotX, pivotY = float64(options.Pivot.X), float64(options.Pivot.Y)
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	offsetX := (dstW - srcW) / 2
	offsetY := (dstH - srcH) / 2

	parallelize(srcH, func(start, end int) {
		// Correct range to include the pixels visible in new bounds
		// Note that cannot be done in parallelize function, otherwise ranges would overlap
		yStart := int((float64(start)/float64(srcH))*float64(dstH)) - offsetY
		yEnd := int((float64(end)/float64(srcH))*float64(dstH)) - offsetY
		xStart := -offsetX
		xEnd := srcW + offsetX

		for x := xStart; x < xEnd; x++ {
			for y := yStart; y < yEnd; y++ {
				dx := float64(x) - pivotX + 0.5
				dy := float64(y) - pivotY + 0.5

				ix := int((math.Cos(radians)*dx - math.Sin(radians)*dy + pivotX))
				iy := int((math.Sin(radians)*dx + math.Cos(radians)*dy + pivotY))

				if ix < 0 || ix >= srcW || iy < 0 || iy >= srcH {
					continue
				}

				srcPos := iy*src.Stride + ix*4
				dstPos := (y+offsetY)*dst.Stride + (x+offsetX)*4

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
