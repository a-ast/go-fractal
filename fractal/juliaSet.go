package fractal

type FloatPoint struct {
	X, Y float64
}

type JuliaSet struct {
	Size          Size
	Complex       complex64
	EscapeRadius  float32
	MaxIterations int
	Scale         float32
	FocalPoint    FloatPoint
}

func (fractal JuliaSet) Render(items chan SpaceItem) {

	limit := fractal.EscapeRadius * fractal.EscapeRadius

	var iteration int
	var x, y float32

	widthScale := fractal.Scale
	heightScale := fractal.Scale * float32(fractal.Size.Width) / float32(fractal.Size.Height)

	if fractal.Size.Width > fractal.Size.Height {
		heightScale = 1 / heightScale
	}

	for i := 0; i < fractal.Size.Width; i++ {
		for j := 0; j < fractal.Size.Height; j++ {
			iteration = 0
			x = interpolate(-widthScale, widthScale, i+int(fractal.FocalPoint.X), fractal.Size.Width)    //- config.FocalPoint.X
			y = interpolate(-heightScale, heightScale, j+int(fractal.FocalPoint.Y), fractal.Size.Height) //- config.FocalPoint.Y

			for x*x+y*y < limit && iteration <= fractal.MaxIterations {
				x, y = x*x-y*y+real(fractal.Complex), 2*x*y+imag(fractal.Complex)

				iteration++
			}

			if iteration > fractal.MaxIterations {
				continue
			}

			items <- SpaceItem{i, j, float32(iteration) / float32(fractal.MaxIterations)}
		}
	}

	close(items)
}

func interpolate(start, end float32, position, size int) float32 {
	return start + float32(position)*(end-start)/float32(size)
}
