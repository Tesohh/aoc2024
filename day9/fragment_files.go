package main

func (fdisk FileDisk) FragmentFiles() FileDisk {
	// possible issue with writing to slice while iterating..
	for i := len(fdisk.Files) - 1; i >= 0; i-- {
		file := &fdisk.Files[i]
		if file.Id == -1 {
			continue
		}

		// look for first possible space
		for j := range fdisk.Files {
			other := &fdisk.Files[j]

			isFreeSpace := other.Id == -1
			isThereEnoughLen := other.Len >= file.Len
			isOtherBeforeFile := other.X < file.X

			if isFreeSpace && isOtherBeforeFile && isThereEnoughLen {
				// fmt.Printf("file: %#v\n", file)
				// fmt.Printf("other: %#v\n", other)
				// fmt.Println("CISSY")

				oldX := file.X
				file.X = other.X
				other.X += file.Len
				other.Len -= file.Len

				// just need to add extra space where file was
				fdisk.Files = append(fdisk.Files, File{
					Id:  -1,
					X:   oldX,
					Len: file.Len,
				})
			}
		}

		// fmt.Println()
		// fmt.Printf("fdisk: %v\n", fdisk)
		// fmt.Println()
	}

	return fdisk
}
