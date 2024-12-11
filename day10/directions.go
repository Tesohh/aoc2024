package main

import aoc "github.com/Tesohh/aoclibgo"

var (
	PointBottom = aoc.Point{
		X: 0,
		Y: 1,
	}
	PointTop = aoc.Point{
		X: 0,
		Y: -1,
	}
	PointRight = aoc.Point{
		X: 1,
		Y: 0,
	}
	PointLeft = aoc.Point{
		X: -1,
		Y: 0,
	}
)

var PointDirections = []aoc.Point{PointBottom, PointTop, PointRight, PointLeft}
