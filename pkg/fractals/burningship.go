package fractals

import "math"

type BurningShip struct {
	Canvas        Canvas
	EscapeRadius  float32
	MaxIterations int
}

func (fractal BurningShip) Render(items chan<- Element) {
	var iteration int
	var xy, p FloatPoint

	window := fractal.Canvas.calculateWindow()

	for i := 0; i < fractal.Canvas.Size.Width; i++ {
		for j := 0; j < fractal.Canvas.Size.Height; j++ {
			iteration = 0

			xy = fractal.Canvas.scaleToWindow(i, j, window)
			p = xy

			for (p.X*p.X)+(p.Y*p.Y) < fractal.EscapeRadius*fractal.EscapeRadius && iteration <= fractal.MaxIterations {
				p = FloatPoint{
					X: (p.X * p.X) - (p.Y * p.Y) + xy.X,
					Y: float32(math.Abs(float64(2*p.X*p.Y))) + xy.Y,
				}

				iteration++
			}

			items <- Element{i, j, float32(iteration) / float32(fractal.MaxIterations)}
		}
	}

	close(items)
}
