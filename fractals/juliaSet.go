package fractals

type JuliaSet struct {
	Size          Size
	Complex       complex64
	EscapeRadius  float32
	MaxIterations int
	Zoom          float32
	Center        FloatPoint
}

func (fractal JuliaSet) Render(items chan Element) {
	var iteration int
	var p FloatPoint

	window := fractal.calculateWindow()

	for i := 0; i < fractal.Size.Width; i++ {
		for j := 0; j < fractal.Size.Height; j++ {
			iteration = 0
			p = fractal.scaleToWindow(i, j, window)

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

func (fractal JuliaSet) calculateWindow() window {
	size := fractal.calculateWindowSize()

	return window{
		FloatPoint{-size.Width, -size.Height},
		FloatPoint{size.Width, size.Height},
	}
}

func (fractal JuliaSet) scaleToWindow(i, j int, window window) FloatPoint {
	return FloatPoint{
		X: interpolate(window.left.X, window.right.X, i, fractal.Size.Width) + fractal.Center.X,
		Y: interpolate(window.left.Y, window.right.Y, j, fractal.Size.Height) + fractal.Center.Y,
	}
}

func (fractal JuliaSet) calculateWindowSize() floatSize {
	if fractal.Size.Width < fractal.Size.Height {
		return floatSize{
			Width:  1 / fractal.Zoom,
			Height: 1 / fractal.Zoom * float32(fractal.Size.Width) / float32(fractal.Size.Height),
		}
	}

	return floatSize{
		Width:  1 / fractal.Zoom,
		Height: 1 / fractal.Zoom * float32(fractal.Size.Height) / float32(fractal.Size.Width),
	}
}

func interpolate(start, end float32, position, size int) float32 {
	return start + float32(position)*(end-start)/float32(size)
}
