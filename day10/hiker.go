package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

type Hiker struct {
	Pos       aoc.Point
	Direction aoc.Point
	Retired   bool
}

func (h Hiker) String() string {
	return fmt.Sprintf("Hiker{(%d, %d) -> (%d, %d) retired: %v}", h.Pos.X, h.Pos.Y, h.Direction.X, h.Direction.Y, h.Retired)
}
