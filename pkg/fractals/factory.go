package fractals

import (
	"errors"
)

func New(fractalType string, width, height int) (Fractal, error) {
	switch fractalType {
	case "julia":
		return JuliaSet{
			Canvas: Canvas{
				Size:   Size{width, height},
				Zoom:   1,
				Center: FloatPoint{0, 0},
			},
			Complex:       0 + 0.8i,
			EscapeRadius:  3.0,
			MaxIterations: 100,
		}, nil

	case "mandelbrot":
		return MandelbrotSet{
			Canvas: Canvas{
				Size:   Size{width, height},
				Zoom:   0.5,
				Center: FloatPoint{-1, 0},
			},
			Complex:       0 + 0.8i,
			EscapeRadius:  3.0,
			MaxIterations: 100,
		}, nil

	default:
		return nil, errors.New("unknown fractal type")
	}
}
