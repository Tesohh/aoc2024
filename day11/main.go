package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	aoc "github.com/Tesohh/aoclibgo"
)

var cache = make(map[uint64][]uint64)

func BlinkStone(x uint64) []uint64 {
	if x == 0 {
		return []uint64{1}
	}

	if v, ok := cache[x]; ok {
		return v
	}

	s := strconv.FormatUint(x, 10)
	l := len(s)
	if l%2 == 0 {
		a, err := strconv.ParseUint(s[:l/2], 10, 64)
		if err != nil {
			log.Fatal("err is not nil")
		}

		b, err := strconv.ParseUint(s[l/2:], 10, 64)
		if err != nil {
			log.Fatal("err is not nil")
		}

		cache[x] = []uint64{a, b}
		return []uint64{a, b}
	}

	cache[x] = []uint64{x * 2024}
	return []uint64{x * 2024}
}

func CountStones(ints []uint64, passes int) int {

	for i := 0; i < passes; i++ {
		stones := []uint64{}
		for _, stone := range ints {
			stones = append(stones, BlinkStone(stone)...)
		}
		ints = stones
		fmt.Printf("Blinks %d, stones %d\n", i, len(stones))
	}

	return len(ints)
}

func main() {
	line := aoc.NewInputFromArgs().Lines()[0]
	nums := bytes.Split(line, []byte{' '})
	ints := []uint64{}
	for _, num := range nums {
		x, _ := strconv.ParseUint(string(num), 10, 64)
		ints = append(ints, x)
	}

	fmt.Printf("CountStones(): %v\n", CountStones(ints, 75))
}
