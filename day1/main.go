package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	s := string(b)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1] // get rid of whitespace line

	idPairs := [2][]int{{}, {}}
	for _, line := range lines {
		pairs := strings.Split(line, "   ")
		atoi0, _ := strconv.Atoi(pairs[0])
		atoi1, _ := strconv.Atoi(pairs[1])
		idPairs[0] = append(idPairs[0], atoi0)
		idPairs[1] = append(idPairs[1], atoi1)
	}

	slices.Sort(idPairs[0])
	slices.Sort(idPairs[1])

	totalDistance := 0

	if len(idPairs[0]) != len(idPairs[1]) {
		panic("unsame length")
	}

	for i := 0; i < len(idPairs[0]); i += 1 {
		totalDistance += int(math.Abs(float64(idPairs[0][i] - idPairs[1][i])))
	}

	similarityScore := 0
	for i := 0; i < len(idPairs[0]); i += 1 {
		id := idPairs[0][i]
		timesAppeared := 0
		for j := 0; j < len(idPairs[1]); j += 1 {
			if idPairs[1][j] == id {
				timesAppeared += 1
			}
		}

		similarityScore += id * timesAppeared
	}

	fmt.Println(totalDistance)
	fmt.Println(similarityScore)
}
