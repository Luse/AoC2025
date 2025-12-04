package main

import (
	"fmt"
	"log"

	"koltrast.net/AOC2025/utils"
)

const (
	EMPTY_FLOOR = '.'
	PAPER_ROLL  = '@'
)

func getNeighbors(lines []string, row, col int) []rune {
	neighbors := []rune{}
	atTop := row == 0
	atBottom := row == len(lines)-1
	atLeft := col == 0
	atRight := col == len(lines[row])-1

	if !atTop {
		if !atLeft {
			neighbors = append(neighbors, rune(lines[row-1][col-1])) // top left
		}
		neighbors = append(neighbors, rune(lines[row-1][col])) // top
		if !atRight {
			neighbors = append(neighbors, rune(lines[row-1][col+1])) // top right
		}
	}

	if !atLeft {
		neighbors = append(neighbors, rune(lines[row][col-1])) // left
	}
	if !atRight {
		neighbors = append(neighbors, rune(lines[row][col+1])) // right
	}

	if !atBottom {
		if !atLeft {
			neighbors = append(neighbors, rune(lines[row+1][col-1])) // bottom left
		}
		neighbors = append(neighbors, rune(lines[row+1][col])) // bottom
		if !atRight {
			neighbors = append(neighbors, rune(lines[row+1][col+1])) // bottom right
		}
	}

	return neighbors
}

func countPaperRolls(neighbors []rune) int {
	count := 0
	for _, neighbor := range neighbors {
		if neighbor == PAPER_ROLL {
			count++
		}
	}
	return count
}

func part1(lines []string) int {
	accessiblePaperRolls := 0
	for row, line := range lines {
		for col, char := range line {
			if char == PAPER_ROLL {
				neighbors := getNeighbors(lines, row, col)
				paperRollCount := countPaperRolls(neighbors)
				if paperRollCount < 4 {
					accessiblePaperRolls++
				}
			}
		}
	}
	return accessiblePaperRolls
}

func part2(lines []string) int {
	totalRemoved := 0

	for {
		removed := 0
		newLines := make([]string, len(lines))
		copy(newLines, lines)

		for row, line := range lines {
			for col, char := range line {
				if char == PAPER_ROLL {
					neighbors := getNeighbors(lines, row, col)
					paperRollCount := countPaperRolls(neighbors)
					if paperRollCount < 4 {
						removed++
						newLines[row] = newLines[row][:col] + string(EMPTY_FLOOR) + newLines[row][col+1:]
					}
				}
			}
		}

		if removed == 0 {
			break
		}

		totalRemoved += removed
		lines = newLines
	}

	return totalRemoved
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
