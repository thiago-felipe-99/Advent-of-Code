// Package main contains the resolutions of the problems of the day 5 of 2015
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func puzzle1() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	quantitesNiceStrings := 0

	for scanner.Scan() {
		line := scanner.Text()
		lenLine := len(line)
		vowels := 0
		nice := false

		for indice, character := range line {
			if character == 'a' || character == 'e' ||
				character == 'i' || character == 'o' || character == 'u' {
				vowels++
			}

			if indice+1 >= lenLine {
				continue
			}

			if line[indice] == line[indice+1] {
				nice = true
			}

			nextCharacter := line[indice : indice+2]
			if nextCharacter == "ab" || nextCharacter == "cd" ||
				nextCharacter == "pq" || nextCharacter == "xy" {
				nice = false

				break
			}
		}

		if nice && vowels >= 3 {
			quantitesNiceStrings++
		}
	}

	return quantitesNiceStrings
}

func main() {
	fmt.Println(puzzle1())
}
