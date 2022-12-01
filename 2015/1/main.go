// Package main contains the resolutions of the problems of the day 1 of 2022
package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func puzzle1() int {
	position := 0

	for _, character := range input {
		if character == '(' {
			position++
		} else if character == ')' {
			position--
		}
	}

	return position
}

func main() {
	fmt.Println(puzzle1())
}
