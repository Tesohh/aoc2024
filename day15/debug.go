package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func Debug(grid aoc.Matrix[byte], robot Robot, level int) {
	if level > 2 {
		fmt.Print(aoc.ClearConsole)
	}

	for y := range grid {
		for x := range grid[y] {
			c := grid[y][x]
			if y == robot.Pos.Y && x == robot.Pos.X {
				fmt.Print(aoc.Red, "@", aoc.Reset)
			} else if c == '#' {
				fmt.Print(aoc.Blue, string(c), aoc.Reset)
			} else if c == '.' {
				fmt.Print(aoc.Faint, string(c), aoc.Reset)
			} else if c == 'O' {
				fmt.Print(aoc.Green, string(c), aoc.Reset)
			} else if c == '[' || c == ']' {
				fmt.Print(aoc.Green, string(c), aoc.Reset)
			}
		}
		fmt.Println()
	}

	if level > 2 {
		fmt.Scanln()
	}

}
