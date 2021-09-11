package main

import "image/color"

type Palette struct {
	colors []color.Color
}

func (palette Palette) GetColor(position int) color.Color {
	return palette.colors[position]
}
