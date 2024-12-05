package main

import "math"

func GetMiddle(queue []int) int {
	return queue[int(math.Floor(float64(len(queue)/2.0)))]
}
