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

func (fdisk FileDisk) Checksum() int {
	checksum := 0

	for _, file := range fdisk.Files {
		if file.Id == -1 {
			continue
		}

		for i := file.X; i-file.X < file.Len; i++ {
			checksum += i * file.Id
		}
	}

	return checksum
}
