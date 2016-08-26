package transform

import (
	"image"
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
// 		result := bild.Rotate(img, 90.0, &bild.RotationOptions{PreserveSize: true, Pivot: &image.Point{0, 0}})
//
func Rotate(img image.Image, angle float64, options *RotationOptions) *image.RGBA {
	src := clone.AsRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()

	// Return early if nothing to do
	absAngle := int(math.Abs(angle) + 0.5)
	if absAngle%360 == 0 {
		return src
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

	// Supersample, currently hard set to 2x
	srcW, srcH = srcW*2, srcH*2
	src = Resize(src, srcW, srcH, NearestNeighbor)
	pivotX, pivotY = pivotX*2, pivotY*2

	// Convert to radians, positive degree maps to clockwise rotation
	angleRadians := -angle * (math.Pi / 180)

	var dstW, dstH int
	if resizeBounds {
		// Reserve larger size in destination image for full image bounds rotation
		// If not preserving size, always take image center as pivot
		pivotX, pivotY = float64(srcW)/2, float64(srcH)/2

		a := math.Abs(float64(srcW) * math.Sin(angleRadians))
		b := math.Abs(float64(srcW) * math.Cos(angleRadians))
		c := math.Abs(float64(srcH) * math.Sin(angleRadians))
		d := math.Abs(float64(srcH) * math.Cos(angleRadians))

		dstW, dstH = int(c+b+0.5), int(a+d+0.5)
	} else {
		dstW, dstH = srcW, srcH
	}
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	// Calculate offsets in case entire image is being displayed
	// Otherwise areas clipped by rotation won't be available
	offsetX := (dstW - srcW) / 2
	offsetY := (dstH - srcH) / 2

	parallel.Parallelize(srcH, func(start, end int) {
		// Correct range to include the pixels visible in new bounds
		// Note that cannot be done in parallelize function input height, otherwise ranges would overlap
		yStart := int((float64(start)/float64(srcH))*float64(dstH)) - offsetY
		yEnd := int((float64(end)/float64(srcH))*float64(dstH)) - offsetY
		xStart := -offsetX
		xEnd := srcW + offsetX

		for y := yStart; y < yEnd; y++ {
			for x := xStart; x < xEnd; x++ {
				dx := float64(x) - pivotX + 0.5
				dy := float64(y) - pivotY + 0.5

				ix := int((math.Cos(angleRadians)*dx - math.Sin(angleRadians)*dy + pivotX))
				iy := int((math.Sin(angleRadians)*dx + math.Cos(angleRadians)*dy + pivotY))

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

	// Downsample to original bounds as part of the Supersampling
	dst = Resize(dst, dstW/2, dstH/2, Linear)

	return dst
}

// FlipH returns a horizontally flipped version of the image.
func FlipH(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	src := clone.AsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()

	parallel.Parallelize(h, func(start, end int) {
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
	src := clone.AsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()

	parallel.Parallelize(h, func(start, end int) {
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
