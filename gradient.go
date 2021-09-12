package main

import (
	"math"
)

type GradientPoint struct {
	Percentage int
	Color      Colour
}

func NewGradientPalette(size int, startColor, endColor Colour) Palette {
	colors := make([]Colour, size)

	for i := 0; i < size; i++ {
		colors[i] = Colour{
			interpolate(startColor.R, endColor.R, i, size),
			interpolate(startColor.G, endColor.G, i, size),
			interpolate(startColor.B, endColor.B, i, size),
		}
	}

	return Palette{colors}
}

func interpolate(start, end, position, size int) int {
	return int(math.Round(float64(start + position*(end-start)/size)))
}
