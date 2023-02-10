# gomp
![Coverage](https://img.shields.io/badge/Coverage-73.3%25-brightgreen)
[![CI](https://github.com/esimov/gomp/actions/workflows/ci.yml/badge.svg)](https://github.com/esimov/gomp/actions/workflows/ci.yml)
[![go.dev reference](https://img.shields.io/badge/pkg.go.dev-reference-007d9c?logo=go)](https://pkg.go.dev/github.com/esimov/gomp)
[![release](https://img.shields.io/badge/release-v1.0.1-blue.svg)](https://github.com/esimov/gomp/releases/tag/v1.0.1)
[![license](https://img.shields.io/github/license/esimov/gomp)](./LICENSE)

Go library for image blending and alpha compositing using advanced features like the Porter-Duff operator and blending modes.

## About

The reason why this package has been developed is because the [**`image/draw`**](https://pkg.go.dev/image/draw) package from Go's standard library defines only one operation: drawing a source image onto the destination image, through an optional image mask. This is performed pixel by pixel and it's based on the classic "[Compositing Digital Images](https://dl.acm.org/doi/pdf/10.1145/964965.808606)" paper by Porter and Duff. This paper presented **12 different** composition operation, but the Draw method uses only two of them: `source over destination` and `source`. 

When dealing with image composition this is simply not enough. This library aims to overcome this deficiency by integrating the missing operators.

| Alpha compositing
|:--:
| ![compositing](https://github.com/esimov/gomp/blob/master/examples/comp/composite.png) |

### Blending modes
For convenience, this package implements also some of the most used blending modes in Photoshop. Similarly to the alpha compositing, blending modes defines the result of compositing a source and a destination but without being constrained to the alpha channel. The implementation follows the blending formulas presented in the W3C document: [Compositing and Blending](https://www.w3.org/TR/compositing-1/#blending). These blending modes are not covered by Porter and Duff, but have been included into this package for convenience.

| Blending modes
|:--:
| ![blending](https://github.com/esimov/gomp/blob/master/examples/blend/blend.png) |

## Installation
```bash
$ go install github.com/esimov/gomp@latest
```

## API
The API of the library is inspired by the [PorterDuff.Mode](https://developer.android.com/reference/android/graphics/PorterDuff.Mode) class from the Android SDK.

### Alpha compositing
```go
op := gomp.InitOp()
op.Set(gomp.SrcOver)
bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
imop.Draw(bmp, source, backdrop, nil)
```

### Blending modes
```go
op := gomp.NewBlend()
op.Set(gomp.Multiply)
bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
imop.Draw(bmp, source, backdrop, op)
```

You can combine the alpha compositing with blending modes at the same step, you just need to replace the last parameter of the `Draw` method with the initialized blending operation.

### Alpha compositing and blending modes combined
```go
imop := gomp.InitOp()
blop := gomp.NewBlend()
blop.Set(gomp.Multiply)
imop.Set(gomp.SrcOver)
bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
imop.Draw(bmp, srcImg, bgr, blop)
```

### Operators
      
| Image compositing | Separable blending modes | Non-separable blending modes 
|:--:|:--:|:--:
| `Clear` | `Normal` | `Hue` |
| `Copy` | `Darken` | `Saturation` |
| `Dst` | `Lighten` | `ColorMode` |
| `SrcOver` | `Multiply` | `Luminosity` |
| `DstOver` | `Screen` |
| `SrcIn` | `Overlay` |
| `DstIn` | `SoftLight` |
| `SrcOut` | `HardLight` |
| `DstOut` | `ColorDodge` |
| `SrcAtop` | `ColorBurn` |
| `DstAtop` | `Difference` |
|  | `Exclusion` |

### Examples
The images used in this document for visualizing the alpha compositing operation and the blending modes have been generated using this library. They can be found in the [examples](https://github.com/esimov/gomp/tree/master/examples) folder.

## Author
* Endre Simo ([@simo_endre](https://twitter.com/simo_endre))

## License
Copyright © 2022 Endre Simo

This software is distributed under the MIT license. See the [LICENSE](https://github.com/esimov/gomp/blob/master/LICENSE) file for the full license text.
