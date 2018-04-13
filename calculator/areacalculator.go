package calculator

type Area struct {
	Unit int
}

func (Area) SquarArea(width int, height int) int {
	return width * height
}
