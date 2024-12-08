package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func CalculateAntinodes(world *World, repeat bool) {
	// make groups
	for _, char := range world.Grid.Traverse() {
		if char.Value == '.' {
			continue
		}

		group, ok := world.Groups[char.Value]
		if !ok {
			group.Points = []aoc.Point{}
			world.Groups[char.Value] = group
		}
		group.Points = append(world.Groups[char.Value].Points, char.Point)
		world.Groups[char.Value] = group
	}

	// get antinodes for each group
	for _, group := range world.Groups {
		for i, antenna := range group.Points {
			for j, other := range group.Points {
				if j == i {
					continue
				}

				if !repeat {
					target := antenna.Mul(aoc.Point{X: 2, Y: 2}).Sub(other)

					if !world.Grid.IsOutOfBounds(target) {
						world.Antinodes[target] = struct{}{}
					}
				} else {
					oldTarget := antenna
					olderTarget := other

					world.Antinodes[antenna] = struct{}{}

					for {
						target := oldTarget.Mul(aoc.Point{X: 2, Y: 2}).Sub(olderTarget)
						olderTarget = oldTarget
						oldTarget = target

						fmt.Printf("target: %v\n", target)
						fmt.Printf("antenna: %v\n", oldTarget)
						fmt.Printf("other: %v\n", olderTarget)
						fmt.Println()

						if !world.Grid.IsOutOfBounds(target) {
							world.Antinodes[target] = struct{}{}
						} else {
							break
						}
					}
				}
			}
		}
	}

	world.Debug()
	fmt.Println("total antinodes: ", len(world.Antinodes))
}

// fmt.Printf("antenna.Mul(aoc.Point{X: 2, Y: 2}): %v\n", antenna.Mul(aoc.Point{X: 2, Y: 2}))
// fmt.Printf("other.Mul(aoc.Point{X: %d, Y: %d}): %v\n", repetitions, repetitions, other.Mul(aoc.Point{X: repetitions, Y: repetitions}))
//
// fmt.Printf("target: %v\n", target)
// fmt.Printf("antenna: %v\n", antenna)
// fmt.Printf("other: %v\n", other)
// fmt.Println()

// if (other == aoc.Point{X: 0, Y: 0}) {
// 	break
// }
