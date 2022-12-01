// Package main contains the resolutions of the problems of the day 1 of 2015
package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func puzzle1() int {
	floor := 0

	for _, character := range input {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}
	}

	return floor
}

func puzzle2() int {
	floor := 0
	position := 0

	for _, character := range input {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}

		position++

		if floor < 0 {
			break
		}
	}

	return position
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
