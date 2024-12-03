package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	instructions := parseInput("./input.txt")

	fmt.Printf("Part 1 (uncorrupted sum):\n%d\n\n", addUncorruptedMuls(instructions))

	fmt.Printf("Part 2 (uncorrupted sum with conditions):\n%d\n\n", addConditionalUncorruptedMuls(instructions))
}

func addConditionalUncorruptedMuls(instructions string) int {
	sum := 0
	do := true

	r := regexp.MustCompile(`don't\(\)|do\(\)|mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := r.FindAllStringSubmatch(instructions, -1)

	if matches == nil {
		return sum
	}

	for _, m := range matches {
		if m[0] == "do()" {
			do = true
			continue
		}

		if m[0] == "don't()" {
			do = false
			continue
		}

		if do == false {
			continue
		}

		l, err := strconv.Atoi(m[1])
		if err != nil {
			panic(fmt.Errorf("error converting read number \"%s\": %v", m[1], err))
		}

		r, err := strconv.Atoi(m[2])
		if err != nil {
			panic(fmt.Errorf("error converting read number \"%s\": %v", m[2], err))
		}

		sum += l * r
	}

	return sum
}

func addUncorruptedMuls(instructions string) int {
	sum := 0

	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := r.FindAllStringSubmatch(instructions, -1)

	if matches == nil {
		return sum
	}

	for _, m := range matches {
		l, err := strconv.Atoi(m[1])
		if err != nil {
			panic(fmt.Errorf("error converting read number \"%s\": %v", m[1], err))
		}

		r, err := strconv.Atoi(m[2])
		if err != nil {
			panic(fmt.Errorf("error converting read number \"%s\": %v", m[2], err))
		}

		sum += l * r
	}

	return sum
}

func parseInput(path string) string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("error reading input file: %v", err))
	}

	return string(c)
}
