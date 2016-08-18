package bild

import (
	"fmt"
	"image"
	"math"
)

// ResampleFilter is used to populate the kernel for resampling images.
// Name is simply an identifier for the filter function.
// Fn is the resample filter function itself.
type ResampleFilter struct {
	Name   string
	Degree float64
	Fn     func(x, y float64) float64
}

// NearestNeighbor is a fast, non-convolution resample filter. It produces
// pixelated results as no interpolation is done when resizing.
var NearestNeighbor ResampleFilter

// Box is a convolution based resample filter that interpolates the values by averaging.
var Box ResampleFilter

// Linear is a convolution based resample filter that interpolates the values linearly.
var Linear ResampleFilter

var Quadratic ResampleFilter

// Gaussian is a convolution based resample filter that interpolates the values
// using a gaussian function.
var Gaussian ResampleFilter

func init() {
	NearestNeighbor = ResampleFilter{
		Name:   "NearestNeighbor",
		Degree: 0,
		Fn:     nil,
	}
	Box = ResampleFilter{
		Name:   "Box",
		Degree: 1,
		Fn: func(x, y float64) float64 {
			if math.Abs(x) < 0.5 {
				return 1
			}
			return 0
		},
	}
	Linear = ResampleFilter{
		Name:   "Linear",
		Degree: 2,
		Fn: func(x, y float64) float64 {
			x = math.Abs(x)
			if x < 1.0 {
				return 1.0 - x
			}
			return 0
		},
	}
	Quadratic = ResampleFilter{
		Name:   "Quadratic",
		Degree: 3,
		Fn: func(x, y float64) float64 {
			x = math.Abs(x)
			if -1 < x && x <= 0 {
				return (x * x * 0.5) + (3 * x / 2) + 1
			} else if 0 < x && x <= 1 {
				return -(x * x) + 1
			} else if 1 < x && x <= 2 {
				return (x * x * 0.5) - (3 * x / 2) + 1
			}
			return 0
		},
	}
	Gaussian = ResampleFilter{
		Name: "Gaussian",
		Fn: func(x, y float64) float64 {
			x, y = 0.5-x, 0.5-y
			return math.Exp(-(x*x/1.0 + y*y/1.0))
		},
	}
}

// Resize returns a new image with its size adjusted to the new width and height. The filter
// param corresponds to the Resample Filter to be used when interpolating between the pixels.
//
// This package includes the following filters: NearestNeighbor, Box, Linear and Gaussian.
//
// Usage example:
//
//		result := Resize(img, 800, 600, bild.NearestNeighbor)
//
func Resize(img image.Image, width, height int, filter ResampleFilter) *image.RGBA {
	if width <= 0 || height <= 0 {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	src := CloneAsRGBA(img)

	var dst *image.RGBA
	if filter.Name == NearestNeighbor.Name {
		// NearestNeighbor is a special case, it's faster to compute without convolution matrix.
		dst = nearestNeighbor(src, width, height)
	} else {
		dst = resampleHorizontal(src, width, filter)
		dst = resampleVertical(dst, height, filter)
	}

	return dst
}

func nearestNeighbor(src *image.RGBA, width, height int) *image.RGBA {
	srcW, srcH := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	scaleX, scaleY := (srcW<<16)/width, (srcH<<16)/height

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*dstStride + x*4
			ipos := ((y*scaleY)>>16)*srcStride + ((x*scaleX)>>16)*4

			dst.Pix[pos+0] = src.Pix[ipos+0]
			dst.Pix[pos+1] = src.Pix[ipos+1]
			dst.Pix[pos+2] = src.Pix[ipos+2]
			dst.Pix[pos+3] = src.Pix[ipos+3]
		}
	}

	return dst
}

// Build the convolution kernel based on the filter selected
func buildKernel(radius float64, filter ResampleFilter) ConvolutionMatrix {
	kernelLength := int(math.Ceil(2*150 + 1))
	kernel := NewKernel(kernelLength, 1)

	fmt.Println(filter.Name)
	for x := 0; x < kernelLength; x++ {
		kernel.Matrix[x] = filter.Fn(float64(x-kernelLength/2)/300, 0)
		fmt.Println(float64(x-kernelLength/2) / float64(2))
	}
	fmt.Println("")

	return kernel.Normalized()
}

func resampleHorizontal(src *image.RGBA, width int, filter ResampleFilter) *image.RGBA {
	srcWidth, srcHeight := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	scaleX := (srcWidth << 16) / width

	dst := image.NewRGBA(image.Rect(0, 0, width, srcHeight))
	dstStride := dst.Stride

	// radius := float64(1<<16) / float64(scaleX)

	k := buildKernel(2, filter)
	kernelLength := k.MaxX()

	fmt.Println(k)

	parallelize(srcHeight, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {

				var r, g, b, a float64
				for kx := 0; kx < kernelLength; kx++ {
					ix := x - kernelLength/2 + kx

					if ix < 0 || ix >= width {
						continue
					}

					ipos := y*srcStride + ((ix*scaleX)>>16)*4
					kvalue := k.At(kx, 0)

					r += float64(src.Pix[ipos+0]) * kvalue
					g += float64(src.Pix[ipos+1]) * kvalue
					b += float64(src.Pix[ipos+2]) * kvalue
					a += float64(src.Pix[ipos+3]) * kvalue
				}

				pos := y*dstStride + x*4
				dst.Pix[pos+0] = uint8(math.Max(math.Min(r, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b, 255), 0))
				dst.Pix[pos+3] = uint8(math.Max(math.Min(a, 255), 0))
			}
		}
	})

	return dst
}

func resampleVertical(src *image.RGBA, height int, filter ResampleFilter) *image.RGBA {
	srcWidth, srcHeight := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	scaleY := (srcHeight << 16) / height

	dst := image.NewRGBA(image.Rect(0, 0, srcWidth, height))
	dstStride := dst.Stride

	// radius := float64(1<<16) / float64(scaleY)
	k := buildKernel(2, filter)
	kernelLength := k.MaxX()

	parallelize(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcWidth; x++ {

				var r, g, b, a float64
				for ky := 0; ky < kernelLength; ky++ {
					iy := y - kernelLength/2 + ky

					if iy < 0 || iy >= height {
						continue
					}

					ipos := ((y*scaleY)>>16)*srcStride + x*4
					kvalue := k.At(ky, 0)

					r += float64(src.Pix[ipos+0]) * kvalue
					g += float64(src.Pix[ipos+1]) * kvalue
					b += float64(src.Pix[ipos+2]) * kvalue
					a += float64(src.Pix[ipos+3]) * kvalue
				}

				pos := y*dstStride + x*4
				dst.Pix[pos+0] = uint8(math.Max(math.Min(r, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b, 255), 0))
				dst.Pix[pos+3] = uint8(math.Max(math.Min(a, 255), 0))
			}
		}
	})

	return dst
}
