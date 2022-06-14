package transform

import (
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

// RotationOptions are the rotation parameters
// ResizeBounds set to false will keep the original image bounds, cutting any
// pixels that go past it when rotating.
// Pivot is the point of anchor for the rotation. Default of center is used if a nil is passed.
// If ResizeBounds is set to true, a center pivot will always be used.
type RotationOptions struct {
	ResizeBounds bool
	Pivot        *image.Point
}

// Rotate returns a rotated image by the provided angle using the pivot as an anchor.
// Parameters angle is in degrees and it's applied clockwise.
// Default parameters are used if a nil *RotationOptions is passed.
//
// Usage example:
//
// 		// Rotate 90.0 degrees clockwise, preserving the image size and the pivot point at the top left corner
// 		result := transform.Rotate(img, 90.0, &transform.RotationOptions{ResizeBounds: true, Pivot: &image.Point{0, 0}})
//
func Rotate(img image.Image, angle float64, options *RotationOptions) *image.RGBA {
	src := clone.AsShallowRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()

	supersample := false
	absAngle := int(math.Abs(angle) + 0.5)
	if absAngle%360 == 0 {
		// Return early if nothing to do
		return src
	} else if absAngle%90 != 0 {
		// Supersampling is required for non-special angles
		// Special angles = 90, 180, 270...
		supersample = true
	}

	// Config defaults
	resizeBounds := false
	// Default pivot position is center of image
	pivotX, pivotY := float64(srcW/2), float64(srcH/2)
	// Get options if provided
	if options != nil {
		resizeBounds = options.ResizeBounds
		if options.Pivot != nil {
			pivotX, pivotY = float64(options.Pivot.X), float64(options.Pivot.Y)
		}
	}

	if supersample {
		// Supersample, currently hard set to 2x
		srcW, srcH = srcW*2, srcH*2
		src = Resize(src, srcW, srcH, NearestNeighbor)
		pivotX, pivotY = pivotX*2, pivotY*2
	}

	// Convert to radians, positive degree maps to clockwise rotation
	angleRadians := -angle * (math.Pi / 180)

	var dstW, dstH int
	var sin, cos = math.Sincos(angleRadians)
	if resizeBounds {
		// Reserve larger size in destination image for full image bounds rotation
		// If not preserving size, always take image center as pivot
		pivotX, pivotY = float64(srcW)/2, float64(srcH)/2

		a := math.Abs(float64(srcW) * sin)
		b := math.Abs(float64(srcW) * cos)
		c := math.Abs(float64(srcH) * sin)
		d := math.Abs(float64(srcH) * cos)

		dstW, dstH = int(c+b+0.5), int(a+d+0.5)
	} else {
		dstW, dstH = srcW, srcH
	}
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	// Calculate offsets in case entire image is being displayed
	// Otherwise areas clipped by rotation won't be available
	offsetX := (dstW - srcW) / 2
	offsetY := (dstH - srcH) / 2

	parallel.Line(srcH, func(start, end int) {
		// Correct range to include the pixels visible in new bounds
		// Note that cannot be done in parallelize function input height, otherwise ranges would overlap
		yStart := int((float64(start)/float64(srcH))*float64(dstH)) - offsetY
		yEnd := int((float64(end)/float64(srcH))*float64(dstH)) - offsetY
		xStart := -offsetX
		xEnd := srcW + offsetX

		for y := yStart; y < yEnd; y++ {
			dy := float64(y) - pivotY + 0.5
			for x := xStart; x < xEnd; x++ {
				dx := float64(x) - pivotX + 0.5

				ix := int((cos*dx - sin*dy + pivotX))
				iy := int((sin*dx + cos*dy + pivotY))

				if ix < 0 || ix >= srcW || iy < 0 || iy >= srcH {
					continue
				}

				red, green, blue, alpha := src.At(ix, iy).RGBA()

				dst.Set(x+offsetX, y+offsetY, color.RGBA64{
					R: uint16(red),
					G: uint16(green),
					B: uint16(blue),
					A: uint16(alpha),
				})
			}
		}
	})

	if supersample {
		// Downsample to original bounds as part of the Supersampling
		dst = Resize(dst, dstW/2, dstH/2, Linear)
	}

	return dst
}

// FlipH returns a horizontally flipped version of the image.
func FlipH(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	src := clone.AsShallowRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()

	parallel.Line(h, func(start, end int) {
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
	src := clone.AsShallowRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()

	parallel.Line(h, func(start, end int) {
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
