package main

import aoc "github.com/Tesohh/aoclibgo"

func ConvertGrid(grid aoc.Matrix[byte]) aoc.Matrix[byte] {
	newGrid := make(aoc.Matrix[byte], 0)

	for y := range grid {
		newGrid = append(newGrid, make([]byte, 0))
		for x := range grid[y] {
			if grid[y][x] == '#' {
				newGrid[y] = append(newGrid[y], '#', '#')
			} else if grid[y][x] == 'O' {
				newGrid[y] = append(newGrid[y], '[', ']')
			} else if grid[y][x] == '.' {
				newGrid[y] = append(newGrid[y], '.', '.')
			} else if grid[y][x] == '@' {
				newGrid[y] = append(newGrid[y], '@', '.')
			}
		}
	}

	return newGrid
}
