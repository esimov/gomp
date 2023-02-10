package gomp

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComp_Basic(t *testing.T) {
	assert := assert.New(t)

	op := InitOp()
	err := op.Set("unsupported_composite_operation")
	assert.Error(err)

	op.Set(Clear)
	assert.Equal(Clear, op.Get())
	assert.NotEqual("unsupported_composite_operation", op.Get())

	op.Set(Dst)
	assert.Equal(Dst, op.Get())
}

func TestComp_Draw(t *testing.T) {
	assert := assert.New(t)
	imop := InitOp()

	transparent := color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	cyan := color.NRGBA{R: 33, G: 150, B: 243, A: 255}
	magenta := color.NRGBA{R: 233, G: 30, B: 99, A: 255}

	rect := image.Rect(0, 0, 10, 10)
	bmp := NewBitmap(rect)
	assert.Greater(bmp.Img.Bounds().Max.X, 0)
	assert.Greater(bmp.Img.Bounds().Max.Y, 0)
	source := image.NewNRGBA(rect)
	backdrop := image.NewNRGBA(rect)

	// No composition operation applied. The SrcOver is the default one.
	draw.Draw(source, image.Rect(0, 4, 6, 10), &image.Uniform{cyan}, image.Point{}, draw.Src)
	draw.Draw(backdrop, image.Rect(4, 0, 10, 6), &image.Uniform{magenta}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)

	// Pick three representative points/pixels from the generated image output.
	// Depending on the applied composition operation the colors of the
	// selected pixels should be the source color, the destination color or transparent.
	topRight := bmp.Img.At(9, 0)
	bottomLeft := bmp.Img.At(0, 9)
	center := bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, cyan)

	// Clear
	imop.Set(Clear)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, transparent)
	assert.EqualValues(center, transparent)

	// Copy
	imop.Set(Copy)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, cyan)

	// Dst
	imop.Set(Dst)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, transparent)
	assert.EqualValues(center, magenta)

	// SrcOver
	imop.Set(SrcOver)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, cyan)

	// DstOver
	imop.Set(DstOver)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, magenta)

	// SrcIn
	imop.Set(SrcIn)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, transparent)
	assert.EqualValues(center, cyan)

	// DstIn
	imop.Set(DstIn)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, transparent)
	assert.EqualValues(center, magenta)

	// SrcOut
	imop.Set(SrcOut)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, transparent)

	// DstOut
	imop.Set(DstOut)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, transparent)
	assert.EqualValues(center, transparent)

	// SrcAtop
	imop.Set(SrcAtop)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, transparent)
	assert.EqualValues(center, cyan)

	// DstAtop
	imop.Set(DstAtop)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, magenta)

	// Xor
	imop.Set(Xor)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, magenta)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, transparent)
	// DstAtop
	imop.Set(DstAtop)
	imop.Draw(bmp, source, backdrop, nil)

	topRight = bmp.Img.At(9, 0)
	bottomLeft = bmp.Img.At(0, 9)
	center = bmp.Img.At(5, 5)

	assert.EqualValues(topRight, transparent)
	assert.EqualValues(bottomLeft, cyan)
	assert.EqualValues(center, magenta)
}
