package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	cp "github.com/a-ast/go-fractal/pkg/colourpicker"
	"github.com/a-ast/go-fractal/pkg/fractals"
)

func SaveItemsToFile(items *fractals.FractalElements, filename string, width, height int, picker cp.ColourPicker, withCenter bool) {

	factory := fractals.ImageFactory{
		Width:      width,
		Height:     height,
		Picker:     picker,
		WithCenter: withCenter,
	}
	image := factory.FromItems(items)

	savePng(filename, image)
}

func SavePaletteToFile(picker cp.ColourPicker, filename string) {
	image := image.NewRGBA(image.Rect(0, 0, len(picker.Colours), 30))

	for i, item := range picker.Colours {

		colour := color.RGBA{uint8(item.R), uint8(item.G), uint8(item.B), 255}
		for j := 0; j < 30; j++ {
			image.Set(i, j, colour)
		}
	}

	savePng(filename, image)
}

func savePng(filename string, image image.Image) {
	f, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer f.Close()

	png.Encode(f, image)
}
