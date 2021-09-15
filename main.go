package main

import (
	"fmt"

	"github.com/a-ast/go-fractal/colour_picker"
	"github.com/a-ast/go-fractal/fractals"
)

func main() {

	width := 800
	height := 400

	colourPicker := colour_picker.Electro

	fractal, err := build("mandelbrot", width, height)
	if err != nil {
		panic(err)
	}

	items := make(chan fractals.Element, width*height)
	go fractal.Render(items)

	SaveItemsToFile(items, "img/fractal.png", width, height, colourPicker, false)

	fmt.Println("Finished Async!")
}
