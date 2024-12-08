package main

import "os"

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

}
