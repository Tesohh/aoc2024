package main

import (
	"fmt"

	aoc "github.com/Tesohh/aoclibgo"
)

func tests() {
	disk := NewDiskFromFile([]byte("2333133121414131402"))
	if disk.String() != "00...111...2...333.44.5555.6666.777.888899" {
		panic("test failed during parsing in NewDiskFromFile")
	}

	if disk.Fragment().String() != "0099811188827773336446555566.............." {
		panic("test failed during fragmentation")
	}

	if disk.Fragment().Checksum() != 1928 {
		panic(fmt.Sprintf("test failed during checksum %v", disk.Fragment().Checksum()))
	}

	// if disk.IntoFileDisk().FragmentFiles().String() != "00992111777.44.333....5555.6666.....8888.." {
	// 	panic(fmt.Sprintf("test failed during FragmentFiles"))
	// }
}

func main() {
	tests()

	file := aoc.NewInputFromArgs().Bytes()

	disk := NewDiskFromFile(file)
	fdisk := disk.IntoFileDisk()
	// fmt.Printf("disk:  %v\n", disk)

	checksum := disk.Fragment().Checksum()
	fmt.Printf("checksum: %v\n", checksum)

	checksum2 := fdisk.FragmentFiles().Checksum()
	fmt.Printf("checksum2: %v\n", checksum2)
}
