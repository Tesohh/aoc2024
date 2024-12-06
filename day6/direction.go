package main

import "fmt"

type Direction uint8

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

func (d *Direction) Turn90Deg() {
	(*d) += 1
	(*d) %= 4
}

func (d Direction) Vector() (x int, y int) {
	switch d {
	case Top:
		return 0, -1
	case Right:
		return 1, 0
	case Bottom:
		return 0, 1
	case Left:
		return -1, 0
	default:
		panic(fmt.Sprintf("unexpected main.Direction: %#v", d))
	}
}

func (d Direction) String() string {
	switch d {
	case Top:
		return "Top"
	case Right:
		return "Right"
	case Bottom:
		return "Bottom"
	case Left:
		return "Left"
	default:
		panic(fmt.Sprintf("unexpected main.Direction: %#v", d))
	}

}

func (d Direction) Char() byte {
	switch d {
	case Top:
		return '^'
	case Right:
		return '>'
	case Bottom:
		return 'v'
	case Left:
		return '<'
	default:
		panic(fmt.Sprintf("unexpected main.Direction: %#v", d))
	}

}
