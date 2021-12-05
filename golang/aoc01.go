package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getWindow(lines []string, index int, length int) int {
	window := 0
	for i := 0; i < length; i++ {
		curVal, err := strconv.Atoi(lines[index+i])
		if err != nil {
			panic(err)
		}
		window += curVal
	}
	return window
}

func main() {

	file, err := os.ReadFile("aoc01.dat")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	increases := 0
	for i := range lines {
		if i+1 < len(lines) {
			curVal, err := strconv.Atoi(lines[i])
			if err != nil {
				panic(err)
			}

			nexVal, err := strconv.Atoi(lines[i+1])
			if err != nil {
				panic(err)
			}

			if curVal < nexVal {
				increases++
			}
		}
	}

	fmt.Printf("part 1 increases: %v\n", increases)

	increases = 0
	for i := range lines {
		if i+3 < len(lines) {
			win1 := getWindow(lines, i, 3)
			win2 := getWindow(lines, i+1, 3)
			if win1 < win2 {
				increases++
			}
		}
	}

	fmt.Printf("part 2 increases: %v\n", increases)
}
