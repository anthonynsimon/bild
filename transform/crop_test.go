package transform

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

type cropReadCounter struct {
	image.Image
	reads int
}

func (img *cropReadCounter) At(x, y int) color.Color {
	img.reads++
	return img.Image.At(x, y)
}

func TestCropReadsOnlyIntersection(t *testing.T) {
	cases := []struct {
		name string
		rect image.Rectangle
	}{
		{"inside", image.Rect(2, 3, 4, 5)},
		{"partial", image.Rect(-1, -1, 2, 2)},
		{"outside", image.Rect(10, 10, 12, 12)},
		{"empty", image.Rect(2, 2, 2, 4)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			src := &cropReadCounter{Image: image.NewNRGBA(image.Rect(0, 0, 6, 6))}
			bounds := src.Bounds().Intersect(c.rect)
			result := Crop(src, c.rect)
			if result.Bounds() != bounds {
				t.Errorf("expected bounds %v, actual %v", bounds, result.Bounds())
			}
			if expected := bounds.Dx() * bounds.Dy(); src.reads != expected {
				t.Errorf("expected %d pixel reads, actual %d", expected, src.reads)
			}
		})
	}
}

func TestCropImageTypes(t *testing.T) {
	bounds := image.Rect(-2, 3, 4, 8)
	rgba := image.NewRGBA(bounds)
	nrgba := image.NewNRGBA(bounds)
	gray := image.NewGray(bounds)
	nrgba64 := image.NewNRGBA64(bounds)
	cmyk := image.NewCMYK(bounds)
	paletted := image.NewPaletted(bounds, color.Palette{color.Transparent, color.Black, color.White, color.NRGBA{R: 200, B: 100, A: 128}})
	ycbcr := image.NewYCbCr(bounds, image.YCbCrSubsampleRatio420)
	for i := range ycbcr.Y {
		ycbcr.Y[i] = uint8(i * 7)
	}
	for i := range ycbcr.Cb {
		ycbcr.Cb[i] = uint8(i * 13)
		ycbcr.Cr[i] = uint8(i * 19)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.NRGBA{R: uint8((x + 2) * 40), G: uint8(y * 20), B: 200, A: uint8((x + 2) * 50)}
			rgba.Set(x, y, c)
			nrgba.Set(x, y, c)
			gray.Set(x, y, c)
			cmyk.Set(x, y, c)
			paletted.Set(x, y, c)
			nrgba64.SetNRGBA64(x, y, color.NRGBA64{R: uint16(x+2) * 10001, G: uint16(y) * 7001, B: 50003, A: uint16(x+2) * 13001})
		}
	}
	cases := []struct {
		name string
		img  image.Image
	}{
		{"RGBA", rgba},
		{"NRGBA", nrgba},
		{"Gray", gray},
		{"NRGBA64", nrgba64},
		{"CMYK", cmyk},
		{"Paletted", paletted},
		{"YCbCr", ycbcr},
		{"RGBA subimage", rgba.SubImage(image.Rect(-1, 4, 4, 7))},
		{"NRGBA subimage", nrgba.SubImage(image.Rect(-1, 4, 4, 7))},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rect := image.Rect(0, 5, 5, 9).Intersect(c.img.Bounds())
			expected := image.NewRGBA(rect)
			for y := rect.Min.Y; y < rect.Max.Y; y++ {
				for x := rect.Min.X; x < rect.Max.X; x++ {
					expected.Set(x, y, c.img.At(x, y))
				}
			}
			result := Crop(c.img, rect)
			if !util.RGBAImageEqual(result, expected) {
				t.Fatalf("expected %#v, actual %#v", expected, result)
			}
			draw.Draw(result, result.Bounds(), image.Transparent, image.Point{}, draw.Src)
			if after := Crop(c.img, rect); !util.RGBAImageEqual(after, expected) {
				t.Error("modifying the crop changed the source image")
			}
		})
	}
}

func BenchmarkCrop(b *testing.B) {
	bounds := image.Rect(0, 0, 2048, 2048)
	cases := []struct {
		name string
		img  image.Image
	}{
		{"RGBA", image.NewRGBA(bounds)},
		{"NRGBA", image.NewNRGBA(bounds)},
		{"YCbCr", image.NewYCbCr(bounds, image.YCbCrSubsampleRatio420)},
	}
	for _, c := range cases {
		for _, size := range []struct {
			name string
			rect image.Rectangle
		}{
			{"small", image.Rect(100, 100, 150, 150)},
			{"full", bounds},
		} {
			b.Run(c.name+"/"+size.name, func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					Crop(c.img, size.rect)
				}
			})
		}
	}
}
