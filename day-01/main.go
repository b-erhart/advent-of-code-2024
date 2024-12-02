package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	left, right := parseInput("./input.txt")

	fmt.Printf("Part 1 (total distance):\n%d\n\n", totalDistance(left, right))

	fmt.Printf("Part 2 (similiarity score):\n%d\n", similiarityScore(left, right))
}

func similiarityScore(left []int, right []int) int {
	if len(left) != len(right) {
		panic("left and right must be equal size")
	}

	appearances := make(map[int]int)
	for _, r := range right {
		appearances[r] += 1
	}

	score := 0
	for _, l := range left {
		score += l * appearances[l]
	}

	return score
}

func totalDistance(left []int, right []int) int {
	if len(left) != len(right) {
		panic("left and right must be equal size")
	}

	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]

		if diff < 0 {
			diff = -diff
		}

		total += diff
	}

	return total
}

func parseInput(path string) ([]int, []int) {
	var left []int
	var right []int

	r, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("error opening input file: %v", err))
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	line := 0
	for scanner.Scan() {
		line++

		if len(scanner.Text()) == 0 {
			continue
		}

		parts := strings.Split(scanner.Text(), "   ")

		if len(parts) != 2 {
			panic(fmt.Sprintf("Line %d: invalid number of parts (%d, must be 2)", line, len(parts)))
		}

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("Line %d, parsing left value: %v", line, err))
		}

		r, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Sprintf("Line %d, parsing right value: %v", line, err))
		}

		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
