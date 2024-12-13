package main

import (
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/Tesohh/aoclibgo"
)

type Machine struct {
	A     aoc.Point
	B     aoc.Point
	Prize aoc.Point
}

func ParseInput(input aoc.Input) Machine {
	buttonRegex := regexp.MustCompile(`Button (\w): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var machine Machine

	for _, line := range input.LineStrings() {
		line = strings.TrimSpace(line)
		if buttonMatch := buttonRegex.FindStringSubmatch(line); buttonMatch != nil {
			button := buttonMatch[1]
			x, _ := strconv.Atoi(buttonMatch[2])
			y, _ := strconv.Atoi(buttonMatch[3])

			switch button {
			case "A":
				machine.A = aoc.Point{X: x, Y: y}
			case "B":
				machine.B = aoc.Point{X: x, Y: y}
			}
		} else if prizeMatch := prizeRegex.FindStringSubmatch(line); prizeMatch != nil {
			x, _ := strconv.Atoi(prizeMatch[1])
			y, _ := strconv.Atoi(prizeMatch[2])
			machine.Prize = aoc.Point{X: x, Y: y}
		}
	}

	return machine
}
