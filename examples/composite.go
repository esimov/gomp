package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/esimov/gomp"
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
	finalImg := image.NewNRGBA(image.Rect(0, 0, 1024, 768))

	gridX, gridY := 0, 0
	size := 256
	for _, op := range imop.Ops {
		imop.Set(op)
		bmp := gomp.NewBitmap(image.Rect(0, 0, size, size))
		imop.Draw(bmp, srcImg, dstImg, nil)

		if gridX == size*4 {
			gridY += size
			gridX = 0
		}

		draw.Draw(finalImg, image.Rectangle{image.Pt(gridX, gridY), image.Pt(gridX+size, gridY+size)}, bmp.Img, image.Point{}, draw.Src)
		gridX += size
	}

	output, _ := os.Create("output.png")
	png.Encode(output, finalImg)
}
