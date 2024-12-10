package main

import (
	"fmt"
	"strconv"

	aoc "github.com/Tesohh/aoclibgo"
)

type Disk struct {
	Cells []int
}

func NewDiskFromFile(file []byte) Disk {
	disk := Disk{
		Cells: []int{},
	}
	id := 0 // for id
	isFile := true

	for _, char := range file {
		times, _ := strconv.Atoi(string(char))
		if isFile {
			disk.Cells = append(disk.Cells, aoc.RepeatSlice(id, times)...)
			id += 1
		} else {
			disk.Cells = append(disk.Cells, aoc.RepeatSlice(-1, times)...)
		}

		isFile = !isFile
	}

	return disk
}

func (d Disk) String() string {
	str := ""
	for _, cell := range d.Cells {
		if cell != -1 {
			str += fmt.Sprint(cell)
		} else {
			str += "."
		}
	}

	return str
}
