package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])

	fileParts := bytes.Split(file, []byte("\n\n"))
	rulesRaw := string(fileParts[0])
	printsRaw := string(fileParts[1])

	rules := []Rule{}
	for _, line := range strings.Split(rulesRaw, "\n") {
		splits := strings.Split(line, "|")
		lhs, _ := strconv.Atoi(splits[0])
		rhs, _ := strconv.Atoi(splits[1])
		rules = append(rules, Rule{lhs: lhs, rhs: rhs})
	}

	queues := [][]int{}
	for _, line := range strings.Split(printsRaw, "\n") {
		queues = append(queues, make([]int, 0))
		for _, page := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(page)
			queues[len(queues)-1] = append(queues[len(queues)-1], num)
		}
	}

	sum1 := 0
	sum2 := 0

	for i, queue := range queues {
		fmt.Println("Queue number", i)
		if CheckQueue(queue, rules, true) {
			sum1 += GetMiddle(queue)
		} else {
			if os.Args[2] == "bogo" {
				sum2 += GetMiddle(BogosortQueue(queue, rules))
			} else {
				sum2 += GetMiddle(SortedQueue(queue, rules))
			}
		}
		fmt.Println()
	}

	fmt.Println("part1", sum1)
	fmt.Println("part2", sum2)
}
