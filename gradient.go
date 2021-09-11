package main

import "image/color"

type GradientPoint struct {
	Percentage int
	Color      color.Color
}

func NewGradientPalette(size uint32, startColor, endColor color.Color) Palette {
	colors := make([]color.Color, size)

	for i := uint32(0); i < size; i++ {

		r1, g1, b1, a1 := startColor.RGBA()
		r2, g2, b2, a2 := endColor.RGBA()

		colors[i] = color.RGBA{
			uint8(i * (r2 - r1) / size),
			uint8(i * (g2 - g1) / size),
			uint8(i * (b2 - b1) / size),
			uint8(i * (a2 - a1) / size),
		}
	}

	return Palette{colors}
}
