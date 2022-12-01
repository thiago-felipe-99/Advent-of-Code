// Package main contains the resolutions of the problems of the day 1 of 2022
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
	maxSum := 0
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > maxSum {
				maxSum = sum
			}

			sum = 0

			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += calories
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return maxSum
}

func main() {
	fmt.Println(puzzle1())
}
