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
}

func main() {
	tests()

	file := aoc.NewInputFromArgs().Bytes()

	disk := NewDiskFromFile(file)

	checksum := disk.Fragment().Checksum()
	fmt.Printf("checksum: %v\n", checksum)
	fmt.Printf("disk: %v\n", disk)
}
