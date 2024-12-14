package aoc

// Assumes that each row has the same length
type Matrix[T any] [][]T

func NewEmptyByteMatrix(width, height int) Matrix[byte] {
	grid := make(Matrix[byte], height)
	for i := 0; i < height; i++ {
		grid[i] = make([]byte, width)
	}
	return grid
}

func (m Matrix[T]) Width() int {
	return len(m[0])
}

func (m Matrix[T]) Height() int {
	return len(m)
}

func (m Matrix[T]) IsOutOfBounds(p Point) bool {
	xOut := p.X < 0 || p.X > len(m)-1
	yOut := p.Y < 0 || p.Y > len(m[0])-1
	return xOut || yOut
}

type MatrixCell[T any] struct {
	Point Point
	Value T
}

func (m Matrix[T]) Traverse() []MatrixCell[T] {
	cells := []MatrixCell[T]{}
	for y, line := range m {
		for x, el := range line {
			cells = append(cells, MatrixCell[T]{
				Point: Point{X: x, Y: y},
				Value: el,
			})
		}
	}

	return cells
}

// always check if out of bounds before this!
func (m Matrix[T]) Point(p Point) T {
	return m[p.Y][p.X]
}
