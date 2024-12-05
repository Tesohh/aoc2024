package main

import (
	"fmt"
	"os"
)

const MAS_SUM = 'M' + 'A' + 'S'

func checkForXMAS(b []byte) bool {
	return b[0]+b[1]+b[2] == MAS_SUM
}

func checkCrossword(grid [][]byte, x, y int, count *int) {
	if grid[x][y] != 'X' {
		return
	}
	canGoLeft := x-3 >= 0
	canGoRight := x+3 < len(grid[y])-1
	canGoTop := y-3 >= 0
	canGoBottom := y+3 < len(grid)-1

	// horizontal
	if canGoRight {
		b := []byte{grid[y][x+1], grid[y][x+2], grid[y][x+3]}
		if string(b) == "MAS" {
			fmt.Println("RIGHT XMAS AT", x, y)
			(*count) += 1
		}
	}
	// FIX:problem should be here
	if canGoLeft {
		b := []byte{grid[y][x-1], grid[y][x-2], grid[y][x-3]}
		if string(b) == "SAM" {
			fmt.Println("LEFT  XMAS AT", x, y)
			(*count) += 1
		}
	}

	// Vertical
	if canGoBottom {
		b := []byte{grid[y+1][x], grid[y+2][x], grid[y+3][x]}
		if string(b) == "MAS" {
			fmt.Println("BOTTM XMAS AT", x, y)
			(*count) += 1
		}
	}
	if canGoTop {
		b := []byte{grid[y-1][x], grid[y-2][x], grid[y-3][x]}
		if string(b) == "MAS" {
			fmt.Println("TOP   XMAS AT", x, y)
			(*count) += 1
		}
	}

	// Diagonal
	if canGoBottom && canGoRight {
		b := []byte{grid[y+1][x+1], grid[y+2][x+2], grid[y+3][x+3]}
		if string(b) == "MAS" {
			fmt.Println("DBR   XMAS AT", x, y)
			(*count) += 1
		}
	}
	if canGoBottom && canGoLeft {
		b := []byte{grid[y+1][x-1], grid[y+2][x-2], grid[y+3][x-3]}
		if string(b) == "MAS" {
			fmt.Println("DBL   XMAS AT", x, y)
			(*count) += 1
		}
	}
	if canGoTop && canGoRight {
		b := []byte{grid[y-1][x+1], grid[y-2][x+2], grid[y-3][x+3]}
		if string(b) == "MAS" {
			fmt.Println("DTR   XMAS AT", x, y)
			(*count) += 1
		}
	}
	if canGoTop && canGoLeft {
		b := []byte{grid[y-1][x-1], grid[y-2][x-2], grid[y-3][x-3]}
		if string(b) == "MAS" {
			fmt.Println("DTL   XMAS AT", x, y)
			(*count) += 1
		}
	}
}

func main() {
	grid := [][]byte{make([]byte, 0)}
	y := 0

	file, _ := os.ReadFile("short.txt")
	for _, char := range file {
		if char == '\n' {
			y += 1
			grid = append(grid, make([]byte, 0))
		} else {
			grid[y] = append(grid[y], char)
		}

	}
	grid = grid[:len(grid)-1]
	count := 0
	for y := range grid {
		fmt.Println(string(grid[y]))
	}
	for y := range grid {

		for x := range grid[y] {
			checkCrossword(grid, x, y, &count)

		}
	}

	fmt.Println(count)
}
