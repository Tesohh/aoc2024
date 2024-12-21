package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func RuneToPoint(r rune) aoc.Point {
	switch r {
	case '^':
		return aoc.Point{X: 0, Y: -1}
	case 'v':
		return aoc.Point{X: 0, Y: 1}
	case '<':
		return aoc.Point{X: -1, Y: 0}
	case '>':
		return aoc.Point{X: 1, Y: 0}
	default:
		fmt.Println("Invalid ByteToPoint parameter:", r)
		return aoc.Point{X: 0, Y: 0}
	}
}
