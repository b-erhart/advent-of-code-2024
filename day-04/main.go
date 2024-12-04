package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	puzzle := parseInput("./input.txt")

	fmt.Printf("Part 1 (num of xmas):\n%d\n\n", numberOfXmas(puzzle))

	fmt.Printf("Part 1 (num of x-mas):\n%d\n\n", numberOfXMas(puzzle))
}

func numberOfXMas(puzzle [][]rune) int {
	count := 0

	for i, row := range puzzle {
		for j, char := range row {
			if char != 'A' {
				continue
			}

			if i-1 < 0 || j-1 < 0 || i+1 >= len(puzzle) || j+1 >= len(puzzle[i]) {
				continue
			}

			topLeftBottomRight := (puzzle[i-1][j-1] == 'M' && puzzle[i+1][j+1] == 'S') || (puzzle[i-1][j-1] == 'S' && puzzle[i+1][j+1] == 'M')
			bottomLeftTopRight := (puzzle[i+1][j-1] == 'M' && puzzle[i-1][j+1] == 'S') || (puzzle[i+1][j-1] == 'S' && puzzle[i-1][j+1] == 'M')

			if topLeftBottomRight && bottomLeftTopRight {
				count++
			}
		}
	}

	return count
}

func numberOfXmas(puzzle [][]rune) int {
	count := 0

	for i, row := range puzzle {
		for j, char := range row {
			if char != 'X' {
				continue
			}

			for iStep := -1; iStep <= 1; iStep++ {
				for jStep := -1; jStep <= 1; jStep++ {
					if iStep == 0 && jStep == 0 {
						continue
					}

					if isXmas(puzzle, i, j, iStep, jStep) {
						count++
					}
				}
			}
		}
	}

	return count
}

func isXmas(puzzle [][]rune, iStart int, jStart int, iStep int, jStep int) bool {
	if iStart+3*iStep < 0 || iStart+3*iStep >= len(puzzle) {
		return false
	}

	if jStart+3*jStep < 0 || jStart+3*jStep >= len(puzzle[iStart]) {
		return false
	}

	x := puzzle[iStart][jStart] == 'X'
	m := puzzle[iStart+1*iStep][jStart+1*jStep] == 'M'
	a := puzzle[iStart+2*iStep][jStart+2*jStep] == 'A'
	s := puzzle[iStart+3*iStep][jStart+3*jStep] == 'S'

	return x && m && a && s
}

func parseInput(path string) [][]rune {
	puzzle := make([][]rune, 0)

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

		puzzle = append(puzzle, []rune(scanner.Text()))
	}

	return puzzle
}
