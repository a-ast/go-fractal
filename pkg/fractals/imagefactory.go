package fractals

import (
	"image"
	"image/color"

	cp "github.com/a-ast/go-fractal/pkg/colourpicker"
)

type ImageFactory struct {
	Width, Height int
	Picker        cp.ColourPicker
	WithCenter    bool
}

func (f ImageFactory) FromItems(items *FractalElements) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, f.Width, f.Height))

	for _, item := range items.Elements {
		paletteColour := f.Picker.Pick(int(1000 * item.Value))
		colour := color.RGBA{
			R: uint8(paletteColour.R),
			G: uint8(paletteColour.G),
			B: uint8(paletteColour.B),
			A: 255,
		}

		img.Set(item.X, item.Y, colour)
	}

	if f.WithCenter {
		drawCenter(*img)
	}

	return img
}

func drawCenter(img image.RGBA) {

	c1 := (img.Bounds().Max.X - img.Bounds().Min.X) / 2
	c2 := (img.Bounds().Max.Y - img.Bounds().Min.Y) / 2

	centerColour := color.RGBA{R: 255, A: 255}
	img.Set(c1, c2, centerColour)

	img.Set(c1-1, c2-1, centerColour)
	img.Set(c1-1, c2+1, centerColour)
	img.Set(c1+1, c2-1, centerColour)
	img.Set(c1+1, c2+1, centerColour)
}
