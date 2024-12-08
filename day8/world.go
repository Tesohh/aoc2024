package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

type World struct {
	Grid      aoc.Matrix[byte]
	Groups    map[byte]AntennaeGroup
	Antinodes map[aoc.Point]struct{} // map to ensure uniqueness
}

func (w World) Debug() {
	for y, line := range w.Grid {
		for x, char := range line {
			if _, ok := w.Antinodes[aoc.Point{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println("")
	}
}
