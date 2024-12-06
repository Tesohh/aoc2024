package main

type Tally struct {
	Steps int
	Cells int
	Done  bool
}

func Patrol(grid [][]byte, debug bool, tally *Tally) int {
	guard := Guard{PositionX: -1, PositionY: -1, Direction: Top}

	for y := range grid {
		for x, char := range grid[y] {
			if char == '^' {
				guard.PositionX = x
				guard.PositionY = y
				grid[y][x] = '.'
			}
		}
	}

	if guard.PositionX == -1 || guard.PositionY == -1 {
		panic("no ^ found")
	}

	cells := 0

	for !(guard.IsNextOutOfBounds(grid)) {
		guard.Step(grid, &cells)
		tally.Steps += 1
		for guard.GetNext(grid) == '#' {
			guard.Direction.Turn90Deg()
		}
		// if guard.GetNext(grid) == '#' {
		// 	guard.Direction.Turn90Deg()
		// }
		if len(grid) < 100 && debug {
			guard.Debug(grid, 1)
		}

		if tally.Steps > 999999 {
			return cells + 1
		}
	}

	tally.Done = true
	tally.Cells = cells + 1
	return cells + 1
}
