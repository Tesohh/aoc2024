package main

import (
	"fmt"
	"slices"
)

func CheckQueue(queue []int, rules []Rule) bool {
	for i, lhs := range queue {
		if i == 0 {
			fmt.Printf("🫧 SKIPPED on number: %d\n", lhs)
			continue
		}

		ruleFound := false
		for _, rule := range rules {
			if rule.lhs == lhs {
				ruleFound = true
				if slices.Contains(queue[0:i], rule.rhs) {
					fmt.Printf("❌ ERROR   on number: %d rule: %s queue: %#v\n", lhs, rule, queue[0:i])
					return false
				} else {
					fmt.Printf("✅ OK      on number: %d rule: %s queue: %#v\n", lhs, rule, queue[0:i])
				}
			}
		}
		if !ruleFound {
			fmt.Printf("🫧 NORULES on number: %d (no rule found)\n", lhs)
		}
	}

	return true
}
