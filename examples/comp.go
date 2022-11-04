package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/esimov/gomp"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	imop := gomp.InitOp()
	dc := gg.NewContext(1024, 768)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	dc.DrawRectangle(0, 0, 1024, 768)
	dc.Fill()

	// Source image
	src := gg.NewContext(256, 256)
	src.DrawRectangle(15, 85, 135, 135)
	src.SetHexColor("#2196f3")
	src.Fill()
	srcImg := gomp.ImgToNRGBA(src.Image())

	// Backdrop image
	bgr := gg.NewContext(256, 256)
	bgr.DrawCircle(165, 85, 75)
	bgr.SetHexColor("#e91e63")
	bgr.Fill()
	bdImg := gomp.ImgToNRGBA(bgr.Image())

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

	for _, op := range imop.Ops {
		if gridX == size*4 {
			gridY += size
			gridX = 0
		}

		var i, j int
		for x := gridX; x < gridX+size; x += cellSize {
			for y := gridY; y < gridY+size; y += cellSize {
				if (i+j)%2 == 0 {
					dc.SetHexColor("#dedede")
				} else {
					dc.SetHexColor("#f3f3f3")
				}
				dc.DrawRectangle(float64(x), float64(y), float64(cellSize), float64(cellSize))
				dc.Fill()
				j++
			}
			i++
		}

		imop.Set(op)
		bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
		imop.Draw(bmp, srcImg, bdImg, nil)

		strw, _ := dc.MeasureString(op)
		dc.DrawImage(bmp.Img, gridX, gridY)
		dc.DrawRectangle(float64(gridX), float64(gridY), float64(gridX+size), float64(gridY+size))
		dc.SetRGB(0.6, 0.6, 0.6)
		dc.Stroke()

		dc.SetRGB(0.2, 0.2, 0.2)
		dc.Stroke()
		dc.DrawString(op, float64(gridX)+(float64(size)/2-strw/2), float64(gridY-10+size))

		gridX += size
	}

	finalImg := dc.Image()
	output, _ := os.Create("out/composite.png")
	png.Encode(output, finalImg)
}
