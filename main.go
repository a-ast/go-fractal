package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/a-ast/go-fractal/colour_picker"
	"github.com/a-ast/go-fractal/space"
)

func SavePng(filename string, image image.Image) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	png.Encode(f, image)
}

func SavePalette(filename string, picker colour_picker.ColourPicker) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	image := image.NewRGBA(image.Rect(0, 0, len(picker.Colours), 30))

	for i, item := range picker.Colours {

		colour := color.RGBA{uint8(item.R), uint8(item.G), uint8(item.B), 255}
		for j := 0; j < 30; j++ {
			image.Set(i, j, colour)
		}
	}

	png.Encode(f, image)
}

func main() {

	width := 1640
	height := 1480

	space := space.NewSpace(width, height)

	config := Config{3.0, -1.0, -1.0, 1.0, 1.0, 100, 0.8, 0}
	drawJuliaSet(config, &space)

	palette := colour_picker.NewGradientPicker(1000,
		colour_picker.Colour{0, 0, 0},
		colour_picker.Colour{0, 0, 0},
		[]colour_picker.GradientPoint{
			{5, colour_picker.Colour{16, 133, 139}},
			{15, colour_picker.Colour{255, 174, 0}},
		},
	)

	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, item := range space.Items {
		paletteColour := palette.Pick(int(1000 * item.Value))

		colour := color.RGBA{uint8(paletteColour.R), uint8(paletteColour.G), uint8(paletteColour.B), 255}
		image.Set(item.X, item.Y, colour)
	}

	SavePalette("img/palette.png", palette)
	SavePng("img/fractal.png", image)
	fmt.Println("Finished!")
}
