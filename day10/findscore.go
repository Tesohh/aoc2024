package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func FindScore(grid aoc.Matrix[int]) {
	trailheads := []Hiker{}
	for _, cell := range grid.Traverse() {
		if cell.Value == 0 {
			trailheads = append(trailheads, Hiker{Pos: cell.Point, Direction: aoc.Point{}})
		}
	}

	totalScore := 0
	totalRating := 0

	for _, trailhead := range trailheads {
		peaks := map[aoc.Point]struct{}{}
		hikers := []Hiker{trailhead}

		// when a hiker has no more possible moves, or he reached a 9, he will be deleted
		for len(hikers) > 0 {
			// Debug(grid, hikers, 3)

			newHikers, newPeaks, newRating := StepHikers(hikers, grid)
			hikers = newHikers
			for _, peak := range newPeaks {
				peaks[peak] = struct{}{}
			}

			fmt.Printf("len(%v) peaks: %v\n", len(peaks), peaks)
			totalScore += len(peaks)
			totalRating += newRating
		}
	}

	fmt.Printf("totalScore: %v\n", totalScore/2) // no idea why, but you have to divide it by 2
	fmt.Printf("branchCount: %v\n", totalRating)
}
