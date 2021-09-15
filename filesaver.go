package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	cp "github.com/a-ast/go-fractal/colour_picker"
	"github.com/a-ast/go-fractal/fractals"
)

func SaveItemsToFile(items chan fractals.Element, filename string, width, height int, picker cp.ColourPicker, withCenter bool) {
	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for item := range items {
		paletteColour := picker.Pick(int(1000 * item.Value))
		colour := color.RGBA{uint8(paletteColour.R), uint8(paletteColour.G), uint8(paletteColour.B), 255}

		image.Set(item.X, item.Y, colour)
	}

	if withCenter == true {
		drawCenter(*image)
	}

	savePng(filename, image)
}

func drawCenter(img image.RGBA) {

	c1 := (img.Bounds().Max.X - img.Bounds().Min.X) / 2
	c2 := (img.Bounds().Max.Y - img.Bounds().Min.Y) / 2

	img.Set(c1, c2, color.RGBA{255, 0, 0, 255})

	img.Set(c1-1, c2-1, color.RGBA{255, 0, 0, 255})
	img.Set(c1-1, c2+1, color.RGBA{255, 0, 0, 255})
	img.Set(c1+1, c2-1, color.RGBA{255, 0, 0, 255})
	img.Set(c1+1, c2+1, color.RGBA{255, 0, 0, 255})
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
		panic(err)
	}

	defer f.Close()

	png.Encode(f, image)
}
