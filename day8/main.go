package main

import (
	aoc "github.com/Tesohh/aoclibgo"
)

func main() {
	input := aoc.NewInputFromArgs()
	grid := input.CharMatrix()

	world := World{
		Grid:      grid,
		Groups:    map[byte]AntennaeGroup{},
		Antinodes: map[aoc.Point]struct{}{},
	}

	CalculateAntinodes(&world, false)

	world = World{
		Grid:      grid,
		Groups:    map[byte]AntennaeGroup{},
		Antinodes: map[aoc.Point]struct{}{},
	}

	CalculateAntinodes(&world, true)
}
