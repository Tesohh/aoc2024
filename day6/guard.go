package main

import (
	"fmt"
	"time"
)

type Guard struct {
	PositionX int
	PositionY int
	Direction Direction
}

func (g Guard) IsOutOfBounds(grid [][]byte) bool {
	isOutOfBoundsX := g.PositionX < 0 || g.PositionX >= len(grid[0])
	isOutOfBoundsY := g.PositionY < 0 || g.PositionX >= len(grid)
	return isOutOfBoundsX || isOutOfBoundsY
}

func (g Guard) IsNextOutOfBounds(grid [][]byte) bool {
	vectorX, vectorY := g.Direction.Vector()
	nextX := g.PositionX + vectorX
	nextY := g.PositionY + vectorY

	isOutOfBoundsX := nextX < 0 || nextX >= len(grid[0])
	isOutOfBoundsY := nextY < 0 || nextY >= len(grid)-1
	return isOutOfBoundsX || isOutOfBoundsY
}

func (g Guard) GetNext(grid [][]byte) byte {
	if g.IsNextOutOfBounds(grid) {
		return 0
	}
	vectorX, vectorY := g.Direction.Vector()
	nextX := g.PositionX + vectorX
	nextY := g.PositionY + vectorY
	return grid[nextY][nextX]
}

func (g *Guard) Step(grid [][]byte, cells *int) {
	if grid[g.PositionY][g.PositionX] != 'X' {
		(*cells) += 1
		grid[g.PositionY][g.PositionX] = 'X'
	}

	vectorX, vectorY := g.Direction.Vector()
	g.PositionX += vectorX
	g.PositionY += vectorY
}

var count = 0

func (g Guard) Debug(grid [][]byte, level int) {
	fmt.Print("\033[H\033[2J")
	vectorX, vectorY := g.Direction.Vector()
	nextX := g.PositionX + vectorX
	nextY := g.PositionY + vectorY

	fmt.Printf("pos.X: %v, pos.Y: %v\n", g.PositionX, g.PositionY)
	fmt.Printf("nextX: %v, nextY: %v\n", nextX, nextY)
	fmt.Printf("direction: %v (%v, %v)\n", g.Direction, vectorX, vectorY)
	fmt.Println("count:", count)
	for y := range grid {
		for x, char := range grid[y] {
			if x == g.PositionX && y == g.PositionY {
				fmt.Print("\033[0;31m" + string(g.Direction.Char()) + "\033[0m")
			} else {
				if char == '.' {
					fmt.Printf("\033[0;30m%c\033[0m", char)
				} else if char == '#' {
					fmt.Printf("\033[0;36m%c\033[0m", char)
				} else if char == 'X' {
					fmt.Printf("\033[0;34m%c\033[0m", char)
				} else {
					fmt.Printf(string(char))
				}
			}
		}
		fmt.Println("")
	}

	if g.IsOutOfBounds(grid) {
		fmt.Println("out of bounds...")
	}
	if g.IsNextOutOfBounds(grid) {
		fmt.Println("next is out of bounds...")
	}

	if level == 1 {
		time.Sleep(50 * time.Millisecond)
	} else if level == 2 {
		time.Sleep(100 * time.Millisecond)
	} else if level == 3 {
		fmt.Scanln()
	}

	count += 1
}
