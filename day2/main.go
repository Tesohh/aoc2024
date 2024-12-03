package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// thanks mat
func checkSafety(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	// figure out the trend
	isIncreasing := nums[0] < nums[1]
	isDecreasing := nums[0] > nums[1]

	// go through every thing, starting from i = 1
	for i := 1; i < len(nums); i += 1 {
		if !checkPairSafety(nums[i-1], nums[i], isIncreasing, isDecreasing) {
			return false
		}
	}

	return true
}

func checkPairSafety(num1 int, num2 int, isIncreasing bool, isDecreasing bool) bool {
	if isIncreasing && num1 > num2 {
		return false
	}
	if isDecreasing && num1 < num2 {
		return false
	}

	levelDiff := math.Abs(float64(num1 - num2))
	if levelDiff < 1 || levelDiff > 3 {
		return false
	}

	return true
}

func checkSafetyDampened(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	safe := false
	for i := 0; i < len(nums); i += 1 {
		c := make([]int, len(nums))
		_ = copy(c, nums)
		if checkSafety(append(c[:i], c[i+1:]...)) {
			safe = true
		}
	}

	return safe
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	successCountPart1 := 0
	successCountPart2 := 0

	for _, bline := range bytes.Split(file, []byte("\n")) {
		line := string(bline)

		nums := []int{}

		for _, strnum := range strings.Split(line, " ") {
			if strnum == "" {
				break
			}
			num, err := strconv.Atoi(strnum)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		if checkSafety(nums) {
			successCountPart1 += 1
		}

		if checkSafetyDampened(nums) {
			successCountPart2 += 1
		}
	}

	fmt.Println(successCountPart1)
	fmt.Println(successCountPart2)
}
