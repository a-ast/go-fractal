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
		Canvas: fractals.Canvas{
			Size:   fractals.Size{width, height},
			Zoom:   1,
			Center: fractals.FloatPoint{0, 0},
		},
		Complex:       0 + 0.8i,
		EscapeRadius:  3.0,
		MaxIterations: 100,
	}

	items := make(chan fractals.Element, width*height)
	go juliaSet.Render(items)

	SaveItemsToFile(items, fmt.Sprintf("img/%v.png", juliaSet.Canvas.Zoom), juliaSet.Canvas.Size, colourPicker)

	fmt.Println("Finished Async!")
}
