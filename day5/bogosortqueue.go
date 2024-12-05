package main

import (
	"math/rand"
	"time"
)

// dumbest shit i've ever done but it works (and not even that bad)
func BogosortQueue(queue []int, rules []Rule) []int {
	newQueue := make([]int, len(queue))
	copy(newQueue, queue)

	for !CheckQueue(newQueue, rules, false) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(newQueue), func(i, j int) { newQueue[i], newQueue[j] = newQueue[j], newQueue[i] })
	}

	return newQueue
}
