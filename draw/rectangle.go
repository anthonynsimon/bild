package draw

import (
	"image"
	"image/color"
	"image/draw"
)

type Rectangle struct {
	TopLeft, BottomRight image.Point
	FillColor            color.Color
}

func NewRectangle(x0, y0, x1, y1 int, color color.Color) Rectangle {
	return Rectangle{
		TopLeft: image.Point{
			X: x0,
			Y: y0,
		},
		BottomRight: image.Point{
			X: x1,
			Y: y1,
		},
		FillColor: color,
	}
}

func FillRectangle(img image.Image, r Rectangle) *image.RGBA {
	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, img.Bounds(), img, image.Point{}, draw.Over)

	rec := image.Rect(r.TopLeft.X, r.TopLeft.Y, r.BottomRight.X, r.BottomRight.Y)
	add := image.Point{}.Add(image.Pt(r.TopLeft.X, r.TopLeft.Y))

	draw.Draw(dst, rec, &image.Uniform{C: r.FillColor}, add, draw.Over)
	return dst
}
