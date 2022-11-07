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

	op.Set(Dest)
	assert.Equal(Dest, op.Get())
}

func TestComp_Ops(t *testing.T) {
	assert := assert.New(t)
	imop := InitOp()

	frontColor := color.RGBA{R: 33, G: 150, B: 243, A: 255}
	backColor := color.RGBA{R: 233, G: 30, B: 99, A: 255}

	rect := image.Rect(0, 0, 1, 1)
	bmp := NewBitmap(rect)
	source := image.NewNRGBA(rect)
	backdrop := image.NewNRGBA(rect)

	// Copy
	imop.Set(Copy)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected := []uint8{33, 150, 243, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Dest
	imop.Set(Dest)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{233, 30, 99, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// SrcOver
	imop.Set(SrcOver)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{33, 150, 243, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// DstOver
	imop.Set(DstOver)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{233, 30, 99, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// SrcIn
	imop.Set(SrcIn)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{33, 150, 243, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// DstIn
	imop.Set(DstIn)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{233, 30, 99, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// SrcOut
	imop.Set(SrcOut)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{0, 0, 0, 0}
	assert.EqualValues(expected, bmp.Img.Pix)

	// DstOut
	imop.Set(DstOut)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{0, 0, 0, 0}
	assert.EqualValues(expected, bmp.Img.Pix)

	// SrcAtop
	imop.Set(SrcAtop)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{33, 150, 243, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// DstAtop
	imop.Set(DstAtop)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{233, 30, 99, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Xor
	imop.Set(Xor)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, nil)
	expected = []uint8{0, 0, 0, 0}
	assert.EqualValues(expected, bmp.Img.Pix)
}
