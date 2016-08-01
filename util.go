package bild

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Format is used to identify the image encoding type
type Format int

const (
	JPEG = iota
	PNG
)

// Open loads and decodes an image from a file
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

// Encode writes an image in the specified format
func Encode(w io.Writer, img image.Image, format Format) error {
	var err error

	switch format {
	case PNG:
		err = png.Encode(w, img)
	case JPEG:
		err = jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	}

	return err
}

// Save creates a file and writes to it an image in the specified format
func Save(filename string, img image.Image, format Format) error {
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	switch format {
	case PNG:
		filename += ".png"
	case JPEG:
		filename += ".jpg"
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return Encode(f, img, format)
}
