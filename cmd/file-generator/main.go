package main

import (
	"fmt"

	cp "github.com/a-ast/go-fractal/pkg/colourpicker"
	"github.com/a-ast/go-fractal/pkg/fractals"
)

func main() {

	width := 800
	height := 400

	cf := fractals.CreateFractal{
		Palette: "",
		Width:   800,
		Height:  400,
	}

	colourPicker := cp.PickerByName("")

	canvas := fractals.Canvas{
		Size:   fractals.Size{Width: cf.Width, Height: cf.Height},
		Zoom:   cf.Zoom,
		Center: fractals.FloatPoint{X: cf.Cx, Y: cf.Cy},
	}

	fractal, err := fractals.NewFractal("julia", canvas, cf)
	if err != nil {
		panic(err)
	}

	items := fractal.Render()
	SaveItemsToFile(items, "img/fractal.png", width, height, colourPicker, false)

	fmt.Println("Finished!")
}
