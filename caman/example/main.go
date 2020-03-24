package main

import (
	"fmt"
	_ "image/png"
	"path/filepath"
	"strings"

	_ "image/jpeg"

	"github.com/anthonynsimon/bild/caman"
	"github.com/anthonynsimon/bild/imgio"
)

func main() {

	images := []string{
		"./anne.jpg",
	}

	for _, filename := range images {
		img, err := imgio.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		filenamePrefix := strings.Split(filepath.Base(filename), ".")[0]

		/*  ---------------Caman Sepia---------------*/
		sepia := caman.Sepia(img, 50)
		saveFilename := fmt.Sprintf("sepia_%s.jpg", filenamePrefix)
		if err := imgio.Save(saveFilename, sepia, imgio.JPEGEncoder(80)); err != nil {
			fmt.Println(err)
		}

		/*  ---------------Caman Channels---------------*/
		channels := caman.Channels(img, map[string]float64{"red": 8, "blue": 5})
		saveFilename = fmt.Sprintf("channels_%s.jpg", filenamePrefix)
		if err := imgio.Save(saveFilename, channels, imgio.JPEGEncoder(80)); err != nil {
			fmt.Println(err)
		}

		/*  ---------------Caman Curves---------------*/
		curves := caman.Curves(img, "r", [2]float64{20, 0}, [2]float64{90, 120}, [2]float64{186, 144}, [2]float64{255, 230})
		curves = caman.Curves(curves, "rgb", [2]float64{20, 0}, [2]float64{90, 120}, [2]float64{186, 144}, [2]float64{255, 230})
		saveFilename = fmt.Sprintf("curves_%s.jpg", filenamePrefix)
		if err := imgio.Save(saveFilename, curves, imgio.JPEGEncoder(80)); err != nil {
			fmt.Println(err)
		}

		/*  ---------------Caman Colorize---------------*/
		colorize := caman.Colorize(img, &caman.RGB{R: 255, G: 105, B: 59}, 10)
		saveFilename = fmt.Sprintf("colorize_%s.jpg", filenamePrefix)
		if err := imgio.Save(saveFilename, colorize, imgio.JPEGEncoder(80)); err != nil {
			fmt.Println(err)
		}
	}
}
