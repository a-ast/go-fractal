package fractals

type MandelbrotSet struct {
	Canvas        Canvas
	EscapeRadius  float32
	MaxIterations int
}

func (fractal MandelbrotSet) Render() *FractalElements {
	var k int
	var c, z FloatPoint

	items := NewFractalElements()
	window := fractal.Canvas.calculateWindow()

	for i := 0; i < fractal.Canvas.Size.Width; i++ {
		for j := 0; j < fractal.Canvas.Size.Height; j++ {
			c = fractal.Canvas.scaleToWindow(i, j, window)
			z = FloatPoint{0, 0}

			for k = 0; k < fractal.MaxIterations; k++ {

				z = FloatPoint{
					X: z.X*z.X - z.Y*z.Y + c.X,
					Y: 2*z.X*z.Y + c.Y,
				}

				if z.X*z.X+z.Y*z.Y >= fractal.EscapeRadius*fractal.EscapeRadius {
					break
				}
			}

			items.Add(Element{i, j, float32(k) / float32(fractal.MaxIterations)})
		}
	}

	return items
}
