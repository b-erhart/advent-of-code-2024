package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := parseInput("./input.txt")

	fmt.Printf("Part 1 (safe reports):\n%d\n\n", numSafeReports(reports))

	fmt.Printf("Part 2 (dampened safe reports):\n%d\n", numDampenedSafeReports(reports))
}

func numDampenedSafeReports(reports [][]int) int {
	safe := 0

	for _, r := range reports {
		if dampenedReportIsSafe(r) {
			safe++
		}
	}

	return safe
}

func dampenedReportIsSafe(report []int) bool {
	if reportIsSafe(report) {
		return true
	}

	for i := range report {
		if reportIsSafe(removeElement(report, i)) {
			return true
		}
	}

	return false
}

func removeElement(slice []int, element int) []int {
	newSlice := make([]int, len(slice)-1)
	for i := 0; i < len(slice); i++ {
		if i < element {
			newSlice[i] = slice[i]
		} else if i > element {
			newSlice[i-1] = slice[i]
		}
	}
	return newSlice
}

func numSafeReports(reports [][]int) int {
	safe := 0
	for _, r := range reports {
		if reportIsSafe(r) {
			safe++
		}
	}

	return safe
}

func reportIsSafe(report []int) bool {
	allIncreasing := true
	allDecreasing := true
	validDifferences := true

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]

		if diff >= 0 {
			allIncreasing = false
		}
		if diff <= 0 {
			allDecreasing = false
		}

		if diff == 0 || diff < -3 || diff > 3 {
			validDifferences = false
		}
	}

	return (allIncreasing || allDecreasing) && validDifferences
}

func parseInput(path string) [][]int {
	var reports [][]int

	r, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("error opening input file: %v", err))
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	line := 0
	for scanner.Scan() {
		line++

		text := strings.Trim(scanner.Text(), " ")

		if len(text) == 0 {
			continue
		}

		parts := strings.Split(text, " ")
		levels := make([]int, len(parts))

		for i, s := range parts {
			levels[i], err = strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("Line %d, parsing value \"%s\": %v", line, s, err))
			}
		}

		reports = append(reports, levels)
	}

	return reports
}
