package main

import (
	"fmt"

	"github.com/a-ast/go-fractal/pkg/colourpicker"
	"github.com/a-ast/go-fractal/pkg/fractals"
)

func main() {

	width := 800
	height := 400

	colourPicker := colourpicker.ArcticSun

	fractal, err := fractals.New("julia", width, height)
	if err != nil {
		panic(err)
	}

	items := fractals.NewFractalElements()

	fractal.Render(items)
	SaveItemsToFile(items, "img/fractal.png", width, height, colourPicker, false)

	fmt.Println("Finished!")
}
