package fractals

type Element struct {
	X, Y  int
	Value float32
}

type Canvas struct {
	Size   Size
	Zoom   float32
	Center FloatPoint
}

// type Fractal interface {
// }

type Size struct {
	Width, Height int
}

type FloatPoint struct {
	X, Y float32
}

type floatSize struct {
	Width, Height float32
}

type window struct {
	left, right FloatPoint
}

func (canvas Canvas) calculateWindow() window {
	size := canvas.calculateWindowSize()

	return window{
		FloatPoint{-size.Width, -size.Height},
		FloatPoint{size.Width, size.Height},
	}
}

func (canvas Canvas) scaleToWindow(i, j int, window window) FloatPoint {
	return FloatPoint{
		X: interpolate(window.left.X, window.right.X, i, canvas.Size.Width) + canvas.Center.X,
		Y: interpolate(window.left.Y, window.right.Y, j, canvas.Size.Height) + canvas.Center.Y,
	}
}

func (canvas Canvas) calculateWindowSize() floatSize {
	if canvas.Size.Width < canvas.Size.Height {
		return floatSize{
			Width:  1 / canvas.Zoom,
			Height: 1 / canvas.Zoom * float32(canvas.Size.Width) / float32(canvas.Size.Height),
		}
	}

	return floatSize{
		Width:  1 / canvas.Zoom,
		Height: 1 / canvas.Zoom * float32(canvas.Size.Height) / float32(canvas.Size.Width),
	}
}

func interpolate(start, end float32, position, size int) float32 {
	return start + float32(position)*(end-start)/float32(size)
}
