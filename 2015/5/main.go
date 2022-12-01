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

		for index, character := range line {
			if character == 'a' || character == 'e' ||
				character == 'i' || character == 'o' || character == 'u' {
				vowels++
			}

			if index+1 >= lenLine {
				continue
			}

			if line[index] == line[index+1] {
				nice = true
			}

			nextCharacter := line[index : index+2]
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

func puzzle2() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	quantitesNiceStrings := 0

	for scanner.Scan() {
		line := scanner.Text()
		lenLine := len(line)
		twoLetters := map[string]struct {
			lastIndex, count int
		}{}
		niceLetterBetween := false
		niceTwoLetters := false

		for index := range line {
			if index+1 >= lenLine {
				continue
			}

			twoLetter := line[index : index+2]
			if value, isMapContainsKey := twoLetters[twoLetter]; isMapContainsKey {
				if value.lastIndex < index-1 {
					value.count++
					value.lastIndex = index
					twoLetters[twoLetter] = value
				}
			} else {
				twoLetters[twoLetter] = struct {
					lastIndex int
					count     int
				}{index, 1}
			}

			if index+2 >= lenLine {
				continue
			}

			if line[index] == line[index+2] {
				niceLetterBetween = true
			}
		}

		for _, value := range twoLetters {
			if value.count >= 2 {
				niceTwoLetters = true
			}
		}

		if niceLetterBetween && niceTwoLetters {
			quantitesNiceStrings++
		}
	}

	return quantitesNiceStrings
}

func main() {
	fmt.Println(puzzle2())
}
