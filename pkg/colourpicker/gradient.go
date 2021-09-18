package colourpicker

import (
	"math"
)

type GradientPoint struct {
	Percentage int
	Color      Colour
}

func NewGradientPicker(size int, points []GradientPoint) ColourPicker {
	colors := make([]Colour, size)

	points = append(points, GradientPoint{100, Colour{0, 0, 0}})

	fromColour := Colour{0, 0, 0}
	fromPosition := 0

	for _, point := range points {
		toPosition := int(math.Round(float64(size * point.Percentage / 100)))

		for i := fromPosition; i < toPosition; i++ {
			colors[i] = interpolatedColour(fromColour, point.Color, i-fromPosition, toPosition-fromPosition)
		}

		fromColour = point.Color
		fromPosition = toPosition
	}

	return ColourPicker{colors}
}

func interpolate(start, end, position, size int) int {
	return int(math.Round(float64(start + position*(end-start)/size)))
}

func interpolatedColour(startColour, endColour Colour, position, size int) Colour {
	return Colour{
		interpolate(startColour.R, endColour.R, position, size),
		interpolate(startColour.G, endColour.G, position, size),
		interpolate(startColour.B, endColour.B, position, size),
	}
}
