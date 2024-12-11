package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func Debug(grid aoc.Matrix[int], hikers []Hiker, level int) {
	// fmt.Print("\033[H\033[2J")
	for y := range grid {
		for x := range grid[y] {
			isHiker := false
			isRetired := false
			for _, hiker := range hikers {
				if (hiker.Pos == aoc.Point{X: x, Y: y}) {
					isHiker = true
					isRetired = hiker.Retired
				}
			}

			if isHiker {
				if isRetired {
					fmt.Print("\033[34m\033[0m")
				} else {
					fmt.Print("\033[31m\033[0m")
				}
			} else if grid[y][x] == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println()
	}

	fmt.Printf("hikers: %v\n", hikers)

	if level > 2 {
		fmt.Scanln()
	}
}
