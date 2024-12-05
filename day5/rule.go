package main

import "fmt"

type Rule struct {
	lhs int
	rhs int
}

func (r Rule) String() string {
	return fmt.Sprintf("(%d | %d)", r.lhs, r.rhs)
}
