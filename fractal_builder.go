package main

import (
	"errors"

	"github.com/a-ast/go-fractal/fractals"
)

func build(fractalType string, width, height int) (fractals.Fractal, error) {
	switch fractalType {
	case "julia":
		return fractals.JuliaSet{
			Canvas: fractals.Canvas{
				Size:   fractals.Size{width, height},
				Zoom:   1000,
				Center: fractals.FloatPoint{0, 0},
			},
			Complex:       0 + 0.8i,
			EscapeRadius:  3.0,
			MaxIterations: 100,
		}, nil

	case "mandelbrot":
		return fractals.MandelbrotSet{
			Canvas: fractals.Canvas{
				Size:   fractals.Size{width, height},
				Zoom:   1,
				Center: fractals.FloatPoint{0, 0},
			},
			Complex:       0 + 0.8i,
			EscapeRadius:  3.0,
			MaxIterations: 100,
		}, nil

	default:
		return nil, errors.New("unknown fractal type")
	}
}
