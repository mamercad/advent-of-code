package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
				panic(lines[i])
			}

			nexVal, err := strconv.Atoi(lines[i+1])
			if err != nil {
				panic(lines[i+1])
			}

			if curVal < nexVal {
				increases++
			}
		}
	}

	fmt.Println(increases)
}
