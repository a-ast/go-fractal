package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func SavePng(filename string, image image.Image) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	png.Encode(f, image)
}

func main() {

	width := 640
	height := 480

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	// bg := color.White
	// draw.Draw(m, m.Bounds(), &image.Uniform{bg}, image.Point{0, 0}, draw.Src)

	config := Config{5.0, -1.0, -1.0, 1.0, 1.0, 200, -0.8, 0.156}
	drawJuliaSet(config, m)

	SavePng("img/fractal.png", m)
	fmt.Println("Finished!")
}
