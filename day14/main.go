package main

import (
	"fmt"
	"strings"

	aoc "github.com/Tesohh/aoclibgo"
)

func main() {
	input := aoc.NewInputFromArgs()

	room := Room{}

	robots := []Robot{}
	for _, line := range input.LineStrings() {
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.Pos.X, &robot.Pos.Y, &robot.Velocity.X, &robot.Velocity.Y)
		robots = append(robots, robot)
	}

	if len(robots) <= 12 {
		room.Grid = aoc.NewEmptyByteMatrix(11, 7)
	} else {
		room.Grid = aoc.NewEmptyByteMatrix(101, 103)
	}

	part1bots := make([]Robot, len(robots))
	copy(part1bots, robots)
	for i := 1; i <= 100; i++ {
		// Debug(room, part1bots, 3)
		part1bots = StepRobots(room, part1bots)
	}

	Debug(room, part1bots, 2)
	fmt.Printf("CountQuadrants(room, part1bots): %v\n", CountQuadrants(room, part1bots))

	fmt.Println("Press enter to go on to part 2")
	fmt.Scanln()

	part2bots := make([]Robot, len(robots))
	copy(part2bots, robots)

	counter := 0
	for {
		s := GetString(room, part2bots)
		if strings.Contains(s, "###########") {
			Debug(room, part2bots, 3)
			fmt.Println("found tree after ", counter)
			fmt.Scanln()
		}
		part2bots = StepRobots(room, part2bots)
		counter += 1
	}
}
