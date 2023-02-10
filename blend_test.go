package gomp

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlend_Basic(t *testing.T) {
	assert := assert.New(t)

	op := NewBlend()
	assert.Empty(op.Get())
	err := op.Set("blend_mode_not_supported")
	assert.Error(err)
	op.Set(Darken)
	assert.Equal(Darken, op.Get())
	op.Set(Lighten)
	assert.Equal(Lighten, op.Get())

	rgb := Color{R: 0xff, G: 0xff, B: 0xff}
	lum := op.Lum(rgb)
	assert.Equal(255.0, lum)

	rgb = Color{R: 0, G: 0, B: 0}
	lum = op.Lum(rgb)
	assert.Equal(0.0, lum)

	rgb = Color{R: 127, G: 127, B: 127}
	lum = op.Lum(rgb)
	assert.Equal(127.0, lum)

	foreground := Color{R: 0xff, G: 0xff, B: 0xff}
	background := Color{R: 0, G: 0, B: 0}

	assert.Equal(0.0, op.Sat(foreground))
	sat := op.SetSat(background, op.Sat(foreground))
	assert.Equal(Color{R: 0, G: 0, B: 0}, sat)
}

func TestBlend_Modes(t *testing.T) {
	// Note: all the expected values are taken by using as reference the results
	// obtained in Photoshop by overlapping two layers and applying the blend mode.
	assert := assert.New(t)

	imop := InitOp()
	blop := NewBlend()

	pinkFront := color.RGBA{R: 214, G: 20, B: 65, A: 255}
	orangeBack := color.RGBA{R: 250, G: 121, B: 17, A: 255}

	rect := image.Rect(0, 0, 1, 1)
	bmp := NewBitmap(rect)
	source := image.NewNRGBA(rect)
	backdrop := image.NewNRGBA(rect)

	imop.Set(SrcOver)

	// Darken
	blop.Set(Darken)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected := []uint8{214, 20, 17, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Multiply
	blop.Set(Multiply)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{209, 9, 4, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Screen
	blop.Set(Screen)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{254, 131, 77, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Overlay
	blop.Set(Overlay)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{253, 18, 8, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// SoftLight
	blop.Set(SoftLight)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{232, 19, 23, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// HardLight
	blop.Set(HardLight)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{251, 67, 9, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// ColorDodge
	blop.Set(ColorDodge)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{255, 131, 22, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// ColorBurn
	blop.Set(ColorBurn)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{249, 0, 0, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Difference
	blop.Set(Difference)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{35, 101, 48, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Exclusion
	blop.Set(Exclusion)
	draw.Draw(source, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{44, 122, 73, 255}
	assert.EqualValues(expected, bmp.Img.Pix)
}

func TestBlend_NonSeparableModes(t *testing.T) {
	assert := assert.New(t)

	imop := InitOp()
	blop := NewBlend()

	frontColor := color.RGBA{R: 250, G: 121, B: 17, A: 255}
	backColor := color.RGBA{R: 214, G: 20, B: 65, A: 255}

	rect := image.Rect(0, 0, 1, 1)
	bmp := NewBitmap(rect)
	source := image.NewNRGBA(rect)
	backdrop := image.NewNRGBA(rect)

	imop.Set(SrcOver)

	// Hue
	blop.Set(Hue)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected := []uint8{255, 97, 133, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Saturation
	blop.Set(Saturation)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{233, 126, 39, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Color
	blop.Set(ColorMode)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{255, 97, 133, 255}
	assert.EqualValues(expected, bmp.Img.Pix)

	// Luminosity
	blop.Set(Luminosity)
	draw.Draw(source, rect, &image.Uniform{frontColor}, image.Point{}, draw.Src)
	draw.Draw(backdrop, rect, &image.Uniform{backColor}, image.Point{}, draw.Src)
	imop.Draw(bmp, source, backdrop, blop)

	expected = []uint8{148, 66, 0, 255}
	assert.EqualValues(expected, bmp.Img.Pix)
}
