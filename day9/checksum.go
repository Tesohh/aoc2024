package main

func (disk Disk) Checksum() int {
	checksum := 0

	for i, cell := range disk.Cells {
		if cell == -1 {
			continue
		}
		checksum += i * cell
	}

	return checksum
}
