package fractals

import "math"

func NewBurningShip(canvas Canvas, er float32, mi int) BurningShip {
	if canvas.Zoom == 0 {
		canvas.Zoom = 5
	}

	if canvas.Center.X == 0 {
		canvas.Center.X = -1.8
	}

	if er == 0 {
		er = 2.0
	}

	if mi == 0 {
		mi = 200
	}

	return BurningShip{Canvas: canvas, EscapeRadius: er, MaxIterations: mi}
}

type BurningShip struct {
	Canvas        Canvas
	EscapeRadius  float32
	MaxIterations int
}

func (fractal BurningShip) Render() *FractalElements {
	var iteration int
	var xy, p FloatPoint

	items := NewFractalElements()
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

			items.Add(Element{i, j, float32(iteration) / float32(fractal.MaxIterations)})
		}
	}

	return items
}
