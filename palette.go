package main

type Colour struct {
	R, G, B int
}

type Palette struct {
	colors []Colour
}

func (palette Palette) GetColor(position int) Colour {
	if position >= len(palette.colors) {
		return palette.colors[len(palette.colors)-1]
	}

	return palette.colors[position]
}
