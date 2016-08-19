package bild

import (
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
		Degree: 0.5,
		Fn: func(x, y float64) float64 {
			if math.Abs(x) < 0.5 {
				return 1
			}
			return 0
		},
	}
	Linear = ResampleFilter{
		Name:   "Linear",
		Degree: 1,
		Fn: func(x, y float64) float64 {
			x = math.Abs(x)
			if x < 1.0 {
				return 1.0 - x
			}
			return 0
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

	// NearestNeighbor is a special case, it's faster to compute without convolution matrix.
	if filter.Name == NearestNeighbor.Name {
		dst = nearestNeighbor(src, width, height)
	} else {
		srcW, srcH := src.Bounds().Max.X, src.Bounds().Max.Y
		scaleX, scaleY := 1<<16/((srcW<<16)/width), 1<<16/((srcH<<16)/height)

		dst = sample(src, width, height, scaleX, scaleY)
		dst = resampleHorizontal(dst, float64(scaleX), filter)
		dst = resampleVertical(dst, float64(scaleY), filter)

		crop := dst.SubImage(image.Rect(scaleX/2, scaleY/2, width+scaleX/2, height+scaleY/2))
		dst = CloneAsRGBA(crop)
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

func sample(src *image.RGBA, width, height, scaleX, scaleY int) *image.RGBA {
	srcW, srcH := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width+scaleX+1, height+scaleY+1))
	dstStride := dst.Stride

	for y := 0; y < srcH; y++ {
		for x := 0; x < srcW; x++ {
			srcPos := y*srcStride + x*4
			dstPos := y*scaleY*dstStride + scaleY*dstStride + x*scaleX*4 + scaleX*4

			dst.Pix[dstPos+0] = src.Pix[srcPos+0]
			dst.Pix[dstPos+1] = src.Pix[srcPos+1]
			dst.Pix[dstPos+2] = src.Pix[srcPos+2]
			dst.Pix[dstPos+3] = src.Pix[srcPos+3]
		}
	}

	return dst
}

// Build the convolution kernel based on the filter selected
func buildKernel(radius float64, filter ResampleFilter) ConvolutionMatrix {
	kernelLength := int(math.Ceil(2*radius*filter.Degree + 1))
	kernel := NewKernel(kernelLength, 1)

	// fmt.Println(filter.Name, "kernelLength:", kernelLength)
	for x := 0; x < kernelLength; x++ {
		kernel.Matrix[x] = filter.Fn(float64(x-kernelLength/2)/radius, 0)
	}

	return kernel
}

func resampleHorizontal(src *image.RGBA, radius float64, filter ResampleFilter) *image.RGBA {
	width, height := src.Bounds().Max.X, src.Bounds().Max.Y
	stride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	k := buildKernel(radius, filter)
	kernelLength := k.MaxX()
	// fmt.Println(k)

	parallelize(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {

				var r, g, b, a float64
				for kx := 0; kx < kernelLength; kx++ {
					ix := x - kernelLength/2 + kx

					if ix < 0 || ix >= width {
						continue
					}

					ipos := y*stride + ix*4
					kvalue := k.At(kx, 0)

					r += float64(src.Pix[ipos+0]) * kvalue
					g += float64(src.Pix[ipos+1]) * kvalue
					b += float64(src.Pix[ipos+2]) * kvalue
					a += float64(src.Pix[ipos+3]) * kvalue
				}

				pos := y*stride + x*4
				dst.Pix[pos+0] = uint8(math.Max(math.Min(r, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b, 255), 0))
				dst.Pix[pos+3] = uint8(math.Max(math.Min(a, 255), 0))
			}
		}
	})

	return dst
}

func resampleVertical(src *image.RGBA, radius float64, filter ResampleFilter) *image.RGBA {
	width, height := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	k := buildKernel(radius, filter)
	kernelLength := k.MaxX()

	parallelize(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {

				var r, g, b, a float64
				for ky := 0; ky < kernelLength; ky++ {
					iy := y - kernelLength/2 + ky

					if iy < 0 || iy >= height {
						continue
					}

					ipos := iy*srcStride + x*4
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
