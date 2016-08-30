/*Package effect provides the functionality to manipulate images to achieve various looks.*/
package effect

import (
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/convolution"
	"github.com/anthonynsimon/bild/parallel"
	"github.com/anthonynsimon/bild/util"
)

// Invert returns a negated version of the image.
func Invert(src image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{255 - c.R, 255 - c.G, 255 - c.B, c.A}
	}

	img := adjust.Apply(src, fn)

	return img
}

// Grayscale returns a copy of the image in Grayscale using the weights
// 0.3R + 0.6G + 0.1B as a heuristic.
func Grayscale(img image.Image) *image.Gray {
	src := clone.AsRGBA(img)
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()

	if bounds.Empty() {
		return &image.Gray{}
	}

	dst := image.NewGray(bounds)

	parallel.Line(srcH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcW; x++ {
				srcPos := y*src.Stride + x*4
				dstPos := y*dst.Stride + x

				var c uint8
				c += uint8(0.3*float64(src.Pix[srcPos+0]) + 0.5)
				c += uint8(0.6*float64(src.Pix[srcPos+1]) + 0.5)
				c += uint8(0.1*float64(src.Pix[srcPos+2]) + 0.5)

				dst.Pix[dstPos] = c
			}
		}
	})

	return dst
}

// EdgeDetection returns a copy of the image with it's edges highlighted.
func EdgeDetection(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return clone.AsRGBA(src)
	}

	length := int(math.Ceil(2*radius + 1))
	k := convolution.NewKernel(length, length)

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			v := -1.0
			if x == length/2 && y == length/2 {
				v = float64(length*length) - 1
			}
			k.Matrix[y*length+x] = v

		}
	}
	return convolution.Convolve(src, k, &convolution.Options{Bias: 0, Wrap: false, CarryAlpha: true})
}

// Emboss returns a copy of the image in which each pixel has been
// replaced either by a highlight or a shadow representation.
func Emboss(src image.Image) *image.RGBA {
	k := convolution.Kernel{
		Matrix: []float64{
			-1, -1, 0,
			-1, 0, 1,
			0, 1, 1,
		},
		Width:  3,
		Height: 3,
	}

	return convolution.Convolve(src, &k, &convolution.Options{Bias: 128, Wrap: false, CarryAlpha: true})
}

// Sharpen returns a sharpened copy of the image by detecting it's edges and adding it to the original.
func Sharpen(src image.Image) *image.RGBA {
	k := convolution.Kernel{
		Matrix: []float64{
			0, -1, 0,
			-1, 5, -1,
			0, -1, 0,
		},
		Width:  3,
		Height: 3,
	}

	return convolution.Convolve(src, &k, &convolution.Options{Bias: 0, Wrap: false, CarryAlpha: true})
}

// Sobel returns an image emphasising edges using an approximation to the Sobel–Feldman operator.
func Sobel(src image.Image) *image.RGBA {

	hk := convolution.Kernel{
		Matrix: []float64{
			1, 2, 1,
			0, 0, 0,
			-1, -2, -1,
		},
		Width:  3,
		Height: 3,
	}

	vk := convolution.Kernel{
		Matrix: []float64{
			-1, 0, 1,
			-2, 0, 2,
			-1, 0, 1,
		},
		Width:  3,
		Height: 3,
	}

	vSobel := convolution.Convolve(src, &vk, &convolution.Options{Bias: 0, Wrap: false, CarryAlpha: true})
	hSobel := convolution.Convolve(src, &hk, &convolution.Options{Bias: 0, Wrap: false, CarryAlpha: true})

	return blend.Add(blend.Multiply(vSobel, vSobel), blend.Multiply(hSobel, hSobel))
}

// Median returns a new image in which each pixel is the median of it's neighbors.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
func Median(img image.Image, radius float64) *image.RGBA {
	fn := func(neighbors []color.RGBA) color.RGBA {
		util.SortRGBA(neighbors, 0, len(neighbors)-1)
		return neighbors[len(neighbors)/2]
	}

	result := spatialFilter(img, radius, fn)

	return result
}

// Dilate picks the local maxima from the neighbors of each pixel and returns the resulting image.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
func Dilate(img image.Image, radius float64) *image.RGBA {
	fn := func(neighbors []color.RGBA) color.RGBA {
		util.SortRGBA(neighbors, 0, len(neighbors)-1)
		return neighbors[len(neighbors)-1]
	}

	result := spatialFilter(img, radius, fn)

	return result
}

// Erode picks the local minima from the neighbors of each pixel and returns the resulting image.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
func Erode(img image.Image, radius float64) *image.RGBA {
	fn := func(neighbors []color.RGBA) color.RGBA {
		util.SortRGBA(neighbors, 0, len(neighbors)-1)
		return neighbors[0]
	}

	result := spatialFilter(img, radius, fn)

	return result
}

// spatialFilter goes through each pixel on an image collecting it's neighbors and picking one
// based on the function provided. The resulting image is then returned.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
// The parameter pickerFn is the function that receives the list of neighbors and returns the selected
// neighbor to be used for the resulting image.
func spatialFilter(img image.Image, radius float64, pickerFn func(neighbors []color.RGBA) color.RGBA) *image.RGBA {
	bounds := img.Bounds()
	src := clone.AsRGBA(img)

	if radius <= 0 {
		return src
	}

	kernelSize := int(2*radius + 1 + 0.5)

	dst := image.NewRGBA(bounds)

	w, h := bounds.Dx(), bounds.Dy()
	neighborsCount := kernelSize * kernelSize

	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {

				neighbors := make([]color.RGBA, neighborsCount)
				i := 0
				for ky := 0; ky < kernelSize; ky++ {
					for kx := 0; kx < kernelSize; kx++ {
						ix := x - kernelSize/2 + kx
						iy := y - kernelSize/2 + ky

						if ix < 0 {
							ix = 0
						} else if ix >= w {
							ix = w - 1
						}

						if iy < 0 {
							iy = 0
						} else if iy >= h {
							iy = h - 1
						}

						ipos := iy*dst.Stride + ix*4
						neighbors[i] = color.RGBA{
							R: src.Pix[ipos+0],
							G: src.Pix[ipos+1],
							B: src.Pix[ipos+2],
							A: src.Pix[ipos+3],
						}
						i++
					}
				}

				c := pickerFn(neighbors)

				pos := y*dst.Stride + x*4
				dst.Pix[pos+0] = c.R
				dst.Pix[pos+1] = c.G
				dst.Pix[pos+2] = c.B
				dst.Pix[pos+3] = c.A
			}
		}
	})

	return dst
}
