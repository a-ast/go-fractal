package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func SavePng(filename string, image image.Image) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	png.Encode(f, image)
}

type SpaceItem struct {
	X, Y  int
	value float32
}

type Space struct {
	Width, Height int
	items         []SpaceItem
}

func NewSpace(width, height int) Space {
	items := make([]SpaceItem, 0)

	return Space{width, height, items}
}

func (space *Space) AddItem(item SpaceItem) {
	space.items = append(space.items, item)
}

func main() {

	width := 1640
	height := 1480

	space := NewSpace(width, height)

	config := Config{5.0, -1.0, -1.0, 1.0, 1.0, 200, -0.8, 0.156}
	drawJuliaSet(config, &space)

	palette := NewGradientPalette(1000,
		Colour{200, 174, 0}, Colour{0, 0, 0},
	)

	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, item := range space.items {
		paletteColour := palette.GetColor(int(1000 * item.value))
		colour := color.RGBA{uint8(paletteColour.R), uint8(paletteColour.G), uint8(paletteColour.B), 255}
		image.Set(item.X, item.Y, colour)
	}

	SavePng("img/fractal.png", image)
	fmt.Println("Finished!")
}
