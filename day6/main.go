package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	grid := bytes.Split(file, []byte{'\n'})

	tally := Tally{}
	part1 := Patrol(grid, false, &tally)
	fmt.Println("part1:", part1)

	file, err = os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	grid = bytes.Split(file, []byte{'\n'})

	part2 := 0
	wg := sync.WaitGroup{}
	for y := range grid {
		for x := range grid[y] {
			go ManipulatedPatrol(grid, x, y, &wg, &part2)
		}
	}
	wg.Wait()
	fmt.Println("part2:", part2)
}
