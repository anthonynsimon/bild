# bild

![bild logo](https://anthonynsimon.github.io/projects/bild/logo.png)  

[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/bild/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/anthonynsimon/bild?status.svg)](https://godoc.org/github.com/anthonynsimon/bild)
[![Build Status](https://travis-ci.org/anthonynsimon/bild.svg?branch=master)](https://travis-ci.org/anthonynsimon/bild)
[![Go Report Card](https://goreportcard.com/badge/github.com/anthonynsimon/bild)](https://goreportcard.com/report/github.com/anthonynsimon/bild)

Simple image processing in Go with parallel processing support.

The aim of this project is simplicity in use and development over high performance, but most algorithms are designed to be efficient and make use of parallelism when available. It is based on standard Go packages to reduce dependency use and development abstractions.

Package bild provides a collection of common image processing functions. The input images must implement the image.Image interface and the functions return an *image.RGBA.

**Notice:** This package is under heavy development and the API might change at any time until a 1.0 version is reached.

 
## Documentation

http://godoc.org/github.com/anthonynsimon/bild


## Install

bild requires Go version 1.4 or greater.

    go get -u github.com/anthonynsimon/bild/...

Notice the '...' at the end, this is to signify that you want all the packages on the repo. In your code, simply import the specific package that you want to use (see example below).


## Basic example:
```go
package main

import (
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func main() {
	img, err := imgio.Open("filename.jpg")
	if err != nil {
		panic(err)
	}

	inverted := effect.Invert(img)
	resized := transform.Resize(inverted, 800, 800, transform.Linear)
	rotated := transform.Rotate(resized, 45, nil)

	if err := imgio.Save("filename", rotated, imgio.PNG); err != nil {
		panic(err)
	}
}

```

# Output examples
## Adjustment
    import "github.com/anthonynsimon/bild/adjust"

### Brightness
    result := adjust.Brightness(img, 0.25)

![example](https://anthonynsimon.github.io/projects/bild/brightness.jpg)  

### Contrast
    result := adjust.Contrast(img, -0.5)

![example](https://anthonynsimon.github.io/projects/bild/contrast.jpg)  

### Gamma
    result := adjust.Gamma(img, 2.2)

![example](https://anthonynsimon.github.io/projects/bild/gamma.jpg)  



## Blend modes
    import "github.com/anthonynsimon/bild/blend"

    result := blend.Multiply(bg, fg)

| Add | Color Burn | Color Dodge |
| :----------: | :---------: | :------: |
| ![](https://anthonynsimon.github.io/projects/bild/add.jpg) | ![](https://anthonynsimon.github.io/projects/bild/colorburn.jpg) | ![](https://anthonynsimon.github.io/projects/bild/colordodge.jpg) |
| **Darken** | **Difference** | **Divide** |
| ![](https://anthonynsimon.github.io/projects/bild/darken.jpg) | ![](https://anthonynsimon.github.io/projects/bild/difference.jpg) | ![](https://anthonynsimon.github.io/projects/bild/divide.jpg) |
| **Exclusion** | **Lighten** | **Linear Burn** |
| ![](https://anthonynsimon.github.io/projects/bild/exclusion.jpg) | ![](https://anthonynsimon.github.io/projects/bild/lighten.jpg) | ![](https://anthonynsimon.github.io/projects/bild/linearburn.jpg) |
| **Linear Light** | **Multiply** | **Normal** |
| ![](https://anthonynsimon.github.io/projects/bild/linearlight.jpg) | ![](https://anthonynsimon.github.io/projects/bild/multiply.jpg) | ![](https://anthonynsimon.github.io/projects/bild/normal.jpg) |
| **Opacity** | **Overlay** | **Screen** |
| ![](https://anthonynsimon.github.io/projects/bild/opacity.jpg) | ![](https://anthonynsimon.github.io/projects/bild/overlay.jpg) | ![](https://anthonynsimon.github.io/projects/bild/screen.jpg) |
| **Soft Light** | **Subtract** | |
| ![](https://anthonynsimon.github.io/projects/bild/softlight.jpg) | ![](https://anthonynsimon.github.io/projects/bild/subtract.jpg) | |


## Blur
    import "github.com/anthonynsimon/bild/blur"

### Box Blur
    result := blur.Box(img, 3.0)

![example](https://anthonynsimon.github.io/projects/bild/boxblur.jpg)  


### Gaussian Blur
    result := blur.Gaussian(img, 3.0)


![example](https://anthonynsimon.github.io/projects/bild/gaussianblur.jpg)  


## Channel
    import "github.com/anthonynsimon/bild/channel"

### Extract Channel
    result := channel.Extract(img, channel.Alpha)

![example](https://anthonynsimon.github.io/projects/bild/extractchannel.jpg)  


## Effect
    import "github.com/anthonynsimon/bild/effect"

### Edge Detection
    result := effect.EdgeDetection(img, 1.0)

![example](https://anthonynsimon.github.io/projects/bild/edgedetection.jpg)  

### Emboss
    result := effect.Emboss(img)

![example](https://anthonynsimon.github.io/projects/bild/emboss.jpg)  

### Grayscale
    result := effect.Grayscale(img)

![example](https://anthonynsimon.github.io/projects/bild/grayscale.jpg)  

### Invert
    result := effect.Invert(img)

![example](https://anthonynsimon.github.io/projects/bild/invert.jpg)  

### Median
    result := effect.Median(img, 10.0)

![example](https://anthonynsimon.github.io/projects/bild/median.jpg)  

### Sharpen
    result := effect.Sharpen(img)

![example](https://anthonynsimon.github.io/projects/bild/sharpen.jpg)  


### Sobel
    result := effect.Sobel(img)

![example](https://anthonynsimon.github.io/projects/bild/sobel.jpg)  


## Histogram
    import "github.com/anthonynsimon/bild/histogram"

### RGBA Histogram
	hist := histogram.NewRGBAHistogram(img)
    result := hist.Image()

![example](https://anthonynsimon.github.io/projects/bild/histogram.png)  


## Noise
    import "github.com/anthonynsimon/bild/noise"

### Uniform colored
	result := noise.Generate(280, 280, &noise.Options{Monochrome: false, NoiseFn: noise.Uniform})

![example](https://anthonynsimon.github.io/projects/bild/noiseuniform.jpg)  


### Binary monochrome
	result := noise.Generate(280, 280, &noise.Options{Monochrome: true, NoiseFn: noise.Binary})

![example](https://anthonynsimon.github.io/projects/bild/noisebinary.jpg)  


### Gaussian monochrome
	result := noise.Generate(280, 280, &noise.Options{Monochrome: true, NoiseFn: noise.Gaussian})

![example](https://anthonynsimon.github.io/projects/bild/noisegaussian.jpg)  


## Segmentation

### Threshold
	result := bild.Threshold(img, 128)

![example](https://anthonynsimon.github.io/projects/bild/threshold.jpg)


## Transform
    import "github.com/anthonynsimon/bild/transform"

### Crop
    // Source image is 280x280
    result := transform.Crop(img, image.Rect(70,70,210,210))

![example](https://anthonynsimon.github.io/projects/bild/crop.jpg)

### FlipH
    result := transform.FlipH(img)

![example](https://anthonynsimon.github.io/projects/bild/fliph.jpg)  

### FlipV
    result := transform.FlipV(img)

![example](https://anthonynsimon.github.io/projects/bild/flipv.jpg) 


### Resize Resampling Filters
    result := transform.Resize(img, 280, 280, transform.Linear)

| Nearest Neighbor | Linear | Gaussian |
|:----------: | :---------: | :------: |
| ![](https://anthonynsimon.github.io/projects/bild/resizenearestneighbor.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizelinear.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizegaussian.jpg) |
| **Mitchell Netravali** | **Catmull Rom** | **Lanczos** |
| ![](https://anthonynsimon.github.io/projects/bild/resizemitchell.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizecatmullrom.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizelanczos.jpg) |


### Rotate
    // Options set to nil will use defaults (ResizeBounds set to false, Pivot at center)
    result := transform.Rotate(img, -45.0, nil)

![example](https://anthonynsimon.github.io/projects/bild/rotation03.gif)

    // If ResizeBounds is set to true, the full rotation bounding area is used
    result := transform.Rotate(img, -45.0, &transform.RotationOptions{ResizeBounds: true})

![example](https://anthonynsimon.github.io/projects/bild/rotation01.gif)

    // Pivot coordinates are set from the top-left corner
    // Notice ResizeBounds being set to default (false)
    result := transform.Rotate(img, -45.0, &transform.RotationOptions{Pivot: &image.Point{0, 0}})

![example](https://anthonynsimon.github.io/projects/bild/rotation02.gif)


## License

This project is licensed under the MIT license. Please read the LICENSE file.

## Contribute

Want to hack on the project? Any kind of contribution is welcome!  
Simply follow the next steps:

- Fork the project.
- Create a new branch.
- Make your changes and write tests when practical.
- Commit your changes to the new branch.
- Send a pull request.
