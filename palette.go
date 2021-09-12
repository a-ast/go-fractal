package main

type Colour struct {
	R, G, B int
}

type Palette struct {
	colors []Colour
}

func (palette Palette) GetColor(position int) Colour {
	return palette.colors[position]
}
