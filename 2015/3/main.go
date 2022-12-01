// Package main contains the resolutions of the problems of the day 2 of 2015
package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type houseLocation struct {
	x, y int
}

func puzzle1() int {
	currentHouse := houseLocation{0, 0}
	housesAlreadyVisited := []houseLocation{currentHouse}

	for _, character := range input {
		switch character {
		case '^':
			currentHouse.y++
		case 'v':
			currentHouse.y--
		case '>':
			currentHouse.x++
		case '<':
			currentHouse.x--
		}

		visited := false

		for _, house := range housesAlreadyVisited {
			if currentHouse.x == house.x && currentHouse.y == house.y {
				visited = true

				break
			}
		}

		if !visited {
			housesAlreadyVisited = append(housesAlreadyVisited, currentHouse)
		}
	}

	return len(housesAlreadyVisited)
}

func main() {
	fmt.Println(puzzle1())
}
