package main

import aoc "github.com/Tesohh/aoclibgo"

func CountQuadrants(room Room, robots []Robot) int {
	w := room.Grid.Width()
	h := room.Grid.Height()
	center := aoc.Point{X: room.Grid.Width() / 2, Y: room.Grid.Height() / 2}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, robot := range robots {
		p := robot.Pos
		isLeft := (p.X >= 0 && p.X < center.X)
		isRight := (p.X > center.X && p.X <= w)
		isTop := (p.Y >= 0 && p.Y < center.Y)
		isBottom := (p.Y > center.Y && p.Y <= h)

		if isLeft && isTop {
			q1 += 1
		} else if isRight && isTop {
			q2 += 1
		} else if isLeft && isBottom {
			q3 += 1
		} else if isRight && isBottom {
			q4 += 1
		}
	}

	return q1 * q2 * q3 * q4
}
