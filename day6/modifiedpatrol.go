package main

import (
	"sync"
)

func ManipulatedPatrol(grid [][]byte, x, y int, wg *sync.WaitGroup, part2counter *int) bool {
	wg.Add(1)
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	if newGrid[y][x] != '^' {
		newGrid[y][x] = '#'
	}

	tally := Tally{}
	Patrol(newGrid, false, &tally)

	wg.Done()
	if !tally.Done {
		(*part2counter) += 1
	}
	return tally.Done
}
