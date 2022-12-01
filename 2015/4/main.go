// Package main contains the resolutions of the problems of the day 4 of 2015
package main

import (
	"crypto/md5"
	"fmt"
)

const input = "yzbqklnj"

func puzzle1() int {
	count := 0

	for {
		data := fmt.Sprintf("%v%d", input, count)
		checksum := fmt.Sprintf("%x", md5.Sum([]byte(data)))

		if checksum[0:5] == "00000" {
			break
		}

		count++
	}

	return count
}

func main() {
	fmt.Println(puzzle1())
}
