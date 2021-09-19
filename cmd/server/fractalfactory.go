package main

import (
	"errors"

	"github.com/a-ast/go-fractal/pkg/fractals"
)

func NewFractal(canvas fractals.Canvas, f CreateFractal) (fractals.Fractal, error) {
	switch f.kind {
	case "julia":
		return fractals.JuliaSet{
			Canvas:        canvas,
			Complex:       complex(f.re, f.im),
			EscapeRadius:  f.er,
			MaxIterations: f.mi,
		}, nil

	case "mandelbrot":
		return fractals.MandelbrotSet{
			Canvas:        canvas,
			EscapeRadius:  f.er,
			MaxIterations: f.mi,
		}, nil

	// z=5&cx=-1.8&mi=200&er=2
	case "burningship":
		return fractals.BurningShip{
			Canvas:        canvas,
			EscapeRadius:  f.er,
			MaxIterations: f.mi,
		}, nil
	default:
		return nil, errors.New("unknown fractal type")
	}
}
