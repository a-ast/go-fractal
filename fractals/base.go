package fractals

type Element struct {
	X, Y  int
	Value float32
}

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
