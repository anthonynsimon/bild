package bild

import (
	"image"
	"math"
)

// ResampleFilter is used to evaluate sample points and interpolate between them.
// Support is the number of points required by the filter per 'side'.
// For example, a support of 1.0 means that the filter will get pixels on
// positions -1 and +1 away from it.
// Fn is the resample filter function to evaluate the samples.
type ResampleFilter struct {
	Support float64
	Fn      func(x float64) float64
}

// NearestNeighbor resampling filter assigns to each point the sample point nearest to it.
var NearestNeighbor ResampleFilter

// Box resampling filter, only let pass values in the x < 0.5 range from sample.
// It produces similar results to the Nearest Neighbor method.
var Box ResampleFilter

// Linear resampling filter interpolates linearly between the two nearest samples per dimension.
var Linear ResampleFilter

func init() {
	NearestNeighbor = ResampleFilter{
		Support: 0,
		Fn:      nil,
	}
	Box = ResampleFilter{
		Support: 0.5,
		Fn: func(x float64) float64 {
			if math.Abs(x) < 0.5 {
				return 1
			}
			return 0
		},
	}
	Linear = ResampleFilter{
		Support: 1.0,
		Fn: func(x float64) float64 {
			x = math.Abs(x)
			if x < 1.0 {
				return 1.0 - x
			}
			return 0
		},
	}
}

// Resize returns a new image with its size adjusted to the new width and height. The filter
// param corresponds to the Resampling Filter to be used when interpolating between the sample points.
//
//
// Usage example:
//
//		result := Resize(img, 800, 600, bild.Linear)
//
func Resize(img image.Image, width, height int, filter ResampleFilter) *image.RGBA {
	if width <= 0 || height <= 0 {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	src := CloneAsRGBA(img)
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

func resampleHorizontal(src *image.RGBA, width int, filter ResampleFilter) *image.RGBA {
	srcWidth, srcHeight := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	delta := float64(srcWidth) / float64(width)
	// Scale must be at least 1. Special case for image size reduction filter radius.
	scale := math.Max(delta, 1.0)

	dst := image.NewRGBA(image.Rect(0, 0, width, srcHeight))
	dstStride := dst.Stride

	filterRadius := math.Ceil(scale * filter.Support)

	parallelize(srcHeight, func(start, end int) {
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
				dst.Pix[dstPos+0] = uint8(clampFloat64((r/sum)+0.5, 0, 255))
				dst.Pix[dstPos+1] = uint8(clampFloat64((g/sum)+0.5, 0, 255))
				dst.Pix[dstPos+2] = uint8(clampFloat64((b/sum)+0.5, 0, 255))
				dst.Pix[dstPos+3] = uint8(clampFloat64((a/sum)+0.5, 0, 255))
			}
		}
	})

	return dst
}

func resampleVertical(src *image.RGBA, height int, filter ResampleFilter) *image.RGBA {
	srcWidth, srcHeight := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	delta := float64(srcHeight) / float64(height)
	scale := math.Max(delta, 1.0)

	dst := image.NewRGBA(image.Rect(0, 0, srcWidth, height))
	dstStride := dst.Stride

	filterRadius := math.Ceil(scale * filter.Support)

	parallelize(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcWidth; x++ {

				iy := (float64(y)+0.5)*delta - 0.5

				istart, iend := int(iy-filterRadius+0.5), int(iy+filterRadius)

				if istart < 0 {
					istart = 0
				}
				if iend >= srcHeight {
					iend = srcHeight - 1
				}

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
				dst.Pix[dstPos+0] = uint8(clampFloat64((r/sum)+0.5, 0, 255))
				dst.Pix[dstPos+1] = uint8(clampFloat64((g/sum)+0.5, 0, 255))
				dst.Pix[dstPos+2] = uint8(clampFloat64((b/sum)+0.5, 0, 255))
				dst.Pix[dstPos+3] = uint8(clampFloat64((a/sum)+0.5, 0, 255))
			}
		}
	})

	return dst
}

func nearestNeighbor(src *image.RGBA, width, height int) *image.RGBA {
	srcW, srcH := src.Bounds().Max.X, src.Bounds().Max.Y
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
