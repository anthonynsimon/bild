package cmd

import (
	"image"
	"log"

	"github.com/anthonynsimon/bild/imgio"
)

func apply(fin, fout string, process func(image.Image) (image.Image, error)) {
	in, err := imgio.Open(fin)
	if err != nil {
		log.Fatal(err)
	}

	result, err := process(in)
	if err != nil {
		log.Fatal(err)
	}

	err = imgio.Save(fout, result, imgio.PNGEncoder())
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

	err = imgio.Save(fout, result, imgio.PNGEncoder())
	if err != nil {
		log.Fatal(err)
	}
}
