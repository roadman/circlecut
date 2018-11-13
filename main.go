package main

import (
	"image"
	"os"

	"image/draw"
	"image/png"
)

func main() {
	origImagePath := "fgosyuten.png"
	cutFrameImagePath := "template/B.png"
	cutImagePath := "cut.png"

	origFile, err := os.Open(origImagePath)
	if err != nil {
		panic(err)
	}
	defer origFile.Close()

	cutFrameFile, err := os.Open(cutFrameImagePath)
	if err != nil {
		panic(err)
	}
	defer cutFrameFile.Close()

	origImg, _, err := image.Decode(origFile)
	if err != nil {
		panic(err)
	}
	cutFrameImg, _, err := image.Decode(cutFrameFile)
	if err != nil {
		panic(err)
	}

	// 合成は左上固定で、最大サイズはcutのサイズ
	// 画像がcutより小さい場合を考慮していない
	startPoint := image.Point{0, 0}
	cutFrameRectangle := image.Rectangle{startPoint, startPoint.Add(cutFrameImg.Bounds().Size())}
	origRectangle := image.Rectangle{startPoint, startPoint.Add(cutFrameImg.Bounds().Size())}

	rgba := image.NewRGBA(origRectangle)
	draw.Draw(rgba, origRectangle, origImg, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, cutFrameRectangle, cutFrameImg, image.Point{0, 0}, draw.Over)

	cutFile, err := os.Create(cutImagePath)
	if err != nil {
		panic(err)
	}
	defer cutFile.Close()

	if err := png.Encode(cutFile, rgba); err != nil {
		panic(err)
	}
}
