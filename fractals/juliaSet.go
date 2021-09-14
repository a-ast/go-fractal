package fractals

type FloatPoint struct {
	X, Y float32
}

type JuliaSet struct {
	Size          Size
	Complex       complex64
	EscapeRadius  float32
	MaxIterations int
	Scale         float32
	FocalPoint    FloatPoint
}

func (fractal JuliaSet) Render(items chan Element) {
	var iteration int
	var x, y float32

	widthScale, heightScale := fractal.sizeScales()

	for i := 0; i < fractal.Size.Width; i++ {
		for j := 0; j < fractal.Size.Height; j++ {
			iteration = 0

			x, y = fractal.scaleToSize(i, j, widthScale, heightScale)

			for x*x+y*y < fractal.EscapeRadius*fractal.EscapeRadius && iteration <= fractal.MaxIterations {
				x, y = x*x-y*y+real(fractal.Complex), 2*x*y+imag(fractal.Complex)
				iteration++
			}

			items <- Element{i, j, float32(iteration) / float32(fractal.MaxIterations)}
		}
	}

	close(items)
}

func interpolate(start, end float32, position, size int) float32 {
	return start + float32(position)*(end-start)/float32(size)
}

func (fractal JuliaSet) scaleToSize(i, j int, widthScale, heightScale float32) (float32, float32) {
	return interpolate(-widthScale, widthScale, i+int(fractal.FocalPoint.X), fractal.Size.Width),
		interpolate(-heightScale, heightScale, j+int(fractal.FocalPoint.Y), fractal.Size.Height)
}

func (fractal JuliaSet) sizeScales() (float32, float32) {
	if fractal.Size.Width < fractal.Size.Height {
		return fractal.Scale,
			fractal.Scale * float32(fractal.Size.Width) / float32(fractal.Size.Height)
	}

	return fractal.Scale,
		fractal.Scale * float32(fractal.Size.Height) / float32(fractal.Size.Width)
}
