# Bild
Simple image processing in Go with parallel processing support

Basic example:
```go
package main

import "github.com/anthonynsimon/bild"

func main() {
	img, err := bild.Open("filename")
	if err != nil {
		panic(err)
	}

	result := bild.Invert(img)

	if err := bild.Save("filename", result, bild.PNG); err != nil {
		panic(err)
	}
}

```
