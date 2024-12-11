package main

import (
	"slices"

	aoc "github.com/Tesohh/aoclibgo"
)

// how to figure out whether to branch or to just move?
// save direction in hiker
// if the direction is the same, just move,
// if not, create a new hiker
// and if he cannot move in his own direction, then delete him

// if grid.Point(next) == currentCell+1 {
// 	if hikers[i].Direction == (aoc.Point{X: 0, Y: 0}) {
// 		hikers[i].Direction = direction
// 	}
//
// 	if hikers[i].Direction == direction {
// 		hikers[i].Pos = next
// 	}
// }

// at every step, find possible branching paths.
// if it's 0: retire
// if it's just 1: change direction to that path
// if it's more than 1: retire current hiker, and branch in every direction.
// then, all hikers move by 1 in their direction
// if they end up out of the map, retire.
// if they end up on a 9, increment total score and retire.
// make sure that 9's are counted only ONCE per trailhead!

func StepHikers(hikers []Hiker, grid aoc.Matrix[int]) ([]Hiker, []aoc.Point, int) {
	peaks := []aoc.Point{}

	// purge any retired hikers
	hikers = slices.DeleteFunc(hikers, func(h Hiker) bool { return h.Retired })

	rating := 0

	newHikers := []Hiker{}
	for i := range hikers {
		currentCell := grid.Point(hikers[i].Pos)

		// find all possible branches
		branches := []aoc.Point{}
		for _, direction := range PointDirections {
			next := hikers[i].Pos.Add(direction)

			if grid.IsOutOfBounds(next) {
				continue
			}

			if grid.Point(next) == currentCell+1 {
				branches = append(branches, direction)
			}
		}

		// manage branching and retirement
		if len(branches) == 0 {
			hikers[i].Retired = true
		} else if len(branches) == 1 {
			hikers[i].Direction = branches[0]
		} else {
			hikers[i].Retired = true
			for _, direction := range branches {
				newHiker := Hiker{
					Pos:       hikers[i].Pos,
					Direction: direction,
				}
				newHikers = append(newHikers, newHiker)
			}
		}
	}

	// add new hikers
	hikers = append(hikers, newHikers...)

	// actually step
	for i := range hikers {
		if hikers[i].Retired {
			continue
		}

		next := hikers[i].Pos.Add(hikers[i].Direction)
		if grid.IsOutOfBounds(next) {
			hikers[i].Retired = true
			continue
		}
		// else if grid.Point(next) == 9 {
		// 			hikers[i].Retired = true
		// 			totalScore += 1
		// 		}
		hikers[i].Pos = next
	}

	for i := range hikers {
		if grid.Point(hikers[i].Pos) == 9 {
			hikers[i].Retired = true
			rating += 1
			peaks = append(peaks, hikers[i].Pos)
		}
	}

	return hikers, peaks, rating
}
