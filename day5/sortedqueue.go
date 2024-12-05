package main

import (
	"fmt"
	"slices"
)

// thanks lallo S
func SortedQueue(queue []int, rules []Rule) []int {
	newQueue := make([]int, len(queue))
	copy(newQueue, queue)

	for {
		changed := false
		for i, page := range newQueue {
			if i == 0 {
				continue
			}

			// try to build a rule so that the previous item comes before current
			hypotheticalRule := Rule{lhs: newQueue[i-1], rhs: page}

			if !slices.Contains(rules, hypotheticalRule) { // if this rule doesn't exist (therefore it's invalid)
				// swap the items...
				newQueue[i-1], newQueue[i] = newQueue[i], newQueue[i-1]
				changed = true
				fmt.Printf("🔁 SWAP    %d with %d due to constructed rule %s\n", queue[i-1], page, hypotheticalRule)
			}
		}

		// repeat the loop until nothing is changed, therefore everything is SORTED
		if !changed {
			break
		}
	}

	return newQueue
}
