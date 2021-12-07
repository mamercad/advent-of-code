package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	return lines
}

func computePositionPart1(lines []string) ([]int, int) {
	coords := []int{0, 0}
	for _, line := range lines {
		s := strings.Split(line, " ")
		direction := s[0]
		distance, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			coords[0] += distance
		case "up":
			coords[1] -= distance
		case "down":
			coords[1] += distance
		}
	}
	return coords, coords[0] * coords[1]
}

func computePositionPart2(lines []string) ([]int, int) {
	coords := []int{0, 0}
	aim := 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		direction := s[0]
		distance, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			coords[0] += distance
			coords[1] += distance * aim
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return coords, coords[0] * coords[1]
}

func main() {
	lines := readInput("aoc02.dat")
	var coords, product = computePositionPart1(lines)
	fmt.Printf("part 1\nposition: [%v, %v]\nproduct: %v\n", coords[0], coords[1], product)
	coords, product = computePositionPart2(lines)
	fmt.Printf("part 2\nposition: [%v, %v]\nproduct: %v\n", coords[0], coords[1], product)
}
