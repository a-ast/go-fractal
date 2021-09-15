package main

import (
	"fmt"

	"github.com/a-ast/go-fractal/colour_picker"
	"github.com/a-ast/go-fractal/fractals"
)

func main() {

	width := 1600
	height := 800

	colourPicker := colour_picker.ArcticSun

	fractal, err := build("julia", width, height)
	if err != nil {
		panic(err)
	}

	items := make(chan fractals.Element, width*height)
	go fractal.Render(items)

	SaveItemsToFile(items, "img/fractal.png", width, height, colourPicker)

	fmt.Println("Finished Async!")
}
