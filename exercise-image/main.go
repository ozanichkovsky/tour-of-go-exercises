package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math/rand"
)

type Image struct{}

func (i Image) At(x, y int) color.Color {
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	return color.RGBA{r, g, b, 255}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

