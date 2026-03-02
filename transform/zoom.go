package transform

import (
	"image"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

// ZoomOptions are the zoom parameters.
// ResizeBounds set to false will keep the original image bounds, cropping any
// pixels that fall outside when zooming in, and leaving transparent pixels when zooming out.
// Pivot is the point of anchor for the zoom. Default of center is used if nil is passed.
// If ResizeBounds is set to true, a center pivot will always be used.
type ZoomOptions struct {
	ResizeBounds bool
	Pivot        *image.Point
}

// Zoom returns a zoomed version of the image by the provided factor using the pivot as an anchor.
// A factor greater than 1.0 zooms in, and a factor less than 1.0 zooms out.
// Default parameters are used if a nil *ZoomOptions is passed.
//
// Usage example:
//
//	// Zoom in 2x, preserving the image size with the pivot at the center
//	result := transform.Zoom(img, 2.0, nil)
//
//	// Zoom in 2x, expanding the canvas to fit the full zoomed image
//	result := transform.Zoom(img, 2.0, &transform.ZoomOptions{ResizeBounds: true})
func Zoom(img image.Image, factor float64, options *ZoomOptions) *image.RGBA {
	src := clone.AsShallowRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()

	if factor == 1.0 {
		return src
	}

	// Config defaults
	resizeBounds := false
	pivotX, pivotY := float64(srcW)/2, float64(srcH)/2
	if options != nil {
		resizeBounds = options.ResizeBounds
		if options.Pivot != nil {
			pivotX, pivotY = float64(options.Pivot.X), float64(options.Pivot.Y)
		}
	}

	var dstW, dstH int
	if resizeBounds {
		pivotX, pivotY = float64(srcW)/2, float64(srcH)/2
		dstW = int(float64(srcW)*factor + 0.5)
		dstH = int(float64(srcH)*factor + 0.5)
	} else {
		dstW, dstH = srcW, srcH
	}
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	// Calculate offsets when resizing bounds to center the zoomed content
	offsetX := (dstW - srcW) / 2
	offsetY := (dstH - srcH) / 2

	invFactor := 1.0 / factor

	parallel.Line(dstH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < dstW; x++ {
				// Map destination pixel back to source using inverse zoom
				// Adjust for offset when bounds are resized
				dx := float64(x-offsetX) - pivotX + 0.5
				dy := float64(y-offsetY) - pivotY + 0.5

				ix := int(dx*invFactor + pivotX)
				iy := int(dy*invFactor + pivotY)

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
