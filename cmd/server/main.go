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
	fractalType := query.GetString("t", "mandelbrot")
	cf := newCreateFractalFromQuery(query)

	if query.Error() != "" {
		http.Error(w, query.Error(), 404)
		return
	}

	colourPicker := cp.PickerByName(cf.Palette)
	canvas := fractals.Canvas{
		Size:   fractals.Size{Width: cf.Width, Height: cf.Height},
		Zoom:   cf.Zoom,
		Center: fractals.FloatPoint{X: cf.Cx, Y: cf.Cy},
	}

	fractal, err := fractals.NewFractal(fractalType, canvas, cf)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	items := fractal.Render()

	imageFactory := fractals.ImageFactory{
		Width:      cf.Width,
		Height:     cf.Height,
		Picker:     colourPicker,
		WithCenter: false,
	}
	img := imageFactory.FromItems(items)

	writeImage(w, img)
}

func writeImage(w http.ResponseWriter, image *image.RGBA) {
	buf := new(bytes.Buffer)
	png.Encode(buf, image)

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}

func handleRequests() {
	http.HandleFunc("/", getFractal)

	log.Println(`
╔═══════════════════════════════════════════════════════════════╗
║ ✨ Go Fractal Server Started                                  ║
╠═══════════════════════════════════════════════════════════════╣
║ 🌐 http://localhost:10000                                     ║
║                                                               ║
║ QUERY PARAMETERS:                                             ║
║ ───────────────────────────────────────────────────────────── ║
║ t  - Fractal type (default: mandelbrot)                       ║
║      Options: julia, mandelbrot, burningship                  ║
║ p  - Colour palette (default: arcticsun)                      ║
║      Options: arcticsun, electro                              ║
║ w  - Image width in pixels (default: 800)                     ║
║ h  - Image height in pixels (default: 400)                    ║
║ z  - Zoom level (default: 1.0)                                ║
║ cx - Center X coordinate (default: 0)                         ║
║ cy - Center Y coordinate (default: 0)                         ║
║ re - Real part for Julia set (default: 0)                     ║
║ im - Imaginary part for Julia set (default: 0.8)              ║
║ er - Escape radius (default: 3.0)                             ║
║ mi - Maximum iterations (default: 100)                        ║
║                                                               ║
║ Example: http://localhost:10000?t=mandelbrot&cx=-0.5&z=0.5    ║
╚═══════════════════════════════════════════════════════════════╝
	`)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func newCreateFractalFromQuery(query Query) fractals.CreateFractal {
	return fractals.CreateFractal{
		Palette: query.GetString("p", "arcticsun"),
		Width:   query.GetInt("w", 800),
		Height:  query.GetInt("h", 400),
		Zoom:    query.GetFloat("z", 0),
		Cx:      query.GetFloat("cx", 0),
		Cy:      query.GetFloat("cy", 0),
		Re:      query.GetFloat("re", 0),
		Im:      query.GetFloat("im", 0.8),
		Er:      query.GetFloat("er", 0),
		Mi:      query.GetInt("mi", 0),
	}
}
