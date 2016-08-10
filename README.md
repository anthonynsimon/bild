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
    bild.Brightness(img, 0.25)

![example](https://anthonynsimon.github.io/projects/bild/brightness.jpg)  



## Blend modes

### Add
        bild.Add(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/add.jpg)  

### ColorBurn
        bild.ColorBurn(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/colorburn.jpg)  

### ColorDodge
        bild.ColorDodge(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/colordodge.jpg) 

### Darken
        bild.Darken(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/darken.jpg)  

### Difference
        bild.Difference(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/difference.jpg)  

### Divide
        bild.Divide(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/divide.jpg)  

### Exclusion
        bild.Exclusion(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/exclusion.jpg)  

### Lighten
        bild.Lighten(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/lighten.jpg)  

### LinearBurn
        bild.LinearBurn(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/linearburn.jpg)  


### LinearLight
        bild.LinearLight(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/linearlight.jpg)  

### Multiply
        bild.Multiply(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/multiply.jpg)  

### Normal
        bild.Normal(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/normal.jpg)  

### Opacity
        bild.Opacity(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/opacity.jpg)  

### Overlay
        bild.Overlay(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/overlay.jpg)  

### Screen
        bild.Screen(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/screen.jpg)  

### SoftLight
        bild.SoftLight(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/softlight.jpg)  

### Subtract
        bild.Subtract(bg, fg)

![example](https://anthonynsimon.github.io/projects/bild/subtract.jpg)  



## Blur

### BoxBlur
        bild.BoxBlur(img, 3.0)

![example](https://anthonynsimon.github.io/projects/bild/boxblur.jpg)  


### GaussianBlur
        bild.GaussianBlur(img, 3.0)


![example](https://anthonynsimon.github.io/projects/bild/gaussianblur.jpg)  



## Effects

### EdgeDetection
        bild.EdgeDetection(img, 1.0)

![example](https://anthonynsimon.github.io/projects/bild/edgedetection.jpg)  

### Emboss
        bild.Emboss(img)

![example](https://anthonynsimon.github.io/projects/bild/emboss.jpg)  

### Grayscale
        bild.Grayscale(img)

![example](https://anthonynsimon.github.io/projects/bild/grayscale.jpg)  

### Invert
        bild.Invert(img)

![example](https://anthonynsimon.github.io/projects/bild/invert.jpg)  

### Median
        bild.Median(img, 10.0)

![example](https://anthonynsimon.github.io/projects/bild/median.jpg)  

### Sharpen
        bild.Sharpen(img)

![example](https://anthonynsimon.github.io/projects/bild/sharpen.jpg)  


### Sobel
        bild.Sobel(img)

![example](https://anthonynsimon.github.io/projects/bild/sobel.jpg)  



## Transform

### FlipH
        bild.FlipH(img)

![example](https://anthonynsimon.github.io/projects/bild/fliph.jpg)  

### FlipV
        bild.FlipV(img)

![example](https://anthonynsimon.github.io/projects/bild/flipv.jpg)  
