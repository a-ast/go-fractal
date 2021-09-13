package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/a-ast/go-fractal/colour_picker"
	"github.com/a-ast/go-fractal/fractal"
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

	width := 1600
	height := 800

	palette := colour_picker.NewGradientPicker(1000,
		colour_picker.Colour{0, 0, 0},
		colour_picker.Colour{0, 0, 0},
		[]colour_picker.GradientPoint{
			{5, colour_picker.Colour{16, 133, 139}},
			{15, colour_picker.Colour{255, 174, 0}},
		},
	)

	juliaSet := fractal.JuliaSet{
		Size:          fractal.Size{width, height},
		Complex:       0 + 0.8i,
		EscapeRadius:  3.0,
		MaxIterations: 100,
		Scale:         1,
		FocalPoint:    fractal.FloatPoint{50, 0},
	}

	items := make(chan fractal.SpaceItem, width*height)
	go juliaSet.Render(items)

	SavePalette("img/palette.png", palette)

	image2 := image.NewRGBA(image.Rect(0, 0, width, height))

	for item := range items {
		paletteColour := palette.Pick(int(1000 * item.Value))

		colour := color.RGBA{uint8(paletteColour.R), uint8(paletteColour.G), uint8(paletteColour.B), 255}
		image2.Set(item.X, item.Y, colour)
	}

	SavePng("img/fractal-a.png", image2)
	fmt.Println("Finished Async!")
}
