package main

import (
	aoc "github.com/Tesohh/aoclibgo"
)

func StepRobots(room Room, robots []Robot) []Robot {
	for i := range robots {
		r := &robots[i]

		next := r.Pos.Add(r.Velocity)

		for {
			isTopOut := next.Y < 0
			isBottomOut := next.Y >= room.Grid.Height()
			isLeftOut := next.X < 0
			isRightOut := next.X >= room.Grid.Width()

			if !isTopOut && !isBottomOut && !isLeftOut && !isRightOut {
				break
			}

			if isTopOut || isBottomOut {
				next.Y = aoc.AbsInt(aoc.AbsInt(next.Y) - room.Grid.Height())
			}

			if isLeftOut || isRightOut {
				next.X = aoc.AbsInt(aoc.AbsInt(next.X) - room.Grid.Width())
			}
		}

		r.Pos = next
	}

	return robots
}
