package pointgen

type LineStatus int

const Visible LineStatus = 0
const Hidden LineStatus = -1

type Point struct {
	X int
	Y int
}

type Line struct {
	A      Point // upper left
	B      Point // lower right
	Width  int
	Status LineStatus
}

func (l *Line) Hide() {
	l.Status = Hidden
}

func (l *Line) Show() {
	l.Status = Visible
}

func (l *Line) IsShown() bool {
	return l.Status == Visible
}
