package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"

	cp "github.com/a-ast/go-fractal/pkg/colourpicker"
	"github.com/a-ast/go-fractal/pkg/fractals"
)

func main() {
	handleRequests()
}

func getFractal(w http.ResponseWriter, r *http.Request) {

	query := NewQuery(r.URL.Query())
	cf := newCreateFractalFromQuery(query)

	if query.Error() != "" {
		http.Error(w, query.Error(), 404)
		return
	}

	colourPicker := cp.PickerByName(cf.palette)

	canvas := fractals.Canvas{
		Size:   fractals.Size{cf.width, cf.height},
		Zoom:   cf.zoom,
		Center: fractals.FloatPoint{cf.cx, cf.cy},
	}

	fractal, err := NewFractal(canvas, cf)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	items := fractals.NewFractalElements()
	fractal.Render(items)

	imageFactory := fractals.ImageFactory{
		Width:      cf.width,
		Height:     cf.height,
		Picker:     colourPicker,
		WithCenter: false,
	}
	image := imageFactory.FromItems(items)

	writeImage(w, image)
}

func writeImage(w http.ResponseWriter, image *image.RGBA) {
	buf := new(bytes.Buffer)
	png.Encode(buf, image)

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}

func handleRequests() {
	http.HandleFunc("/", getFractal)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
