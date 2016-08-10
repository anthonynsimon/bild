# bild

![bild logo](https://anthonynsimon.github.io/projects/bild/logo.png)  

[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/bild/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/anthonynsimon/bild?status.svg)](https://godoc.org/github.com/anthonynsimon/bild)
[![Build Status](https://travis-ci.org/anthonynsimon/bild.svg?branch=master)](https://travis-ci.org/anthonynsimon/bild)
[![Go Report Card](https://goreportcard.com/badge/github.com/anthonynsimon/bild)](https://goreportcard.com/report/github.com/anthonynsimon/bild)

    import "github.com/anthonynsimon/bild"

Simple image processing in Go with parallel processing support.

Package bild provides a collection of common image processing functions. The input images must implement the image.Image interface and the functions return an *image.RGBA.

The aim of this project is simplicity in use and development over high performance, but most algorithms are designed to be efficient and make use of parallelism when available. It is based on standard Go packages to reduce dependecy use and development abstractions.

Basic example:
```go
package main

import "github.com/anthonynsimon/bild"

func main() {
	img, err := bild.Open("filename")
	if err != nil {
		panic(err)
	}

	result := bild.Invert(img)

	if err := bild.Save("filename", result, bild.PNG); err != nil {
		panic(err)
	}
}

```


## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Add(bg image.Image, fg image.Image) *image.RGBA](#Add)
* [func BoxBlur(src image.Image, radius float64) *image.RGBA](#BoxBlur)
* [func Brightness(src image.Image, change float64) *image.RGBA](#Brightness)
* [func CloneAsRGBA(src image.Image) *image.RGBA](#CloneAsRGBA)
* [func ColorBurn(bg image.Image, fg image.Image) *image.RGBA](#ColorBurn)
* [func ColorDodge(bg image.Image, fg image.Image) *image.RGBA](#ColorDodge)
* [func Convolute(img image.Image, k ConvolutionMatrix, o *ConvolutionOptions) *image.RGBA](#Convolute)
* [func Darken(bg image.Image, fg image.Image) *image.RGBA](#Darken)
* [func Difference(bg image.Image, fg image.Image) *image.RGBA](#Difference)
* [func Divide(bg image.Image, fg image.Image) *image.RGBA](#Divide)
* [func EdgeDetection(src image.Image, radius float64) *image.RGBA](#EdgeDetection)
* [func Emboss(src image.Image) *image.RGBA](#Emboss)
* [func Encode(w io.Writer, img image.Image, format Format) error](#Encode)
* [func Exclusion(bg image.Image, fg image.Image) *image.RGBA](#Exclusion)
* [func FlipH(img image.Image) *image.RGBA](#FlipH)
* [func FlipV(img image.Image) *image.RGBA](#FlipV)
* [func GaussianBlur(src image.Image, radius float64) *image.RGBA](#GaussianBlur)
* [func Grayscale(src image.Image) *image.RGBA](#Grayscale)
* [func Invert(src image.Image) *image.RGBA](#Invert)
* [func Lighten(bg image.Image, fg image.Image) *image.RGBA](#Lighten)
* [func LinearBurn(bg image.Image, fg image.Image) *image.RGBA](#LinearBurn)
* [func LinearLight(bg image.Image, fg image.Image) *image.RGBA](#LinearLight)
* [func Median(img image.Image, size int) *image.RGBA](#Median)
* [func Multiply(bg image.Image, fg image.Image) *image.RGBA](#Multiply)
* [func Opacity(bg image.Image, fg image.Image, percent float64) *image.RGBA](#Opacity)
* [func Open(filename string) (image.Image, error)](#Open)
* [func Overlay(bg image.Image, fg image.Image) *image.RGBA](#Overlay)
* [func Save(filename string, img image.Image, format Format) error](#Save)
* [func Screen(bg image.Image, fg image.Image) *image.RGBA](#Screen)
* [func Sharpen(src image.Image) *image.RGBA](#Sharpen)
* [func Sobel(src image.Image) *image.RGBA](#Sobel)
* [func SoftLight(bg image.Image, fg image.Image) *image.RGBA](#SoftLight)
* [func Subtract(bg image.Image, fg image.Image) *image.RGBA](#Subtract)
* [type ConvolutionMatrix](#ConvolutionMatrix)
* [type ConvolutionOptions](#ConvolutionOptions)
* [type Format](#Format)
* [type Kernel](#Kernel)
  * [func NewKernel(length int) *Kernel](#NewKernel)
  * [func (k *Kernel) At(x, y int) float64](#Kernel.At)
  * [func (k *Kernel) Normalized() ConvolutionMatrix](#Kernel.Normalized)
  * [func (k *Kernel) SideLength() int](#Kernel.SideLength)
  * [func (k *Kernel) String() string](#Kernel.String)
* [type RGBAF64](#RGBAF64)
  * [func NewRGBAF64(r, g, b, a uint8) RGBAF64](#NewRGBAF64)
  * [func (c *RGBAF64) Clamp()](#RGBAF64.Clamp)

## <a name="pkg-constants">Constants</a>
``` go
const (
    JPEG = iota
    PNG
)
```
Supported image encoding types




## <a name="Add">func</a> Add
``` go
func Add(bg image.Image, fg image.Image) *image.RGBA
```
Add combines the foreground and background images by adding their values and
returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/add.jpg)  


## <a name="BoxBlur">func</a> BoxBlur
``` go
func BoxBlur(src image.Image, radius float64) *image.RGBA
```
BoxBlur returns a blurred (average) version of the image.
Radius must be larger than 0.

![example](https://anthonynsimon.github.io/projects/bild/boxblur.jpg)  


## <a name="Brightness">func</a> Brightness
``` go
func Brightness(src image.Image, change float64) *image.RGBA
```
Brightness returns a copy of the image with the adjusted brightness.
Change is the normalized amount of change to be applied (range -1.0 to 1.0).

![example](https://anthonynsimon.github.io/projects/bild/brightness.jpg)  


## <a name="CloneAsRGBA">func</a> CloneAsRGBA
``` go
func CloneAsRGBA(src image.Image) *image.RGBA
```
CloneAsRGBA returns an RGBA copy of the supplied image.


## <a name="ColorBurn">func</a> ColorBurn
``` go
func ColorBurn(bg image.Image, fg image.Image) *image.RGBA
```
ColorBurn combines the foreground and background images by dividing the inverted
background by the foreground image and then inverting the result which is then returned.

![example](https://anthonynsimon.github.io/projects/bild/colorburn.jpg)  


## <a name="ColorDodge">func</a> ColorDodge
``` go
func ColorDodge(bg image.Image, fg image.Image) *image.RGBA
```
ColorDodge combines the foreground and background images by dividing background by the
inverted foreground image and returns the result.

![example](https://anthonynsimon.github.io/projects/bild/colordodge.jpg)  


## <a name="Convolute">func</a> Convolute
``` go
func Convolute(img image.Image, k ConvolutionMatrix, o *ConvolutionOptions) *image.RGBA
```
Convolute applies a convolution matrix (kernel) to an image with the supplied options.



## <a name="Darken">func</a> Darken
``` go
func Darken(bg image.Image, fg image.Image) *image.RGBA
```
Darken combines the foreground and background images by picking the darkest value per channel
for each pixel. The result is then returned.

![example](https://anthonynsimon.github.io/projects/bild/darken.jpg)  


## <a name="Difference">func</a> Difference
``` go
func Difference(bg image.Image, fg image.Image) *image.RGBA
```
Difference calculates the absolute difference between the foreground and background images and
returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/difference.jpg)  


## <a name="Divide">func</a> Divide
``` go
func Divide(bg image.Image, fg image.Image) *image.RGBA
```
Divide combines the foreground and background images by diving the values from the background
by the foreground and returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/divide.jpg)  


## <a name="EdgeDetection">func</a> EdgeDetection
``` go
func EdgeDetection(src image.Image, radius float64) *image.RGBA
```
EdgeDetection returns a copy of the image with it's edges highlighted.

![example](https://anthonynsimon.github.io/projects/bild/edgedetection.jpg)  


## <a name="Emboss">func</a> Emboss
``` go
func Emboss(src image.Image) *image.RGBA
```
Emboss returns a copy of the image in which each pixel has been
replaced either by a highlight or a shadow representation.

![example](https://anthonynsimon.github.io/projects/bild/emboss.jpg)  


## <a name="Encode">func</a> Encode
``` go
func Encode(w io.Writer, img image.Image, format Format) error
```
Encode writes an image in the specified format.

Usage example:


	// Encode an image to a writer in PNG format,
	// returns an error if something went wrong
	err := Encode(outFile, img, bild.PNG)



## <a name="Exclusion">func</a> Exclusion
``` go
func Exclusion(bg image.Image, fg image.Image) *image.RGBA
```
Exclusion combines the foreground and background images applying the Exclusion blend mode and
returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/exclusion.jpg)  


## <a name="FlipH">func</a> FlipH
``` go
func FlipH(img image.Image) *image.RGBA
```
FlipH returns a horizontally flipped version of the image.

![example](https://anthonynsimon.github.io/projects/bild/fliph.jpg)  


## <a name="FlipV">func</a> FlipV
``` go
func FlipV(img image.Image) *image.RGBA
```
FlipV returns a vertically flipped version of the image.

![example](https://anthonynsimon.github.io/projects/bild/flipv.jpg)  


## <a name="GaussianBlur">func</a> GaussianBlur
``` go
func GaussianBlur(src image.Image, radius float64) *image.RGBA
```
GaussianBlur returns a smoothly blurred version of the image using
a Gaussian function. Radius must be larger than 0.

![example](https://anthonynsimon.github.io/projects/bild/gaussianblur.jpg)  


## <a name="Grayscale">func</a> Grayscale
``` go
func Grayscale(src image.Image) *image.RGBA
```
Grayscale returns a copy of the image in Grayscale using the weights
0.3R + 0.6G + 0.1B as a heuristic.

![example](https://anthonynsimon.github.io/projects/bild/grayscale.jpg)  


## <a name="Invert">func</a> Invert
``` go
func Invert(src image.Image) *image.RGBA
```
Invert returns a negated version of the image.

![example](https://anthonynsimon.github.io/projects/bild/invert.jpg)  


## <a name="Lighten">func</a> Lighten
``` go
func Lighten(bg image.Image, fg image.Image) *image.RGBA
```
Lighten combines the foreground and background images by picking the brightest value per channel
for each pixel. The result is then returned.

![example](https://anthonynsimon.github.io/projects/bild/lighten.jpg)  


## <a name="LinearBurn">func</a> LinearBurn
``` go
func LinearBurn(bg image.Image, fg image.Image) *image.RGBA
```
LinearBurn combines the foreground and background images by adding them and
then subtracting 255 (1.0 in normalized scale). The resulting image is then returned.

![example](https://anthonynsimon.github.io/projects/bild/linearburn.jpg)  


## <a name="LinearLight">func</a> LinearLight
``` go
func LinearLight(bg image.Image, fg image.Image) *image.RGBA
```
LinearLight combines the foreground and background images by a mix of a Linear Dodge and
Linear Burn operation. The resulting image is then returned.

![example](https://anthonynsimon.github.io/projects/bild/linearlight.jpg)  


## <a name="Median">func</a> Median
``` go
func Median(img image.Image, size int) *image.RGBA
```
Median returns a new image in which each pixel is the median of it's neighbors.
Size sets the amount of neighbors to be searched.

![example](https://anthonynsimon.github.io/projects/bild/median.jpg)  


## <a name="Multiply">func</a> Multiply
``` go
func Multiply(bg image.Image, fg image.Image) *image.RGBA
```
Multiply combines the foreground and background images by multiplying their
normalized values and returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/multiply.jpg)  


## <a name="Opacity">func</a> Opacity
``` go
func Opacity(bg image.Image, fg image.Image, percent float64) *image.RGBA
```
Opacity returns an image which blends the two input images by the percentage provided.
Percent must be of range 0 <= percent <= 1.0

![example](https://anthonynsimon.github.io/projects/bild/opacity.jpg)  


## <a name="Open">func</a> Open
``` go
func Open(filename string) (image.Image, error)
```
Open loads and decodes an image from a file and returns it.

Usage example:


	// Encode an image to a writer in PNG format,
	// returns an error if something went wrong
	img, err := Open("exampleName")



## <a name="Overlay">func</a> Overlay
``` go
func Overlay(bg image.Image, fg image.Image) *image.RGBA
```
Overlay combines the foreground and background images by using Multiply when channel values < 0.5
or using Screen otherwise and returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/overlay.jpg)  


## <a name="Save">func</a> Save
``` go
func Save(filename string, img image.Image, format Format) error
```
Save creates a file and writes to it an image in the specified format

Usage example:


	// Save an image to a file in PNG format,
	// returns an error if something went wrong
	err := Save("exampleName", img, bild.PNG)



## <a name="Screen">func</a> Screen
``` go
func Screen(bg image.Image, fg image.Image) *image.RGBA
```
Screen combines the foreground and background images by inverting, multiplying and inverting the output.
The result is a brighter image which is then returned.

![example](https://anthonynsimon.github.io/projects/bild/screen.jpg)  


## <a name="Sharpen">func</a> Sharpen
``` go
func Sharpen(src image.Image) *image.RGBA
```
Sharpen returns a sharpened copy of the image by detecting it's edges and adding it to the original.

![example](https://anthonynsimon.github.io/projects/bild/sharpen.jpg)  


## <a name="Sobel">func</a> Sobel
``` go
func Sobel(src image.Image) *image.RGBA
```
Sobel returns an image emphasising edges using an approximation to the Sobel–Feldman operator.

![example](https://anthonynsimon.github.io/projects/bild/sobel.jpg)  


## <a name="SoftLight">func</a> SoftLight
``` go
func SoftLight(bg image.Image, fg image.Image) *image.RGBA
```
SoftLight combines the foreground and background images by using Pegtop's Soft Light formula and
returns the resulting image.

![example](https://anthonynsimon.github.io/projects/bild/softlight.jpg)  


## <a name="Subtract">func</a> Subtract
``` go
func Subtract(bg image.Image, fg image.Image) *image.RGBA
```
Subtract combines the foreground and background images by Subtracting the background from the
foreground. The result is then returned.

![example](https://anthonynsimon.github.io/projects/bild/subtract.jpg)  


## <a name="ConvolutionMatrix">type</a> ConvolutionMatrix
``` go
type ConvolutionMatrix interface {
    At(x, y int) float64
    Normalized() ConvolutionMatrix
    SideLength() int
}
```
ConvolutionMatrix interface.
At returns the matrix value at position x, y.
Normalized returns a new matrix with normalized values.
SideLength returns the matrix side length.










## <a name="ConvolutionOptions">type</a> ConvolutionOptions
``` go
type ConvolutionOptions struct {
    Bias       float64
    Wrap       bool
    CarryAlpha bool
}
```
ConvolutionOptions are the convolute function parameters.
Bias is added to each RGB channel after convoluting. Range is -255 to 255.
Wrap sets if indices outside of image dimensions should be taken from the opposite side.
CarryAlpha sets if the alpha should be taken from the source image without convoluting










## <a name="Format">type</a> Format
``` go
type Format int
```
Format is used to identify the image encoding type










## <a name="Kernel">type</a> Kernel
``` go
type Kernel struct {
    Matrix float64
    Stride int
}
```
Kernel to be used as a convolution matrix.







### <a name="NewKernel">func</a> NewKernel
``` go
func NewKernel(length int) *Kernel
```
NewKernel returns a kernel of the provided length.





### <a name="Kernel.At">func</a> (\*Kernel) At
``` go
func (k *Kernel) At(x, y int) float64
```
At returns the matrix value at position x, y.




### <a name="Kernel.Normalized">func</a> (\*Kernel) Normalized
``` go
func (k *Kernel) Normalized() ConvolutionMatrix
```
Normalized returns a new Kernel with normalized values.




### <a name="Kernel.SideLength">func</a> (\*Kernel) SideLength
``` go
func (k *Kernel) SideLength() int
```
SideLength returns the matrix side length.




### <a name="Kernel.String">func</a> (\*Kernel) String
``` go
func (k *Kernel) String() string
```
String returns the string representation of the matrix.




## <a name="RGBAF64">type</a> RGBAF64
``` go
type RGBAF64 struct {
    R, G, B, A float64
}
```
RGBAF64 represents an RGBA color using the range 0.0 to 1.0 with a float64 for each channel.







### <a name="NewRGBAF64">func</a> NewRGBAF64
``` go
func NewRGBAF64(r, g, b, a uint8) RGBAF64
```
NewRGBAF64 returns a new RGBAF64 color based on the provided uint8 values.





### <a name="RGBAF64.Clamp">func</a> (\*RGBAF64) Clamp
``` go
func (c *RGBAF64) Clamp()
```
Clamp limits the channel values of the RGBAF64 color to the range 0.0 to 1.0.
