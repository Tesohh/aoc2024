package main

import aoc "github.com/Tesohh/aoclibgo"

func GPSValue(grid aoc.Matrix[byte]) int {
	total := 0

	for _, cell := range grid.Traverse() {
		if cell.Value != 'O' {
			continue
		}

		total += 100*cell.Point.Y + cell.Point.X
	}

	return total
}
