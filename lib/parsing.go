package aoc

import (
	"bytes"
	"os"
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
	return strs
}

// Commonly used for maps
func (i Input) CharMatrix() Matrix[byte] {
	return Matrix[byte](i.Lines())
}

func NewInputFromArgs() Input {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	return Input{file: file}
}
