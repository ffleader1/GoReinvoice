package fpdfpoint

type LineStatus int

const Visible LineStatus = 0
const Hidden LineStatus = -1

type Point struct {
	X float64
	Y float64
}

func (p Point) Translation(x, y float64) Point {
	p.X += x
	p.Y += y
	return p
}

type Line struct {
	A      Point // upper left
	B      Point // lower right
	Width  float64
	Status LineStatus
}

func (l Line) Hide() Line {
	l.Status = Hidden
	return l
}

func (l Line) Show() Line {
	l.Status = Visible
	return l
}

func (l Line) IsShown() bool {
	return l.Status == Visible
}

func (l Line) Translation(x, y float64) Line {
	l.A = l.A.Translation(x, y)
	l.B = l.B.Translation(x, y)
	return l
}
