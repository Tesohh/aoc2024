package main

import (
	aoc "github.com/Tesohh/aoclibgo"
)

func main() {
	input := aoc.NewInputFromArgs()
	grid := input.IntMatrix()

	FindScore(grid)
}
