package main

import (
	"slices"
	"strconv"
	"strings"

	aoc "github.com/Tesohh/aoclibgo"
)

type FileDisk struct {
	Files []File
}

func (d Disk) IntoFileDisk() FileDisk {
	fdisk := FileDisk{}

	lastCell := d.Cells[0]
	file := File{Id: d.Cells[0], X: 0, Len: 0}
	for i, cell := range d.Cells {
		if lastCell == cell {
			file.Len += 1
		} else {
			fdisk.Files = append(fdisk.Files, file)

			file.Id = cell
			file.X = i
			file.Len = 1
		}

		lastCell = cell
	}

	fdisk.Files = append(fdisk.Files, file)

	return fdisk
}

func (d FileDisk) String() string {
	s := ""
	sorted := make([]File, len(d.Files))
	copy(sorted, d.Files)
	slices.SortFunc(sorted, func(a, b File) int {
		return a.X - b.X
	})

	for _, file := range sorted {
		if file.Id == -1 {
			s += string(aoc.RepeatSlice('.', file.Len))
			continue
		}
		s += strings.Join(aoc.RepeatSlice(strconv.Itoa(file.Id), file.Len), "")
	}

	return s
}
