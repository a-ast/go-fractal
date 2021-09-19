package main

type CreateFractal struct {
	kind   string
	width  int
	height int
	zoom   float32
	cx     float32
	cy     float32
	re     float32
	im     float32
	er     float32
	mi     int
}

func newCreateFractalFromQuery(query Query) CreateFractal {
	return CreateFractal{
		kind:   query.GetString("t", "burningship"),
		width:  query.GetInt("w", 800),
		height: query.GetInt("h", 400),
		zoom:   query.GetFloat("z", 1.0),
		cx:     query.GetFloat("cx", 0),
		cy:     query.GetFloat("cy", 0),
		re:     query.GetFloat("re", 0),
		im:     query.GetFloat("im", 0.8),
		er:     query.GetFloat("er", 3.0),
		mi:     query.GetInt("mi", 100),
	}
}
