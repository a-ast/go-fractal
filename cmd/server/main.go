package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"

	"github.com/a-ast/go-fractal/pkg/colourpicker"
	"github.com/a-ast/go-fractal/pkg/fractals"
)

func main() {
	handleRequests()
}

func getFractal(w http.ResponseWriter, r *http.Request) {

	width := 800
	height := 400

	colourPicker := colourpicker.Electro

	fractal, err := fractals.New("mandelbrot", width, height)
	if err != nil {
		panic(err)
	}

	items := make(chan fractals.Element, width*height)
	go fractal.Render(items)

	factory := fractals.ImageFactory{
		Width:      width,
		Height:     height,
		Picker:     colourPicker,
		WithCenter: true,
	}
	image := factory.FromItems(items)

	buf := new(bytes.Buffer)
	png.Encode(buf, image)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(buf.Bytes())
}

func handleRequests() {
	http.HandleFunc("/", getFractal)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
