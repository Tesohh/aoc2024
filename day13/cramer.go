package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

// solves the following system
//
// { a_x * n + b_y * m = expect_x
// { a_y * n + b_x * m = expect_y
//
// checking whether the results are correct is up to you. (or use CheckedCramer)
// (there is no mechanism to check if the system is solvable)
func Cramer(a aoc.Point, b aoc.Point, expect aoc.Point) (n, m int) {
	determinant := (a.X * b.Y) - (b.X * a.Y)

	n = ((expect.X * b.Y) - (expect.Y * b.X)) / determinant
	m = ((a.X * expect.Y) - (a.Y * expect.X)) / determinant

	return n, m
}

// same as Cramer but returns whether the results are correct too
func CheckedCramer(a aoc.Point, b aoc.Point, expect aoc.Point, debug bool) (solvable bool, n, m int) {
	n, m = Cramer(a, b, expect)

	notNegative := n < 0 || m < 0
	xIsExpected := a.X*n+b.X*m != expect.X
	yIsExpected := a.Y*n+b.Y*m != expect.Y

	if debug {
		fmt.Printf("notNegative: %v\n", notNegative)
		fmt.Printf("xIsExpected: %v\n", xIsExpected)
		fmt.Printf("yIsExpected: %v\n", yIsExpected)
	}

	solvable = !notNegative && !xIsExpected && !yIsExpected

	return solvable, n, m
}
