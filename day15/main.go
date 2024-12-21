package main

import (
	"fmt"
	"strings"

	aoc "github.com/Tesohh/aoclibgo"
)

func main() {
	input := aoc.NewInputFromArgs()
	parts := input.Parts()

	grid := parts[0].CharMatrix()
	moves := strings.Join(parts[1].LineStrings(), "\n")

	// the robot won't "physically" be in the grid
	pos := aoc.Point{}
	for _, cell := range grid.Traverse() {
		if cell.Value == '@' {
			pos = cell.Point
			grid[cell.Point.Y][cell.Point.X] = '.'
			break
		}
	}

	robot := Robot{
		Pos:   pos,
		Moves: moves,
	}

	part1grid := make(aoc.Matrix[byte], len(grid))
	copy(part1grid, grid)

	grid = Part1(part1grid, robot)
	fmt.Printf("GPSValue(grid): %v\n", GPSValue(grid))

	part2gridfake := make(aoc.Matrix[byte], len(grid))
	copy(part2gridfake, grid)

	part2grid := ConvertGrid(part2gridfake)
	Debug(part2grid, robot, 2)

	grid = Part2(part2grid, robot)
}
