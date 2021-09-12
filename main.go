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

func SavePalette(filename string, picker Palette) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	image := image.NewRGBA(image.Rect(0, 0, len(picker.colors), 30))

	for i, item := range picker.colors {

		colour := color.RGBA{uint8(item.R), uint8(item.G), uint8(item.B), 255}
		for j := 0; j < 30; j++ {
			image.Set(i, j, colour)
		}
	}

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

	config := Config{3.0, -1.0, -1.0, 1.0, 1.0, 100, 0.8, 0}
	drawJuliaSet(config, &space)

	palette := NewGradientPalette(1000,
		Colour{0, 0, 0},
		Colour{0, 0, 0},
		[]GradientPoint{
			{5, Colour{16, 133, 139}},
			{15, Colour{255, 174, 0}},
		},
	)

	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, item := range space.items {
		paletteColour := palette.GetColor(int(1000 * item.value))

		colour := color.RGBA{uint8(paletteColour.R), uint8(paletteColour.G), uint8(paletteColour.B), 255}
		image.Set(item.X, item.Y, colour)
	}

	SavePalette("img/palette.png", palette)
	SavePng("img/fractal.png", image)
	fmt.Println("Finished!")
}
