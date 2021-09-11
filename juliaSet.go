package main

import (
	"image"
	"image/color"
)

type Config struct {
	EscapeRadius  float64
	MinX, MinY    float64
	MaxX, MaxY    float64
	MaxIterations int
	Imag, Real    float64
}

func drawJuliaSet(config Config, image *image.RGBA) {
	width := image.Bounds().Dx()
	height := image.Bounds().Dy()

	var widthFactor, heightFactor float64

	widthFactor = 1 / float64(width-1)
	heightFactor = 1 / float64(height-1)

	limit := config.EscapeRadius * config.EscapeRadius

	var iteration int
	var x, y, z0, z1 float64

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			x = config.MinX + float64(i)*((config.MaxX-config.MinX)*widthFactor)
			y = config.MinY + float64(j)*((config.MaxX-config.MinX)*heightFactor)

			iteration = 0
			z0 = x
			z1 = y

			for x*x+y*y < limit && iteration <= config.MaxIterations {
				z1 = 2*z0*z1 + config.Imag
				z0 = x*x - y*y + config.Real

				x = z0
				y = z1

				iteration++
			}

			if iteration > config.MaxIterations {
				continue
			}

			// r := math.Log(float64(iteration)) / math.Log(float64(config.MaxIterations)) * 255

			image.Set(i, j, color.RGBA{uint8(iteration), uint8(10 * iteration), 0, 255})
			//image.Set(i, j, color.RGBA{100, 100, 22, 0})
		}
	}
}
