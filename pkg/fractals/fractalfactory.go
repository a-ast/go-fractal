package fractals

import "errors"

func NewFractal(fractalType string, canvas Canvas, f CreateFractal) (Fractal, error) {
	switch fractalType {
	case "julia":
		return NewJuliaSet(canvas, complex(f.Re, f.Im), f.Er, f.Mi), nil
	case "mandelbrot":
		return NewMandelbrotSet(canvas, f.Er, f.Mi), nil
	case "burningship":
		return NewBurningShip(canvas, f.Er, f.Mi), nil
	default:
		return nil, errors.New("unknown fractal type")
	}
}
