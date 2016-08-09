# Bild

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


## Constants
``` go
const (
    JPEG = iota
    PNG
)
```
Supported image encoding types



## func Add
``` go
func Add(bg image.Image, fg image.Image) *image.RGBA
```
Add combines the foreground and background images by adding their values and
returns the resulting image.


## func BoxBlur
``` go
func BoxBlur(src image.Image, radius float64) *image.RGBA
```
BoxBlur returns a blurred (average) version of the image.
Radius must be larger than 0.


## func Brightness
``` go
func Brightness(src image.Image, change float64) *image.RGBA
```
Brightness returns a copy of the image with the adjusted brightness.
Change is the normalized amount of change to be applied (range -1.0 to 1.0).


## func CloneAsRGBA
``` go
func CloneAsRGBA(src image.Image) *image.RGBA
```
CloneAsRGBA returns an RGBA copy of the supplied image.


## func ColorBurn
``` go
func ColorBurn(bg image.Image, fg image.Image) *image.RGBA
```
ColorBurn combines the foreground and background images by dividing the inverted
background by the foreground image and then inverting the result which is then returned.


## func ColorDodge
``` go
func ColorDodge(bg image.Image, fg image.Image) *image.RGBA
```
ColorDodge combines the foreground and background images by dividing background by the
inverted foreground image and returns the result.


## func Convolute
``` go
func Convolute(img image.Image, k ConvolutionMatrix, o *ConvolutionOptions) *image.RGBA
```
Convolute applies a convolution matrix (kernel) to an image with the supplied options.


## func Darken
``` go
func Darken(bg image.Image, fg image.Image) *image.RGBA
```
Darken combines the foreground and background images by picking the darkest value per channel
for each pixel. The result is then returned.


## func Difference
``` go
func Difference(bg image.Image, fg image.Image) *image.RGBA
```
Difference calculates the absolute difference between the foreground and background images and
returns the resulting image.


## func Divide
``` go
func Divide(bg image.Image, fg image.Image) *image.RGBA
```
Divide combines the foreground and background images by diving the values from the background
by the foreground and returns the resulting image.


## func EdgeDetection
``` go
func EdgeDetection(src image.Image, radius float64) *image.RGBA
```
EdgeDetection returns a copy of the image with it's edges highlighted.


## func Emboss
``` go
func Emboss(src image.Image) *image.RGBA
```
Emboss returns a copy of the image in which each pixel has been
replaced either by a highlight or a shadow representation.


## func Encode
``` go
func Encode(w io.Writer, img image.Image, format Format) error
```
Encode writes an image in the specified format.

Usage example:


	// Encode an image to a writer in PNG format,
	// returns an error if something went wrong
	err := Encode(outFile, img, bild.PNG)


## func Exclusion
``` go
func Exclusion(bg image.Image, fg image.Image) *image.RGBA
```
Exclusion combines the foreground and background images applying the Exclusion blend mode and
returns the resulting image.


## func FlipH
``` go
func FlipH(img image.Image) *image.RGBA
```
FlipH returns a horizontally flipped version of the image.


## func FlipV
``` go
func FlipV(img image.Image) *image.RGBA
```
FlipV returns a vertically flipped version of the image.


## func GaussianBlur
``` go
func GaussianBlur(src image.Image, radius float64) *image.RGBA
```
GaussianBlur returns a smoothly blurred version of the image using
a Gaussian function. Radius must be larger than 0.


## func Grayscale
``` go
func Grayscale(src image.Image) *image.RGBA
```
Grayscale returns a copy of the image in Grayscale using the weights
0.3R + 0.6G + 0.1B as a heuristic.


## func Invert
``` go
func Invert(src image.Image) *image.RGBA
```
Invert returns a negated version of the image.


## func Lighten
``` go
func Lighten(bg image.Image, fg image.Image) *image.RGBA
```
Lighten combines the foreground and background images by picking the brightest value per channel
for each pixel. The result is then returned.


## func LinearBurn
``` go
func LinearBurn(bg image.Image, fg image.Image) *image.RGBA
```
LinearBurn combines the foreground and background images by adding them and
then subtracting 255 (1.0 in normalized scale). The resulting image is then returned.


## func LinearLight
``` go
func LinearLight(bg image.Image, fg image.Image) *image.RGBA
```
LinearLight combines the foreground and background images by a mix of a Linear Dodge and
Linear Burn operation. The resulting image is then returned.


## func Median
``` go
func Median(img image.Image, size int) *image.RGBA
```
Median returns a new image in which each pixel is the median of it's neighbors.
Size sets the amount of neighbors to be searched.


## func Multiply
``` go
func Multiply(bg image.Image, fg image.Image) *image.RGBA
```
Multiply combines the foreground and background images by multiplying their
normalized values and returns the resulting image.


## func Opacity
``` go
func Opacity(bg image.Image, fg image.Image, percent float64) *image.RGBA
```
Opacity returns an image which blends the two input images by the percentage provided.
Percent must be of range 0 <= percent <= 1.0


## func Open
``` go
func Open(filename string) (image.Image, error)
```
Open loads and decodes an image from a file and returns it.

Usage example:


	// Encode an image to a writer in PNG format,
	// returns an error if something went wrong
	img, err := Open("exampleName")


## func Overlay
``` go
func Overlay(bg image.Image, fg image.Image) *image.RGBA
```
Overlay combines the foreground and background images by using Multiply when channel values < 0.5
or using Screen otherwise and returns the resulting image.


## func Save
``` go
func Save(filename string, img image.Image, format Format) error
```
Save creates a file and writes to it an image in the specified format

Usage example:


	// Save an image to a file in PNG format,
	// returns an error if something went wrong
	err := Save("exampleName", img, bild.PNG)


## func Screen
``` go
func Screen(bg image.Image, fg image.Image) *image.RGBA
```
Screen combines the foreground and background images by inverting, multiplying and inverting the output.
The result is a brighter image which is then returned.


## func Sobel
``` go
func Sobel(src image.Image) *image.RGBA
```
Sobel returns an image emphasising edges using an approximation to the Sobel–Feldman operator.


## func SoftLight
``` go
func SoftLight(bg image.Image, fg image.Image) *image.RGBA
```
SoftLight combines the foreground and background images by using Pegtop's Soft Light formula and
returns the resulting image.


## func Subtract
``` go
func Subtract(bg image.Image, fg image.Image) *image.RGBA
```
Subtract combines the foreground and background images by Subtracting the background from the
foreground. The result is then returned.



## type ConvolutionMatrix
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











## type ConvolutionOptions
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











## type Format
``` go
type Format int
```
Format is used to identify the image encoding type











## type Kernel
``` go
type Kernel struct {
    Matrix []float64
    Stride int
}
```
Kernel to be used as a convolution matrix.









### func NewKernel
``` go
func NewKernel(length int) *Kernel
```
NewKernel returns a kernel of the provided length.




### func (\*Kernel) At
``` go
func (k *Kernel) At(x, y int) float64
```
At returns the matrix value at position x, y.



### func (\*Kernel) Normalized
``` go
func (k *Kernel) Normalized() ConvolutionMatrix
```
Normalized returns a new Kernel with normalized values.



### func (\*Kernel) SideLength
``` go
func (k *Kernel) SideLength() int
```
SideLength returns the matrix side length.



### func (\*Kernel) String
``` go
func (k *Kernel) String() string
```
String returns the string representation of the matrix.