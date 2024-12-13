package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

var TenTrillion = 10000000000000

const A_COST = 3
const B_COST = 1

func main() {
	// Button A: X+27, Y+18
	// Button B: X+17, Y+44
	// Prize: X=8105, Y=9698

	input := aoc.NewInputFromArgs()
	parts := input.Parts()

	totalCoinsPart1 := 0

	for _, part := range parts {
		machine := ParseInput(part)
		solvable, n, m := CheckedCramer(machine.A, machine.B, machine.Prize, false)
		if solvable {
			totalCoinsPart1 += A_COST*n + B_COST*m
		}
	}

	totalCoinsPart2 := 0
	for _, part := range parts {
		machine := ParseInput(part)
		machine.Prize = machine.Prize.Add(aoc.Point{X: TenTrillion, Y: TenTrillion})
		solvable, n, m := CheckedCramer(machine.A, machine.B, machine.Prize, false)
		if solvable {
			totalCoinsPart2 += A_COST*n + B_COST*m
		}
	}

	fmt.Printf("totalCoins: %v\n", totalCoinsPart1)
	fmt.Printf("totalCoins: %v\n", totalCoinsPart2)
}
