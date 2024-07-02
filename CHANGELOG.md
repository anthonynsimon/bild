# Changelog

## next
- 

## 0.14.0
- [PR-100:](https://github.com/anthonynsimon/bild/pull/101) Update dependencies
- Update toolchain version

## 0.13.0
- [PR-85:](https://github.com/anthonynsimon/bild/pull/85) Up to 20% less allocations and 90% less bytes allocated.
- Minor documentation improvements.

## 0.12.0
- [PR-74:](https://github.com/anthonynsimon/bild/pull/74) Add Perlin noise function
- [PR-77:](https://github.com/anthonynsimon/bild/pull/77) Performance improvements for the image adjustment function
- [PR-81:](https://github.com/anthonynsimon/bild/pull/81) Make blend() exported to allow for custom blending implementations
- [PR-82:](https://github.com/anthonynsimon/bild/pull/82) Fix rotate panic
- Minor additional fixes and documentation improvements.

## 0.11.1
- [PR-71:](https://github.com/anthonynsimon/bild/pull/71) Gaussian blur is up to ~20x faster.

## 0.11.0
- bild now comes with a built-in CLI
- Added extract multiple channels functionality
- Minor fixes and performance improvements

## 0.10.0
- New feature effect.UnsharpMask
- Changed paint.FloodFill fuzz parameter to tolerance based. This is a - breaking change.


## 0.9.0
- New feature paint.FloodFill
- New feature transform.Translate

## 0.8.7
- Significant performance optimisations for Resize, Rotate, Convolve and Spatial Filtering functions. Most effects and blurs are indirectly benefited from this.

## 0.7.0
- New feature transform.Shear
- New feature adjust.Hue and adjust.Saturation
- New features effect.Dilate and effect.Erode

## 0.6.0
- New noise package, now you can generate Binary, Uniform and Gaussian noise (colored and monochrome).

## 0.5.0
- Major code refactor. Breaking changes as all APIs have been decentralised into sub-packages.

## 0.4.0
- Initial open source release.
- Release before major code refactor. Package bild contains all APIs in this release.
