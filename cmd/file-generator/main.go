package main

import (
	"fmt"

	"github.com/a-ast/go-fractal/pkg/colourpicker"
	"github.com/a-ast/go-fractal/pkg/fractals"
)

func main() {

	width := 800
	height := 400

	colourPicker := colourpicker.Electro

	fractal, err := fractals.New("mandelbrot", width, height)
	if err != nil {
		panic(err)
	}

	items := make(chan fractals.Element, width*height)
	go fractal.Render(items)

	SaveItemsToFile(items, "img/fractal.png", width, height, colourPicker, false)

	fmt.Println("Finished Async!")
}
