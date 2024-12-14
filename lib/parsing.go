package aoc

import (
	"bytes"
	"os"
	"strconv"
)

type Input struct {
	file []byte
}

func (i Input) Parts() []Input {
	inputs := []Input{}
	for _, part := range bytes.Split(i.file, []byte("\n\n")) {
		inputs = append(inputs, Input{file: part})
	}

	return inputs
}

func (i Input) Lines() [][]byte {
	lines := bytes.Split(i.file, []byte("\n"))
	lines = lines[:len(lines)-1]
	return lines
}

func (i Input) LineStrings() []string {
	strs := []string{}
	for _, line := range bytes.Split(i.file, []byte("\n")) {
		strs = append(strs, string(line))
	}
	strs = strs[:len(strs)-1]
	return strs
}

func (i Input) Bytes() []byte {
	return i.file
}

// Commonly used for maps
func (i Input) CharMatrix() Matrix[byte] {
	return Matrix[byte](i.Lines())
}

func (i Input) IntMatrix() Matrix[int] {
	chargrid := i.CharMatrix()
	grid := Matrix[int]{}

	for y := range chargrid {
		grid = append(grid, make([]int, len(chargrid[y])))
		for x, char := range chargrid[y] {
			if char == '.' {
				grid[y][x] = -1
			} else {
				num, _ := strconv.Atoi(string(char))
				grid[y][x] = num
			}
		}
	}

	return grid
}

func NewInputFromArgs() Input {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	return Input{file: file}
}
