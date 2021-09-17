package fractals

type JuliaSet struct {
	Canvas        Canvas
	Complex       complex64
	EscapeRadius  float32
	MaxIterations int
	Center        FloatPoint
}

func (fractal JuliaSet) Render(items chan Element) {
	var iteration int
	var p FloatPoint

	window := fractal.Canvas.calculateWindow()

	for i := 0; i < fractal.Canvas.Size.Width; i++ {
		for j := 0; j < fractal.Canvas.Size.Height; j++ {
			iteration = 0
			p = fractal.Canvas.scaleToWindow(i, j, window)

			for (p.X*p.X)+(p.Y*p.Y) < fractal.EscapeRadius*fractal.EscapeRadius && iteration <= fractal.MaxIterations {
				p = FloatPoint{
					X: (p.X * p.X) - (p.Y * p.Y) + real(fractal.Complex),
					Y: 2*(p.X*p.Y) + imag(fractal.Complex),
				}
				iteration++
			}

			items <- Element{i, j, float32(iteration) / float32(fractal.MaxIterations)}
		}
	}

	close(items)
}
