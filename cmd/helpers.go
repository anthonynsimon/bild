package cmd

import (
	"image"
	"log"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
)

var jpgExtensions = []string{".jpg", ".jpeg"}
var pngExtensions = []string{".png"}

func resolveEncoder(outputfile string, defaultEncoding imgio.Encoder) imgio.Encoder {
	lower := strings.ToLower(outputfile)

	for _, ext := range jpgExtensions {
		if strings.HasSuffix(lower, ext) {
			return imgio.JPEGEncoder(100)
		}
	}

	for _, ext := range pngExtensions {
		if strings.HasSuffix(lower, ext) {
			return imgio.PNGEncoder()
		}
	}

	return defaultEncoding
}

func apply(fin, fout string, process func(image.Image) (image.Image, error)) {
	in, err := imgio.Open(fin)
	if err != nil {
		log.Fatal(err)
	}

	result, err := process(in)
	if err != nil {
		log.Fatal(err)
	}

	encoder := resolveEncoder(fout, imgio.PNGEncoder())
	err = imgio.Save(fout, result, encoder)
	if err != nil {
		log.Fatal(err)
	}
}

func apply2(fin1, fin2, fout string, process func(image.Image, image.Image) (image.Image, error)) {
	in1, err := imgio.Open(fin1)
	if err != nil {
		log.Fatal(err)
	}

	in2, err := imgio.Open(fin2)
	if err != nil {
		log.Fatal(err)
	}

	result, err := process(in1, in2)
	if err != nil {
		log.Fatal(err)
	}

	encoder := resolveEncoder(fout, imgio.PNGEncoder())
	err = imgio.Save(fout, result, encoder)
	if err != nil {
		log.Fatal(err)
	}
}
