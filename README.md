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

### Add
    result := bild.Add(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/add.jpg)  

### ColorBurn
    result := bild.ColorBurn(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/colorburn.jpg)  

### ColorDodge
    result := bild.ColorDodge(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/colordodge.jpg) 

### Darken
    result := bild.Darken(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/darken.jpg)  

### Difference
    result := bild.Difference(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/difference.jpg)  

### Divide
    result := bild.Divide(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/divide.jpg)  

### Exclusion
    result := bild.Exclusion(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/exclusion.jpg)  

### Lighten
    result := bild.Lighten(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/lighten.jpg)  

### LinearBurn
    result := bild.LinearBurn(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/linearburn.jpg)  


### LinearLight
    result := bild.LinearLight(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/linearlight.jpg)  

### Multiply
    result := bild.Multiply(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/multiply.jpg)  

### Normal
    result := bild.Normal(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/normal.jpg)  

### Opacity
    result := bild.Opacity(bg, fg, 0.5)

![example](https://anthonynsimon.github.io/projects/bild/opacity.jpg)  

### Overlay
    result := bild.Overlay(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/overlay.jpg)  

### Screen
    result := bild.Screen(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/screen.jpg)  

### SoftLight
    result := bild.SoftLight(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/softlight.jpg)  

### Subtract
    result := bild.Subtract(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/subtract.jpg)  



## Blur

### BoxBlur
    result := bild.BoxBlur(img, 3.0)

![example](https://anthonynsimon.github.io/projects/bild/boxblur.jpg)  


### GaussianBlur
    result := bild.GaussianBlur(img, 3.0)


![example](https://anthonynsimon.github.io/projects/bild/gaussianblur.jpg)  



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



## Transform

### FlipH
    result := bild.FlipH(img)

![example](https://anthonynsimon.github.io/projects/bild/fliph.jpg)  

### FlipV
    result := bild.FlipV(img)

![example](https://anthonynsimon.github.io/projects/bild/flipv.jpg)  

# License

This project is licensed under the MIT license. Please read the LICENSE file.

# Contribute

Want to hack on the project? Any kind of contribution is welcome!  
Simply follow the next steps:

- Fork the project.
- Create a new branch.
- Make your changes and write tests when practical.
- Commit your changes to the new branch.
- Send a pull request, by the way you are awesome.
