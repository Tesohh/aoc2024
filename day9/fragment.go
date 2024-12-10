package main

import (
	"slices"
)

func (disk Disk) Fragment() Disk {
	emptyCount := 0
	for _, cell := range disk.Cells {
		if cell == -1 {
			emptyCount += 1
		}
	}

	for i := len(disk.Cells) - 1; i >= 0; i-- {
		for j, other := range disk.Cells {
			if other == -1 {
				disk.Cells[i], disk.Cells[j] = disk.Cells[j], disk.Cells[i]
			}
		}

		// not so elegant, could be done better, but that's ok
		firstEmpty := slices.Index(disk.Cells, -1)
		if len(disk.Cells[firstEmpty:]) == emptyCount {
			break
		}
	}

	return disk
}
