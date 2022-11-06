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

	frontground := Color{R: 0xff, G: 0xff, B: 0xff}
	background := Color{R: 0, G: 0, B: 0}

	assert.Equal(0.0, op.Sat(frontground))
	sat := op.SetSat(background, op.Sat(frontground))
	assert.Equal(Color{R: 0, G: 0, B: 0}, sat)
}

func TestBlend_Modes(t *testing.T) {
	assert := assert.New(t)
	imop := InitOp()
	blop := NewBlend()

	pinkFront := color.RGBA{R: 214, G: 20, B: 65, A: 255}
	orangeBack := color.RGBA{R: 250, G: 121, B: 17, A: 255}

	rect := image.Rect(0, 0, 1, 1)
	bmp := NewBitmap(rect)
	src := image.NewNRGBA(rect)
	dst := image.NewNRGBA(rect)

	imop.Set(SrcOver)
	// Darken
	blop.Set(Darken)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp := []uint8{214, 20, 17, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// Multiply
	blop.Set(Multiply)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{209, 9, 4, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// Screen
	blop.Set(Screen)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{254, 131, 77, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// Overlay
	blop.Set(Overlay)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{253, 18, 8, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// SoftLight
	blop.Set(SoftLight)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{232, 19, 23, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// HardLight
	blop.Set(HardLight)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{251, 67, 9, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// ColorDodge
	blop.Set(ColorDodge)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{255, 131, 22, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// ColorBurn
	blop.Set(ColorBurn)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{249, 0, 0, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// Difference
	blop.Set(Difference)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{35, 101, 48, 255}
	assert.EqualValues(exp, bmp.Img.Pix)

	// Exclusion
	blop.Set(Exclusion)
	draw.Draw(src, rect, &image.Uniform{pinkFront}, image.Point{}, draw.Src)
	draw.Draw(dst, rect, &image.Uniform{orangeBack}, image.Point{}, draw.Src)
	imop.Draw(bmp, src, dst, blop)

	exp = []uint8{44, 122, 73, 255}
	assert.EqualValues(exp, bmp.Img.Pix)
}

func TestBlend_NonSeparableModes(t *testing.T) {
	assert := assert.New(t)
	op := NewBlend()

	pinkFront := Color{R: 214, G: 20, B: 65}
	orangeBack := Color{R: 250, G: 121, B: 17}

	rsn, gsn, bsn, asn := pinkFront.R/255, pinkFront.G/255, pinkFront.G/255, 1.0
	rbn, gbn, bbn, abn := orangeBack.R/255, orangeBack.G/255, orangeBack.G/255, 1.0

	// Hue
	sat := op.SetSat(orangeBack, op.Sat(pinkFront))
	rgb := op.SetLum(sat, op.Lum(pinkFront))
	rn, gn, bn := alphaCompose(op, rsn, gsn, bsn, asn, rbn, gbn, bbn, abn, rgb)

	exp := []uint32{20, 99, 163}
	actual := []uint32{rn, gn, bn}
	assert.EqualValues(exp, actual)

	// Saturation
	sat = op.SetSat(pinkFront, op.Sat(orangeBack))
	rgb = op.SetLum(sat, op.Lum(pinkFront))
	rn, gn, bn = alphaCompose(op, rsn, gsn, bsn, asn, rbn, gbn, bbn, abn, rgb)

	exp = []uint32{1, 122, 94}
	actual = []uint32{rn, gn, bn}
	assert.EqualValues(exp, actual)

	// Color
	rgb = op.SetLum(orangeBack, op.Lum(pinkFront))
	rn, gn, bn = alphaCompose(op, rsn, gsn, bsn, asn, rbn, gbn, bbn, abn, rgb)

	exp = []uint32{31, 97, 150}
	actual = []uint32{rn, gn, bn}
	assert.EqualValues(exp, actual)

	// Luminosity
	rgb = op.SetLum(pinkFront, op.Lum(orangeBack))
	rn, gn, bn = alphaCompose(op, rsn, gsn, bsn, asn, rbn, gbn, bbn, abn, rgb)

	exp = []uint32{1, 219, 168}
	actual = []uint32{rn, gn, bn}
	assert.EqualValues(exp, actual)
}

func alphaCompose(op *Blend, rsn, gsn, bsn, asn, rbn, gbn, bbn, abn float64, rgb Color) (uint32, uint32, uint32) {
	a := asn + abn - asn*abn
	rn := op.AlphaCompose(abn, asn, a, rbn*255, rsn*255, rgb.R*255)
	gn := op.AlphaCompose(abn, asn, a, gbn*255, gsn*255, rgb.G*255)
	bn := op.AlphaCompose(abn, asn, a, bbn*255, bsn*255, rgb.B*255)
	rn, gn, bn = rn/255, gn/255, bn/255

	return uint32(rn), uint32(gn), uint32(bn)
}
