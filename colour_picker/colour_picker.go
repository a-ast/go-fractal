package colour_picker

type Colour struct {
	R, G, B int
}

type ColourPicker struct {
	Colours []Colour
}

func (palette ColourPicker) Pick(position int) Colour {
	if position >= len(palette.Colours) {
		return palette.Colours[len(palette.Colours)-1]
	}

	return palette.Colours[position]
}
