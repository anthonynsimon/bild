package rectangle

import (
	"errors"
	"image"
	"image/color"

	"github.com/anthonynsimon/bild/parallel"
)

//AddRectangle adds rectangle to the passes image
//it uses image to operate on
//definations of arguments are :
//	 img  : input image
//   xLeft : x value of  top left most corner of rectangle to be drawn
//   yLeft : y value of top left most corner of rectangle to be drawn
//   xRight : x value of bottom right most corner of rectangle to be drawn
//   yRight : y value of bottom right  most corner of rectangle to be drawn
// 	 c : color for the inner rectangle
func AddRectangle(img *image.RGBA, xLeft, yLeft, xRight, yRight int, c color.Color) error {

	height, width := img.Rect.Dy(), img.Rect.Dx()
	if xRight > width || yRight > height {
		return errors.New("bottom right corner values exceed the limit")
	} else if xLeft < 0 || yLeft < 0 {
		return errors.New("top left corner values preceed the limit")
	} else if xLeft > xRight || yLeft < yRight {
		return errors.New("values must be in correct order")
	}
	height, width = yLeft-yRight, xRight-xLeft
	parallel.Line(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				img.Set(x+xLeft, y+yRight, c)
			}
		}
	})

	return nil
}
