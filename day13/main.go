package main

import (
	"fmt"
	"math"

	aoc "github.com/Tesohh/aoclibgo"
)

var TenTrillion = int(math.Pow(10, 13))

func main() {
	a := aoc.Point{X: 94, Y: 34}
	b := aoc.Point{X: 22, Y: 67}
	expected := aoc.Point{X: 8400, Y: 5400}

	fmt.Println(CheckedCramer(a, b, expected, true))
	// fmt.Println(CheckedCramer(aoc.Point{3, 2}, aoc.Point{2, 1}, aoc.Point{43, 28}))
}
