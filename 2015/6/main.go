// Package main contains the resolutions of the problems of the day 6 of 2015
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const gridLen = 1000

func getPoint(point string) (x, y int) {
	points := strings.Split(point, ",")

	x, err := strconv.Atoi(points[0])
	if err != nil {
		panic(err)
	}

	y, err = strconv.Atoi(points[1])
	if err != nil {
		panic(err)
	}

	return x, y
}

func puzzle1() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	grid := [gridLen][gridLen]bool{}
	sum := 0

	var function func(bool) bool

	turnOn := func(_ bool) bool { return true }
	turnOff := func(_ bool) bool { return false }
	toggle := func(value bool) bool { return !value }

	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")

		if commands[0] == "toggle" {
			function = toggle
		} else if commands[1] == "on" {
			function = turnOn
		} else {
			function = turnOff
		}

		xStart, yStart := getPoint(commands[len(commands)-3])
		xEnd, yEnd := getPoint(commands[len(commands)-1])

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				grid[x][y] = function(grid[x][y])
			}
		}
	}

	for x := 0; x < gridLen; x++ {
		for y := 0; y < gridLen; y++ {
			if grid[x][y] {
				sum++
			}
		}
	}

	return sum
}

func puzzle2() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	grid := [gridLen][gridLen]int{}
	sum := 0

	var function func(int) int

	turnOn := func(brightness int) int { return brightness + 1 }
	turnOff := func(brightness int) int {
		if brightness <= 0 {
			return 0
		}
		return brightness - 1
	}
	toggle := func(brightness int) int { return brightness + 2 }

	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")

		if commands[0] == "toggle" {
			function = toggle
		} else if commands[1] == "on" {
			function = turnOn
		} else {
			function = turnOff
		}

		xStart, yStart := getPoint(commands[len(commands)-3])
		xEnd, yEnd := getPoint(commands[len(commands)-1])

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				grid[x][y] = function(grid[x][y])
			}
		}
	}

	for x := 0; x < gridLen; x++ {
		for y := 0; y < gridLen; y++ {
			sum += grid[x][y]
		}
	}

	return sum
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
