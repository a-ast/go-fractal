package space

type SpaceItem struct {
	X, Y  int
	Value float32
}

type Space struct {
	Width, Height int
	Items         []SpaceItem
}

func NewSpace(width, height int) Space {
	items := make([]SpaceItem, 0)

	return Space{width, height, items}
}

func (space *Space) AddItem(item SpaceItem) {
	space.Items = append(space.Items, item)
}
