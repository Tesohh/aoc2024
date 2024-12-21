package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func Part2(grid aoc.Matrix[byte], robot Robot) aoc.Matrix[byte] {
	for pc, moveChar := range robot.Moves {
		Debug(grid, robot, 2)
		fmt.Println("move", pc, string(moveChar))

		direction := RuneToPoint(moveChar)

		next := robot.Pos.Add(direction)
		if grid.IsOutOfBounds(next) {
			panic("out of bonuds!")
		}

		// vertical movement (recursive collision detection)
		if direction.Y != 0 {

		}

		// horizontal movement
		if direction.X != 0 {

		}
	}

	return grid
}
