package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	h, w int
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x ^ y), uint8((x + y) / 2), 255}
}

func main() {
	m := &Image{64, 64}
	pic.ShowImage(m)
}
