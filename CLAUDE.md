# CLAUDE.md

Pure Go image processing library (`github.com/anthonynsimon/bild`) with a built-in CLI. Minimal dependencies — just cobra for CLI and `golang.org/x/image` for BMP support.

## Build & Test

```bash
make build    # Build binary to dist/
make test     # Run all tests (60s timeout, verbose)
make fmt      # Format code (go fmt)
make race     # Run tests with race detector
make bench    # Run benchmarks
```

CI runs `make build && make test` across Go 1.21–1.26 on Ubuntu and macOS. Arch differences can cause issues in floating point values, but it's generally acceptable as long as it's effectively imperceptible to the human eye.

## Project Structure

Each image processing domain is its own importable package:

- `adjust/` — brightness, contrast, gamma, hue, saturation
- `blend/` — 16 blending modes (multiply, overlay, screen, etc.)
- `blur/` — box and gaussian blur
- `channel/` — RGBA channel extraction
- `clone/` — image cloning and padding
- `convolution/` — kernel convolution operations
- `effect/` — grayscale, sepia, sharpen, sobel, invert, median, erode, dilate, emboss, edge detection
- `histogram/` — histogram analysis and visualization
- `imgio/` — image I/O (PNG, JPEG, BMP)
- `noise/` — noise generation (uniform, binary, gaussian, perlin)
- `paint/` — flood fill
- `segment/` — thresholding
- `transform/` — rotate, resize (7 resampling filters), crop, flip, translate, shear, zoom

Supporting packages: `parallel/` (goroutine work splitting), `fcolor/` (float64 RGBA), `math/` (clamp, min/max), `util/` (color conversions, comparison helpers), `perlin/`.

## CLI

- `main.go` → `cmd.Execute()`
- `cmd/root.go` — registers all subcommands via `init()`
- `cmd/helpers.go` — `apply()` single-image pipeline, `apply2()` two-image pipeline, parsers, encoder resolution
- `cmd/<category>.go` — one file per command group, mirrors the library packages

## Patterns

- All processing functions take `image.Image` and return `*image.RGBA`
- Pixel-level parallelism via `parallel.Line()` which splits rows across CPUs
- Tests are table-driven with manual pixel data; comparison via `util.RGBAImageEqual()` / `util.RGBAImageApproxEqual()`

## When Adding Features

- New library functions go in the appropriate domain package
- New CLI commands follow the pattern in existing `cmd/` files
- Update `CHANGELOG.md` and `README.md`
