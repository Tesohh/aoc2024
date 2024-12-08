package aoc

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func (p Point) Invert() Point {
	p.X *= -1
	p.Y *= -1
	return p
}

func (p Point) Add(other Point) Point {
	p.X += other.X
	p.Y += other.Y
	return p
}

func (p Point) Sub(other Point) Point {
	p.X -= other.X
	p.Y -= other.Y
	return p
}

func (p Point) Mul(other Point) Point {
	p.X *= other.X
	p.Y *= other.Y
	return p
}

func (p Point) Div(other Point) Point {
	p.X /= other.X
	p.Y /= other.Y
	return p
}
