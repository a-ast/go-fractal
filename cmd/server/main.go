package main

import (
	"bytes"
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

	kind := query.GetString("t", "burningship")
	width := query.GetInt("w", 800)
	height := query.GetInt("h", 400)
	zoom := query.GetFloat("z", 1.0)
	cx := query.GetFloat("cx", 0)
	cy := query.GetFloat("cy", 0)
	re := query.GetFloat("re", 0)
	im := query.GetFloat("im", 0.8)
	er := query.GetFloat("er", 3.0)
	mi := query.GetInt("mi", 100)

	if query.Error() != "" {
		http.Error(w, query.Error(), 404)
		return
	}

	colourPicker := cp.NewGradientPicker(1000,
		[]cp.GradientPoint{
			{0, cp.Colour{0, 0, 139}},
			{15, cp.Colour{255, 174, 0}},
			{25, cp.Colour{174, 174, 100}},
		},
	)
	canvas := fractals.Canvas{
		Size:   fractals.Size{width, height},
		Zoom:   zoom,
		Center: fractals.FloatPoint{cx, cy},
	}

	var fractal fractals.Fractal

	switch kind {
	case "julia":
		fractal = fractals.JuliaSet{
			Canvas:        canvas,
			Complex:       complex(re, im),
			EscapeRadius:  er,
			MaxIterations: mi,
		}
	case "mandelbrot":
		fractal = fractals.MandelbrotSet{
			Canvas:        canvas,
			EscapeRadius:  er,
			MaxIterations: mi,
		}
	// z=5&cx=-1.8&mi=200&er=2
	case "burningship":
		fractal = fractals.BurningShip{
			Canvas:        canvas,
			EscapeRadius:  er,
			MaxIterations: mi,
		}
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

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}

func handleRequests() {
	http.HandleFunc("/", getFractal)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
