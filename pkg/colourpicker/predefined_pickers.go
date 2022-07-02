package colourpicker

var (
	ArcticSun ColourPicker = NewGradientPicker(1000,
		[]GradientPoint{
			{5, Colour{16, 133, 139}},
			{15, Colour{255, 174, 0}},
		},
	)

	Electro ColourPicker = NewGradientPicker(1000,
		[]GradientPoint{
			{5, Colour{173, 3, 252}},
			{15, Colour{3, 244, 252}},
		},
	)
)

func PickerByName(name string) ColourPicker {
	switch name {
	case "electro":
		return NewGradientPicker(1000,
			[]GradientPoint{
				{5, Colour{173, 3, 252}},
				{15, Colour{3, 244, 252}},
			},
		)

	case "jenny":
		return NewGradientPicker(1000,
			[]GradientPoint{
				{5, Colour{8, 160, 75}},
				{10, Colour{82, 157, 255}},
				{15, Colour{252, 199, 255}},
			},
		)

	//case "arcticsun":
	default:
		return NewGradientPicker(1000,
			[]GradientPoint{
				{5, Colour{16, 133, 139}},
				{15, Colour{255, 174, 0}},
			},
		)
	}
}
