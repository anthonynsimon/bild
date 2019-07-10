# bild

![bild logo](https://anthonynsimon.com/projects/bild/logo.png)  

[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/bild/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/anthonynsimon/bild?status.svg)](https://godoc.org/github.com/anthonynsimon/bild)
[![Build Status](https://travis-ci.org/anthonynsimon/bild.svg?branch=master)](https://travis-ci.org/anthonynsimon/bild)
[![Go Report Card](https://goreportcard.com/badge/github.com/anthonynsimon/bild)](https://goreportcard.com/report/github.com/anthonynsimon/bild)

A collection of parallel image processing algorithms in pure Go.

The aim of this project is simplicity in use and development over high performance, but most algorithms are designed to be efficient and make use of parallelism when available. It is based on standard Go packages to reduce dependency use and development abstractions.
 
## Documentation

http://godoc.org/github.com/anthonynsimon/bild


## Install

bild requires Go version 1.7 or greater.
```bash
go get -u github.com/anthonynsimon/bild/...
```

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

    // Or imgio.JPEGEncoder(95) as encoder for JPG with quality of 95%
    if err := imgio.Save("filename.png", rotated, imgio.PNGEncoder()); err != nil {
        panic(err)
    }
}

```

# Output examples
## Adjustment
    import "github.com/anthonynsimon/bild/adjust"

### Brightness
    result := adjust.Brightness(img, 0.25)

![example](https://anthonynsimon.com/projects/bild/brightness.jpg)  

### Contrast
    result := adjust.Contrast(img, -0.5)

![example](https://anthonynsimon.com/projects/bild/contrast.jpg)  

### Gamma
    result := adjust.Gamma(img, 2.2)

![example](https://anthonynsimon.com/projects/bild/gamma.jpg)  


### Hue
    result := adjust.Hue(img, -42)

![example](https://anthonynsimon.com/projects/bild/hue.jpg)  

### Saturation
    result := adjust.Saturation(img, 0.5)

![example](https://anthonynsimon.com/projects/bild/saturation.jpg)  



## Blend modes
    import "github.com/anthonynsimon/bild/blend"

    result := blend.Multiply(bg, fg)

| Add | Color Burn | Color Dodge |
| :----------: | :---------: | :------: |
| ![](https://anthonynsimon.com/projects/bild/add.jpg) | ![](https://anthonynsimon.com/projects/bild/colorburn.jpg) | ![](https://anthonynsimon.com/projects/bild/colordodge.jpg) |
| **Darken** | **Difference** | **Divide** |
| ![](https://anthonynsimon.com/projects/bild/darken.jpg) | ![](https://anthonynsimon.com/projects/bild/difference.jpg) | ![](https://anthonynsimon.com/projects/bild/divide.jpg) |
| **Exclusion** | **Lighten** | **Linear Burn** |
| ![](https://anthonynsimon.com/projects/bild/exclusion.jpg) | ![](https://anthonynsimon.com/projects/bild/lighten.jpg) | ![](https://anthonynsimon.com/projects/bild/linearburn.jpg) |
| **Linear Light** | **Multiply** | **Normal** |
| ![](https://anthonynsimon.com/projects/bild/linearlight.jpg) | ![](https://anthonynsimon.com/projects/bild/multiply.jpg) | ![](https://anthonynsimon.com/projects/bild/normal.jpg) |
| **Opacity** | **Overlay** | **Screen** |
| ![](https://anthonynsimon.com/projects/bild/opacity.jpg) | ![](https://anthonynsimon.com/projects/bild/overlay.jpg) | ![](https://anthonynsimon.com/projects/bild/screen.jpg) |
| **Soft Light** | **Subtract** | |
| ![](https://anthonynsimon.com/projects/bild/softlight.jpg) | ![](https://anthonynsimon.com/projects/bild/subtract.jpg) | |


## Blur
    import "github.com/anthonynsimon/bild/blur"

### Box Blur
    result := blur.Box(img, 3.0)

![example](https://anthonynsimon.com/projects/bild/boxblur.jpg)  


### Gaussian Blur
    result := blur.Gaussian(img, 3.0)


![example](https://anthonynsimon.com/projects/bild/gaussianblur.jpg)  


## Channel
    import "github.com/anthonynsimon/bild/channel"

### Extract Channel
    result := channel.Extract(img, channel.Alpha)

![example](https://anthonynsimon.com/projects/bild/extractchannel.jpg)  


## Effect
    import "github.com/anthonynsimon/bild/effect"

### Dilate
    result := effect.Dilate(img, 3)

![example](https://anthonynsimon.com/projects/bild/dilate.jpg)     

### Edge Detection
    result := effect.EdgeDetection(img, 1.0)

![example](https://anthonynsimon.com/projects/bild/edgedetection.jpg)  

### Emboss
    result := effect.Emboss(img)

![example](https://anthonynsimon.com/projects/bild/emboss.jpg)  

### Erode
    result := effect.Erode(img, 3)

![example](https://anthonynsimon.com/projects/bild/erode.jpg)   

### Grayscale
    result := effect.Grayscale(img)

![example](https://anthonynsimon.com/projects/bild/grayscale.jpg)  

### Invert
    result := effect.Invert(img)

![example](https://anthonynsimon.com/projects/bild/invert.jpg)  

### Median
    result := effect.Median(img, 10.0)

![example](https://anthonynsimon.com/projects/bild/median.jpg)  

### Sepia
    result := effect.Sepia(img)

![example](https://anthonynsimon.com/projects/bild/sepia.jpg)  

### Sharpen
    result := effect.Sharpen(img)

![example](https://anthonynsimon.com/projects/bild/sharpen.jpg)  

### Sobel
    result := effect.Sobel(img)

![example](https://anthonynsimon.com/projects/bild/sobel.jpg)  

### Unsharp Mask
    result := effect.UnsharpMask(img, 0.6, 1.2)

![example](https://anthonynsimon.com/projects/bild/unsharpmask.jpg)  


## Histogram
    import "github.com/anthonynsimon/bild/histogram"

### RGBA Histogram
    hist := histogram.NewRGBAHistogram(img)
    result := hist.Image()

![example](https://anthonynsimon.com/projects/bild/histogram.png)  


## Noise
    import "github.com/anthonynsimon/bild/noise"

### Uniform colored
    result := noise.Generate(280, 280, &noise.Options{Monochrome: false, NoiseFn: noise.Uniform})

![example](https://anthonynsimon.com/projects/bild/noiseuniform.jpg)  


### Binary monochrome
    result := noise.Generate(280, 280, &noise.Options{Monochrome: true, NoiseFn: noise.Binary})

![example](https://anthonynsimon.com/projects/bild/noisebinary.jpg)  


### Gaussian monochrome
    result := noise.Generate(280, 280, &noise.Options{Monochrome: true, NoiseFn: noise.Gaussian})

![example](https://anthonynsimon.com/projects/bild/noisegaussian.jpg)  


## Paint
    import "github.com/anthonynsimon/bild/paint"

### Flood Fill
    // Fuzz is the percentage of maximum color distance that is tolerated
    result := paint.FloodFill(img, image.Point{240, 0}, color.RGBA{255, 0, 0, 255}, 15)

![example](https://anthonynsimon.com/projects/bild/floodfill.jpg) 


## Segmentation
    import "github.com/anthonynsimon/bild/segment"

### Threshold
    result := segment.Threshold(img, 128)

![example](https://anthonynsimon.com/projects/bild/threshold.jpg)


## Transform
    import "github.com/anthonynsimon/bild/transform"

### Crop
    // Source image is 280x280
    result := transform.Crop(img, image.Rect(70,70,210,210))

![example](https://anthonynsimon.com/projects/bild/crop.jpg)

### FlipH
    result := transform.FlipH(img)

![example](https://anthonynsimon.com/projects/bild/fliph.jpg)  

### FlipV
    result := transform.FlipV(img)

![example](https://anthonynsimon.com/projects/bild/flipv.jpg) 


### Resize Resampling Filters
    result := transform.Resize(img, 280, 280, transform.Linear)

| Nearest Neighbor | Linear | Gaussian |
|:----------: | :---------: | :------: |
| ![](https://anthonynsimon.com/projects/bild/resizenearestneighbor.jpg) | ![](https://anthonynsimon.com/projects/bild/resizelinear.jpg) | ![](https://anthonynsimon.com/projects/bild/resizegaussian.jpg) |
| **Mitchell Netravali** | **Catmull Rom** | **Lanczos** |
| ![](https://anthonynsimon.com/projects/bild/resizemitchell.jpg) | ![](https://anthonynsimon.com/projects/bild/resizecatmullrom.jpg) | ![](https://anthonynsimon.com/projects/bild/resizelanczos.jpg) |


### Rotate
    // Options set to nil will use defaults (ResizeBounds set to false, Pivot at center)
    result := transform.Rotate(img, -45.0, nil)

![example](https://anthonynsimon.com/projects/bild/rotation03.gif)

    // If ResizeBounds is set to true, the full rotation bounding area is used
    result := transform.Rotate(img, -45.0, &transform.RotationOptions{ResizeBounds: true})

![example](https://anthonynsimon.com/projects/bild/rotation01.gif)

    // Pivot coordinates are set from the top-left corner
    // Notice ResizeBounds being set to default (false)
    result := transform.Rotate(img, -45.0, &transform.RotationOptions{Pivot: &image.Point{0, 0}})

![example](https://anthonynsimon.com/projects/bild/rotation02.gif)

### Shear Horizontal
    result := transform.ShearH(img, 30)

![example](https://anthonynsimon.com/projects/bild/shearh.jpg)  

### Shear Vertical
    result := transform.ShearV(img, 30)

![example](https://anthonynsimon.com/projects/bild/shearv.jpg) 

### Translate
    result := transform.Translate(img, 80, 0)

![example](https://anthonynsimon.com/projects/bild/translate.jpg) 


## License

This project is licensed under the MIT license. Please read the LICENSE file.

## Contribute

Want to hack on the project? Any kind of contribution is welcome!  
Simply follow the next steps:

- Fork the project.
- Create a new branch.
- Make your changes and write tests when practical.
- Commit your changes to the new branch.
- Send a pull request, it will be reviewed shortly.

In case you want to add a feature, please create a new issue and briefly explain what the feature would consist of. 
For bugs or requests, before creating an issue please check if one has already been created for it.
