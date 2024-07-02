package transform

import (
	"image"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/math/f64"
	"github.com/anthonynsimon/bild/parallel"
)

// Resize returns a new image with its size adjusted to the new width and height. The filter
// param corresponds to the Resampling Filter to be used when interpolating between the sample points.
//
// Usage example:
//
//	result := transform.Resize(img, 800, 600, transform.Linear)
func Resize(img image.Image, width, height int, filter ResampleFilter) *image.RGBA {
	if width <= 0 || height <= 0 || img.Bounds().Empty() {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	src := clone.AsShallowRGBA(img)
	var dst *image.RGBA

	// NearestNeighbor is a special case, it's faster to compute without convolution matrix.
	if filter.Support <= 0 {
		dst = nearestNeighbor(src, width, height)
	} else {
		dst = resampleHorizontal(src, width, filter)
		dst = resampleVertical(dst, height, filter)
	}

	return dst
}

// Crop returns a new image which contains the intersection between the rect and the image provided as params.
// Only the intersection is returned. If a rect larger than the image is provided, no fill is done to
// the 'empty' area.
//
// Usage example:
//
//	result := transform.Crop(img, image.Rect(0, 0, 512, 256))
func Crop(img image.Image, rect image.Rectangle) *image.RGBA {
	src := clone.AsShallowRGBA(img)
	return clone.AsRGBA(src.SubImage(rect))
}

func resampleHorizontal(src *image.RGBA, width int, filter ResampleFilter) *image.RGBA {
	srcWidth, srcHeight := src.Bounds().Dx(), src.Bounds().Dy()
	srcStride := src.Stride

	delta := float64(srcWidth) / float64(width)
	// Scale must be at least 1. Special case for image size reduction filter radius.
	scale := math.Max(delta, 1.0)

	dst := image.NewRGBA(image.Rect(0, 0, width, srcHeight))
	dstStride := dst.Stride

	filterRadius := math.Ceil(scale * filter.Support)

	parallel.Line(srcHeight, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				// value of x from src
				ix := (float64(x)+0.5)*delta - 0.5
				istart, iend := int(ix-filterRadius+0.5), int(ix+filterRadius)

				if istart < 0 {
					istart = 0
				}
				if iend >= srcWidth {
					iend = srcWidth - 1
				}

				var r, g, b, a float64
				var sum float64
				for kx := istart; kx <= iend; kx++ {

					srcPos := y*srcStride + kx*4
					// normalize the sample position to be evaluated by the filter
					normPos := (float64(kx) - ix) / scale
					fValue := filter.Fn(normPos)

					r += float64(src.Pix[srcPos+0]) * fValue
					g += float64(src.Pix[srcPos+1]) * fValue
					b += float64(src.Pix[srcPos+2]) * fValue
					a += float64(src.Pix[srcPos+3]) * fValue
					sum += fValue
				}

				dstPos := y*dstStride + x*4
				dst.Pix[dstPos+0] = uint8(f64.Clamp((r/sum)+0.5, 0, 255))
				dst.Pix[dstPos+1] = uint8(f64.Clamp((g/sum)+0.5, 0, 255))
				dst.Pix[dstPos+2] = uint8(f64.Clamp((b/sum)+0.5, 0, 255))
				dst.Pix[dstPos+3] = uint8(f64.Clamp((a/sum)+0.5, 0, 255))
			}
		}
	})

	return dst
}

func resampleVertical(src *image.RGBA, height int, filter ResampleFilter) *image.RGBA {
	srcWidth, srcHeight := src.Bounds().Dx(), src.Bounds().Dy()
	srcStride := src.Stride

	delta := float64(srcHeight) / float64(height)
	scale := math.Max(delta, 1.0)

	dst := image.NewRGBA(image.Rect(0, 0, srcWidth, height))
	dstStride := dst.Stride

	filterRadius := math.Ceil(scale * filter.Support)

	parallel.Line(height, func(start, end int) {
		for y := start; y < end; y++ {
			iy := (float64(y)+0.5)*delta - 0.5

			istart, iend := int(iy-filterRadius+0.5), int(iy+filterRadius)

			if istart < 0 {
				istart = 0
			}
			if iend >= srcHeight {
				iend = srcHeight - 1
			}

			for x := 0; x < srcWidth; x++ {
				var r, g, b, a float64
				var sum float64
				for ky := istart; ky <= iend; ky++ {

					srcPos := ky*srcStride + x*4
					normPos := (float64(ky) - iy) / scale
					fValue := filter.Fn(normPos)

					r += float64(src.Pix[srcPos+0]) * fValue
					g += float64(src.Pix[srcPos+1]) * fValue
					b += float64(src.Pix[srcPos+2]) * fValue
					a += float64(src.Pix[srcPos+3]) * fValue
					sum += fValue
				}

				dstPos := y*dstStride + x*4
				dst.Pix[dstPos+0] = uint8(f64.Clamp((r/sum)+0.5, 0, 255))
				dst.Pix[dstPos+1] = uint8(f64.Clamp((g/sum)+0.5, 0, 255))
				dst.Pix[dstPos+2] = uint8(f64.Clamp((b/sum)+0.5, 0, 255))
				dst.Pix[dstPos+3] = uint8(f64.Clamp((a/sum)+0.5, 0, 255))
			}
		}
	})

	return dst
}

func nearestNeighbor(src *image.RGBA, width, height int) *image.RGBA {
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	dx := float64(srcW) / float64(width)
	dy := float64(srcH) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*dstStride + x*4
			ipos := int((float64(y)+0.5)*dy)*srcStride + int((float64(x)+0.5)*dx)*4

			dst.Pix[pos+0] = src.Pix[ipos+0]
			dst.Pix[pos+1] = src.Pix[ipos+1]
			dst.Pix[pos+2] = src.Pix[ipos+2]
			dst.Pix[pos+3] = src.Pix[ipos+3]
		}
	}

	return dst
}
