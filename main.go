package main

import (
	"fmt"

	"github.com/a-ast/go-fractal/colour_picker"
	"github.com/a-ast/go-fractal/fractals"
)

func main() {

	width := 200
	height := 200

	colourPicker := colour_picker.ArcticSun

	juliaSet := fractals.JuliaSet{
		Size:          fractals.Size{width, height},
		Complex:       0 + 0.8i,
		EscapeRadius:  3.0,
		MaxIterations: 100,
		Zoom:          1,
		Center:        fractals.FloatPoint{0, 0},
	}

	items := make(chan fractals.Element, width*height)
	go juliaSet.Render(items)

	SaveItemsToFile(items, fmt.Sprintf("img/%v.png", juliaSet.Zoom), juliaSet.Size, colourPicker)

	fmt.Println("Finished Async!")
}
