// Package main contains the resolutions of the problems of the day 2 of 2015
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

func puzzle1() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	dimensions := [3]int{0, 0, 0}
	sumTotalPaper := 0

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "x")
		for index := range dimensions {
			dimension, err := strconv.Atoi(values[index])
			if err != nil {
				panic(err)
			}

			dimensions[index] = dimension
		}

		side0 := dimensions[0] * dimensions[1]
		side1 := dimensions[0] * dimensions[2]
		side2 := dimensions[1] * dimensions[2]

		sumPaper := 2*side0 + 2*side1 + 2*side2 //nolint: gomnd

		if side0 < side1 { //nolint: nestif
			if side0 < side2 {
				sumPaper += side0
			} else {
				sumPaper += side2
			}
		} else {
			if side1 < side2 {
				sumPaper += side1
			} else {
				sumPaper += side2
			}
		}

		sumTotalPaper += sumPaper
	}

	return sumTotalPaper
}

func puzzle2() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	dimensions := [3]int{0, 0, 0}
	sumTotalRibbon := 0

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "x")
		for index := range dimensions {
			dimension, err := strconv.Atoi(values[index])
			if err != nil {
				panic(err)
			}

			dimensions[index] = dimension
		}

		sumRibbon := dimensions[0] * dimensions[1] * dimensions[2]

		//nolint: nestif,gomnd
		if dimensions[0] < dimensions[1] {
			if dimensions[1] < dimensions[2] {
				sumRibbon += 2*dimensions[0] + 2*dimensions[1]
			} else {
				sumRibbon += 2*dimensions[0] + 2*dimensions[2]
			}
		} else {
			if dimensions[0] < dimensions[2] {
				sumRibbon += 2*dimensions[1] + 2*dimensions[0]
			} else {
				sumRibbon += 2*dimensions[1] + 2*dimensions[2]
			}
		}

		sumTotalRibbon += sumRibbon
	}

	return sumTotalRibbon
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
