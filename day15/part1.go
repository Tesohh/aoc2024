package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func Part1(grid aoc.Matrix[byte], robot Robot) aoc.Matrix[byte] {
	for pc, moveChar := range robot.Moves {
		// Debug(grid, robot, 2)
		fmt.Println("move", pc, string(moveChar))

		direction := RuneToPoint(moveChar)

		next := robot.Pos.Add(direction)
		if grid.IsOutOfBounds(next) {
			panic("out of bonuds!")
		}

		nextChar := grid.Point(next)
		if nextChar == '#' {
			continue
		} else if nextChar == '.' {
			robot.Pos = next
			continue
		}

		nextNext := next
		canMoveBox := false
		for {
			nextNext = nextNext.Add(direction)

			if grid.IsOutOfBounds(nextNext) {
				// panic("out of bonuds!")
				break
			}

			if grid.Point(nextNext) == '.' {
				canMoveBox = true
				break
			} else if grid.Point(nextNext) == '#' {
				break
			}
		}

		if canMoveBox {
			grid[next.Y][next.X] = '.'
			grid[nextNext.Y][nextNext.X] = 'O'
			robot.Pos = next
		}
	}

	return grid
}
