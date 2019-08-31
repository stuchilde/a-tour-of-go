/*
	Images
	Package image defines the Image interface:

	package image

	type Image interface {
			ColorModel() color.Model
			Bounds() Rectangle
			At(x, y int) color.Color
	}
	Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

	(See the documentation for all the details.)

	The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package
*/

package main

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x ^ y), uint8((x + y) / 2), 255}
}
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(2, 8).RGBA())
}
