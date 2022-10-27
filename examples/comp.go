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
	in, err := os.Open("src.png")
	if err != nil {
		log.Fatalf("cannot open the source file: %s", err)
	}

	out, err := os.Open("dst.png")
	if err != nil {
		log.Fatalf("cannot open the destination file: %s", err)
	}

	src, err := png.Decode(in)
	if err != nil {
		log.Fatalf("cannot decode the source image: %s", err)
	}

	dst, err := png.Decode(out)
	if err != nil {
		log.Fatalf("cannot decode the destination image: %s", err)
	}
	srcImg := gomp.ImgToNRGBA(src)
	dstImg := gomp.ImgToNRGBA(dst)

	imop := gomp.InitOp()
	dc := gg.NewContext(1024, 768)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	dc.DrawRectangle(0, 0, 1024, 768)
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

	for _, op := range imop.Ops {
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

		imop.Set(op)
		bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
		imop.Draw(bmp, srcImg, dstImg, nil)

		strw, _ := dc.MeasureString(op)
		dc.DrawImage(bmp.Img, gridX, gridY)
		dc.DrawRectangle(float64(gridX), float64(gridY), float64(gridX+size), float64(gridY+size))
		dc.SetRGB(0.6, 0.6, 0.6)
		dc.Stroke()

		dc.SetRGB(0.2, 0.2, 0.2)
		dc.Stroke()
		dc.DrawString(op, float64(gridX)+(float64(size)/2-strw/2), float64(gridY-2+size))

		gridX += size
	}

	finalImg := dc.Image()
	output, _ := os.Create("output.png")
	png.Encode(output, finalImg)
}
