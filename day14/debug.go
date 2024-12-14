package main

import (
	"fmt"
	"strings"

	aoc "github.com/Tesohh/aoclibgo"
)

func Debug(room Room, robots []Robot, level int) {
	grid := aoc.NewEmptyByteMatrix(room.Grid.Width(), room.Grid.Height())

	for _, robot := range robots {
		grid[robot.Pos.Y][robot.Pos.X] += 1
	}

	fmt.Print(aoc.ClearConsole)

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] > 0 {
				fmt.Print(aoc.Red, grid[y][x], aoc.Reset)
			} else {
				fmt.Print(aoc.Gray, string('.'), aoc.Reset)
			}
		}
		fmt.Println()
	}

	if level > 2 {
		fmt.Scanln()
	}

}

func GetString(room Room, robots []Robot) string {
	grid := aoc.NewEmptyByteMatrix(room.Grid.Width(), room.Grid.Height())

	for _, robot := range robots {
		grid[robot.Pos.Y][robot.Pos.X] += 1
	}

	builder := strings.Builder{}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] > 0 {
				builder.WriteByte('#')
			} else {
				builder.WriteByte('.')
			}
		}
		builder.WriteByte('\n')
	}

	return builder.String()
}
