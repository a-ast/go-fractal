package colourpicker

var (
	ArcticSun ColourPicker = NewGradientPicker(1000,
		Colour{0, 0, 0},
		Colour{0, 0, 0},
		[]GradientPoint{
			{5, Colour{16, 133, 139}},
			{15, Colour{255, 174, 0}},
		},
	)

	Electro ColourPicker = NewGradientPicker(1000,
		Colour{0, 0, 0},
		Colour{0, 0, 0},
		[]GradientPoint{
			{5, Colour{173, 3, 252}},
			{15, Colour{3, 244, 252}},
		},
	)
)
