package fractals

func NewJuliaSet(canvas Canvas, c complex64, er float32, mi int) JuliaSet {

	if canvas.Zoom == 0 {
		canvas.Zoom = 1
	}

	if imag(c) == 0 {
		c = complex(real(c), 0.8)
	}

	if er == 0 {
		er = 3.0
	}

	if mi == 0 {
		mi = 100
	}

	return JuliaSet{Canvas: canvas, Complex: c, EscapeRadius: er, MaxIterations: mi}
}

type JuliaSet struct {
	Canvas        Canvas
	Complex       complex64
	EscapeRadius  float32
	MaxIterations int
}

func (fractal JuliaSet) Render() *FractalElements {
	var iteration int
	var p FloatPoint

	items := NewFractalElements()
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

			items.Add(Element{i, j, float32(iteration) / float32(fractal.MaxIterations)})
		}
	}

	return items
}
