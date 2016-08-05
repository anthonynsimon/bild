# Bild

[![GoDoc](https://godoc.org/github.com/anthonynsimon/bild?status.svg)](https://godoc.org/github.com/anthonynsimon/bild)
[![Build Status](https://travis-ci.org/anthonynsimon/bild.svg?branch=master)](https://travis-ci.org/anthonynsimon/bild)

Simple image processing in Go with parallel processing support  

Blend Modes:
- Add
- Multiply
- Overlay
- Soft Light
- Screen
- Difference
- Divide
- Color Burn
- Exclusion
- Color Dodge
- Linear Burn
- Linear Light
- Substract
- Opacity
- Darken
- Lighten

Effects:
- Emboss
- Sobel
- Median
- Grayscale
- Edge Detection
- Invert

Blurs:
- Gaussian
- Box

Adjustment:
- Brightness

Transform:
- FlipH
- FlipV

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
