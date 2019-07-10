/*Package effect provides the functionality to manipulate images to achieve various looks.*/
package effect

import (
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/convolution"
	"github.com/anthonynsimon/bild/math/f64"
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

				c := 0.3*float64(src.Pix[srcPos+0]) + 0.6*float64(src.Pix[srcPos+1]) + 0.1*float64(src.Pix[srcPos+2])
				dst.Pix[dstPos] = uint8(c + 0.5)
			}
		}
	})

	return dst
}

// Sepia returns a copy of the image in Sepia tone.
func Sepia(img image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		// Cache values as float64
		var fc [3]float64
		fc[0] = float64(c.R)
		fc[1] = float64(c.G)
		fc[2] = float64(c.B)

		// Calculate out color based on heuristic
		outRed := fc[0]*0.393 + fc[1]*0.769 + fc[2]*0.189
		outGreen := fc[0]*0.349 + fc[1]*0.686 + fc[2]*0.168
		outBlue := fc[0]*0.272 + fc[1]*0.534 + fc[2]*0.131

		// Clamp ceiled values before returning
		return color.RGBA{
			R: uint8(f64.Clamp(outRed+0.5, 0, 255)),
			G: uint8(f64.Clamp(outGreen+0.5, 0, 255)),
			B: uint8(f64.Clamp(outBlue+0.5, 0, 255)),
			A: c.A,
		}
	}

	dst := adjust.Apply(img, fn)

	return dst
}

// EdgeDetection returns a copy of the image with its edges highlighted.
func EdgeDetection(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return image.NewRGBA(src.Bounds())
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
	return convolution.Convolve(src, k, &convolution.Options{Bias: 0, Wrap: false, KeepAlpha: true})
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

	return convolution.Convolve(src, &k, &convolution.Options{Bias: 128, Wrap: false, KeepAlpha: true})
}

// Sharpen returns a sharpened copy of the image by detecting its edges and adding it to the original.
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

	return convolution.Convolve(src, &k, &convolution.Options{Bias: 0, Wrap: false})
}

// UnsharpMask returns a copy of the image with its high-frecuency components amplified.
// Parameter radius corresponds to the radius to be samples per pixel.
// Parameter amount is the normalized strength of the effect. A value of 0.0 will leave
// the image untouched and a value of 1.0 will fully apply the unsharp mask.
func UnsharpMask(img image.Image, radius, amount float64) *image.RGBA {
	amount = f64.Clamp(amount, 0, 10)

	blurred := blur.Gaussian(img, 5*radius) // scale radius by matching factor

	bounds := img.Bounds()
	src := clone.AsRGBA(img)
	dst := image.NewRGBA(bounds)
	w, h := bounds.Dx(), bounds.Dy()

	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				pos := y*dst.Stride + x*4

				r := float64(src.Pix[pos+0])
				g := float64(src.Pix[pos+1])
				b := float64(src.Pix[pos+2])
				a := float64(src.Pix[pos+3])

				rBlur := float64(blurred.Pix[pos+0])
				gBlur := float64(blurred.Pix[pos+1])
				bBlur := float64(blurred.Pix[pos+2])
				aBlur := float64(blurred.Pix[pos+3])

				r = r + (r-rBlur)*amount
				g = g + (g-gBlur)*amount
				b = b + (b-bBlur)*amount
				a = a + (a-aBlur)*amount

				dst.Pix[pos+0] = uint8(f64.Clamp(r, 0, 255))
				dst.Pix[pos+1] = uint8(f64.Clamp(g, 0, 255))
				dst.Pix[pos+2] = uint8(f64.Clamp(b, 0, 255))
				dst.Pix[pos+3] = uint8(f64.Clamp(a, 0, 255))
			}
		}
	})

	return dst
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

	vSobel := convolution.Convolve(src, &vk, &convolution.Options{Bias: 0, Wrap: false, KeepAlpha: true})
	hSobel := convolution.Convolve(src, &hk, &convolution.Options{Bias: 0, Wrap: false, KeepAlpha: true})

	return blend.Add(blend.Multiply(vSobel, vSobel), blend.Multiply(hSobel, hSobel))
}

// Median returns a new image in which each pixel is the median of its neighbors.
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

// spatialFilter goes through each pixel on an image collecting its neighbors and picking one
// based on the function provided. The resulting image is then returned.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
// The parameter pickerFn is the function that receives the list of neighbors and returns the selected
// neighbor to be used for the resulting image.
func spatialFilter(img image.Image, radius float64, pickerFn func(neighbors []color.RGBA) color.RGBA) *image.RGBA {
	if radius <= 0 {
		return clone.AsRGBA(img)
	}

	padding := int(radius + 0.5)
	src := clone.Pad(img, padding, padding, clone.EdgeExtend)

	kernelSize := int(2*radius + 1 + 0.5)

	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	w, h := bounds.Dx(), bounds.Dy()
	neighborsCount := kernelSize * kernelSize

	parallel.Line(h, func(start, end int) {
		for y := start + padding; y < end+padding; y++ {
			for x := padding; x < w+padding; x++ {

				neighbors := make([]color.RGBA, neighborsCount)
				i := 0
				for ky := 0; ky < kernelSize; ky++ {
					for kx := 0; kx < kernelSize; kx++ {
						ix := x - kernelSize>>1 + kx
						iy := y - kernelSize>>1 + ky

						ipos := iy*src.Stride + ix*4
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

				pos := (y-padding)*dst.Stride + (x-padding)*4
				dst.Pix[pos+0] = c.R
				dst.Pix[pos+1] = c.G
				dst.Pix[pos+2] = c.B
				dst.Pix[pos+3] = c.A
			}
		}
	})

	return dst
}
