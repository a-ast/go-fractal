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
	image := image.NewRGBA(image.Rect(0, 0, f.Width, f.Height))

	for _, item := range items.Elements {
		paletteColour := f.Picker.Pick(int(1000 * item.Value))
		colour := color.RGBA{uint8(paletteColour.R), uint8(paletteColour.G), uint8(paletteColour.B), 255}

		image.Set(item.X, item.Y, colour)
	}

	if f.WithCenter {
		drawCenter(*image)
	}

	return image
}

func drawCenter(img image.RGBA) {

	c1 := (img.Bounds().Max.X - img.Bounds().Min.X) / 2
	c2 := (img.Bounds().Max.Y - img.Bounds().Min.Y) / 2

	img.Set(c1, c2, color.RGBA{255, 0, 0, 255})

	img.Set(c1-1, c2-1, color.RGBA{255, 0, 0, 255})
	img.Set(c1-1, c2+1, color.RGBA{255, 0, 0, 255})
	img.Set(c1+1, c2-1, color.RGBA{255, 0, 0, 255})
	img.Set(c1+1, c2+1, color.RGBA{255, 0, 0, 255})
}
