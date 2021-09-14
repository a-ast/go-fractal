package colour_picker

var (
	ArcticSun ColourPicker = NewGradientPicker(1000,
		Colour{0, 0, 0},
		Colour{0, 0, 0},
		[]GradientPoint{
			{5, Colour{16, 133, 139}},
			{15, Colour{255, 174, 0}},
		},
	)
)
