package cmd

import (
	"errors"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

var jpgExtensions = []string{".jpg", ".jpeg"}
var pngExtensions = []string{".png"}
var bmpExtensions = []string{".bmp"}

var (
	// ErrWrongSize is thrown when the provided size string does not match the expected form.
	errWrongSize = errors.New("size must be of form [width]x[height], i.e. 400x200")
	// errWrongRect is thrown when the provided rect string does not match the expected form.
	errWrongRect = errors.New("rect must be of form [x0]x[y0]+[x1]x[y1], i.e. 0x0+512x256")
	// errUnknownFilter is thrown when an unknown resample filter name is provided.
	errUnknownFilter = errors.New("unknown filter, options: nearestneighbor, box, linear, gaussian, mitchellnetravali, catmullrom, lanczos")
)

type size struct {
	Width  int
	Height int
}

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

	for _, ext := range bmpExtensions {
		if strings.HasSuffix(lower, ext) {
			return imgio.BMPEncoder()
		}
	}

	return defaultEncoding
}

func apply(fin, fout string, process func(image.Image) (image.Image, error)) {
	in, err := imgio.Open(fin)
	exitIfNotNil(err)

	result, err := process(in)
	exitIfNotNil(err)

	encoder := resolveEncoder(fout, imgio.PNGEncoder())
	err = imgio.Save(fout, result, encoder)
	exitIfNotNil(err)
}

func apply2(fin1, fin2, fout string, process func(image.Image, image.Image) (image.Image, error)) {
	in1, err := imgio.Open(fin1)
	exitIfNotNil(err)

	in2, err := imgio.Open(fin2)
	exitIfNotNil(err)

	result, err := process(in1, in2)
	exitIfNotNil(err)

	encoder := resolveEncoder(fout, imgio.PNGEncoder())
	err = imgio.Save(fout, result, encoder)
	exitIfNotNil(err)
}

func exitIfNotNil(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseSizeStr(sizestr string) (*size, error) {
	parts := strings.Split(sizestr, "x")
	if len(parts) != 2 {
		return nil, errWrongSize
	}

	w, err := strconv.Atoi(parts[0])
	if err != nil || w < 0 {
		return nil, errWrongSize
	}

	h, err := strconv.Atoi(parts[1])
	if err != nil || h < 0 {
		return nil, errWrongSize
	}

	return &size{
		Width:  w,
		Height: h,
	}, nil
}

func parseRectStr(rectstr string) (image.Rectangle, error) {
	parts := strings.SplitN(rectstr, "+", 2)
	if len(parts) != 2 {
		return image.Rectangle{}, errWrongRect
	}

	min, err := parseSizeStr(parts[0])
	if err != nil {
		return image.Rectangle{}, errWrongRect
	}

	max, err := parseSizeStr(parts[1])
	if err != nil {
		return image.Rectangle{}, errWrongRect
	}

	return image.Rect(min.Width, min.Height, max.Width, max.Height), nil
}

func parseResampleFilter(name string) (transform.ResampleFilter, error) {
	switch strings.ToLower(name) {
	case "nearestneighbor":
		return transform.NearestNeighbor, nil
	case "box":
		return transform.Box, nil
	case "linear":
		return transform.Linear, nil
	case "gaussian":
		return transform.Gaussian, nil
	case "mitchellnetravali":
		return transform.MitchellNetravali, nil
	case "catmullrom":
		return transform.CatmullRom, nil
	case "lanczos":
		return transform.Lanczos, nil
	default:
		return transform.ResampleFilter{}, errUnknownFilter
	}
}
