package main

import (
	"image"

	"github.com/anthonynsimon/bild"

	"os"

	"path/filepath"
)

// InDir is the relative directory which will be searched for input files
const InDir = "in"

// OutDir is the relative directory where the output files will be written
const OutDir = "out"

func main() {
	dir, err := os.Open(InDir)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		loadApplySave(file.Name(), bild.Box)
	}
}

func loadApplySave(filename string, fn func(image.Image) *image.NRGBA) {
	img, err := bild.Open(filepath.Join(InDir, filename))
	if err != nil {
		panic(err)
	}

	result := fn(img)

	if err := bild.Save(filepath.Join(OutDir, filename), result, bild.PNG); err != nil {
		panic(err)
	}
}
