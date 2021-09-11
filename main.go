package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func SavePng(filename string, image image.Image) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = png.Encode(f, image)

	fmt.Println(err)
}

func SaveJpeg(filename string, image image.Image) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	jpeg.Encode(f, image, &jpeg.Options{})
}

func drawSome(image *image.RGBA) {

	width := image.Bounds().Dx()
	height := image.Bounds().Dy()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			image.Set(i, j, color.RGBA{255, 0, 22, 0})
		}
	}
}

func main() {

	width := 640
	height := 480

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	bg := color.White
	draw.Draw(m, m.Bounds(), &image.Uniform{bg}, image.Point{0, 0}, draw.Src)

	// for x := 0; x < width; x++ {
	// 	for y := 0; y < height; y++ {
	// 		m.Set(x, y, color.RGBA{100, 100, 22, 0})
	// 	}
	// }

	config := Config{5.0, -1.0, -1.0, 1.0, 1.0, 200, -0.8, 0.156}
	drawJuliaSet(config, m)

	SavePng("blue.png", m)
	SaveJpeg("blue.jpeg", m)
	fmt.Println("Finished!")
}
