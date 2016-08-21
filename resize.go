package bild

import (
	"fmt"
	"image"
	"math"
)

// ResampleFilter is used to populate the kernel for resampling images.
// Name is simply an identifier for the filter function.
// Support is the number of pixels to be used from each side.
// For example, a support of 1.0 means that the filter will get pixels on
// "positions" -1 and +1 away from it.
// Fn is the resample filter function itself.
type ResampleFilter struct {
	Name    string
	Support float64
	Fn      func(x float64) float64
}

// NearestNeighbor is a fast, non-convolution resample filter. It produces
// pixelated results as no interpolation is done when resizing.
var NearestNeighbor ResampleFilter

// Box is a convolution based resample filter that interpolates the values by averaging.
var Box ResampleFilter

// Linear is a convolution based resample filter that interpolates the values linearly.
var Linear ResampleFilter

var Gaussian ResampleFilter

func init() {
	NearestNeighbor = ResampleFilter{
		Name:    "NearestNeighbor",
		Support: 0,
		Fn:      nil,
	}
	Box = ResampleFilter{
		Name:    "Box",
		Support: 0.5,
		Fn: func(x float64) float64 {
			if math.Abs(x) < 0.5 {
				return 1
			}
			return 0
		},
	}
	Linear = ResampleFilter{
		Name:    "Linear",
		Support: 1.0,
		Fn: func(x float64) float64 {
			x = math.Abs(x)
			if x < 1.0 {
				return 1.0 - x
			}
			return 0
		},
	}
	Gaussian = ResampleFilter{
		Name:    "Gaussian",
		Support: 2.0,
		Fn: func(x float64) float64 {
			x = math.Abs(x)
			if x < 2.0 {
				return math.Exp(-2 * x * x)
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
		scaleX, scaleY := (float64(srcW) / float64(width)), (float64(srcH) / float64(height))

		// if scaleX < 1.0 {
		// 	scaleX = 1.0
		// }
		// if scaleY < 1.0 {
		// 	scaleY = 1.0
		// }

		// dst = nearestNeighbor(src, width, height)
		dst = sample(src, width, height, scaleX, scaleY)
		dst = resampleHorizontal(dst, scaleX, filter)
		dst = resampleVertical(dst, scaleY, filter)
	}

	return dst
}

// Build the convolution kernel based on the filter selected
func buildKernel(scale float64, filter ResampleFilter) ConvolutionMatrix {
	step := 1.0 / scale
	fmt.Println(step, scale)

	kernelLength := int(step*2*filter.Support + 1 + 0.5)
	kernel := NewKernel(kernelLength, 1)

	if kernelLength == 1 {
		kernel.Matrix[0] = 1
	} else {
		for x := 0; x < kernelLength; x++ {
			u := (float64(x)/float64(kernelLength-1))*filter.Support*2 - filter.Support
			fmt.Println(u)
			kernel.Matrix[x] = filter.Fn(u)
		}
	}

	fmt.Println(kernel)

	return kernel
}

func sample(src *image.RGBA, width, height int, scaleX, scaleY float64) *image.RGBA {
	srcW, srcH := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	dx, dy := int(1/scaleX+0.5), int(1/scaleY+0.5)

	for y := 0; y < srcH; y++ {
		for x := 0; x < srcW; x++ {
			srcPos := y*srcStride + x*4
			dstPos := y*dx*dstStride + x*dy*4
			// dstPos := y*dx*dstStride + dstStride*dy/2 + x*dy*4 + 4*dx/2

			// yPos := (float64(y) / float64(srcH-1))
			// xPos := (float64(x) / float64(srcW-1))
			// dstPosY := int((yPos*float64(height-1) + 0.5)) * dstStride
			// dstPosX := int((xPos*float64(width-1) + 0.5)) * 4
			// dstPos := dstPosX + dstPosY

			dst.Pix[dstPos+0] = src.Pix[srcPos+0]
			dst.Pix[dstPos+1] = src.Pix[srcPos+1]
			dst.Pix[dstPos+2] = src.Pix[srcPos+2]
			dst.Pix[dstPos+3] = src.Pix[srcPos+3]
		}
	}

	return dst
}

func resampleHorizontal(src *image.RGBA, scale float64, filter ResampleFilter) *image.RGBA {
	width, height := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	k := buildKernel(scale, filter)
	kernelLength := k.MaxX()

	parallelize(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {

				var r, g, b, a float64
				for kx := 0; kx < kernelLength; kx++ {
					ix := (x - (kernelLength / 2) + kx)

					if ix < 0 {
						ix = 0
					} else if ix >= width {
						ix = width - 1
					}
					// if ix < 0 || ix >= width {
					// 	continue
					// }

					ipos := y*srcStride + ix*4
					kvalue := k.At(kx, 0)
					// fmt.Println(ipos, kvalue)

					r += float64(src.Pix[ipos+0]) * kvalue
					g += float64(src.Pix[ipos+1]) * kvalue
					b += float64(src.Pix[ipos+2]) * kvalue
					a += float64(src.Pix[ipos+3]) * kvalue
				}

				pos := y*dstStride + x*4
				dst.Pix[pos+0] = uint8(math.Max(math.Min(r+0.5, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g+0.5, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b+0.5, 255), 0))
				dst.Pix[pos+3] = uint8(math.Max(math.Min(a+0.5, 255), 0))
			}
		}
	})

	return dst
}

func resampleVertical(src *image.RGBA, scale float64, filter ResampleFilter) *image.RGBA {
	width, height := src.Bounds().Max.X, src.Bounds().Max.Y
	srcStride := src.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	k := buildKernel(scale, filter)
	kernelLength := k.MaxX()

	parallelize(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {

				var r, g, b, a float64
				for ky := 0; ky < kernelLength; ky++ {
					iy := y - (kernelLength / 2) + ky

					if iy < 0 {
						iy = 0
					} else if iy >= height {
						iy = height - 1
					}

					// if iy < 0 || iy >= height {
					// 	continue
					// }

					ipos := iy*srcStride + x*4
					kvalue := k.At(ky, 0)

					r += float64(src.Pix[ipos+0]) * kvalue
					g += float64(src.Pix[ipos+1]) * kvalue
					b += float64(src.Pix[ipos+2]) * kvalue
					a += float64(src.Pix[ipos+3]) * kvalue
				}

				pos := y*dstStride + x*4
				dst.Pix[pos+0] = uint8(math.Max(math.Min(r+0.5, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g+0.5, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b+0.5, 255), 0))
				dst.Pix[pos+3] = uint8(math.Max(math.Min(a+0.5, 255), 0))
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
