# bild

![bild logo](https://s10.postimg.org/4x22ndng9/bild.png)  

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




## <a name="Add">func</a> [Add](/src/target/blend.go?s=599:651#L5)
``` go
func Add(bg image.Image, fg image.Image) *image.RGBA
```
Add combines the foreground and background images by adding their values and
returns the resulting image.



## <a name="BoxBlur">func</a> [BoxBlur](/src/target/blur.go?s=137:194#L1)
``` go
func BoxBlur(src image.Image, radius float64) *image.RGBA
```
BoxBlur returns a blurred (average) version of the image.
Radius must be larger than 0.



## <a name="Brightness">func</a> [Brightness](/src/target/adjustment.go?s=210:270#L1)
``` go
func Brightness(src image.Image, change float64) *image.RGBA
```
Brightness returns a copy of the image with the adjusted brightness.
Change is the normalized amount of change to be applied (range -1.0 to 1.0).



## <a name="CloneAsRGBA">func</a> [CloneAsRGBA](/src/target/helpers.go?s=131:176#L1)
``` go
func CloneAsRGBA(src image.Image) *image.RGBA
```
CloneAsRGBA returns an RGBA copy of the supplied image.



## <a name="ColorBurn">func</a> [ColorBurn](/src/target/blend.go?s=3758:3816#L122)
``` go
func ColorBurn(bg image.Image, fg image.Image) *image.RGBA
```
ColorBurn combines the foreground and background images by dividing the inverted
background by the foreground image and then inverting the result which is then returned.



## <a name="ColorDodge">func</a> [ColorDodge](/src/target/blend.go?s=4748:4807#L166)
``` go
func ColorDodge(bg image.Image, fg image.Image) *image.RGBA
```
ColorDodge combines the foreground and background images by dividing background by the
inverted foreground image and returns the result.



## <a name="Convolute">func</a> [Convolute](/src/target/convolution.go?s=1931:2018#L72)
``` go
func Convolute(img image.Image, k ConvolutionMatrix, o *ConvolutionOptions) *image.RGBA
```
Convolute applies a convolution matrix (kernel) to an image with the supplied options.



## <a name="Darken">func</a> [Darken](/src/target/blend.go?s=7116:7171#L256)
``` go
func Darken(bg image.Image, fg image.Image) *image.RGBA
```
Darken combines the foreground and background images by picking the darkest value per channel
for each pixel. The result is then returned.



## <a name="Difference">func</a> [Difference](/src/target/blend.go?s=2903:2962#L92)
``` go
func Difference(bg image.Image, fg image.Image) *image.RGBA
```
Difference calculates the absolute difference between the foreground and background images and
returns the resulting image.



## <a name="Divide">func</a> [Divide](/src/target/blend.go?s=3335:3390#L107)
``` go
func Divide(bg image.Image, fg image.Image) *image.RGBA
```
Divide combines the foreground and background images by diving the values from the background
by the foreground and returns the resulting image.



## <a name="EdgeDetection">func</a> [EdgeDetection](/src/target/effects.go?s=783:846#L31)
``` go
func EdgeDetection(src image.Image, radius float64) *image.RGBA
```
EdgeDetection returns a copy of the image with it's edges highlighted.



## <a name="Emboss">func</a> [Emboss](/src/target/effects.go?s=1372:1412#L54)
``` go
func Emboss(src image.Image) *image.RGBA
```
Emboss returns a copy of the image in which each pixel has been
replaced either by a highlight or a shadow representation.



## <a name="Encode">func</a> [Encode](/src/target/util.go?s=898:960#L41)
``` go
func Encode(w io.Writer, img image.Image, format Format) error
```
Encode writes an image in the specified format.

Usage example:


	// Encode an image to a writer in PNG format,
	// returns an error if something went wrong
	err := Encode(outFile, img, bild.PNG)



## <a name="Exclusion">func</a> [Exclusion](/src/target/blend.go?s=4300:4358#L150)
``` go
func Exclusion(bg image.Image, fg image.Image) *image.RGBA
```
Exclusion combines the foreground and background images applying the Exclusion blend mode and
returns the resulting image.



## <a name="FlipH">func</a> [FlipH](/src/target/transform.go?s=92:131#L1)
``` go
func FlipH(img image.Image) *image.RGBA
```
FlipH returns a horizontally flipped version of the image.



## <a name="FlipV">func</a> [FlipV](/src/target/transform.go?s=737:776#L22)
``` go
func FlipV(img image.Image) *image.RGBA
```
FlipV returns a vertically flipped version of the image.



## <a name="GaussianBlur">func</a> [GaussianBlur](/src/target/blur.go?s=636:698#L19)
``` go
func GaussianBlur(src image.Image, radius float64) *image.RGBA
```
GaussianBlur returns a smoothly blurred version of the image using
a Gaussian function. Radius must be larger than 0.



## <a name="Grayscale">func</a> [Grayscale](/src/target/effects.go?s=401:444#L12)
``` go
func Grayscale(src image.Image) *image.RGBA
```
Grayscale returns a copy of the image in Grayscale using the weights
0.3R + 0.6G + 0.1B as a heuristic.



## <a name="Invert">func</a> [Invert](/src/target/effects.go?s=108:148#L1)
``` go
func Invert(src image.Image) *image.RGBA
```
Invert returns a negated version of the image.



## <a name="Lighten">func</a> [Lighten](/src/target/blend.go?s=7538:7594#L271)
``` go
func Lighten(bg image.Image, fg image.Image) *image.RGBA
```
Lighten combines the foreground and background images by picking the brightest value per channel
for each pixel. The result is then returned.



## <a name="LinearBurn">func</a> [LinearBurn](/src/target/blend.go?s=5185:5244#L181)
``` go
func LinearBurn(bg image.Image, fg image.Image) *image.RGBA
```
LinearBurn combines the foreground and background images by adding them and
then subtracting 255 (1.0 in normalized scale). The resulting image is then returned.



## <a name="LinearLight">func</a> [LinearLight](/src/target/blend.go?s=5604:5664#L196)
``` go
func LinearLight(bg image.Image, fg image.Image) *image.RGBA
```
LinearLight combines the foreground and background images by a mix of a Linear Dodge and
Linear Burn operation. The resulting image is then returned.



## <a name="Median">func</a> [Median](/src/target/effects.go?s=2553:2603#L98)
``` go
func Median(img image.Image, size int) *image.RGBA
```
Median returns a new image in which each pixel is the median of it's neighbors.
Size sets the amount of neighbors to be searched.



## <a name="Multiply">func</a> [Multiply](/src/target/blend.go?s=976:1033#L20)
``` go
func Multiply(bg image.Image, fg image.Image) *image.RGBA
```
Multiply combines the foreground and background images by multiplying their
normalized values and returns the resulting image.



## <a name="Opacity">func</a> [Opacity](/src/target/blend.go?s=6604:6677#L239)
``` go
func Opacity(bg image.Image, fg image.Image, percent float64) *image.RGBA
```
Opacity returns an image which blends the two input images by the percentage provided.
Percent must be of range 0 <= percent <= 1.0



## <a name="Open">func</a> [Open](/src/target/util.go?s=457:504#L19)
``` go
func Open(filename string) (image.Image, error)
```
Open loads and decodes an image from a file and returns it.

Usage example:


	// Encode an image to a writer in PNG format,
	// returns an error if something went wrong
	img, err := Open("exampleName")



## <a name="Overlay">func</a> [Overlay](/src/target/blend.go?s=1388:1444#L35)
``` go
func Overlay(bg image.Image, fg image.Image) *image.RGBA
```
Overlay combines the foreground and background images by using Multiply when channel values < 0.5
or using Screen otherwise and returns the resulting image.



## <a name="Save">func</a> [Save](/src/target/util.go?s=1358:1422#L61)
``` go
func Save(filename string, img image.Image, format Format) error
```
Save creates a file and writes to it an image in the specified format

Usage example:


	// Save an image to a file in PNG format,
	// returns an error if something went wrong
	err := Save("exampleName", img, bild.PNG)



## <a name="Screen">func</a> [Screen](/src/target/blend.go?s=2496:2551#L77)
``` go
func Screen(bg image.Image, fg image.Image) *image.RGBA
```
Screen combines the foreground and background images by inverting, multiplying and inverting the output.
The result is a brighter image which is then returned.



## <a name="Sharpen">func</a> [Sharpen](/src/target/effects.go?s=1680:1721#L65)
``` go
func Sharpen(src image.Image) *image.RGBA
```
Sharpen returns a sharpened copy of the image by detecting it's edges and adding it to the original.



## <a name="Sobel">func</a> [Sobel](/src/target/effects.go?s=1984:2023#L76)
``` go
func Sobel(src image.Image) *image.RGBA
```
Sobel returns an image emphasising edges using an approximation to the Sobel–Feldman operator.



## <a name="SoftLight">func</a> [SoftLight](/src/target/blend.go?s=2012:2070#L63)
``` go
func SoftLight(bg image.Image, fg image.Image) *image.RGBA
```
SoftLight combines the foreground and background images by using Pegtop's Soft Light formula and
returns the resulting image.



## <a name="Subtract">func</a> [Subtract](/src/target/blend.go?s=6217:6274#L224)
``` go
func Subtract(bg image.Image, fg image.Image) *image.RGBA
```
Subtract combines the foreground and background images by Subtracting the background from the
foreground. The result is then returned.




## <a name="ConvolutionMatrix">type</a> [ConvolutionMatrix](/src/target/convolution.go?s=236:344#L3)
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










## <a name="ConvolutionOptions">type</a> [ConvolutionOptions](/src/target/convolution.go?s=1751:1839#L65)
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










## <a name="Format">type</a> [Format](/src/target/util.go?s=156:171#L4)
``` go
type Format int
```
Format is used to identify the image encoding type










## <a name="Kernel">type</a> [Kernel](/src/target/convolution.go?s=542:594#L15)
``` go
type Kernel struct {
    Matrix []float64
    Stride int
}
```
Kernel to be used as a convolution matrix.







### <a name="NewKernel">func</a> [NewKernel](/src/target/convolution.go?s=400:434#L10)
``` go
func NewKernel(length int) *Kernel
```
NewKernel returns a kernel of the provided length.





### <a name="Kernel.At">func</a> (\*Kernel) [At](/src/target/convolution.go?s=1063:1100#L44)
``` go
func (k *Kernel) At(x, y int) float64
```
At returns the matrix value at position x, y.




### <a name="Kernel.Normalized">func</a> (\*Kernel) [Normalized](/src/target/convolution.go?s=655:702#L21)
``` go
func (k *Kernel) Normalized() ConvolutionMatrix
```
Normalized returns a new Kernel with normalized values.




### <a name="Kernel.SideLength">func</a> (\*Kernel) [SideLength](/src/target/convolution.go?s=958:991#L39)
``` go
func (k *Kernel) SideLength() int
```
SideLength returns the matrix side length.




### <a name="Kernel.String">func</a> (\*Kernel) [String](/src/target/convolution.go?s=1196:1228#L49)
``` go
func (k *Kernel) String() string
```
String returns the string representation of the matrix.




## <a name="RGBAF64">type</a> [RGBAF64](/src/target/color.go?s=110:153#L1)
``` go
type RGBAF64 struct {
    R, G, B, A float64
}
```
RGBAF64 represents an RGBA color using the range 0.0 to 1.0 with a float64 for each channel.







### <a name="NewRGBAF64">func</a> [NewRGBAF64](/src/target/color.go?s=233:274#L1)
``` go
func NewRGBAF64(r, g, b, a uint8) RGBAF64
```
NewRGBAF64 returns a new RGBAF64 color based on the provided uint8 values.





### <a name="RGBAF64.Clamp">func</a> (\*RGBAF64) [Clamp](/src/target/color.go?s=506:531#L9)
``` go
func (c *RGBAF64) Clamp()
```
Clamp limits the channel values of the RGBAF64 color to the range 0.0 to 1.0.