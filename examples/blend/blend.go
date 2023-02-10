package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/esimov/gomp"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	in, err := os.Open("sample.png")
	if err != nil {
		log.Fatalf("cannot open the source file: %s", err)
	}

	src, err := png.Decode(in)
	if err != nil {
		log.Fatalf("cannot decode the source image: %s", err)
	}

	srcImg := gomp.ImgToNRGBA(src)

	bgr := image.NewNRGBA(src.Bounds())
	col := color.RGBA{R: 0xf4, G: 0x7a, B: 0x03, A: 0xff}
	draw.Draw(bgr, bgr.Bounds(), &image.Uniform{col}, image.Point{}, draw.Src)

	imop := gomp.InitOp()
	blop := gomp.NewBlend()

	dc := gg.NewContext(1024, 1024)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	dc.DrawRectangle(0, 0, 1024, 1024)
	dc.Fill()

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("cannot parse the font: %s", err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 20,
	})
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SetFontFace(face)

	gridX, gridY := 0, 0
	size := 256
	cellSize := 32

	for _, op := range blop.Modes {
		if gridX == size*4 {
			gridY += size
			gridX = 0
		}

		var i, j int
		for x := gridX; x < gridX+size; x += cellSize {
			for y := gridY; y < gridY+size; y += cellSize {
				if (i+j)%2 == 0 {
					dc.SetHexColor("dedede")
				} else {
					dc.SetHexColor("f3f3f3")
				}
				dc.DrawRectangle(float64(x), float64(y), float64(cellSize), float64(cellSize))
				dc.Fill()
				j++
			}
			i++
		}

		blop.Set(op)
		imop.Set(gomp.SrcOver)
		bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
		imop.Draw(bmp, srcImg, bgr, blop)

		dx, _ := dc.MeasureString(op)
		dc.DrawImage(bmp.Img, gridX, gridY)
		dc.DrawRectangle(float64(gridX), float64(gridY), float64(gridX+size), float64(gridY+size))
		dc.SetRGB(0.7, 0.7, 0.7)
		dc.Stroke()

		dc.SetRGB(1, 1, 1)
		dc.Stroke()
		opName := strings.ReplaceAll(op, "_", " ")
		dc.DrawString(opName, float64(gridX)+(float64(size)/2-dx/2), float64(gridY-5+size))

		gridX += size
	}

	finalImg := dc.Image()
	output, _ := os.Create("blend.png")
	png.Encode(output, finalImg)
}
