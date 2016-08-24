# bild

![bild logo](https://anthonynsimon.github.io/projects/bild/logo.png)  

[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/bild/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/anthonynsimon/bild?status.svg)](https://godoc.org/github.com/anthonynsimon/bild)
[![Build Status](https://travis-ci.org/anthonynsimon/bild.svg?branch=master)](https://travis-ci.org/anthonynsimon/bild)
[![Go Report Card](https://goreportcard.com/badge/github.com/anthonynsimon/bild)](https://goreportcard.com/report/github.com/anthonynsimon/bild)

    import "github.com/anthonynsimon/bild"

Simple image processing in Go with parallel processing support.

Package bild provides a collection of common image processing functions. The input images must implement the image.Image interface and the functions return an *image.RGBA.

The aim of this project is simplicity in use and development over high performance, but most algorithms are designed to be efficient and make use of parallelism when available. It is based on standard Go packages to reduce dependency use and development abstractions.

## Install

bild requires Go version 1.4 or greater.

    go get -u github.com/anthonynsimon/bild
    
## Documentation

http://godoc.org/github.com/anthonynsimon/bild

## Basic example:
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

# Output examples
## Adjustment

### Brightness
    result := bild.Brightness(img, 0.25)

![example](https://anthonynsimon.github.io/projects/bild/brightness.jpg)  

### Contrast
    result := bild.Contrast(img, -0.5)

![example](https://anthonynsimon.github.io/projects/bild/contrast.jpg)  

### Gamma
    result := bild.Gamma(img, 2.2)

![example](https://anthonynsimon.github.io/projects/bild/gamma.jpg)  



## Blend modes
    result := bild.Multiply(bg, fg)

| Add | Color Burn | Color Dodge |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/add.jpg) | ![](https://anthonynsimon.github.io/projects/bild/colorburn.jpg) | ![](https://anthonynsimon.github.io/projects/bild/colordodge.jpg) |

| Darken | Difference | Divide |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/darken.jpg) | ![](https://anthonynsimon.github.io/projects/bild/difference.jpg) | ![](https://anthonynsimon.github.io/projects/bild/divide.jpg) |

| Exclusion | Lighten | Linear Burn |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/exclusion.jpg) | ![](https://anthonynsimon.github.io/projects/bild/lighten.jpg) | ![](https://anthonynsimon.github.io/projects/bild/linearburn.jpg) |

| Linear Light | Multiply | Normal |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/linearlight.jpg) | ![](https://anthonynsimon.github.io/projects/bild/multiply.jpg) | ![](https://anthonynsimon.github.io/projects/bild/normal.jpg) |

| Opacity | Overlay | Screen |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/opacity.jpg) | ![](https://anthonynsimon.github.io/projects/bild/overlay.jpg) | ![](https://anthonynsimon.github.io/projects/bild/screen.jpg) |

| Soft Light | Subtract |
|---------- | --------- |
| ![](https://anthonynsimon.github.io/projects/bild/softlight.jpg) | ![](https://anthonynsimon.github.io/projects/bild/subtract.jpg) |


## Blur

### BoxBlur
    result := bild.BoxBlur(img, 3.0)

![example](https://anthonynsimon.github.io/projects/bild/boxblur.jpg)  


### GaussianBlur
    result := bild.GaussianBlur(img, 3.0)


![example](https://anthonynsimon.github.io/projects/bild/gaussianblur.jpg)  


## Channel

### ExtractChannel
    result := bild.ExtractChannel(img, bild.Alpha)

![example](https://anthonynsimon.github.io/projects/bild/extractchannel.jpg)  


## Effects

### EdgeDetection
    result := bild.EdgeDetection(img, 1.0)

![example](https://anthonynsimon.github.io/projects/bild/edgedetection.jpg)  

### Emboss
    result := bild.Emboss(img)

![example](https://anthonynsimon.github.io/projects/bild/emboss.jpg)  

### Grayscale
    result := bild.Grayscale(img)

![example](https://anthonynsimon.github.io/projects/bild/grayscale.jpg)  

### Invert
    result := bild.Invert(img)

![example](https://anthonynsimon.github.io/projects/bild/invert.jpg)  

### Median
    result := bild.Median(img, 10.0)

![example](https://anthonynsimon.github.io/projects/bild/median.jpg)  

### Sharpen
    result := bild.Sharpen(img)

![example](https://anthonynsimon.github.io/projects/bild/sharpen.jpg)  


### Sobel
    result := bild.Sobel(img)

![example](https://anthonynsimon.github.io/projects/bild/sobel.jpg)  


## Resize

### Crop
    // Source image is 280x280
    result := bild.Crop(img, image.Rect(70,70,210,210))

![example](https://anthonynsimon.github.io/projects/bild/crop.jpg)

### Resize Resampling Filters
    result := bild.Resize(img, 280, 280, bild.Linear)

| Nearest Neighbor | Linear | Gaussian |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/resizenearestneighbor.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizelinear.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizegaussian.jpg) |

| Mitchell Netravali | Catmull Rom | Lanczos |
|---------- | --------- | ------ |
| ![](https://anthonynsimon.github.io/projects/bild/resizemitchell.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizecatmullrom.jpg) | ![](https://anthonynsimon.github.io/projects/bild/resizelanczos.jpg) |


## Transform

### FlipH
    result := bild.FlipH(img)

![example](https://anthonynsimon.github.io/projects/bild/fliph.jpg)  

### FlipV
    result := bild.FlipV(img)

![example](https://anthonynsimon.github.io/projects/bild/flipv.jpg) 

### Rotate
    // Options set to nil will use defaults (ResizeBounds set to false, Pivot at center)
    result := bild.Rotate(img, -45.0, nil)

![example](https://anthonynsimon.github.io/projects/bild/rotation03.gif)

    // If ResizeBounds is set to true, the full rotation bounding area is used
    result := bild.Rotate(img, -45.0, &bild.RotationOptions{ResizeBounds: true})

![example](https://anthonynsimon.github.io/projects/bild/rotation01.gif)

    // Pivot coordinates are set from the top-left corner
    // Notice ResizeBounds being set to default (false)
    result := bild.Rotate(img, -45.0, &bild.RotationOptions{Pivot: &image.Point{0, 0}})

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
- Send a pull request, by the way you are awesome.
