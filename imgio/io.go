/*Package imgio provides basic image file input/output.*/
package imgio

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/bmp"
)

// Encoder encodes the provided image and writes it
type Encoder func(io.Writer, image.Image) error

// Open loads and decodes an image from a file and returns it.
//
// Usage example:
//
//	// Decodes an image from a file with the given filename
//	// returns an error if something went wrong
//	img, err := Open("exampleName")
func Open(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// JPEGEncoder returns an encoder to JPEG given the argument 'quality'
func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}

// PNGEncoder returns an encoder to PNG
func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}

// BMPEncoder returns an encoder to BMP
func BMPEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return bmp.Encode(w, img)
	}
}

// Save creates a file and writes to it an image using the provided encoder.
//
// Usage example:
//
//	// Save an image to a file in PNG format,
//	// returns an error if something went wrong
//	err := Save("exampleName", img, imgio.JPEGEncoder(100))
func Save(filename string, img image.Image, encoder Encoder) error {
	// filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return encoder(f, img)
}
