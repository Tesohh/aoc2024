package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re1 = regexp.MustCompile(`(?m)mul\([0-9]*\,[0-9]*\)`)                   // unethical regex
var re2 = regexp.MustCompile(`(?m)mul\([0-9]*\,[0-9]*\)|do\(\)|don\'t\(\)`) // dishonourable and despicable regex

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	total1 := 0
	total2 := 0

	for _, match := range re1.FindAll(b, -1) {
		stripped := strings.ReplaceAll(string(match), "mul(", "")
		stripped = stripped[:len(stripped)-1]
		values := strings.Split(stripped, ",")

		lhs, _ := strconv.Atoi(values[0])
		rhs, _ := strconv.Atoi(values[1])
		total1 += lhs * rhs
	}

	do := true
	for _, match := range re2.FindAll(b, -1) {
		s := string(match)
		if s == "do()" {
			do = true
			continue
		} else if s == "don't()" {
			do = false
			continue
		}

		if do {
			stripped := strings.ReplaceAll(s, "mul(", "")
			stripped = stripped[:len(stripped)-1]
			values := strings.Split(stripped, ",")

			lhs, _ := strconv.Atoi(values[0])
			rhs, _ := strconv.Atoi(values[1])
			total2 += lhs * rhs
		}
	}

	fmt.Printf("total1: %v\n", total1)
	fmt.Printf("total2: %v\n", total2)
}
